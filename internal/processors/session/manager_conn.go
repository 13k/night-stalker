package session

import (
	"context"

	"github.com/faceit/go-steam"
	"github.com/faceit/go-steam/netutil"
	"golang.org/x/xerrors"
)

func (p *Manager) connect() error {
	if p.conn != nil && !p.conn.IsReady() {
		p.log.Warn("called connect() with connection")
		return nil
	}

	server, err := p.randomServer()

	if err != nil {
		return xerrors.Errorf("error loading steam server: %w", err)
	}

	var addr *netutil.PortAddr

	if server != nil {
		addr = netutil.ParsePortAddr(server.Address)

		if addr == nil {
			return xerrors.Errorf("error parsing server address: %w", &ErrInvalidServerAddress{
				Address: server.Address,
			})
		}
	}

	p.conn, err = newConnection(p.ctx, connOptions{
		Log:         p.log,
		Address:     addr,
		Credentials: p.options.Credentials,
		MachineHash: steam.SentryHash(p.login.MachineHash),
		LoginKey:    p.login.LoginKey,
	})

	if err != nil {
		return xerrors.Errorf("connection error: %w", err)
	}

	return nil
}

func (p *Manager) disconnect() {
	if p.conn == nil {
		p.log.Warn("called disconnect() without connection")
		return
	}

	p.closeSession()
	p.conn.Close(context.Background())
	p.conn = nil
}
