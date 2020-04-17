package rtstats

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"sync"
	"time"

	"cirello.io/oversight"
	gsdota2 "github.com/13k/geyser/dota2"
	"github.com/panjf2000/ants/v2"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsbussub "github.com/13k/night-stalker/internal/bus/subscribers"
	nsctx "github.com/13k/night-stalker/internal/context"
	nsdb "github.com/13k/night-stalker/internal/db"
	nserr "github.com/13k/night-stalker/internal/errors"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nsrt "github.com/13k/night-stalker/internal/runtime"
	nsm "github.com/13k/night-stalker/models"
)

const (
	processorName   = "rtstats"
	flusherCap      = 10
	flusherInterval = 5 * time.Second
)

type MonitorOptions struct {
	Log             *nslog.Logger
	Bus             *nsbus.Bus
	PoolSize        int
	Interval        time.Duration
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Monitor)(nil)

type Monitor struct {
	options       MonitorOptions
	ctx           context.Context
	log           *nslog.Logger
	db            *nsdb.DB
	workerPool    *ants.Pool
	api           *gsdota2.Client
	apiMatchStats *gsdota2.DOTA2MatchStats
	bus           *nsbus.Bus
	liveMatches   *nsbussub.LiveMatches
	activeReqs    *sync.Map
	flusher       *flusher
}

func NewMonitor(options MonitorOptions) *Monitor {
	log := options.Log.WithPackage(processorName)

	return &Monitor{
		options:     options,
		log:         log,
		bus:         options.Bus,
		liveMatches: nsbussub.NewLiveMatchesSubscriber(options.Bus),
		activeReqs:  &sync.Map{},
		flusher: newFlusher(&flusherOptions{
			Log:      log,
			Bus:      options.Bus,
			Interval: flusherInterval,
			Cap:      flusherCap,
		}),
	}
}

func (p *Monitor) ChildSpec() oversight.ChildProcessSpecification {
	var shutdown oversight.Shutdown

	if p.options.ShutdownTimeout > 0 {
		shutdown = oversight.Timeout(p.options.ShutdownTimeout)
	} else {
		shutdown = oversight.Infinity()
	}

	return oversight.ChildProcessSpecification{
		Name:     processorName,
		Start:    p.Start,
		Restart:  oversight.Transient(),
		Shutdown: shutdown,
	}
}

func (p *Monitor) Start(ctx context.Context) (err error) {
	defer nsrt.RecoverError(p.log, &err)

	err = p.start(ctx)

	if err != nil {
		p.handleError(err)
	}

	return err
}

func (p *Monitor) start(ctx context.Context) error {
	if err := p.setupContext(ctx); err != nil {
		return xerrors.Errorf("error setting up context: %w", err)
	}

	if err := p.setupAPI(); err != nil {
		return xerrors.Errorf("error setting up API: %w", err)
	}

	if err := p.setupWorkerPool(); err != nil {
		return xerrors.Errorf("error setting up worker pool: %w", err)
	}

	p.liveMatches.Start(p.ctx)
	p.flusher.Start(p.ctx)

	go p.flusherLoop()

	return p.loop()
}

func (p *Monitor) stop(t *time.Ticker) {
	t.Stop()
	p.teardownWorkerPool()
	p.ctx = nil
	p.log.Warn("stop")
}

func (p *Monitor) setupContext(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextDatabase)
	}

	if p.api = nsctx.GetDotaAPI(ctx); p.api == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextDotaAPI)
	}

	p.ctx = ctx

	return nil
}

func (p *Monitor) setupAPI() error {
	if p.apiMatchStats != nil {
		return nil
	}

	var err error

	if p.apiMatchStats, err = p.api.DOTA2MatchStats(); err != nil {
		return xerrors.Errorf("error creating API interface: %w", err)
	}

	return nil
}

func (p *Monitor) setupWorkerPool() error {
	if p.workerPool != nil {
		return nil
	}

	var err error

	if p.workerPool, err = ants.NewPool(p.options.PoolSize); err != nil {
		return xerrors.Errorf("error starting worker pool: %w", err)
	}

	return nil
}

func (p *Monitor) teardownWorkerPool() {
	if p.workerPool != nil {
		p.workerPool.Release()
		p.workerPool = nil
	}
}

func (p *Monitor) loop() error {
	t := time.NewTicker(p.options.Interval)

	defer p.stop(t)

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case <-t.C:
			p.tick()
		}
	}
}

// flusher closes the errors channel when finished
func (p *Monitor) flusherLoop() {
	for {
		err, ok := <-p.flusher.Errors()

		if !ok {
			return
		}

		if err != nil {
			p.handleError(xerrors.Errorf("error flushing live match stats: %w", err))
		}
	}
}

func (p *Monitor) tick() {
	liveMatches := p.liveMatches.Get()

	if len(liveMatches) == 0 {
		return
	}

	p.log.
		WithField("count", len(liveMatches)).
		Debug("requesting stats")

	for _, liveMatch := range liveMatches {
		if err := p.enqueueWorker(liveMatch); err != nil {
			p.handleError(xerrors.Errorf("error enqueuing worker: %w", err))
		}
	}
}

func (p *Monitor) enqueueWorker(liveMatch *nsm.LiveMatch) error {
	if p.ctx.Err() != nil {
		return &errWorkerSubmitFailure{
			LiveMatch: liveMatch,
			Err:       nserr.Wrap("error enqueuing worker", p.ctx.Err()),
		}
	}

	w := &worker{
		ctx:          p.ctx,
		db:           p.db,
		api:          p.apiMatchStats,
		activeReqs:   p.activeReqs,
		liveMatch:    liveMatch,
		results:      p.flusher,
		errorHandler: p.handleError,
	}

	err := p.workerPool.Submit(w.Run)

	if err != nil {
		return &errWorkerSubmitFailure{
			LiveMatch: liveMatch,
			Err:       nserr.Wrap("error enqueuing worker", err),
		}
	}

	return nil
}

func (p *Monitor) handleError(err error) {
	// safe to ignore
	if xerrors.Is(err, context.Canceled) {
		return
	}

	// safe to ignore
	if e := (&errInvalidResponse{}); xerrors.As(err, &e) {
		return
	}

	if e := (&errRequestInProgress{}); xerrors.As(err, &e) {
		p.log.WithOFields(
			"match_id", e.LiveMatch.MatchID,
			"server_id", e.LiveMatch.ServerID.ToUint64(),
		).Warn("request in progress")

		return
	}

	if e := (&errInvalidResponseStatus{}); xerrors.As(err, &e) {
		p.log.WithOFields(
			"match_id", e.LiveMatch.MatchID,
			"server_id", e.LiveMatch.ServerID.ToUint64(),
			"status", e.Response.StatusCode(),
		).Warn("server returned invalid status")

		return
	}

	msg := fmt.Sprintf("%s error", processorName)
	l := p.log

	if e := (&errWorkerSubmitFailure{}); xerrors.As(err, &e) {
		msg = "submit worker failure"
		l = l.WithOFields(
			"match_id", e.LiveMatch.MatchID,
			"server_id", e.LiveMatch.ServerID.ToUint64(),
		)
	} else if e := (&errWorkerPanic{}); xerrors.As(err, &e) {
		msg = "recovered worker panic"
		l = l.WithOFields(
			"match_id", e.LiveMatch.MatchID,
			"server_id", e.LiveMatch.ServerID.ToUint64(),
			"panic", e.Value,
		)
	} else if e := (&errRequestFailure{}); xerrors.As(err, &e) {
		msg = "request failure"
		l = l.WithOFields(
			"match_id", e.LiveMatch.MatchID,
			"server_id", e.LiveMatch.ServerID.ToUint64(),
		)

		fpath, eErr := handleRequestFailureError(e)

		if eErr != nil {
			l.WithError(eErr).Error("error handling request failure")
		} else if fpath != "" {
			l = l.WithField("error_file", fpath)
		}
	} else if e := (&errStatsSaveFailure{}); xerrors.As(err, &e) {
		msg = "stats save failure"
		l = l.WithOFields(
			"match_id", e.LiveMatch.MatchID,
			"server_id", e.LiveMatch.ServerID.ToUint64(),
		)
	}

	l.WithError(err).Error(msg)
}

func handleRequestFailureError(reqErr *errRequestFailure) (string, error) {
	if reqErr.Response == nil {
		return "", nil
	}

	if reqErr.Response.Request == nil {
		return "", nil
	}

	if reqErr.Response.Request.RawRequest == nil {
		return "", nil
	}

	if reqErr.Response.RawResponse == nil {
		return "", nil
	}

	errorsDir := filepath.Join(os.TempDir(), "rtstats_errors")
	err := os.MkdirAll(errorsDir, 0777)

	if err != nil {
		return "", xerrors.Errorf("error creating errors directory: %w", err)
	}

	dump, err := dumpHTTPTransaction(
		reqErr.Response.Request.RawRequest,
		reqErr.Response.RawResponse,
		reqErr.Response.Body(),
	)

	if err != nil {
		return "", xerrors.Errorf("error dumping response: %w", err)
	}

	filename := fmt.Sprintf("%d.err", reqErr.LiveMatch.MatchID)
	fpath := filepath.Join(errorsDir, filename)

	if err := ioutil.WriteFile(fpath, dump, 0666); err != nil {
		return "", xerrors.Errorf("error creating error file: %w", err)
	}

	return fpath, nil
}

func dumpHTTPTransaction(req *http.Request, res *http.Response, resBody []byte) ([]byte, error) {
	dump := &bytes.Buffer{}

	reqDump, err := httputil.DumpRequestOut(req, true)

	if err != nil {
		return nil, err
	}

	if _, err = dump.Write(reqDump); err != nil {
		return nil, err
	}

	if _, err = dump.WriteString("\n-----\n"); err != nil {
		return nil, err
	}

	resDump, err := httputil.DumpResponse(res, false)

	if err != nil {
		return nil, err
	}

	if _, err := dump.Write(resDump); err != nil {
		return nil, err
	}

	if resBody != nil {
		if _, err := dump.Write(resBody); err != nil {
			return nil, err
		}
	}

	return dump.Bytes(), nil
}
