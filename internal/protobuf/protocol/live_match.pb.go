// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.6.1
// source: protocol/live_match.proto

package protocol

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LiveMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId                    uint64               `protobuf:"varint,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	ServerId                   uint64               `protobuf:"varint,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	LobbyId                    uint64               `protobuf:"varint,3,opt,name=lobby_id,json=lobbyId,proto3" json:"lobby_id,omitempty"`
	LobbyType                  LobbyType            `protobuf:"varint,4,opt,name=lobby_type,json=lobbyType,proto3,enum=ns.protocol.LobbyType" json:"lobby_type,omitempty"`
	LeagueId                   uint64               `protobuf:"varint,5,opt,name=league_id,json=leagueId,proto3" json:"league_id,omitempty"`
	SeriesId                   uint64               `protobuf:"varint,6,opt,name=series_id,json=seriesId,proto3" json:"series_id,omitempty"`
	GameMode                   GameMode             `protobuf:"varint,7,opt,name=game_mode,json=gameMode,proto3,enum=ns.protocol.GameMode" json:"game_mode,omitempty"`
	GameState                  GameState            `protobuf:"varint,8,opt,name=game_state,json=gameState,proto3,enum=ns.protocol.GameState" json:"game_state,omitempty"`
	GameTimestamp              uint32               `protobuf:"varint,9,opt,name=game_timestamp,json=gameTimestamp,proto3" json:"game_timestamp,omitempty"`
	GameTime                   int32                `protobuf:"varint,10,opt,name=game_time,json=gameTime,proto3" json:"game_time,omitempty"`
	AverageMmr                 uint32               `protobuf:"varint,11,opt,name=average_mmr,json=averageMmr,proto3" json:"average_mmr,omitempty"`
	Delay                      uint32               `protobuf:"varint,12,opt,name=delay,proto3" json:"delay,omitempty"`
	Spectators                 uint32               `protobuf:"varint,13,opt,name=spectators,proto3" json:"spectators,omitempty"`
	SortScore                  float64              `protobuf:"fixed64,14,opt,name=sort_score,json=sortScore,proto3" json:"sort_score,omitempty"`
	RadiantLead                int32                `protobuf:"varint,15,opt,name=radiant_lead,json=radiantLead,proto3" json:"radiant_lead,omitempty"`
	RadiantScore               uint32               `protobuf:"varint,16,opt,name=radiant_score,json=radiantScore,proto3" json:"radiant_score,omitempty"`
	RadiantTeamId              uint64               `protobuf:"varint,17,opt,name=radiant_team_id,json=radiantTeamId,proto3" json:"radiant_team_id,omitempty"`
	RadiantTeamName            string               `protobuf:"bytes,18,opt,name=radiant_team_name,json=radiantTeamName,proto3" json:"radiant_team_name,omitempty"`
	RadiantTeamTag             string               `protobuf:"bytes,19,opt,name=radiant_team_tag,json=radiantTeamTag,proto3" json:"radiant_team_tag,omitempty"`
	RadiantTeamLogo            uint64               `protobuf:"varint,20,opt,name=radiant_team_logo,json=radiantTeamLogo,proto3" json:"radiant_team_logo,omitempty"`
	RadiantTeamLogoUrl         string               `protobuf:"bytes,21,opt,name=radiant_team_logo_url,json=radiantTeamLogoUrl,proto3" json:"radiant_team_logo_url,omitempty"`
	RadiantNetWorth            uint32               `protobuf:"varint,22,opt,name=radiant_net_worth,json=radiantNetWorth,proto3" json:"radiant_net_worth,omitempty"`
	DireScore                  uint32               `protobuf:"varint,23,opt,name=dire_score,json=direScore,proto3" json:"dire_score,omitempty"`
	DireTeamId                 uint64               `protobuf:"varint,24,opt,name=dire_team_id,json=direTeamId,proto3" json:"dire_team_id,omitempty"`
	DireTeamName               string               `protobuf:"bytes,25,opt,name=dire_team_name,json=direTeamName,proto3" json:"dire_team_name,omitempty"`
	DireTeamTag                string               `protobuf:"bytes,26,opt,name=dire_team_tag,json=direTeamTag,proto3" json:"dire_team_tag,omitempty"`
	DireTeamLogo               uint64               `protobuf:"varint,27,opt,name=dire_team_logo,json=direTeamLogo,proto3" json:"dire_team_logo,omitempty"`
	DireTeamLogoUrl            string               `protobuf:"bytes,28,opt,name=dire_team_logo_url,json=direTeamLogoUrl,proto3" json:"dire_team_logo_url,omitempty"`
	DireNetWorth               uint32               `protobuf:"varint,29,opt,name=dire_net_worth,json=direNetWorth,proto3" json:"dire_net_worth,omitempty"`
	BuildingState              uint32               `protobuf:"varint,30,opt,name=building_state,json=buildingState,proto3" json:"building_state,omitempty"`
	WeekendTourneyTournamentId uint32               `protobuf:"varint,31,opt,name=weekend_tourney_tournament_id,json=weekendTourneyTournamentId,proto3" json:"weekend_tourney_tournament_id,omitempty"`
	WeekendTourneyDivision     uint32               `protobuf:"varint,32,opt,name=weekend_tourney_division,json=weekendTourneyDivision,proto3" json:"weekend_tourney_division,omitempty"`
	WeekendTourneySkillLevel   uint32               `protobuf:"varint,33,opt,name=weekend_tourney_skill_level,json=weekendTourneySkillLevel,proto3" json:"weekend_tourney_skill_level,omitempty"`
	WeekendTourneyBracketRound uint32               `protobuf:"varint,34,opt,name=weekend_tourney_bracket_round,json=weekendTourneyBracketRound,proto3" json:"weekend_tourney_bracket_round,omitempty"`
	ActivateTime               *timestamp.Timestamp `protobuf:"bytes,35,opt,name=activate_time,json=activateTime,proto3" json:"activate_time,omitempty"`
	DeactivateTime             *timestamp.Timestamp `protobuf:"bytes,36,opt,name=deactivate_time,json=deactivateTime,proto3" json:"deactivate_time,omitempty"`
	LastUpdateTime             *timestamp.Timestamp `protobuf:"bytes,37,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
	Players                    []*LiveMatch_Player  `protobuf:"bytes,100,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *LiveMatch) Reset() {
	*x = LiveMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_live_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveMatch) ProtoMessage() {}

func (x *LiveMatch) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_live_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveMatch.ProtoReflect.Descriptor instead.
func (*LiveMatch) Descriptor() ([]byte, []int) {
	return file_protocol_live_match_proto_rawDescGZIP(), []int{0}
}

func (x *LiveMatch) GetMatchId() uint64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *LiveMatch) GetServerId() uint64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *LiveMatch) GetLobbyId() uint64 {
	if x != nil {
		return x.LobbyId
	}
	return 0
}

func (x *LiveMatch) GetLobbyType() LobbyType {
	if x != nil {
		return x.LobbyType
	}
	return LobbyType_LOBBY_TYPE_CASUAL_MATCH
}

func (x *LiveMatch) GetLeagueId() uint64 {
	if x != nil {
		return x.LeagueId
	}
	return 0
}

func (x *LiveMatch) GetSeriesId() uint64 {
	if x != nil {
		return x.SeriesId
	}
	return 0
}

func (x *LiveMatch) GetGameMode() GameMode {
	if x != nil {
		return x.GameMode
	}
	return GameMode_GAME_MODE_NONE
}

func (x *LiveMatch) GetGameState() GameState {
	if x != nil {
		return x.GameState
	}
	return GameState_GAME_STATE_INIT
}

func (x *LiveMatch) GetGameTimestamp() uint32 {
	if x != nil {
		return x.GameTimestamp
	}
	return 0
}

func (x *LiveMatch) GetGameTime() int32 {
	if x != nil {
		return x.GameTime
	}
	return 0
}

func (x *LiveMatch) GetAverageMmr() uint32 {
	if x != nil {
		return x.AverageMmr
	}
	return 0
}

func (x *LiveMatch) GetDelay() uint32 {
	if x != nil {
		return x.Delay
	}
	return 0
}

func (x *LiveMatch) GetSpectators() uint32 {
	if x != nil {
		return x.Spectators
	}
	return 0
}

func (x *LiveMatch) GetSortScore() float64 {
	if x != nil {
		return x.SortScore
	}
	return 0
}

func (x *LiveMatch) GetRadiantLead() int32 {
	if x != nil {
		return x.RadiantLead
	}
	return 0
}

func (x *LiveMatch) GetRadiantScore() uint32 {
	if x != nil {
		return x.RadiantScore
	}
	return 0
}

func (x *LiveMatch) GetRadiantTeamId() uint64 {
	if x != nil {
		return x.RadiantTeamId
	}
	return 0
}

func (x *LiveMatch) GetRadiantTeamName() string {
	if x != nil {
		return x.RadiantTeamName
	}
	return ""
}

func (x *LiveMatch) GetRadiantTeamTag() string {
	if x != nil {
		return x.RadiantTeamTag
	}
	return ""
}

func (x *LiveMatch) GetRadiantTeamLogo() uint64 {
	if x != nil {
		return x.RadiantTeamLogo
	}
	return 0
}

func (x *LiveMatch) GetRadiantTeamLogoUrl() string {
	if x != nil {
		return x.RadiantTeamLogoUrl
	}
	return ""
}

func (x *LiveMatch) GetRadiantNetWorth() uint32 {
	if x != nil {
		return x.RadiantNetWorth
	}
	return 0
}

func (x *LiveMatch) GetDireScore() uint32 {
	if x != nil {
		return x.DireScore
	}
	return 0
}

func (x *LiveMatch) GetDireTeamId() uint64 {
	if x != nil {
		return x.DireTeamId
	}
	return 0
}

func (x *LiveMatch) GetDireTeamName() string {
	if x != nil {
		return x.DireTeamName
	}
	return ""
}

func (x *LiveMatch) GetDireTeamTag() string {
	if x != nil {
		return x.DireTeamTag
	}
	return ""
}

func (x *LiveMatch) GetDireTeamLogo() uint64 {
	if x != nil {
		return x.DireTeamLogo
	}
	return 0
}

func (x *LiveMatch) GetDireTeamLogoUrl() string {
	if x != nil {
		return x.DireTeamLogoUrl
	}
	return ""
}

func (x *LiveMatch) GetDireNetWorth() uint32 {
	if x != nil {
		return x.DireNetWorth
	}
	return 0
}

func (x *LiveMatch) GetBuildingState() uint32 {
	if x != nil {
		return x.BuildingState
	}
	return 0
}

func (x *LiveMatch) GetWeekendTourneyTournamentId() uint32 {
	if x != nil {
		return x.WeekendTourneyTournamentId
	}
	return 0
}

func (x *LiveMatch) GetWeekendTourneyDivision() uint32 {
	if x != nil {
		return x.WeekendTourneyDivision
	}
	return 0
}

func (x *LiveMatch) GetWeekendTourneySkillLevel() uint32 {
	if x != nil {
		return x.WeekendTourneySkillLevel
	}
	return 0
}

func (x *LiveMatch) GetWeekendTourneyBracketRound() uint32 {
	if x != nil {
		return x.WeekendTourneyBracketRound
	}
	return 0
}

func (x *LiveMatch) GetActivateTime() *timestamp.Timestamp {
	if x != nil {
		return x.ActivateTime
	}
	return nil
}

func (x *LiveMatch) GetDeactivateTime() *timestamp.Timestamp {
	if x != nil {
		return x.DeactivateTime
	}
	return nil
}

func (x *LiveMatch) GetLastUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdateTime
	}
	return nil
}

func (x *LiveMatch) GetPlayers() []*LiveMatch_Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type LiveMatches struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matches []*LiveMatch `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
}

func (x *LiveMatches) Reset() {
	*x = LiveMatches{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_live_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveMatches) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveMatches) ProtoMessage() {}

func (x *LiveMatches) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_live_match_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveMatches.ProtoReflect.Descriptor instead.
func (*LiveMatches) Descriptor() ([]byte, []int) {
	return file_protocol_live_match_proto_rawDescGZIP(), []int{1}
}

func (x *LiveMatches) GetMatches() []*LiveMatch {
	if x != nil {
		return x.Matches
	}
	return nil
}

type LiveMatchesChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op     CollectionOp `protobuf:"varint,1,opt,name=op,proto3,enum=ns.protocol.CollectionOp" json:"op,omitempty"`
	Change *LiveMatches `protobuf:"bytes,2,opt,name=change,proto3" json:"change,omitempty"`
}

func (x *LiveMatchesChange) Reset() {
	*x = LiveMatchesChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_live_match_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveMatchesChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveMatchesChange) ProtoMessage() {}

func (x *LiveMatchesChange) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_live_match_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveMatchesChange.ProtoReflect.Descriptor instead.
func (*LiveMatchesChange) Descriptor() ([]byte, []int) {
	return file_protocol_live_match_proto_rawDescGZIP(), []int{2}
}

func (x *LiveMatchesChange) GetOp() CollectionOp {
	if x != nil {
		return x.Op
	}
	return CollectionOp_COLLECTION_OP_UNSPECIFIED
}

func (x *LiveMatchesChange) GetChange() *LiveMatches {
	if x != nil {
		return x.Change
	}
	return nil
}

type LiveMatch_Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId       uint32   `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name            string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PersonaName     string   `protobuf:"bytes,3,opt,name=persona_name,json=personaName,proto3" json:"persona_name,omitempty"`
	AvatarUrl       string   `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	AvatarMediumUrl string   `protobuf:"bytes,5,opt,name=avatar_medium_url,json=avatarMediumUrl,proto3" json:"avatar_medium_url,omitempty"`
	AvatarFullUrl   string   `protobuf:"bytes,6,opt,name=avatar_full_url,json=avatarFullUrl,proto3" json:"avatar_full_url,omitempty"`
	IsPro           bool     `protobuf:"varint,7,opt,name=is_pro,json=isPro,proto3" json:"is_pro,omitempty"`
	HeroId          uint64   `protobuf:"varint,8,opt,name=hero_id,json=heroId,proto3" json:"hero_id,omitempty"`
	PlayerSlot      uint32   `protobuf:"varint,9,opt,name=player_slot,json=playerSlot,proto3" json:"player_slot,omitempty"`
	Team            GameTeam `protobuf:"varint,10,opt,name=team,proto3,enum=ns.protocol.GameTeam" json:"team,omitempty"`
	Level           uint32   `protobuf:"varint,11,opt,name=level,proto3" json:"level,omitempty"`
	Kills           uint32   `protobuf:"varint,12,opt,name=kills,proto3" json:"kills,omitempty"`
	Deaths          uint32   `protobuf:"varint,13,opt,name=deaths,proto3" json:"deaths,omitempty"`
	Assists         uint32   `protobuf:"varint,14,opt,name=assists,proto3" json:"assists,omitempty"`
	Denies          uint32   `protobuf:"varint,15,opt,name=denies,proto3" json:"denies,omitempty"`
	LastHits        uint32   `protobuf:"varint,16,opt,name=last_hits,json=lastHits,proto3" json:"last_hits,omitempty"`
	Gold            uint32   `protobuf:"varint,17,opt,name=gold,proto3" json:"gold,omitempty"`
	NetWorth        uint32   `protobuf:"varint,18,opt,name=net_worth,json=netWorth,proto3" json:"net_worth,omitempty"`
	Label           string   `protobuf:"bytes,19,opt,name=label,proto3" json:"label,omitempty"`
	Slug            string   `protobuf:"bytes,20,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *LiveMatch_Player) Reset() {
	*x = LiveMatch_Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_live_match_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveMatch_Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveMatch_Player) ProtoMessage() {}

func (x *LiveMatch_Player) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_live_match_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveMatch_Player.ProtoReflect.Descriptor instead.
func (*LiveMatch_Player) Descriptor() ([]byte, []int) {
	return file_protocol_live_match_proto_rawDescGZIP(), []int{0, 0}
}

func (x *LiveMatch_Player) GetAccountId() uint32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *LiveMatch_Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LiveMatch_Player) GetPersonaName() string {
	if x != nil {
		return x.PersonaName
	}
	return ""
}

func (x *LiveMatch_Player) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *LiveMatch_Player) GetAvatarMediumUrl() string {
	if x != nil {
		return x.AvatarMediumUrl
	}
	return ""
}

func (x *LiveMatch_Player) GetAvatarFullUrl() string {
	if x != nil {
		return x.AvatarFullUrl
	}
	return ""
}

func (x *LiveMatch_Player) GetIsPro() bool {
	if x != nil {
		return x.IsPro
	}
	return false
}

func (x *LiveMatch_Player) GetHeroId() uint64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *LiveMatch_Player) GetPlayerSlot() uint32 {
	if x != nil {
		return x.PlayerSlot
	}
	return 0
}

func (x *LiveMatch_Player) GetTeam() GameTeam {
	if x != nil {
		return x.Team
	}
	return GameTeam_GAME_TEAM_UNKNOWN
}

func (x *LiveMatch_Player) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *LiveMatch_Player) GetKills() uint32 {
	if x != nil {
		return x.Kills
	}
	return 0
}

func (x *LiveMatch_Player) GetDeaths() uint32 {
	if x != nil {
		return x.Deaths
	}
	return 0
}

func (x *LiveMatch_Player) GetAssists() uint32 {
	if x != nil {
		return x.Assists
	}
	return 0
}

func (x *LiveMatch_Player) GetDenies() uint32 {
	if x != nil {
		return x.Denies
	}
	return 0
}

func (x *LiveMatch_Player) GetLastHits() uint32 {
	if x != nil {
		return x.LastHits
	}
	return 0
}

func (x *LiveMatch_Player) GetGold() uint32 {
	if x != nil {
		return x.Gold
	}
	return 0
}

func (x *LiveMatch_Player) GetNetWorth() uint32 {
	if x != nil {
		return x.NetWorth
	}
	return 0
}

func (x *LiveMatch_Player) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *LiveMatch_Player) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

var File_protocol_live_match_proto protoreflect.FileDescriptor

var file_protocol_live_match_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x5f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x11, 0x0a, 0x09,
	0x4c, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0a,
	0x6c, 0x6f, 0x62, 0x62, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c,
	0x6f, 0x62, 0x62, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x12, 0x32, 0x0a,
	0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x35, 0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x67,
	0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x6d, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x6d, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x64, 0x65,
	0x6c, 0x61, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x61, 0x64, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x6c, 0x65,
	0x61, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x61, 0x64, 0x69, 0x61, 0x6e,
	0x74, 0x4c, 0x65, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x61, 0x64, 0x69, 0x61, 0x6e, 0x74,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x61,
	0x64, 0x69, 0x61, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x61,
	0x64, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0d, 0x72, 0x61, 0x64, 0x69, 0x61, 0x6e, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x61, 0x64, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x65,
	0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72,
	0x61, 0x64, 0x69, 0x61, 0x6e, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x72, 0x61, 0x64, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x74,
	0x61, 0x67, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x61, 0x64, 0x69, 0x61, 0x6e,
	0x74, 0x54, 0x65, 0x61, 0x6d, 0x54, 0x61, 0x67, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x61, 0x64, 0x69,
	0x61, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0f, 0x72, 0x61, 0x64, 0x69, 0x61, 0x6e, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x31, 0x0a, 0x15, 0x72, 0x61, 0x64, 0x69, 0x61, 0x6e, 0x74, 0x5f,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x61, 0x64, 0x69, 0x61, 0x6e, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x4c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x61, 0x64, 0x69, 0x61,
	0x6e, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x5f, 0x77, 0x6f, 0x72, 0x74, 0x68, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x72, 0x61, 0x64, 0x69, 0x61, 0x6e, 0x74, 0x4e, 0x65, 0x74, 0x57, 0x6f,
	0x72, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x72, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x64, 0x69, 0x72, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69,
	0x72, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x69,
	0x72, 0x65, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x1a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x54, 0x61, 0x67, 0x12, 0x24,
	0x0a, 0x0e, 0x64, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6c, 0x6f, 0x67, 0x6f,
	0x18, 0x1b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x64, 0x69, 0x72, 0x65, 0x54, 0x65, 0x61, 0x6d,
	0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x2b, 0x0a, 0x12, 0x64, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x64, 0x69, 0x72, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x6f, 0x55, 0x72,
	0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x69, 0x72, 0x65, 0x5f, 0x6e, 0x65, 0x74, 0x5f, 0x77, 0x6f,
	0x72, 0x74, 0x68, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x69, 0x72, 0x65, 0x4e,
	0x65, 0x74, 0x57, 0x6f, 0x72, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x41,
	0x0a, 0x1d, 0x77, 0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x77, 0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x65, 0x79, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x38, 0x0a, 0x18, 0x77, 0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x75,
	0x72, 0x6e, 0x65, 0x79, 0x5f, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x16, 0x77, 0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x65, 0x79, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x1b, 0x77,
	0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x5f, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x18, 0x77, 0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x41, 0x0a, 0x1d, 0x77, 0x65,
	0x65, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x5f, 0x62, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x22, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x1a, 0x77, 0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x42, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x3f, 0x0a,
	0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x23,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x43,
	0x0a, 0x0f, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x24, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0e, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x25, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x18, 0x64, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x1a, 0xbb, 0x04, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55,
	0x72, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6d, 0x65, 0x64,
	0x69, 0x75, 0x6d, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x55, 0x72, 0x6c, 0x12, 0x26,
	0x0a, 0x0f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46,
	0x75, 0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x12, 0x17, 0x0a,
	0x07, 0x68, 0x65, 0x72, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x68, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65,
	0x61, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x61, 0x74, 0x68, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x64, 0x65, 0x61, 0x74, 0x68, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x6e, 0x69, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x64, 0x65, 0x6e, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x68, 0x69, 0x74, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x48, 0x69, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x6f, 0x6c, 0x64, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x67, 0x6f, 0x6c, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x74,
	0x5f, 0x77, 0x6f, 0x72, 0x74, 0x68, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6e, 0x65,
	0x74, 0x57, 0x6f, 0x72, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6c, 0x75, 0x67, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67,
	0x22, 0x3f, 0x0a, 0x0b, 0x4c, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12,
	0x30, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c,
	0x69, 0x76, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x22, 0x70, 0x0a, 0x11, 0x4c, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x52, 0x02, 0x6f,
	0x70, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x4c, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x06, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x31, 0x33, 0x6b, 0x2f, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x2d, 0x73, 0x74, 0x61, 0x6c,
	0x6b, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_live_match_proto_rawDescOnce sync.Once
	file_protocol_live_match_proto_rawDescData = file_protocol_live_match_proto_rawDesc
)

func file_protocol_live_match_proto_rawDescGZIP() []byte {
	file_protocol_live_match_proto_rawDescOnce.Do(func() {
		file_protocol_live_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_live_match_proto_rawDescData)
	})
	return file_protocol_live_match_proto_rawDescData
}

var file_protocol_live_match_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protocol_live_match_proto_goTypes = []interface{}{
	(*LiveMatch)(nil),           // 0: ns.protocol.LiveMatch
	(*LiveMatches)(nil),         // 1: ns.protocol.LiveMatches
	(*LiveMatchesChange)(nil),   // 2: ns.protocol.LiveMatchesChange
	(*LiveMatch_Player)(nil),    // 3: ns.protocol.LiveMatch.Player
	(LobbyType)(0),              // 4: ns.protocol.LobbyType
	(GameMode)(0),               // 5: ns.protocol.GameMode
	(GameState)(0),              // 6: ns.protocol.GameState
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(CollectionOp)(0),           // 8: ns.protocol.CollectionOp
	(GameTeam)(0),               // 9: ns.protocol.GameTeam
}
var file_protocol_live_match_proto_depIdxs = []int32{
	4,  // 0: ns.protocol.LiveMatch.lobby_type:type_name -> ns.protocol.LobbyType
	5,  // 1: ns.protocol.LiveMatch.game_mode:type_name -> ns.protocol.GameMode
	6,  // 2: ns.protocol.LiveMatch.game_state:type_name -> ns.protocol.GameState
	7,  // 3: ns.protocol.LiveMatch.activate_time:type_name -> google.protobuf.Timestamp
	7,  // 4: ns.protocol.LiveMatch.deactivate_time:type_name -> google.protobuf.Timestamp
	7,  // 5: ns.protocol.LiveMatch.last_update_time:type_name -> google.protobuf.Timestamp
	3,  // 6: ns.protocol.LiveMatch.players:type_name -> ns.protocol.LiveMatch.Player
	0,  // 7: ns.protocol.LiveMatches.matches:type_name -> ns.protocol.LiveMatch
	8,  // 8: ns.protocol.LiveMatchesChange.op:type_name -> ns.protocol.CollectionOp
	1,  // 9: ns.protocol.LiveMatchesChange.change:type_name -> ns.protocol.LiveMatches
	9,  // 10: ns.protocol.LiveMatch.Player.team:type_name -> ns.protocol.GameTeam
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_protocol_live_match_proto_init() }
func file_protocol_live_match_proto_init() {
	if File_protocol_live_match_proto != nil {
		return
	}
	file_protocol_collections_proto_init()
	file_protocol_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protocol_live_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protocol_live_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveMatches); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protocol_live_match_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveMatchesChange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protocol_live_match_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveMatch_Player); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protocol_live_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_live_match_proto_goTypes,
		DependencyIndexes: file_protocol_live_match_proto_depIdxs,
		MessageInfos:      file_protocol_live_match_proto_msgTypes,
	}.Build()
	File_protocol_live_match_proto = out.File
	file_protocol_live_match_proto_rawDesc = nil
	file_protocol_live_match_proto_goTypes = nil
	file_protocol_live_match_proto_depIdxs = nil
}
