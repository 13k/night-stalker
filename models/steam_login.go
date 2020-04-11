package models

import (
	"database/sql"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
)

var SteamLoginTable = NewTable("steam_logins")

type SteamLogin struct {
	ID `db:"id" goqu:"defaultifempty"`

	Username                  string                 `db:"username"`
	SteamID                   nspb.SteamID           `db:"steam_id"`
	AccountFlags              nspb.SteamAccountFlags `db:"account_flags"`
	MachineHash               nssql.Base64Bytes      `db:"machine_hash"`
	UniqueID                  uint32                 `db:"unique_id"`
	LoginKey                  string                 `db:"login_key"`
	WebAuthNonce              string                 `db:"web_auth_nonce"`
	WebSessionID              string                 `db:"web_session_id"`
	WebAuthToken              string                 `db:"web_auth_token"`
	WebAuthTokenSecure        string                 `db:"web_auth_token_secure"`
	SuspendedUntil            sql.NullTime           `db:"suspended_until"`
	GameVersion               uint32                 `db:"game_version"`
	LocationCountry           string                 `db:"location_country"`
	LocationLatitude          float32                `db:"location_latitude"`
	LocationLongitude         float32                `db:"location_longitude"`
	CellID                    uint32                 `db:"cell_id"`
	CellIDPingThreshold       uint32                 `db:"cell_id_ping_threshold"`
	EmailDomain               string                 `db:"email_domain"`
	VanityURL                 string                 `db:"vanity_url"`
	OutOfGameHeartbeatSeconds int32                  `db:"out_of_game_heartbeat_seconds"`
	InGameHeartbeatSeconds    int32                  `db:"in_game_heartbeat_seconds"`
	PublicIP                  uint32                 `db:"public_ip"`
	ServerTime                uint32                 `db:"server_time"`
	SteamTicket               []byte                 `db:"steam_ticket"`
	UsePics                   bool                   `db:"use_pics"`
	CountryCode               string                 `db:"country_code"`
	ParentalSettings          []byte                 `db:"parental_settings"`
	ParentalSettingSignature  []byte                 `db:"parental_setting_signature"`
	LoginFailuresToMigrate    int32                  `db:"login_failures_to_migrate"`
	DisconnectsToMigrate      int32                  `db:"disconnects_to_migrate"`
	OgsDataReportTimeWindow   int32                  `db:"ogs_data_report_time_window"`
	ClientInstanceID          uint64                 `db:"client_instance_id"`
	ForceClientUpdateCheck    bool                   `db:"force_client_update_check"`

	Timestamps
	SoftDelete
}
