// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/protocol/live_match.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LiveMatch struct {
	MatchId                    uint64               `protobuf:"varint,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	ServerSteamId              uint64               `protobuf:"varint,2,opt,name=server_steam_id,json=serverSteamId,proto3" json:"server_steam_id,omitempty"`
	LobbyId                    uint64               `protobuf:"varint,3,opt,name=lobby_id,json=lobbyId,proto3" json:"lobby_id,omitempty"`
	LobbyType                  LobbyType            `protobuf:"varint,4,opt,name=lobby_type,json=lobbyType,proto3,enum=protocol.LobbyType" json:"lobby_type,omitempty"`
	LeagueId                   uint64               `protobuf:"varint,5,opt,name=league_id,json=leagueId,proto3" json:"league_id,omitempty"`
	SeriesId                   uint64               `protobuf:"varint,6,opt,name=series_id,json=seriesId,proto3" json:"series_id,omitempty"`
	GameMode                   GameMode             `protobuf:"varint,7,opt,name=game_mode,json=gameMode,proto3,enum=protocol.GameMode" json:"game_mode,omitempty"`
	GameState                  GameState            `protobuf:"varint,8,opt,name=game_state,json=gameState,proto3,enum=protocol.GameState" json:"game_state,omitempty"`
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
	XXX_NoUnkeyedLiteral       struct{}             `json:"-"`
	XXX_unrecognized           []byte               `json:"-"`
	XXX_sizecache              int32                `json:"-"`
}

func (m *LiveMatch) Reset()         { *m = LiveMatch{} }
func (m *LiveMatch) String() string { return proto.CompactTextString(m) }
func (*LiveMatch) ProtoMessage()    {}
func (*LiveMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fd5bbb67bd93734, []int{0}
}

func (m *LiveMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiveMatch.Unmarshal(m, b)
}
func (m *LiveMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiveMatch.Marshal(b, m, deterministic)
}
func (m *LiveMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiveMatch.Merge(m, src)
}
func (m *LiveMatch) XXX_Size() int {
	return xxx_messageInfo_LiveMatch.Size(m)
}
func (m *LiveMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_LiveMatch.DiscardUnknown(m)
}

var xxx_messageInfo_LiveMatch proto.InternalMessageInfo

func (m *LiveMatch) GetMatchId() uint64 {
	if m != nil {
		return m.MatchId
	}
	return 0
}

func (m *LiveMatch) GetServerSteamId() uint64 {
	if m != nil {
		return m.ServerSteamId
	}
	return 0
}

func (m *LiveMatch) GetLobbyId() uint64 {
	if m != nil {
		return m.LobbyId
	}
	return 0
}

func (m *LiveMatch) GetLobbyType() LobbyType {
	if m != nil {
		return m.LobbyType
	}
	return LobbyType_LobbyTypeCasualMatch
}

func (m *LiveMatch) GetLeagueId() uint64 {
	if m != nil {
		return m.LeagueId
	}
	return 0
}

func (m *LiveMatch) GetSeriesId() uint64 {
	if m != nil {
		return m.SeriesId
	}
	return 0
}

func (m *LiveMatch) GetGameMode() GameMode {
	if m != nil {
		return m.GameMode
	}
	return GameMode_GameModeNone
}

func (m *LiveMatch) GetGameState() GameState {
	if m != nil {
		return m.GameState
	}
	return GameState_GameStateInit
}

func (m *LiveMatch) GetGameTimestamp() uint32 {
	if m != nil {
		return m.GameTimestamp
	}
	return 0
}

func (m *LiveMatch) GetGameTime() int32 {
	if m != nil {
		return m.GameTime
	}
	return 0
}

func (m *LiveMatch) GetAverageMmr() uint32 {
	if m != nil {
		return m.AverageMmr
	}
	return 0
}

func (m *LiveMatch) GetDelay() uint32 {
	if m != nil {
		return m.Delay
	}
	return 0
}

func (m *LiveMatch) GetSpectators() uint32 {
	if m != nil {
		return m.Spectators
	}
	return 0
}

func (m *LiveMatch) GetSortScore() float64 {
	if m != nil {
		return m.SortScore
	}
	return 0
}

func (m *LiveMatch) GetRadiantLead() int32 {
	if m != nil {
		return m.RadiantLead
	}
	return 0
}

func (m *LiveMatch) GetRadiantScore() uint32 {
	if m != nil {
		return m.RadiantScore
	}
	return 0
}

func (m *LiveMatch) GetRadiantTeamId() uint64 {
	if m != nil {
		return m.RadiantTeamId
	}
	return 0
}

func (m *LiveMatch) GetRadiantTeamName() string {
	if m != nil {
		return m.RadiantTeamName
	}
	return ""
}

func (m *LiveMatch) GetRadiantTeamTag() string {
	if m != nil {
		return m.RadiantTeamTag
	}
	return ""
}

func (m *LiveMatch) GetRadiantTeamLogo() uint64 {
	if m != nil {
		return m.RadiantTeamLogo
	}
	return 0
}

func (m *LiveMatch) GetRadiantTeamLogoUrl() string {
	if m != nil {
		return m.RadiantTeamLogoUrl
	}
	return ""
}

func (m *LiveMatch) GetRadiantNetWorth() uint32 {
	if m != nil {
		return m.RadiantNetWorth
	}
	return 0
}

func (m *LiveMatch) GetDireScore() uint32 {
	if m != nil {
		return m.DireScore
	}
	return 0
}

func (m *LiveMatch) GetDireTeamId() uint64 {
	if m != nil {
		return m.DireTeamId
	}
	return 0
}

func (m *LiveMatch) GetDireTeamName() string {
	if m != nil {
		return m.DireTeamName
	}
	return ""
}

func (m *LiveMatch) GetDireTeamTag() string {
	if m != nil {
		return m.DireTeamTag
	}
	return ""
}

func (m *LiveMatch) GetDireTeamLogo() uint64 {
	if m != nil {
		return m.DireTeamLogo
	}
	return 0
}

func (m *LiveMatch) GetDireTeamLogoUrl() string {
	if m != nil {
		return m.DireTeamLogoUrl
	}
	return ""
}

func (m *LiveMatch) GetDireNetWorth() uint32 {
	if m != nil {
		return m.DireNetWorth
	}
	return 0
}

func (m *LiveMatch) GetBuildingState() uint32 {
	if m != nil {
		return m.BuildingState
	}
	return 0
}

func (m *LiveMatch) GetWeekendTourneyTournamentId() uint32 {
	if m != nil {
		return m.WeekendTourneyTournamentId
	}
	return 0
}

func (m *LiveMatch) GetWeekendTourneyDivision() uint32 {
	if m != nil {
		return m.WeekendTourneyDivision
	}
	return 0
}

func (m *LiveMatch) GetWeekendTourneySkillLevel() uint32 {
	if m != nil {
		return m.WeekendTourneySkillLevel
	}
	return 0
}

func (m *LiveMatch) GetWeekendTourneyBracketRound() uint32 {
	if m != nil {
		return m.WeekendTourneyBracketRound
	}
	return 0
}

func (m *LiveMatch) GetActivateTime() *timestamp.Timestamp {
	if m != nil {
		return m.ActivateTime
	}
	return nil
}

func (m *LiveMatch) GetDeactivateTime() *timestamp.Timestamp {
	if m != nil {
		return m.DeactivateTime
	}
	return nil
}

func (m *LiveMatch) GetLastUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *LiveMatch) GetPlayers() []*LiveMatch_Player {
	if m != nil {
		return m.Players
	}
	return nil
}

type LiveMatch_Player struct {
	AccountId            uint32   `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PersonaName          string   `protobuf:"bytes,3,opt,name=persona_name,json=personaName,proto3" json:"persona_name,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	AvatarMediumUrl      string   `protobuf:"bytes,5,opt,name=avatar_medium_url,json=avatarMediumUrl,proto3" json:"avatar_medium_url,omitempty"`
	AvatarFullUrl        string   `protobuf:"bytes,6,opt,name=avatar_full_url,json=avatarFullUrl,proto3" json:"avatar_full_url,omitempty"`
	IsPro                bool     `protobuf:"varint,7,opt,name=is_pro,json=isPro,proto3" json:"is_pro,omitempty"`
	HeroId               uint64   `protobuf:"varint,8,opt,name=hero_id,json=heroId,proto3" json:"hero_id,omitempty"`
	PlayerSlot           uint32   `protobuf:"varint,9,opt,name=player_slot,json=playerSlot,proto3" json:"player_slot,omitempty"`
	Team                 GameTeam `protobuf:"varint,10,opt,name=team,proto3,enum=protocol.GameTeam" json:"team,omitempty"`
	Level                uint32   `protobuf:"varint,11,opt,name=level,proto3" json:"level,omitempty"`
	Kills                uint32   `protobuf:"varint,12,opt,name=kills,proto3" json:"kills,omitempty"`
	Deaths               uint32   `protobuf:"varint,13,opt,name=deaths,proto3" json:"deaths,omitempty"`
	Assists              uint32   `protobuf:"varint,14,opt,name=assists,proto3" json:"assists,omitempty"`
	Denies               uint32   `protobuf:"varint,15,opt,name=denies,proto3" json:"denies,omitempty"`
	LastHits             uint32   `protobuf:"varint,16,opt,name=last_hits,json=lastHits,proto3" json:"last_hits,omitempty"`
	Gold                 uint32   `protobuf:"varint,17,opt,name=gold,proto3" json:"gold,omitempty"`
	NetWorth             uint32   `protobuf:"varint,18,opt,name=net_worth,json=netWorth,proto3" json:"net_worth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LiveMatch_Player) Reset()         { *m = LiveMatch_Player{} }
func (m *LiveMatch_Player) String() string { return proto.CompactTextString(m) }
func (*LiveMatch_Player) ProtoMessage()    {}
func (*LiveMatch_Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fd5bbb67bd93734, []int{0, 0}
}

func (m *LiveMatch_Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiveMatch_Player.Unmarshal(m, b)
}
func (m *LiveMatch_Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiveMatch_Player.Marshal(b, m, deterministic)
}
func (m *LiveMatch_Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiveMatch_Player.Merge(m, src)
}
func (m *LiveMatch_Player) XXX_Size() int {
	return xxx_messageInfo_LiveMatch_Player.Size(m)
}
func (m *LiveMatch_Player) XXX_DiscardUnknown() {
	xxx_messageInfo_LiveMatch_Player.DiscardUnknown(m)
}

var xxx_messageInfo_LiveMatch_Player proto.InternalMessageInfo

func (m *LiveMatch_Player) GetAccountId() uint32 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *LiveMatch_Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LiveMatch_Player) GetPersonaName() string {
	if m != nil {
		return m.PersonaName
	}
	return ""
}

func (m *LiveMatch_Player) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *LiveMatch_Player) GetAvatarMediumUrl() string {
	if m != nil {
		return m.AvatarMediumUrl
	}
	return ""
}

func (m *LiveMatch_Player) GetAvatarFullUrl() string {
	if m != nil {
		return m.AvatarFullUrl
	}
	return ""
}

func (m *LiveMatch_Player) GetIsPro() bool {
	if m != nil {
		return m.IsPro
	}
	return false
}

func (m *LiveMatch_Player) GetHeroId() uint64 {
	if m != nil {
		return m.HeroId
	}
	return 0
}

func (m *LiveMatch_Player) GetPlayerSlot() uint32 {
	if m != nil {
		return m.PlayerSlot
	}
	return 0
}

func (m *LiveMatch_Player) GetTeam() GameTeam {
	if m != nil {
		return m.Team
	}
	return GameTeam_GameTeamUnknown
}

func (m *LiveMatch_Player) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *LiveMatch_Player) GetKills() uint32 {
	if m != nil {
		return m.Kills
	}
	return 0
}

func (m *LiveMatch_Player) GetDeaths() uint32 {
	if m != nil {
		return m.Deaths
	}
	return 0
}

func (m *LiveMatch_Player) GetAssists() uint32 {
	if m != nil {
		return m.Assists
	}
	return 0
}

func (m *LiveMatch_Player) GetDenies() uint32 {
	if m != nil {
		return m.Denies
	}
	return 0
}

func (m *LiveMatch_Player) GetLastHits() uint32 {
	if m != nil {
		return m.LastHits
	}
	return 0
}

func (m *LiveMatch_Player) GetGold() uint32 {
	if m != nil {
		return m.Gold
	}
	return 0
}

func (m *LiveMatch_Player) GetNetWorth() uint32 {
	if m != nil {
		return m.NetWorth
	}
	return 0
}

func init() {
	proto.RegisterType((*LiveMatch)(nil), "protocol.LiveMatch")
	proto.RegisterType((*LiveMatch_Player)(nil), "protocol.LiveMatch.Player")
}

func init() { proto.RegisterFile("internal/protocol/live_match.proto", fileDescriptor_5fd5bbb67bd93734) }

var fileDescriptor_5fd5bbb67bd93734 = []byte{
	// 1052 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x41, 0x6f, 0x1b, 0x37,
	0x13, 0x85, 0x12, 0x59, 0x96, 0x68, 0x4b, 0x8e, 0x99, 0xd8, 0x61, 0xe4, 0xcf, 0xb1, 0xe2, 0xc4,
	0x86, 0xf0, 0x15, 0x90, 0x51, 0xb7, 0x87, 0x5e, 0x8a, 0xa2, 0x6d, 0xd0, 0x56, 0x80, 0x1d, 0x04,
	0x6b, 0x05, 0x3d, 0x2e, 0x28, 0x71, 0xb2, 0x26, 0xcc, 0x5d, 0x0a, 0x24, 0x57, 0x81, 0xee, 0x3d,
	0xf7, 0x37, 0x17, 0x33, 0xdc, 0x95, 0x2c, 0x3b, 0x40, 0x4e, 0x22, 0xdf, 0xbc, 0xf7, 0x48, 0xcd,
	0x72, 0x66, 0xd8, 0xa9, 0x2e, 0x02, 0xb8, 0x42, 0x9a, 0x8b, 0xb9, 0xb3, 0xc1, 0xce, 0xac, 0xb9,
	0x30, 0x7a, 0x01, 0x69, 0x2e, 0xc3, 0xec, 0x76, 0x44, 0x18, 0x6f, 0xd7, 0xa1, 0xfe, 0x49, 0x66,
	0x6d, 0x66, 0x20, 0x72, 0xa7, 0xe5, 0xe7, 0x8b, 0xa0, 0x73, 0xf0, 0x41, 0xe6, 0xf3, 0x48, 0xed,
	0x1f, 0x3f, 0xb6, 0x83, 0xa2, 0xcc, 0x7d, 0x0c, 0x9f, 0xfe, 0xb3, 0xcf, 0x3a, 0x57, 0x7a, 0x01,
	0xd7, 0xe8, 0xce, 0x5f, 0xb1, 0x36, 0x1d, 0x93, 0x6a, 0x25, 0x1a, 0x83, 0xc6, 0xb0, 0x99, 0x6c,
	0xd3, 0x7e, 0xac, 0xf8, 0x39, 0xdb, 0xf3, 0xe0, 0x16, 0xe0, 0x52, 0x1f, 0x40, 0xe6, 0xc8, 0x78,
	0x42, 0x8c, 0x6e, 0x84, 0x6f, 0x10, 0x1d, 0x2b, 0xb4, 0x30, 0x76, 0x3a, 0x5d, 0x22, 0xe1, 0x69,
	0xb4, 0xa0, 0xfd, 0x58, 0xf1, 0x4b, 0xc6, 0x62, 0x28, 0x2c, 0xe7, 0x20, 0x9a, 0x83, 0xc6, 0xb0,
	0x77, 0xf9, 0x7c, 0x54, 0x5f, 0x6b, 0x74, 0x85, 0xb1, 0xc9, 0x72, 0x0e, 0x49, 0xc7, 0xd4, 0x4b,
	0x7e, 0xc4, 0x3a, 0x06, 0x64, 0x56, 0x02, 0xfa, 0x6d, 0x91, 0x5f, 0x3b, 0x02, 0x63, 0x85, 0x41,
	0x0f, 0x4e, 0x83, 0xc7, 0x60, 0x2b, 0x06, 0x23, 0x30, 0x56, 0xfc, 0x82, 0x75, 0x32, 0x99, 0x43,
	0x9a, 0x5b, 0x05, 0x62, 0x9b, 0x0e, 0xe3, 0xeb, 0xc3, 0xfe, 0x94, 0x39, 0x5c, 0x5b, 0x05, 0x49,
	0x3b, 0xab, 0x56, 0x78, 0x3d, 0x12, 0xf8, 0x20, 0x03, 0x88, 0xf6, 0xc3, 0xeb, 0xa1, 0xe2, 0x06,
	0x43, 0x09, 0xf9, 0xd2, 0x92, 0x9f, 0xb1, 0x1e, 0x69, 0x56, 0x59, 0x17, 0x9d, 0x41, 0x63, 0xd8,
	0x4d, 0xba, 0x88, 0x4e, 0x6a, 0x10, 0x2f, 0xba, 0xa2, 0x09, 0x36, 0x68, 0x0c, 0xb7, 0xe2, 0xb9,
	0xc8, 0xe0, 0x27, 0x6c, 0x47, 0x2e, 0xc0, 0xc9, 0x0c, 0xd2, 0x3c, 0x77, 0x62, 0x87, 0x0c, 0x58,
	0x05, 0x5d, 0xe7, 0x8e, 0xbf, 0x60, 0x5b, 0x0a, 0x8c, 0x5c, 0x8a, 0x5d, 0x0a, 0xc5, 0x0d, 0x7f,
	0xcd, 0x98, 0x9f, 0xc3, 0x2c, 0xc8, 0x60, 0x9d, 0x17, 0xdd, 0xa8, 0x5a, 0x23, 0xfc, 0x98, 0x31,
	0x6f, 0x5d, 0x48, 0xfd, 0xcc, 0x3a, 0x10, 0xbd, 0x41, 0x63, 0xd8, 0x48, 0x3a, 0x88, 0xdc, 0x20,
	0xc0, 0xdf, 0xb0, 0x5d, 0x27, 0x95, 0x96, 0x45, 0x48, 0x0d, 0x48, 0x25, 0xf6, 0xe8, 0x56, 0x3b,
	0x15, 0x76, 0x05, 0x52, 0xf1, 0xb7, 0xac, 0x5b, 0x53, 0xa2, 0xc9, 0x33, 0x3a, 0xa4, 0xd6, 0x45,
	0x9f, 0x73, 0xb6, 0x57, 0x93, 0xea, 0x77, 0xb1, 0x1f, 0xdf, 0x45, 0x05, 0x4f, 0xe2, 0xbb, 0xf8,
	0x3f, 0xdb, 0xdf, 0xe0, 0x15, 0x32, 0x07, 0xc1, 0x07, 0x8d, 0x61, 0x27, 0xd9, 0xbb, 0xc7, 0xfc,
	0x20, 0x73, 0xe0, 0x43, 0xf6, 0x6c, 0x83, 0x1b, 0x64, 0x26, 0x9e, 0x13, 0xb5, 0x77, 0x8f, 0x3a,
	0x91, 0xd9, 0x23, 0x57, 0x63, 0x33, 0x2b, 0x5e, 0xd0, 0xf9, 0xf7, 0x5d, 0xaf, 0x6c, 0x66, 0xf9,
	0xf7, 0xec, 0xe0, 0x11, 0x37, 0x2d, 0x9d, 0x11, 0x07, 0x64, 0xcd, 0x1f, 0xf0, 0x3f, 0x39, 0x73,
	0xdf, 0xbe, 0x80, 0x90, 0x7e, 0xb1, 0x2e, 0xdc, 0x8a, 0x43, 0xca, 0x42, 0x6d, 0xff, 0x01, 0xc2,
	0xdf, 0x08, 0x63, 0xbe, 0x95, 0x76, 0x50, 0xa5, 0xea, 0x25, 0x91, 0x3a, 0x88, 0xc4, 0x3c, 0x0d,
	0xd8, 0x2e, 0x85, 0xeb, 0x24, 0x09, 0xba, 0x24, 0x49, 0xaa, 0x0c, 0xbd, 0x63, 0xbd, 0x35, 0x83,
	0xd2, 0xf3, 0x8a, 0x2e, 0xb6, 0x5b, 0x73, 0x28, 0x37, 0xa7, 0xac, 0xbb, 0x66, 0x61, 0x62, 0xfa,
	0x44, 0xda, 0xa9, 0x49, 0x98, 0x95, 0x0d, 0x27, 0x4a, 0xc9, 0x11, 0x9d, 0xb6, 0x72, 0xa2, 0x7c,
	0x7c, 0xc7, 0xf8, 0x26, 0x8b, 0x92, 0xf1, 0xbf, 0xf8, 0x49, 0xee, 0x33, 0x31, 0x13, 0xb5, 0xe5,
	0x3a, 0x0d, 0xc7, 0xf1, 0x31, 0x20, 0xba, 0xca, 0xc1, 0x19, 0xeb, 0x4d, 0x4b, 0x6d, 0x94, 0x2e,
	0xb2, 0xaa, 0x8c, 0x5e, 0xc7, 0x72, 0xa8, 0xd1, 0x58, 0x35, 0xbf, 0xb2, 0xe3, 0x2f, 0x00, 0x77,
	0x50, 0xa8, 0x34, 0xd8, 0xd2, 0x15, 0xb0, 0x8c, 0xbf, 0x32, 0x87, 0x22, 0x60, 0x72, 0x4e, 0x48,
	0xd5, 0xaf, 0x48, 0x93, 0xc8, 0x99, 0xac, 0x28, 0x63, 0xc5, 0x7f, 0x62, 0xe2, 0xa1, 0x85, 0xd2,
	0x0b, 0xed, 0xb5, 0x2d, 0xc4, 0x80, 0xd4, 0x87, 0x9b, 0xea, 0xf7, 0x55, 0x94, 0xff, 0xcc, 0x8e,
	0x1e, 0x2a, 0xfd, 0x9d, 0x36, 0x26, 0x35, 0xb0, 0x00, 0x23, 0xde, 0x90, 0x58, 0x6c, 0x8a, 0x6f,
	0x90, 0x70, 0x85, 0xf1, 0xaf, 0xdd, 0x7d, 0xea, 0xe4, 0xec, 0x0e, 0x42, 0xea, 0x6c, 0x59, 0x28,
	0x71, 0xfa, 0xb5, 0xbb, 0xff, 0x16, 0x29, 0x09, 0x32, 0xf8, 0x2f, 0xac, 0x2b, 0x67, 0x41, 0x2f,
	0x64, 0xa8, 0x3a, 0xc2, 0xdb, 0x41, 0x63, 0xb8, 0x73, 0xd9, 0x1f, 0xc5, 0x5e, 0x3e, 0xaa, 0x7b,
	0xf9, 0x68, 0xd5, 0x40, 0x92, 0xdd, 0x5a, 0x40, 0x1d, 0xe3, 0x77, 0xb6, 0xa7, 0x60, 0xd3, 0xe2,
	0xdd, 0x37, 0x2d, 0x7a, 0x6b, 0x09, 0x99, 0xbc, 0x67, 0xcf, 0x8c, 0xf4, 0x21, 0x2d, 0xe7, 0x6a,
	0xe5, 0x72, 0xf6, 0x6d, 0x17, 0xd4, 0x7c, 0x22, 0x09, 0xb9, 0xfc, 0xc8, 0xb6, 0xe7, 0x46, 0x2e,
	0xc1, 0x79, 0xa1, 0x06, 0x4f, 0x49, 0xbc, 0x6e, 0xe8, 0xf5, 0x5c, 0x19, 0x7d, 0x24, 0x4a, 0x52,
	0x53, 0xfb, 0xff, 0x36, 0x59, 0x2b, 0x62, 0x58, 0x36, 0x72, 0x36, 0xb3, 0x65, 0xfc, 0xf0, 0x8d,
	0x58, 0x36, 0x15, 0x32, 0x56, 0x9c, 0xb3, 0x26, 0x95, 0xc2, 0x13, 0x7a, 0x96, 0xb4, 0xc6, 0xd6,
	0x35, 0x07, 0xe7, 0x6d, 0x21, 0x63, 0x99, 0x3c, 0x8d, 0x15, 0x50, 0x61, 0x54, 0x25, 0xe8, 0xba,
	0x90, 0x41, 0x3a, 0x7a, 0xd3, 0x4d, 0x22, 0x74, 0x22, 0x52, 0xd5, 0x75, 0x15, 0xce, 0x41, 0xe9,
	0x32, 0x27, 0xd6, 0x56, 0x7c, 0xf9, 0x31, 0x70, 0x4d, 0x38, 0x72, 0xcf, 0x59, 0x05, 0xa5, 0x9f,
	0x4b, 0x63, 0x88, 0xd9, 0x22, 0x66, 0x37, 0xc2, 0x7f, 0x94, 0xc6, 0x20, 0xef, 0x80, 0xb5, 0xb4,
	0x4f, 0xe7, 0xce, 0xd2, 0xb0, 0x69, 0x27, 0x5b, 0xda, 0x7f, 0x74, 0x96, 0xbf, 0x64, 0xdb, 0xb7,
	0xe0, 0x2c, 0xfe, 0xb9, 0x36, 0x15, 0x61, 0x0b, 0xb7, 0x63, 0x85, 0x6d, 0x3f, 0xa6, 0x23, 0xf5,
	0xc6, 0x86, 0x6a, 0x6e, 0xb0, 0x08, 0xdd, 0x18, 0x1b, 0xf8, 0x39, 0x6b, 0x62, 0x69, 0xd2, 0xbc,
	0x78, 0x34, 0xbb, 0xb0, 0x36, 0x13, 0x8a, 0xe3, 0x78, 0x88, 0x4f, 0x37, 0x4e, 0x8e, 0xb8, 0x41,
	0x14, 0x1f, 0xad, 0xaf, 0x87, 0x06, 0x6d, 0xf8, 0x21, 0x6b, 0x29, 0x90, 0xe1, 0xb6, 0x1e, 0x18,
	0xd5, 0x8e, 0x0b, 0xb6, 0x2d, 0xbd, 0xd7, 0x3e, 0x78, 0x9a, 0x14, 0xdd, 0xa4, 0xde, 0x46, 0x45,
	0xa1, 0xc1, 0xd3, 0x84, 0x20, 0x05, 0xee, 0x68, 0x30, 0xe3, 0xf3, 0xb9, 0xd5, 0xc1, 0x57, 0x83,
	0xa1, 0x8d, 0xc0, 0x5f, 0x3a, 0x78, 0xfc, 0x6a, 0x99, 0x35, 0x71, 0x12, 0x74, 0x13, 0x5a, 0xa3,
	0x60, 0xdd, 0x3c, 0x78, 0x14, 0x14, 0x55, 0xe3, 0x98, 0xb6, 0xe8, 0xcf, 0xfd, 0xf0, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x64, 0x7b, 0x63, 0x0b, 0xfd, 0x08, 0x00, 0x00,
}
