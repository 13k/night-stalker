package models

var SteamServerTable = NewTable("steam_servers")

type SteamServer struct {
	ID `db:"id" goqu:"defaultifempty"`

	Address string `db:"address"`

	Timestamps
	SoftDelete
}
