package session

import (
	"context"

	"github.com/faceit/go-steam"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nsdota2 "github.com/13k/night-stalker/internal/dota2"
	nsproc "github.com/13k/night-stalker/internal/processors"
)

func (p *Manager) setup(ctx context.Context) error {
	if err := p.setupContext(ctx); err != nil {
		return xerrors.Errorf("error setting up context: %w", err)
	}

	if err := p.loadLogin(); err != nil {
		return xerrors.Errorf("error loading login info: %w", err)
	}

	if p.isSuspended() {
		return xerrors.Errorf("fatal error: %w", &nsdota2.ErrClientSuspended{
			Until: p.login.SuspendedUntil.Time,
		})
	}

	if err := p.setupSupervisor(); err != nil {
		return xerrors.Errorf("error setting up supervisor: %w", err)
	}

	p.busSubscribe()

	return nil
}

func (p *Manager) setupContext(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextDatabase)
	}

	if p.steam = nsctx.GetSteam(ctx); p.steam == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextSteamClient)
	}

	if p.dota = nsctx.GetDota(ctx); p.dota == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextDotaClient)
	}

	p.ctx, p.cancel = context.WithCancel(ctx)

	return nil
}

func (p *Manager) setupSupervisor() error {
	addr, err := p.randomServerAddr()

	if err != nil {
		return xerrors.Errorf("error loading random steam server addr: %w", err)
	}

	p.supervisor = newSupervisor(supervisorOptions{
		Log:         p.log,
		Bus:         p.bus,
		Addr:        addr,
		Credentials: p.options.Credentials,
		MachineHash: steam.SentryHash(p.login.MachineHash),
		LoginKey:    p.login.LoginKey,
	})

	return nil
}

func (p *Manager) busSubscribe() {
	if p.busSubSteamEvents == nil {
		p.busSubSteamEvents = p.bus.Sub(nsbus.TopicSteamEvents)
	}

	if p.busSubSessionSteam == nil {
		p.busSubSessionSteam = p.bus.Sub(nsbus.TopicSessionSteam)
	}

	if p.busSubSessionDota == nil {
		p.busSubSessionDota = p.bus.Sub(nsbus.TopicSessionDota)
	}
}

func (p *Manager) busUnsubscribe() {
	if p.busSubSteamEvents != nil {
		p.bus.Unsub(p.busSubSteamEvents)
		p.busSubSteamEvents = nil
	}

	if p.busSubSessionSteam != nil {
		p.bus.Unsub(p.busSubSessionSteam)
		p.busSubSessionSteam = nil
	}

	if p.busSubSessionDota != nil {
		p.bus.Unsub(p.busSubSessionDota)
		p.busSubSessionDota = nil
	}
}
