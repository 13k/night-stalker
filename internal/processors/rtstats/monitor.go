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
	"github.com/jinzhu/gorm"
	"github.com/panjf2000/ants/v2"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsbussub "github.com/13k/night-stalker/internal/bus/subscribers"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nsrt "github.com/13k/night-stalker/internal/runtime"
	"github.com/13k/night-stalker/models"
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
	db            *gorm.DB
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

func (p *Monitor) workerFunc(w *worker) func() {
	return func() {
		var err error
		var stats *models.LiveMatchStats

		defer func() {
			if v := recover(); v != nil {
				err = xerrors.Errorf("worker panic: %w", &errWorkerPanic{
					LiveMatch: w.liveMatch,
					Value:     v,
				})
			}

			if err != nil {
				p.handleError(xerrors.Errorf("worker error: %w", err))
			}
		}()

		stats, err = w.Run(p.ctx)

		if err == nil && stats != nil {
			p.flusher.Add(stats)
		}
	}
}

func (p *Monitor) enqueueWorker(liveMatch *models.LiveMatch) error {
	if p.ctx.Err() != nil {
		return xerrors.Errorf("error enqueuing worker: %w", &errWorkerSubmitFailure{
			LiveMatch: liveMatch,
			Err:       p.ctx.Err(),
		})
	}

	w := &worker{
		db:         p.db,
		api:        p.apiMatchStats,
		activeReqs: p.activeReqs,
		liveMatch:  liveMatch,
	}

	err := p.workerPool.Submit(p.workerFunc(w))

	if err != nil {
		return xerrors.Errorf("error enqueuing worker: %w", &errWorkerSubmitFailure{
			LiveMatch: liveMatch,
			Err:       err,
		})
	}

	return nil
}

func (p *Monitor) handleError(err error) {
	if xerrors.Is(err, context.Canceled) {
		// safe to ignore
		return
	}

	if e := (&errInvalidResponse{}); xerrors.As(err, &e) {
		// safe to ignore
		return
	}

	if e := (&errRequestInProgress{}); xerrors.As(err, &e) {
		p.log.WithOFields(
			"match_id", e.LiveMatch.MatchID,
			"server_id", e.LiveMatch.ServerID.ToUint64(),
		).Warn("request in progress")

		return
	}

	if e := (&errWorkerSubmitFailure{}); xerrors.As(err, &e) {
		p.log.WithOFields(
			"match_id", e.LiveMatch.MatchID,
			"server_id", e.LiveMatch.ServerID.ToUint64(),
		).WithError(e.Err).Error("error submitting worker")
	} else if e := (&errWorkerPanic{}); xerrors.As(err, &e) {
		p.log.WithOFields(
			"match_id", e.LiveMatch.MatchID,
			"server_id", e.LiveMatch.ServerID.ToUint64(),
			"panic", e.Value,
		).Error("recovered worker panic")
	} else if e := (&errRequestFailure{}); xerrors.As(err, &e) {
		l := p.log.WithOFields(
			"match_id", e.LiveMatch.MatchID,
			"server_id", e.LiveMatch.ServerID.ToUint64(),
		).WithError(e.Err)

		fpath, eErr := handleRequestFailureError(e)

		if eErr != nil {
			l.WithError(eErr).Error("error handling request failure error")
			p.log.Errorx(eErr)
		} else if fpath != "" {
			l = l.WithField("error_file", fpath)
		}

		l.Error("request failed")
	} else if e := (&errStatsSaveFailure{}); xerrors.As(err, &e) {
		p.log.WithOFields(
			"match_id", e.LiveMatch.MatchID,
			"server_id", e.LiveMatch.ServerID.ToUint64(),
		).WithError(e.Err).Error("error saving stats")
	} else {
		p.log.WithError(err).Error("rtstats error")
	}

	p.log.Errorx(err)
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
