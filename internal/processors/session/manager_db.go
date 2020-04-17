package session

import (
	"database/sql"
	"time"

	"github.com/faceit/go-steam"
	"github.com/faceit/go-steam/netutil"
	d2pb "github.com/paralin/go-dota2/protocol"
	"golang.org/x/xerrors"

	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

func (p *Manager) loadLogin() error {
	p.login.Username = p.options.Credentials.Username

	_, err := p.db.M().FindBy(p.ctx, p.login, "username", p.options.Credentials.Username)

	if err != nil && !xerrors.Is(err, sql.ErrNoRows) {
		return xerrors.Errorf("error loading login: %w", err)
	}

	return nil
}

func (p *Manager) isSuspended() bool {
	return p.login.SuspendedUntil.Valid && time.Now().Before(p.login.SuspendedUntil.Time)
}

func (p *Manager) updateLogin(update *nsm.SteamLogin) error {
	p.log.Trace("update login")

	if p.login.AssignPartial(update) {
		q := p.db.
			Q().
			Select().
			Eq(nsm.SteamLoginTable.Col("username"), p.login.Username)

		if _, err := p.db.M().Upsert(p.ctx, p.login, q); err != nil {
			return xerrors.Errorf("error updating login: %w", err)
		}
	}

	return nil
}

func (p *Manager) saveLoginKey(uniqueID uint32, loginKey string) error {
	update := &nsm.SteamLogin{
		UniqueID: uniqueID,
		LoginKey: loginKey,
	}

	if err := p.updateLogin(update); err != nil {
		return xerrors.Errorf("error saving steam login key: %w", err)
	}

	return nil
}

func (p *Manager) saveMachineAuthToken(hash []byte) error {
	update := &nsm.SteamLogin{MachineHash: hash}

	if err := p.updateLogin(update); err != nil {
		return xerrors.Errorf("error saving steam machine auth token: %w", err)
	}

	return nil
}

func (p *Manager) saveWebSessionID(sessionID string) error {
	update := &nsm.SteamLogin{WebSessionID: sessionID}

	if err := p.updateLogin(update); err != nil {
		return xerrors.Errorf("error saving steam web session: %w", err)
	}

	return nil
}

func (p *Manager) saveWebAuth(authToken, authSecret string) error {
	update := &nsm.SteamLogin{
		WebAuthToken:       authToken,
		WebAuthTokenSecure: authSecret,
	}

	if err := p.updateLogin(update); err != nil {
		return xerrors.Errorf("error saving steam web auth: %w", err)
	}

	return nil
}

func (p *Manager) saveAccountDetails(ev *steam.LoggedOnEvent) error {
	update := &nsm.SteamLogin{
		SteamID:                   ev.ClientSteamId,
		AccountFlags:              nspb.SteamAccountFlags(ev.AccountFlags),
		WebAuthNonce:              ev.Body.GetWebapiAuthenticateUserNonce(),
		CellID:                    ev.Body.GetCellId(),
		CellIDPingThreshold:       ev.Body.GetCellIdPingThreshold(),
		EmailDomain:               ev.Body.GetEmailDomain(),
		VanityURL:                 ev.Body.GetVanityUrl(),
		OutOfGameHeartbeatSeconds: ev.Body.GetOutOfGameHeartbeatSeconds(),
		InGameHeartbeatSeconds:    ev.Body.GetInGameHeartbeatSeconds(),
		PublicIP:                  ev.Body.GetPublicIp(),
		ServerTime:                ev.Body.GetRtime32ServerTime(),
		SteamTicket:               ev.Body.GetSteam2Ticket(),
		UsePics:                   ev.Body.GetUsePics(),
		CountryCode:               ev.Body.GetIpCountryCode(),
		ParentalSettings:          ev.Body.GetParentalSettings(),
		ParentalSettingSignature:  ev.Body.GetParentalSettingSignature(),
		LoginFailuresToMigrate:    ev.Body.GetCountLoginfailuresToMigrate(),
		DisconnectsToMigrate:      ev.Body.GetCountDisconnectsToMigrate(),
		OgsDataReportTimeWindow:   ev.Body.GetOgsDataReportTimeWindow(),
		ClientInstanceID:          ev.Body.GetClientInstanceId(),
		ForceClientUpdateCheck:    ev.Body.GetForceClientUpdateCheck(),
	}

	if err := p.updateLogin(update); err != nil {
		return xerrors.Errorf("error saving account details: %w", err)
	}

	return nil
}

func (p *Manager) saveDotaWelcome(welcome *d2pb.CMsgClientWelcome) error {
	update := &nsm.SteamLogin{
		GameVersion:       welcome.GetVersion(),
		LocationCountry:   welcome.GetLocation().GetCountry(),
		LocationLatitude:  welcome.GetLocation().GetLatitude(),
		LocationLongitude: welcome.GetLocation().GetLongitude(),
	}

	if err := p.updateLogin(update); err != nil {
		return xerrors.Errorf("error saving dota welcome: %w", err)
	}

	return nil
}

func (p *Manager) saveDotaClientSuspended(until time.Time) error {
	t := sql.NullTime{Time: until, Valid: !until.IsZero()}
	update := &nsm.SteamLogin{SuspendedUntil: t}

	if err := p.updateLogin(update); err != nil {
		return xerrors.Errorf("error saving dota client suspended info: %w", err)
	}

	return nil
}

func (p *Manager) randomServer() (*nsm.SteamServer, error) {
	server := &nsm.SteamServer{}
	exists, err := p.db.M().Random(p.ctx, server)

	if err != nil {
		return nil, xerrors.Errorf("error loading random server: %w", err)
	}

	if !exists {
		return nil, nil
	}

	return server, nil
}

func (p *Manager) randomServerAddr() (*netutil.PortAddr, error) {
	server, err := p.randomServer()

	if err != nil {
		return nil, xerrors.Errorf("error loading random steam server: %w", err)
	}

	if server == nil {
		return nil, nil
	}

	addr := netutil.ParsePortAddr(server.Address)

	if addr == nil {
		return nil, xerrors.Errorf("error parsing server address: %w", &ErrInvalidServerAddress{
			Address: server.Address,
		})
	}

	return addr, nil
}

func (p *Manager) saveServers(addresses []*netutil.PortAddr) error {
	dbs := nsdbda.NewSaver(p.db)

	if _, err := dbs.UpsertSteamServersAddresses(p.ctx, addresses); err != nil {
		return xerrors.Errorf("error saving servers: %w", err)
	}

	return nil
}
