// Code generated by protoc-gen-go. DO NOT EDIT.
// source: live_match.proto

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
	return fileDescriptor_6d0b6910c4ef4a19, []int{0}
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
	return LobbyType_LOBBY_TYPE_CASUAL_MATCH
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
	return GameMode_GAME_MODE_NONE
}

func (m *LiveMatch) GetGameState() GameState {
	if m != nil {
		return m.GameState
	}
	return GameState_GAME_STATE_INIT
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
	Label                string   `protobuf:"bytes,19,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LiveMatch_Player) Reset()         { *m = LiveMatch_Player{} }
func (m *LiveMatch_Player) String() string { return proto.CompactTextString(m) }
func (*LiveMatch_Player) ProtoMessage()    {}
func (*LiveMatch_Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d0b6910c4ef4a19, []int{0, 0}
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
	return GameTeam_GAME_TEAM_UNKNOWN
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

func (m *LiveMatch_Player) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func init() {
	proto.RegisterType((*LiveMatch)(nil), "protocol.LiveMatch")
	proto.RegisterType((*LiveMatch_Player)(nil), "protocol.LiveMatch.Player")
}

func init() { proto.RegisterFile("live_match.proto", fileDescriptor_6d0b6910c4ef4a19) }

var fileDescriptor_6d0b6910c4ef4a19 = []byte{
	// 1051 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xd1, 0x6f, 0xdb, 0xb6,
	0x13, 0x86, 0x1b, 0xc7, 0xb1, 0xe9, 0xd8, 0x49, 0xd8, 0x26, 0x65, 0x9d, 0x5f, 0x1a, 0x35, 0x6d,
	0x02, 0xe3, 0x37, 0xc0, 0xc1, 0xb2, 0x3d, 0xec, 0x65, 0x18, 0xb6, 0x15, 0xdb, 0x0c, 0x24, 0x45,
	0xa1, 0xb8, 0xd8, 0xa3, 0x40, 0x9b, 0x57, 0x45, 0x08, 0x25, 0x1a, 0x24, 0xe5, 0xc2, 0xff, 0xca,
	0x9e, 0xf7, 0x87, 0x0e, 0x77, 0x94, 0xec, 0x38, 0x29, 0xd0, 0x27, 0xeb, 0xbe, 0xfb, 0xbe, 0x4f,
	0xf4, 0x89, 0x77, 0xc7, 0xf6, 0x75, 0xb6, 0x80, 0x24, 0x97, 0x7e, 0x76, 0x37, 0x9a, 0x5b, 0xe3,
	0x0d, 0x6f, 0xd3, 0xcf, 0xcc, 0xe8, 0xc1, 0x69, 0x6a, 0x4c, 0xaa, 0xe1, 0x92, 0x80, 0x69, 0xf9,
	0xf9, 0xd2, 0x67, 0x39, 0x38, 0x2f, 0xf3, 0x79, 0xa0, 0x0e, 0xba, 0x50, 0x94, 0xb9, 0x0b, 0xc1,
	0xd9, 0x3f, 0x07, 0xac, 0x73, 0x9d, 0x2d, 0xe0, 0x06, 0xbd, 0xf8, 0x2b, 0xd6, 0x26, 0xd3, 0x24,
	0x53, 0xa2, 0x11, 0x35, 0x86, 0xcd, 0x78, 0x87, 0xe2, 0xb1, 0xe2, 0x17, 0x6c, 0xcf, 0x81, 0x5d,
	0x80, 0x4d, 0x9c, 0x07, 0x99, 0x23, 0xe3, 0x19, 0x31, 0x7a, 0x01, 0xbe, 0x45, 0x74, 0xac, 0xd0,
	0x42, 0x9b, 0xe9, 0x74, 0x89, 0x84, 0xad, 0x60, 0x41, 0xf1, 0x58, 0xf1, 0x2b, 0xc6, 0x42, 0xca,
	0x2f, 0xe7, 0x20, 0x9a, 0x51, 0x63, 0xd8, 0xbf, 0x7a, 0x3e, 0xaa, 0x0f, 0x3e, 0xba, 0xc6, 0xdc,
	0x64, 0x39, 0x87, 0xb8, 0xa3, 0xeb, 0x47, 0x7e, 0xcc, 0x3a, 0x1a, 0x64, 0x5a, 0x02, 0xfa, 0x6d,
	0x93, 0x5f, 0x3b, 0x00, 0x63, 0x85, 0x49, 0x07, 0x36, 0x03, 0x87, 0xc9, 0x56, 0x48, 0x06, 0x60,
	0xac, 0xf8, 0x25, 0xeb, 0xa4, 0x32, 0x87, 0x24, 0x37, 0x0a, 0xc4, 0x0e, 0xbd, 0x8c, 0xaf, 0x5f,
	0xf6, 0xa7, 0xcc, 0xe1, 0xc6, 0x28, 0x88, 0xdb, 0x69, 0xf5, 0x84, 0xc7, 0x23, 0x81, 0xf3, 0xd2,
	0x83, 0x68, 0x3f, 0x3e, 0x1e, 0x2a, 0x6e, 0x31, 0x15, 0x93, 0x2f, 0x3d, 0xf2, 0x73, 0xd6, 0x27,
	0xcd, 0xaa, 0xc6, 0xa2, 0x13, 0x35, 0x86, 0xbd, 0xb8, 0x87, 0xe8, 0xa4, 0x06, 0xf1, 0xa0, 0x2b,
	0x9a, 0x60, 0x51, 0x63, 0xb8, 0x1d, 0xde, 0x8b, 0x0c, 0x7e, 0xca, 0xba, 0x72, 0x01, 0x56, 0xa6,
	0x90, 0xe4, 0xb9, 0x15, 0x5d, 0x32, 0x60, 0x15, 0x74, 0x93, 0x5b, 0xfe, 0x82, 0x6d, 0x2b, 0xd0,
	0x72, 0x29, 0x76, 0x29, 0x15, 0x02, 0xfe, 0x9a, 0x31, 0x37, 0x87, 0x99, 0x97, 0xde, 0x58, 0x27,
	0x7a, 0x41, 0xb5, 0x46, 0xf8, 0x09, 0x63, 0xce, 0x58, 0x9f, 0xb8, 0x99, 0xb1, 0x20, 0xfa, 0x51,
	0x63, 0xd8, 0x88, 0x3b, 0x88, 0xdc, 0x22, 0xc0, 0xdf, 0xb0, 0x5d, 0x2b, 0x55, 0x26, 0x0b, 0x9f,
	0x68, 0x90, 0x4a, 0xec, 0xd1, 0xa9, 0xba, 0x15, 0x76, 0x0d, 0x52, 0xf1, 0xb7, 0xac, 0x57, 0x53,
	0x82, 0xc9, 0x3e, 0xbd, 0xa4, 0xd6, 0x05, 0x9f, 0x0b, 0xb6, 0x57, 0x93, 0xea, 0x7b, 0x71, 0x10,
	0xee, 0x45, 0x05, 0x4f, 0xc2, 0xbd, 0xf8, 0x3f, 0x3b, 0xd8, 0xe0, 0x15, 0x32, 0x07, 0xc1, 0xa3,
	0xc6, 0xb0, 0x13, 0xef, 0x3d, 0x60, 0x7e, 0x90, 0x39, 0xf0, 0x21, 0xdb, 0xdf, 0xe0, 0x7a, 0x99,
	0x8a, 0xe7, 0x44, 0xed, 0x3f, 0xa0, 0x4e, 0x64, 0xfa, 0xc4, 0x55, 0x9b, 0xd4, 0x88, 0x17, 0xf4,
	0xfe, 0x87, 0xae, 0xd7, 0x26, 0x35, 0xfc, 0x7b, 0x76, 0xf8, 0x84, 0x9b, 0x94, 0x56, 0x8b, 0x43,
	0xb2, 0xe6, 0x8f, 0xf8, 0x9f, 0xac, 0x7e, 0x68, 0x5f, 0x80, 0x4f, 0xbe, 0x18, 0xeb, 0xef, 0xc4,
	0x11, 0x55, 0xa1, 0xb6, 0xff, 0x00, 0xfe, 0x6f, 0x84, 0xb1, 0xde, 0x2a, 0xb3, 0x50, 0x95, 0xea,
	0x25, 0x91, 0x3a, 0x88, 0x84, 0x3a, 0x45, 0x6c, 0x97, 0xd2, 0x75, 0x91, 0x04, 0x1d, 0x92, 0x24,
	0x55, 0x85, 0xde, 0xb1, 0xfe, 0x9a, 0x41, 0xe5, 0x79, 0x45, 0x07, 0xdb, 0xad, 0x39, 0x54, 0x9b,
	0x33, 0xd6, 0x5b, 0xb3, 0xb0, 0x30, 0x03, 0x22, 0x75, 0x6b, 0x12, 0x56, 0x65, 0xc3, 0x89, 0x4a,
	0x72, 0x4c, 0x6f, 0x5b, 0x39, 0x51, 0x3d, 0xbe, 0x63, 0x7c, 0x93, 0x45, 0xc5, 0xf8, 0x5f, 0xf8,
	0x24, 0x0f, 0x99, 0x58, 0x89, 0xda, 0x72, 0x5d, 0x86, 0x93, 0x70, 0x19, 0x10, 0x5d, 0xd5, 0xe0,
	0x9c, 0xf5, 0xa7, 0x65, 0xa6, 0x55, 0x56, 0xa4, 0x55, 0x1b, 0xbd, 0x0e, 0xed, 0x50, 0xa3, 0xa1,
	0x6b, 0x7e, 0x65, 0x27, 0x5f, 0x00, 0xee, 0xa1, 0x50, 0x89, 0x37, 0xa5, 0x2d, 0x60, 0x19, 0x7e,
	0x65, 0x0e, 0x85, 0xc7, 0xe2, 0x9c, 0x92, 0x6a, 0x50, 0x91, 0x26, 0x81, 0x33, 0x59, 0x51, 0xc6,
	0x8a, 0xff, 0xc4, 0xc4, 0x63, 0x0b, 0x95, 0x2d, 0x32, 0x97, 0x99, 0x42, 0x44, 0xa4, 0x3e, 0xda,
	0x54, 0xbf, 0xaf, 0xb2, 0xfc, 0x67, 0x76, 0xfc, 0x58, 0xe9, 0xee, 0x33, 0xad, 0x13, 0x0d, 0x0b,
	0xd0, 0xe2, 0x0d, 0x89, 0xc5, 0xa6, 0xf8, 0x16, 0x09, 0xd7, 0x98, 0xff, 0xda, 0xd9, 0xa7, 0x56,
	0xce, 0xee, 0xc1, 0x27, 0xd6, 0x94, 0x85, 0x12, 0x67, 0x5f, 0x3b, 0xfb, 0x6f, 0x81, 0x12, 0x23,
	0x83, 0xff, 0xc2, 0x7a, 0x72, 0xe6, 0xb3, 0x85, 0xf4, 0xd5, 0x44, 0x78, 0x1b, 0x35, 0x86, 0xdd,
	0xab, 0xc1, 0x28, 0x4c, 0xee, 0x51, 0x3d, 0xb9, 0x47, 0xab, 0x01, 0x12, 0xef, 0xd6, 0x02, 0x9a,
	0x18, 0xbf, 0xb3, 0x3d, 0x05, 0x9b, 0x16, 0xef, 0xbe, 0x69, 0xd1, 0x5f, 0x4b, 0xc8, 0xe4, 0x3d,
	0xdb, 0xd7, 0xd2, 0xf9, 0xa4, 0x9c, 0xab, 0x95, 0xcb, 0xf9, 0xb7, 0x5d, 0x50, 0xf3, 0x89, 0x24,
	0xe4, 0xf2, 0x23, 0xdb, 0x99, 0x6b, 0xb9, 0x04, 0xeb, 0x84, 0x8a, 0xb6, 0x48, 0xbc, 0x1e, 0xe8,
	0xf5, 0x5e, 0x19, 0x7d, 0x24, 0x4a, 0x5c, 0x53, 0x07, 0xff, 0x36, 0x59, 0x2b, 0x60, 0xd8, 0x36,
	0x72, 0x36, 0x33, 0x65, 0xf8, 0xf0, 0x8d, 0xd0, 0x36, 0x15, 0x32, 0x56, 0x9c, 0xb3, 0x26, 0xb5,
	0xc2, 0x33, 0xba, 0x96, 0xf4, 0x8c, 0xa3, 0x6b, 0x0e, 0xd6, 0x99, 0x42, 0x86, 0x36, 0xd9, 0x0a,
	0x1d, 0x50, 0x61, 0xd4, 0x25, 0xe8, 0xba, 0x90, 0x5e, 0x5a, 0xba, 0xd3, 0x4d, 0x22, 0x74, 0x02,
	0x52, 0xf5, 0x75, 0x95, 0xce, 0x41, 0x65, 0x65, 0x4e, 0xac, 0xed, 0x70, 0xf3, 0x43, 0xe2, 0x86,
	0x70, 0xe4, 0x5e, 0xb0, 0x0a, 0x4a, 0x3e, 0x97, 0x5a, 0x13, 0xb3, 0x45, 0xcc, 0x5e, 0x80, 0xff,
	0x28, 0xb5, 0x46, 0xde, 0x21, 0x6b, 0x65, 0x2e, 0x99, 0x5b, 0x43, 0xcb, 0xa6, 0x1d, 0x6f, 0x67,
	0xee, 0xa3, 0x35, 0xfc, 0x25, 0xdb, 0xb9, 0x03, 0x6b, 0xf0, 0xcf, 0xb5, 0xa9, 0x09, 0x5b, 0x18,
	0x8e, 0x15, 0x8e, 0xfd, 0x50, 0x8e, 0xc4, 0x69, 0xe3, 0xab, 0xbd, 0xc1, 0x02, 0x74, 0xab, 0x8d,
	0xe7, 0x17, 0xac, 0x89, 0xad, 0x49, 0xfb, 0xe2, 0xc9, 0xee, 0xc2, 0xde, 0x8c, 0x29, 0x8f, 0xeb,
	0x21, 0x5c, 0xdd, 0xb0, 0x39, 0x42, 0x80, 0x28, 0x5e, 0x5a, 0x57, 0x2f, 0x0d, 0x0a, 0xf8, 0x11,
	0x6b, 0x29, 0x90, 0xfe, 0xae, 0x5e, 0x18, 0x55, 0xc4, 0x05, 0xdb, 0x91, 0xce, 0x65, 0xce, 0x3b,
	0xda, 0x14, 0xbd, 0xb8, 0x0e, 0x83, 0xa2, 0xc8, 0xc0, 0xd1, 0x86, 0x20, 0x05, 0x46, 0xb4, 0x98,
	0xf1, 0xfa, 0xdc, 0x65, 0xde, 0x55, 0x8b, 0xa1, 0x8d, 0xc0, 0x5f, 0x99, 0x77, 0xf8, 0xd5, 0x52,
	0xa3, 0xc3, 0x26, 0xe8, 0xc5, 0xf4, 0x8c, 0x82, 0xf5, 0xf0, 0xe0, 0x41, 0x50, 0xd4, 0x83, 0x03,
	0xff, 0x83, 0x9c, 0x82, 0xae, 0xc6, 0x7c, 0x08, 0xa6, 0x2d, 0xfa, 0xcb, 0x3f, 0xfc, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x52, 0xc4, 0xe5, 0xa7, 0xef, 0x08, 0x00, 0x00,
}
