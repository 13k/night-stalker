package rtstats

import (
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2/protocol"

	nsjson "github.com/13k/night-stalker/internal/json"
)

type apiResult struct {
	*protocol.CMsgDOTARealtimeGameStatsTerse

	Teams []*apiResultTeamDetails `json:"teams"`
}

func (r *apiResult) ToProto() *protocol.CMsgDOTARealtimeGameStatsTerse {
	pb := r.CMsgDOTARealtimeGameStatsTerse

	for _, team := range r.Teams {
		pb.Teams = append(pb.Teams, team.ToProto())
	}

	return pb
}

type apiResultTeamDetails struct {
	*protocol.CMsgDOTARealtimeGameStatsTerse_TeamDetails

	TeamLogo nsjson.StringUint `json:"team_logo"`
}

func (t *apiResultTeamDetails) ToProto() *protocol.CMsgDOTARealtimeGameStatsTerse_TeamDetails {
	pb := t.CMsgDOTARealtimeGameStatsTerse_TeamDetails
	pb.TeamLogo = proto.Uint64(t.TeamLogo.Uint64())
	return pb
}
