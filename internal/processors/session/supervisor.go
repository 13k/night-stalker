package session

import (
	"context"
	"time"

	"cirello.io/oversight"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nslog "github.com/13k/night-stalker/internal/logger"
	nscomm "github.com/13k/night-stalker/internal/processors/comm"
	nsgc "github.com/13k/night-stalker/internal/processors/gc"
	nslm "github.com/13k/night-stalker/internal/processors/livematches"
	nsmdtl "github.com/13k/night-stalker/internal/processors/matchdetails"
	nsrts "github.com/13k/night-stalker/internal/processors/rtstats"
	nstv "github.com/13k/night-stalker/internal/processors/tvgames"
)

type supervisorOptions struct {
	Log                   *nslog.Logger
	Bus                   *nsbus.Bus
	ShutdownTimeout       time.Duration
	TVGamesInterval       time.Duration
	RealtimeStatsPoolSize int
	RealtimeStatsInterval time.Duration
	MatchInfoPoolSize     int
	MatchInfoInterval     time.Duration
}

type supervisor struct {
	log    *nslog.Logger
	root   *oversight.Tree
	waitCh chan struct{}
}

func newSupervisor(options supervisorOptions) *supervisor {
	dispatcherSpec := nsgc.NewDispatcher(nsgc.DispatcherOptions{
		Log:             options.Log,
		Bus:             options.Bus,
		ShutdownTimeout: options.ShutdownTimeout,
	}).ChildSpec()

	tvGamesSpec := nstv.NewWatcher(nstv.WatcherOptions{
		Log:             options.Log,
		Bus:             options.Bus,
		Interval:        options.TVGamesInterval,
		ShutdownTimeout: options.ShutdownTimeout,
	}).ChildSpec()

	rtStatsSpec := nsrts.NewMonitor(nsrts.MonitorOptions{
		Log:             options.Log,
		Bus:             options.Bus,
		PoolSize:        options.RealtimeStatsPoolSize,
		Interval:        options.RealtimeStatsInterval,
		ShutdownTimeout: options.ShutdownTimeout,
	}).ChildSpec()

	matchInfoSpec := nsmdtl.NewMonitor(nsmdtl.MonitorOptions{
		Log:             options.Log,
		Bus:             options.Bus,
		PoolSize:        options.MatchInfoPoolSize,
		Interval:        options.MatchInfoInterval,
		ShutdownTimeout: options.ShutdownTimeout,
	}).ChildSpec()

	chatSpec := nscomm.NewChat(nscomm.ChatOptions{
		Log:             options.Log,
		Bus:             options.Bus,
		ShutdownTimeout: options.ShutdownTimeout,
	}).ChildSpec()

	liveMatchesSpec := nslm.NewCollector(nslm.CollectorOptions{
		Log:             options.Log,
		Bus:             options.Bus,
		ShutdownTimeout: options.ShutdownTimeout,
	}).ChildSpec()

	log := options.Log.WithPackage("supervisor")

	tree := oversight.New(
		oversight.NeverHalt(),
		oversight.WithRestartStrategy(oversight.OneForOne()),
		oversight.WithLogger(log.OversightLogger()),
		oversight.Process(
			dispatcherSpec,
			tvGamesSpec,
			rtStatsSpec,
			matchInfoSpec,
			chatSpec,
			liveMatchesSpec,
		),
	)

	return &supervisor{
		root:   tree,
		log:    log,
		waitCh: make(chan struct{}),
	}
}

func (s *supervisor) start(ctx context.Context) {
	defer s.finished()

	if err := s.root.Start(ctx); err != nil {
		s.log.WithError(err).Error("supervisor error")
	}
}

func (s *supervisor) wait() {
	if s.waitCh != nil {
		<-s.waitCh
	}
}

func (s *supervisor) finished() {
	if s.waitCh != nil {
		close(s.waitCh)
	}
}
