package models

var SteamServerModel Model = (*SteamServer)(nil)

type SteamServerID uint64

// SteamServer ...
type SteamServer struct {
	ID      SteamServerID `gorm:"column:id;primary_key"`
	Address string        `gorm:"column:address;size:255;unique_index;not null"`
	Timestamps
}

func (*SteamServer) TableName() string {
	return "steam_servers"
}
