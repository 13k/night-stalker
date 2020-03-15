package session

import (
	"context"

	"golang.org/x/xerrors"

	nsctx "github.com/13k/night-stalker/internal/context"
	nsproc "github.com/13k/night-stalker/internal/processors"
)

func (p *Manager) setupContext(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextDatabase)
	}

	p.ctx = ctx

	return nil
}

func (p *Manager) setupSupervisor() {
	if p.supervisor != nil {
		return
	}

	p.supervisor = newSupervisor(supervisorOptions{
		Log:                   p.options.Log,
		Bus:                   p.bus,
		ShutdownTimeout:       p.options.ShutdownTimeout,
		TVGamesInterval:       p.options.TVGamesInterval,
		RealtimeStatsPoolSize: p.options.RealtimeStatsPoolSize,
		RealtimeStatsInterval: p.options.RealtimeStatsInterval,
		MatchInfoPoolSize:     p.options.MatchInfoPoolSize,
		MatchInfoInterval:     p.options.MatchInfoInterval,
	})
}
