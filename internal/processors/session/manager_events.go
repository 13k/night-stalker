package session

import (
	"time"

	"github.com/faceit/go-steam"
	d2events "github.com/paralin/go-dota2/events"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
)

func (p *Manager) handleSteamSessionChange(sessmsg *nsbus.SteamSessionChangeMessage) {
	if sessmsg.StateTransition.UnreadyToReady {
		p.log.Info("steam connected")
	} else if sessmsg.StateTransition.ReadyToUnready {
		p.log.Warn("steam disconnected")
	}
}

func (p *Manager) handleDotaSessionChange(sessmsg *nsbus.DotaSessionChangeMessage) {
	if sessmsg.StateTransition.UnreadyToReady {
		p.log.Info("dota connected")
	} else if sessmsg.StateTransition.ReadyToUnready {
		p.log.Warn("dota disconnected")
	}
}

func (p *Manager) handleSteamEvent(ev interface{}) error {
	var err error

	switch ev := ev.(type) {
	case *steam.ClientCMListEvent:
		err = p.onSteamServerList(ev)
	case *steam.LoggedOnEvent:
		err = p.onSteamLogOn(ev)
	case *steam.AccountInfoEvent:
	case *steam.PersonaStateEvent:
	case *steam.LoginKeyEvent:
		err = p.onSteamLoginKey(ev)
	case *steam.MachineAuthUpdateEvent:
		err = p.onSteamMachineAuth(ev)
	case *steam.WebSessionIdEvent:
		err = p.onSteamWebSession(ev)
	case *steam.WebLoggedOnEvent:
		err = p.onSteamWebLogOn(ev)
	case *d2events.ClientSuspended:
		err = p.onDotaClientSuspended(ev)
	case *d2events.ClientWelcomed:
		err = p.onDotaWelcome(ev)
	}

	return err
}

func (p *Manager) onSteamServerList(ev *steam.ClientCMListEvent) error {
	p.log.WithField("count", len(ev.Addresses)).Debug("received server list")

	if err := p.saveServers(ev.Addresses); err != nil {
		return xerrors.Errorf("error saving steam servers: %w", err)
	}

	return nil
}

func (p *Manager) onSteamLoginKey(ev *steam.LoginKeyEvent) error {
	p.log.Debug("received login key")

	if err := p.saveLoginKey(ev.UniqueId, ev.LoginKey); err != nil {
		return xerrors.Errorf("error saving steam login key: %w", err)
	}

	return nil
}

func (p *Manager) onSteamMachineAuth(ev *steam.MachineAuthUpdateEvent) error {
	p.log.Debug("received machine hash")

	if err := p.saveMachineAuthToken(ev.Hash); err != nil {
		return xerrors.Errorf("error saving steam machine auth token: %w", err)
	}

	return nil
}

func (p *Manager) onSteamWebSession(_ *steam.WebSessionIdEvent) error {
	p.log.Debug("received web session id")

	if err := p.saveWebSessionID(p.steam.Web.SessionId); err != nil {
		return xerrors.Errorf("error saving steam machine auth token: %w", err)
	}

	return nil
}

func (p *Manager) onSteamWebLogOn(_ *steam.WebLoggedOnEvent) error {
	p.log.Debug("web logged on")

	if err := p.saveWebAuth(p.steam.Web.SteamLogin, p.steam.Web.SteamLoginSecure); err != nil {
		return xerrors.Errorf("error saving steam web auth: %w", err)
	}

	return nil
}

func (p *Manager) onSteamLogOn(ev *steam.LoggedOnEvent) error {
	p.log.Debug("logged on")

	if err := p.saveAccountDetails(ev); err != nil {
		return xerrors.Errorf("error saving account details: %w", err)
	}

	return nil
}

func (p *Manager) onDotaWelcome(e *d2events.ClientWelcomed) error {
	if err := p.saveDotaWelcome(e.Welcome); err != nil {
		return xerrors.Errorf("error saving dota welcome: %w", err)
	}

	return nil
}

func (p *Manager) onDotaClientSuspended(ev *d2events.ClientSuspended) error {
	until := time.Unix(int64(ev.GetTimeEnd()), 0)

	if err := p.saveDotaClientSuspended(until); err != nil {
		return xerrors.Errorf("error saving dota client suspended info: %w", err)
	}

	return nil
}
