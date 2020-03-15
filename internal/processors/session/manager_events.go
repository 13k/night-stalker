package session

import (
	"github.com/faceit/go-steam"
	"github.com/paralin/go-dota2/events"
	"golang.org/x/xerrors"

	"github.com/13k/night-stalker/models"
)

func (p *Manager) handleEvent(ev interface{}) (err error) {
	switch e := ev.(type) {
	case *steam.ClientCMListEvent:
		err = p.onSteamServerList(e)
	case *steam.LoggedOnEvent:
		err = p.onSteamLogOn(e)
	case *steam.AccountInfoEvent:
	case *steam.PersonaStateEvent:
	case *steam.LoginKeyEvent:
		err = p.onSteamLoginKey(e)
	case *steam.MachineAuthUpdateEvent:
		err = p.onSteamMachineAuth(e)
	case *SteamWebSessionIDEvent:
		err = p.onSteamWebSession(e)
	case *SteamWebLoggedOnEvent:
		err = p.onSteamWebLogOn(e)
	case *events.ClientSuspended:
		err = p.onDotaClientSuspended(e)
	case *events.ClientWelcomed:
		err = p.onDotaWelcome(e)
	case steam.FatalErrorEvent:
		p.log.WithError(e).Error("steam fatal error")
	default:
		err = p.busPubEvent(ev)
	}

	if err != nil {
		return xerrors.Errorf("session error: %w", err)
	}

	return
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

func (p *Manager) onSteamWebSession(ev *SteamWebSessionIDEvent) error {
	p.log.Debug("received web session id")

	if err := p.saveWebSessionID(ev.SessionID); err != nil {
		return xerrors.Errorf("error saving steam machine auth token: %w", err)
	}

	return nil
}

func (p *Manager) onSteamWebLogOn(ev *SteamWebLoggedOnEvent) error {
	p.log.Debug("web logged on")

	if err := p.saveWebAuth(ev.AuthToken, ev.AuthSecret); err != nil {
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

func (p *Manager) onDotaWelcome(e *events.ClientWelcomed) error {
	if err := p.saveDotaWelcome(e.Welcome); err != nil {
		return xerrors.Errorf("error saving dota welcome: %w", err)
	}

	return nil
}

func (p *Manager) onDotaClientSuspended(ev *events.ClientSuspended) error {
	p.closeSession()

	until := models.NullUnixTimestamp(int64(ev.GetTimeEnd()))

	if err := p.saveDotaClientSuspended(until); err != nil {
		return xerrors.Errorf("error saving dota client suspended info: %w", err)
	}

	return xerrors.Errorf("dota error: %w", &ErrDotaClientSuspended{Until: until})
}
