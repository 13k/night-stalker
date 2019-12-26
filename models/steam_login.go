package models

import (
	"time"

	nssql "github.com/13k/night-stalker/internal/sql"
	"github.com/faceit/go-steam/steamid"
)

var SteamLoginModel = (*SteamLogin)(nil)

type SteamLoginID uint64

// SteamLogin ...
type SteamLogin struct {
	ID                        SteamLoginID      `gorm:"column:id;primary_key"`
	Username                  string            `gorm:"column:username;size:255;unique_index;not null"`
	SteamID                   steamid.SteamId   `gorm:"column:steam_id;unique_index"`
	AccountFlags              uint32            `gorm:"column:account_flags"`
	MachineHash               nssql.Base64Bytes `gorm:"column:machine_hash;type:varchar(255)"`
	UniqueID                  uint32            `gorm:"column:unique_id"`
	LoginKey                  string            `gorm:"column:login_key;size:255"`
	WebAuthNonce              string            `gorm:"column:web_auth_nonce;size:255"`
	WebSessionID              string            `gorm:"column:web_session_id;size:255"`
	WebAuthToken              string            `gorm:"column:web_auth_token;size:255"`
	WebAuthTokenSecure        string            `gorm:"column:web_auth_token_secure;size:255"`
	SuspendedUntil            *time.Time        `gorm:"column:suspended_until"`
	GameVersion               uint32            `gorm:"column:game_version"`
	LocationCountry           string            `gorm:"column:location_country;size:255"`
	LocationLatitude          float32           `gorm:"column:location_latitude"`
	LocationLongitude         float32           `gorm:"column:location_longitude"`
	CellID                    uint32            `gorm:"column:cell_id"`
	CellIDPingThreshold       uint32            `gorm:"column:cell_id_ping_threshold"`
	EmailDomain               string            `gorm:"column:email_domain;size:255"`
	VanityURL                 string            `gorm:"column:vanity_url"`
	OutOfGameHeartbeatSeconds int32             `gorm:"column:out_of_game_heartbeat_seconds"`
	InGameHeartbeatSeconds    int32             `gorm:"column:in_game_heartbeat_seconds"`
	PublicIP                  uint32            `gorm:"column:public_ip"`
	ServerTime                uint32            `gorm:"column:server_time"`
	SteamTicket               []byte            `gorm:"column:steam_ticket"`
	UsePics                   bool              `gorm:"column:use_pics"`
	CountryCode               string            `gorm:"column:country_code;size:255"`
	ParentalSettings          []byte            `gorm:"column:parental_settings"`
	ParentalSettingSignature  []byte            `gorm:"column:parental_setting_signature"`
	LoginFailuresToMigrate    int32             `gorm:"column:login_failures_to_migrate"`
	DisconnectsToMigrate      int32             `gorm:"column:disconnects_to_migrate"`
	OgsDataReportTimeWindow   int32             `gorm:"column:ogs_data_report_time_window"`
	ClientInstanceID          uint64            `gorm:"column:client_instance_id"`
	ForceClientUpdateCheck    bool              `gorm:"column:force_client_update_check"`
	Timestamps
}
