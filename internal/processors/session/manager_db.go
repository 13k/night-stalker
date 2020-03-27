package session

import (
	"database/sql"
	"time"

	"github.com/faceit/go-steam"
	"github.com/faceit/go-steam/netutil"
	"github.com/jinzhu/gorm"
	"github.com/paralin/go-dota2/protocol"
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

func (p *Manager) loadLogin() error {
	dbres := p.db.
		Where(&models.SteamLogin{Username: p.options.Credentials.Username}).
		FirstOrInit(p.login)

	if dbres.Error != nil {
		return xerrors.Errorf("error loading login: %w", dbres.Error)
	}

	return nil
}

func (p *Manager) isSuspended() bool {
	return p.login.SuspendedUntil.Valid && time.Now().Before(p.login.SuspendedUntil.Time)
}

func (p *Manager) updateLogin(update *models.SteamLogin) error {
	p.log.Trace("update login")

	dbres := p.db.
		Where(&models.SteamLogin{Username: p.login.Username}).
		Assign(update).
		FirstOrCreate(p.login)

	if dbres.Error != nil {
		return xerrors.Errorf("error updating login: %w", dbres.Error)
	}

	return nil
}

func (p *Manager) saveLoginKey(uniqueID uint32, loginKey string) error {
	update := &models.SteamLogin{
		UniqueID: uniqueID,
		LoginKey: loginKey,
	}

	if err := p.updateLogin(update); err != nil {
		return xerrors.Errorf("error saving steam login key: %w", err)
	}

	return nil
}

func (p *Manager) saveMachineAuthToken(hash []byte) error {
	update := &models.SteamLogin{MachineHash: hash}

	if err := p.updateLogin(update); err != nil {
		return xerrors.Errorf("error saving steam machine auth token: %w", err)
	}

	return nil
}

func (p *Manager) saveWebSessionID(sessionID string) error {
	update := &models.SteamLogin{WebSessionID: sessionID}

	if err := p.updateLogin(update); err != nil {
		return xerrors.Errorf("error saving steam web session: %w", err)
	}

	return nil
}

func (p *Manager) saveWebAuth(authToken, authSecret string) error {
	update := &models.SteamLogin{
		WebAuthToken:       authToken,
		WebAuthTokenSecure: authSecret,
	}

	if err := p.updateLogin(update); err != nil {
		return xerrors.Errorf("error saving steam web auth: %w", err)
	}

	return nil
}

func (p *Manager) saveAccountDetails(ev *steam.LoggedOnEvent) error {
	update := &models.SteamLogin{
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

func (p *Manager) saveDotaWelcome(welcome *protocol.CMsgClientWelcome) error {
	update := &models.SteamLogin{
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
	update := &models.SteamLogin{SuspendedUntil: t}

	if err := p.updateLogin(update); err != nil {
		return xerrors.Errorf("error saving dota client suspended info: %w", err)
	}

	return nil
}

func (p *Manager) randomServer() (*models.SteamServer, error) {
	server := &models.SteamServer{}
	result := p.db.Order("random()").Take(&server)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, xerrors.Errorf("error loading random server: %w", err)
	}

	return server, nil
}

func (p *Manager) randomServerAddr() (*netutil.PortAddr, error) {
	server, err := p.randomServer()

	if err != nil {
		return nil, xerrors.Errorf("error loading random steam server: %w", err)
	}

	var addr *netutil.PortAddr

	if server != nil {
		addr = netutil.ParsePortAddr(server.Address)

		if addr == nil {
			return nil, xerrors.Errorf("error parsing server address: %w", &ErrInvalidServerAddress{
				Address: server.Address,
			})
		}
	}

	return addr, nil
}

func (p *Manager) saveServers(addresses []*netutil.PortAddr) error {
	var err error

	tx := p.db.Begin()

	for _, addr := range addresses {
		server := &models.SteamServer{}
		err = tx.
			Where(&models.SteamServer{Address: addr.String()}).
			FirstOrCreate(server).
			Error

		if err != nil {
			break
		}
	}

	if err != nil {
		tx.Rollback()
		return xerrors.Errorf("error saving servers: %w", err)
	}

	return tx.Commit().Error
}
