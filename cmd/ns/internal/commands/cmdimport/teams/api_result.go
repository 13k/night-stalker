package teams

import (
	nsjson "github.com/13k/night-stalker/internal/json"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

type apiResult []*apiResultEntry

type apiResultEntry struct {
	TeamID        nspb.TeamID      `json:"team_id,omitempty"`
	Name          string           `json:"name,omitempty"`
	Tag           string           `json:"tag,omitempty"`
	Rating        float32          `json:"rating,omitempty"`
	Wins          uint32           `json:"wins,omitempty"`
	Losses        uint32           `json:"losses,omitempty"`
	LogoURL       string           `json:"logo_url,omitempty"`
	LastMatchTime *nsjson.UnixTime `json:"last_match_time,omitempty"`
}
