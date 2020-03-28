package session

import (
	"context"
	"time"

	"cirello.io/oversight"
	"github.com/faceit/go-steam"
	"github.com/faceit/go-steam/netutil"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nslog "github.com/13k/night-stalker/internal/logger"
)

type supervisorOptions struct {
	Log             *nslog.Logger
	Bus             *nsbus.Bus
	Addr            *netutil.PortAddr
	Credentials     *Credentials
	MachineHash     steam.SentryHash
	LoginKey        string
	ShutdownTimeout time.Duration
}

type supervisor struct {
	log  *nslog.Logger
	tree *oversight.Tree
}

func newSupervisor(options supervisorOptions) *supervisor {
	log := options.Log.WithPackage("supervisor")

	steamSession := newSteamSession((steamSessionOptions)(options))

	dotaSession := newDotaSession(dotaSessionOptions{
		Log:             options.Log,
		Bus:             options.Bus,
		ShutdownTimeout: options.ShutdownTimeout,
	})

	tree := oversight.New(
		oversight.NeverHalt(),
		oversight.WithRestartStrategy(oversight.OneForOne()),
		oversight.WithLogger(log.OversightLogger()),
		oversight.Process(
			steamSession.ChildSpec(),
			dotaSession.ChildSpec(),
		),
	)

	return &supervisor{
		log:  log,
		tree: tree,
	}
}

func (s *supervisor) Start(ctx context.Context) error {
	return s.tree.Start(ctx)
}
