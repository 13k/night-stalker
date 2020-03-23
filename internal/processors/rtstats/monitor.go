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
	geyserd2 "github.com/13k/geyser/dota2"
	"github.com/jinzhu/gorm"
	"github.com/panjf2000/ants/v2"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nscol "github.com/13k/night-stalker/internal/collections"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
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
	options               MonitorOptions
	ctx                   context.Context
	log                   *nslog.Logger
	db                    *gorm.DB
	workerPool            *ants.Pool
	api                   *geyserd2.Client
	apiMatchStats         *geyserd2.DOTA2MatchStats
	bus                   *nsbus.Bus
	busLiveMatchesReplace *nsbus.Subscription
	liveMatchesMtx        sync.RWMutex
	liveMatches           nscol.LiveMatches
	activeReqs            *sync.Map
	flusher               *flusher
}

func NewMonitor(options MonitorOptions) *Monitor {
	p := &Monitor{
		options:    options,
		log:        options.Log.WithPackage(processorName),
		bus:        options.Bus,
		activeReqs: &sync.Map{},
	}

	p.busSubscribe()

	return p
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

	p.busSubscribe()
	p.setupFlusher()

	go p.flusherLoop()

	return p.loop()
}

func (p *Monitor) stop(t *time.Ticker) {
	t.Stop()
	p.busUnsubscribe()
	p.teardownWorkerPool()
	p.teardownFlusher()
	p.ctx = nil
	p.log.Warn("stop")
}

func (p *Monitor) busSubscribe() {
	if p.busLiveMatchesReplace == nil {
		p.busLiveMatchesReplace = p.bus.Sub(nsbus.TopicLiveMatchesReplace)
	}
}

func (p *Monitor) busUnsubscribe() {
	if p.busLiveMatchesReplace != nil {
		p.bus.Unsub(p.busLiveMatchesReplace)
		p.busLiveMatchesReplace = nil
	}
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

func (p *Monitor) setupFlusher() {
	p.flusher = newFlusher(&flusherOptions{
		Log:      p.log,
		Bus:      p.bus,
		Interval: flusherInterval,
		Cap:      flusherCap,
	})

	p.flusher.Start(p.ctx)
}

func (p *Monitor) teardownFlusher() {
	p.flusher = nil
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
		case busmsg, ok := <-p.busLiveMatchesReplace.C:
			if !ok {
				return xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: p.busLiveMatchesReplace,
				})
			}

			if msg, ok := busmsg.Payload.(*nsbus.LiveMatchesChangeMessage); ok {
				p.handleLiveMatchesChange(msg)
			}
		}
	}
}

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

func (p *Monitor) handleLiveMatchesChange(msg *nsbus.LiveMatchesChangeMessage) {
	if msg.Op != nspb.CollectionOp_COLLECTION_OP_REPLACE {
		p.log.WithField("op", msg.Op.String()).Warn("ignored live matches change message")
		return
	}

	p.log.
		WithField("count", len(msg.Matches)).
		Debug("received live matches")

	p.liveMatchesMtx.Lock()
	defer p.liveMatchesMtx.Unlock()
	p.liveMatches = msg.Matches
}

func (p *Monitor) tick() {
	p.liveMatchesMtx.RLock()
	defer p.liveMatchesMtx.RUnlock()

	if len(p.liveMatches) == 0 {
		return
	}

	p.log.
		WithField("count", len(p.liveMatches)).
		Debug("requesting stats")

	for _, liveMatch := range p.liveMatches {
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
