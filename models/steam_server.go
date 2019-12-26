package models

var SteamServerModel = (*SteamServer)(nil)

type SteamServerID uint64

// SteamServer ...
type SteamServer struct {
	ID      SteamServerID `gorm:"column:id;primary_key"`
	Address string        `gorm:"column:address;size:255;unique_index;not null"`
	Timestamps
}
