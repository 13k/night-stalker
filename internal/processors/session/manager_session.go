package session

import (
	"golang.org/x/xerrors"

	nsctx "github.com/13k/night-stalker/internal/context"
)

func (p *Manager) startSession() error {
	if p.session != nil && !p.session.IsReady() {
		p.log.Warn("called startSession() with live session")
		return nil
	}

	if p.conn == nil {
		return xerrors.New("not connected")
	}

	var err error

	p.session, err = newSession(p.ctx, p.conn, sessionOptions{
		Log: p.log,
	})

	if err != nil {
		return xerrors.Errorf("error starting dota session: %w", err)
	}

	return nil
}

func (p *Manager) closeSession() {
	if p.session == nil {
		p.log.Warn("called closeSession() without session")
		return
	}

	p.session.Close()
	p.session = nil

	if p.supervisor != nil {
		p.supervisor.Wait()
		p.supervisor = nil
	}
}

func (p *Manager) startSupervisor() {
	if p.conn == nil {
		p.log.Warn("called startSupervisor() without connection")
		return
	}

	if p.session == nil {
		p.log.Warn("called startSupervisor() without session")
		return
	}

	p.setupSupervisor()

	ctx := p.session.Context()
	ctx = nsctx.WithSteam(ctx, p.conn.Client())
	ctx = nsctx.WithDota(ctx, p.session.Client())

	go p.supervisor.Start(ctx)
}
