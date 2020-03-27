package rtstats

import (
	d2pb "github.com/paralin/go-dota2/protocol"
	"google.golang.org/protobuf/proto"

	nsjson "github.com/13k/night-stalker/internal/json"
)

type apiResult struct {
	*d2pb.CMsgDOTARealtimeGameStatsTerse

	Teams []*apiResultTeamDetails `json:"teams"`
}

func (r *apiResult) ToProto() *d2pb.CMsgDOTARealtimeGameStatsTerse {
	pb := r.CMsgDOTARealtimeGameStatsTerse

	for _, team := range r.Teams {
		pb.Teams = append(pb.Teams, team.ToProto())
	}

	return pb
}

type apiResultTeamDetails struct {
	*d2pb.CMsgDOTARealtimeGameStatsTerse_TeamDetails

	TeamLogo nsjson.StringUint `json:"team_logo"`
}

func (t *apiResultTeamDetails) ToProto() *d2pb.CMsgDOTARealtimeGameStatsTerse_TeamDetails {
	pb := t.CMsgDOTARealtimeGameStatsTerse_TeamDetails
	pb.TeamLogo = proto.Uint64(t.TeamLogo.Uint64())
	return pb
}
