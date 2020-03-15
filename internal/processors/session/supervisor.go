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
	log  *nslog.Logger
	root *oversight.Tree
	wait chan struct{}
}

func newSupervisor(options supervisorOptions) *supervisor {
	dispatcher := nsgc.NewDispatcher(nsgc.DispatcherOptions{
		Log:             options.Log,
		Bus:             options.Bus,
		ShutdownTimeout: options.ShutdownTimeout,
	})

	tvGames := nstv.NewWatcher(nstv.WatcherOptions{
		Log:             options.Log,
		Bus:             options.Bus,
		Interval:        options.TVGamesInterval,
		ShutdownTimeout: options.ShutdownTimeout,
	})

	rtStats := nsrts.NewMonitor(nsrts.MonitorOptions{
		Log:             options.Log,
		Bus:             options.Bus,
		PoolSize:        options.RealtimeStatsPoolSize,
		Interval:        options.RealtimeStatsInterval,
		ShutdownTimeout: options.ShutdownTimeout,
	})

	matchDetails := nsmdtl.NewMonitor(nsmdtl.MonitorOptions{
		Log:             options.Log,
		Bus:             options.Bus,
		PoolSize:        options.MatchInfoPoolSize,
		Interval:        options.MatchInfoInterval,
		ShutdownTimeout: options.ShutdownTimeout,
	})

	chat := nscomm.NewChat(nscomm.ChatOptions{
		Log:             options.Log,
		Bus:             options.Bus,
		ShutdownTimeout: options.ShutdownTimeout,
	})

	liveMatches := nslm.NewCollector(nslm.CollectorOptions{
		Log:             options.Log,
		Bus:             options.Bus,
		ShutdownTimeout: options.ShutdownTimeout,
	})

	log := options.Log.WithPackage(processorName).WithPackage("supervisor")

	tree := oversight.New(
		oversight.NeverHalt(),
		oversight.WithRestartStrategy(oversight.OneForOne()),
		oversight.WithLogger(log.OversightLogger()),
		oversight.Process(
			chat.ChildSpec(),
			liveMatches.ChildSpec(),
			dispatcher.ChildSpec(),
			tvGames.ChildSpec(),
			rtStats.ChildSpec(),
			matchDetails.ChildSpec(),
		),
	)

	return &supervisor{
		root: tree,
		log:  log,
	}
}

func (s *supervisor) Start(ctx context.Context) {
	s.wait = make(chan struct{})

	defer func() {
		close(s.wait)
		s.wait = nil
		s.log.Trace("finished")
	}()

	s.log.Trace("starting")

	if err := s.root.Start(ctx); err != nil {
		s.log.WithError(err).Error("supervisor error")
	}
}

func (s *supervisor) Wait() {
	if s.wait != nil {
		s.log.Trace("waiting finish")
		<-s.wait
		s.log.Trace("wait ended")
	}
}
