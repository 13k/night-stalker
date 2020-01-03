// source: enums.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var proto = {};

goog.exportSymbol('protocol.BuildingType', null, proto);
goog.exportSymbol('protocol.FantasyRole', null, proto);
goog.exportSymbol('protocol.GameMode', null, proto);
goog.exportSymbol('protocol.GameState', null, proto);
goog.exportSymbol('protocol.GameTeam', null, proto);
goog.exportSymbol('protocol.LaneType', null, proto);
goog.exportSymbol('protocol.LobbyType', null, proto);
goog.exportSymbol('protocol.MatchOutcome', null, proto);
/**
 * @enum {number}
 */
proto.protocol.LobbyType = {
  LOBBY_TYPE_CASUAL_MATCH: 0,
  LOBBY_TYPE_PRACTICE: 1,
  LOBBY_TYPE_COOP_BOT_MATCH: 4,
  LOBBY_TYPE_LEGACY_TEAM_MATCH: 5,
  LOBBY_TYPE_LEGACY_SOLO_QUEUE_MATCH: 6,
  LOBBY_TYPE_COMPETITIVE_MATCH: 7,
  LOBBY_TYPE_CASUAL_1V1_MATCH: 8,
  LOBBY_TYPE_WEEKEND_TOURNEY: 9,
  LOBBY_TYPE_LOCAL_BOT_MATCH: 10,
  LOBBY_TYPE_SPECTATOR: 11,
  LOBBY_TYPE_EVENT_MATCH: 12
};

/**
 * @enum {number}
 */
proto.protocol.GameMode = {
  GAME_MODE_NONE: 0,
  GAME_MODE_AP: 1,
  GAME_MODE_CM: 2,
  GAME_MODE_RD: 3,
  GAME_MODE_SD: 4,
  GAME_MODE_AR: 5,
  GAME_MODE_INTRO: 6,
  GAME_MODE_HW: 7,
  GAME_MODE_REVERSE_CM: 8,
  GAME_MODE_XMAS: 9,
  GAME_MODE_TUTORIAL: 10,
  GAME_MODE_MO: 11,
  GAME_MODE_LP: 12,
  GAME_MODE_POOL1: 13,
  GAME_MODE_FH: 14,
  GAME_MODE_CUSTOM: 15,
  GAME_MODE_CD: 16,
  GAME_MODE_BD: 17,
  GAME_MODE_ABILITY_DRAFT: 18,
  GAME_MODE_EVENT: 19,
  GAME_MODE_ARDM: 20,
  GAME_MODE_1V1_MID: 21,
  GAME_MODE_ALL_DRAFT: 22,
  GAME_MODE_TURBO: 23,
  GAME_MODE_MUTATION: 24,
  GAME_MODE_COACHES_CHALLENGE: 25
};

/**
 * @enum {number}
 */
proto.protocol.GameState = {
  GAME_STATE_INIT: 0,
  GAME_STATE_WAIT_FOR_PLAYERS_TO_LOAD: 1,
  GAME_STATE_HERO_SELECTION: 2,
  GAME_STATE_STRATEGY_TIME: 3,
  GAME_STATE_PRE_GAME: 4,
  GAME_STATE_GAME_IN_PROGRESS: 5,
  GAME_STATE_POST_GAME: 6,
  GAME_STATE_DISCONNECT: 7,
  GAME_STATE_TEAM_SHOWCASE: 8,
  GAME_STATE_CUSTOM_GAME_SETUP: 9,
  GAME_STATE_WAIT_FOR_MAP_TO_LOAD: 10,
  GAME_STATE_LAST: 11
};

/**
 * @enum {number}
 */
proto.protocol.GameTeam = {
  GAME_TEAM_UNKNOWN: 0,
  GAME_TEAM_GOODGUYS: 2,
  GAME_TEAM_BADGUYS: 3,
  GAME_TEAM_NEUTRALS: 4,
  GAME_TEAM_NOTEAM: 5,
  GAME_TEAM_CUSTOM1: 6,
  GAME_TEAM_CUSTOM2: 7,
  GAME_TEAM_CUSTOM3: 8,
  GAME_TEAM_CUSTOM4: 9,
  GAME_TEAM_CUSTOM5: 10,
  GAME_TEAM_CUSTOM6: 11,
  GAME_TEAM_CUSTOM7: 12,
  GAME_TEAM_CUSTOM8: 13
};

/**
 * @enum {number}
 */
proto.protocol.BuildingType = {
  BUILDING_TYPE_TOWER: 0,
  BUILDING_TYPE_BARRACKS: 1,
  BUILDING_TYPE_ANCIENT: 2
};

/**
 * @enum {number}
 */
proto.protocol.FantasyRole = {
  FANTASY_ROLE_UNDEFINED: 0,
  FANTASY_ROLE_CORE: 1,
  FANTASY_ROLE_SUPPORT: 2,
  FANTASY_ROLE_OFFLANE: 3,
  FANTASY_ROLE_MID: 4
};

/**
 * @enum {number}
 */
proto.protocol.LaneType = {
  LANE_TYPE_UNKNOWN: 0,
  LANE_TYPE_SAFE: 1,
  LANE_TYPE_OFF: 2,
  LANE_TYPE_MID: 3,
  LANE_TYPE_JUNGLE: 4,
  LANE_TYPE_ROAM: 5
};

/**
 * @enum {number}
 */
proto.protocol.MatchOutcome = {
  MATCH_OUTCOME_UNKNOWN: 0,
  MATCH_OUTCOME_RAD_VICTORY: 2,
  MATCH_OUTCOME_DIRE_VICTORY: 3,
  MATCH_OUTCOME_NOT_SCORED_POOR_NETWORK_CONDITIONS: 64,
  MATCH_OUTCOME_NOT_SCORED_LEAVER: 65,
  MATCH_OUTCOME_NOT_SCORED_SERVER_CRASH: 66,
  MATCH_OUTCOME_NOT_SCORED_NEVER_STARTED: 67,
  MATCH_OUTCOME_NOT_SCORED_CANCELED: 68
};

goog.object.extend(exports, proto);
