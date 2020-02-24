/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
import * as $protobuf from "protobufjs/minimal";

// Common aliases
const $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const protocol = $root.protocol = (() => {

    /**
     * Namespace protocol.
     * @exports protocol
     * @namespace
     */
    const protocol = {};

    protocol.LiveMatch = (function() {

        /**
         * Properties of a LiveMatch.
         * @memberof protocol
         * @interface ILiveMatch
         * @property {Long|null} [match_id] LiveMatch match_id
         * @property {Long|null} [server_steam_id] LiveMatch server_steam_id
         * @property {Long|null} [lobby_id] LiveMatch lobby_id
         * @property {protocol.LobbyType|null} [lobby_type] LiveMatch lobby_type
         * @property {Long|null} [league_id] LiveMatch league_id
         * @property {Long|null} [series_id] LiveMatch series_id
         * @property {protocol.GameMode|null} [game_mode] LiveMatch game_mode
         * @property {protocol.GameState|null} [game_state] LiveMatch game_state
         * @property {number|null} [game_timestamp] LiveMatch game_timestamp
         * @property {number|null} [game_time] LiveMatch game_time
         * @property {number|null} [average_mmr] LiveMatch average_mmr
         * @property {number|null} [delay] LiveMatch delay
         * @property {number|null} [spectators] LiveMatch spectators
         * @property {number|null} [sort_score] LiveMatch sort_score
         * @property {number|null} [radiant_lead] LiveMatch radiant_lead
         * @property {number|null} [radiant_score] LiveMatch radiant_score
         * @property {Long|null} [radiant_team_id] LiveMatch radiant_team_id
         * @property {string|null} [radiant_team_name] LiveMatch radiant_team_name
         * @property {string|null} [radiant_team_tag] LiveMatch radiant_team_tag
         * @property {Long|null} [radiant_team_logo] LiveMatch radiant_team_logo
         * @property {string|null} [radiant_team_logo_url] LiveMatch radiant_team_logo_url
         * @property {number|null} [radiant_net_worth] LiveMatch radiant_net_worth
         * @property {number|null} [dire_score] LiveMatch dire_score
         * @property {Long|null} [dire_team_id] LiveMatch dire_team_id
         * @property {string|null} [dire_team_name] LiveMatch dire_team_name
         * @property {string|null} [dire_team_tag] LiveMatch dire_team_tag
         * @property {Long|null} [dire_team_logo] LiveMatch dire_team_logo
         * @property {string|null} [dire_team_logo_url] LiveMatch dire_team_logo_url
         * @property {number|null} [dire_net_worth] LiveMatch dire_net_worth
         * @property {number|null} [building_state] LiveMatch building_state
         * @property {number|null} [weekend_tourney_tournament_id] LiveMatch weekend_tourney_tournament_id
         * @property {number|null} [weekend_tourney_division] LiveMatch weekend_tourney_division
         * @property {number|null} [weekend_tourney_skill_level] LiveMatch weekend_tourney_skill_level
         * @property {number|null} [weekend_tourney_bracket_round] LiveMatch weekend_tourney_bracket_round
         * @property {google.protobuf.ITimestamp|null} [activate_time] LiveMatch activate_time
         * @property {google.protobuf.ITimestamp|null} [deactivate_time] LiveMatch deactivate_time
         * @property {google.protobuf.ITimestamp|null} [last_update_time] LiveMatch last_update_time
         * @property {Array.<protocol.LiveMatch.IPlayer>|null} [players] LiveMatch players
         */

        /**
         * Constructs a new LiveMatch.
         * @memberof protocol
         * @classdesc Represents a LiveMatch.
         * @implements ILiveMatch
         * @constructor
         * @param {protocol.ILiveMatch=} [properties] Properties to set
         */
        function LiveMatch(properties) {
            this.players = [];
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * LiveMatch match_id.
         * @member {Long} match_id
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.match_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * LiveMatch server_steam_id.
         * @member {Long} server_steam_id
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.server_steam_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * LiveMatch lobby_id.
         * @member {Long} lobby_id
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.lobby_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * LiveMatch lobby_type.
         * @member {protocol.LobbyType} lobby_type
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.lobby_type = 0;

        /**
         * LiveMatch league_id.
         * @member {Long} league_id
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.league_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * LiveMatch series_id.
         * @member {Long} series_id
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.series_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * LiveMatch game_mode.
         * @member {protocol.GameMode} game_mode
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.game_mode = 0;

        /**
         * LiveMatch game_state.
         * @member {protocol.GameState} game_state
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.game_state = 0;

        /**
         * LiveMatch game_timestamp.
         * @member {number} game_timestamp
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.game_timestamp = 0;

        /**
         * LiveMatch game_time.
         * @member {number} game_time
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.game_time = 0;

        /**
         * LiveMatch average_mmr.
         * @member {number} average_mmr
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.average_mmr = 0;

        /**
         * LiveMatch delay.
         * @member {number} delay
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.delay = 0;

        /**
         * LiveMatch spectators.
         * @member {number} spectators
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.spectators = 0;

        /**
         * LiveMatch sort_score.
         * @member {number} sort_score
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.sort_score = 0;

        /**
         * LiveMatch radiant_lead.
         * @member {number} radiant_lead
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.radiant_lead = 0;

        /**
         * LiveMatch radiant_score.
         * @member {number} radiant_score
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.radiant_score = 0;

        /**
         * LiveMatch radiant_team_id.
         * @member {Long} radiant_team_id
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.radiant_team_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * LiveMatch radiant_team_name.
         * @member {string} radiant_team_name
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.radiant_team_name = "";

        /**
         * LiveMatch radiant_team_tag.
         * @member {string} radiant_team_tag
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.radiant_team_tag = "";

        /**
         * LiveMatch radiant_team_logo.
         * @member {Long} radiant_team_logo
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.radiant_team_logo = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * LiveMatch radiant_team_logo_url.
         * @member {string} radiant_team_logo_url
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.radiant_team_logo_url = "";

        /**
         * LiveMatch radiant_net_worth.
         * @member {number} radiant_net_worth
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.radiant_net_worth = 0;

        /**
         * LiveMatch dire_score.
         * @member {number} dire_score
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.dire_score = 0;

        /**
         * LiveMatch dire_team_id.
         * @member {Long} dire_team_id
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.dire_team_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * LiveMatch dire_team_name.
         * @member {string} dire_team_name
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.dire_team_name = "";

        /**
         * LiveMatch dire_team_tag.
         * @member {string} dire_team_tag
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.dire_team_tag = "";

        /**
         * LiveMatch dire_team_logo.
         * @member {Long} dire_team_logo
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.dire_team_logo = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * LiveMatch dire_team_logo_url.
         * @member {string} dire_team_logo_url
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.dire_team_logo_url = "";

        /**
         * LiveMatch dire_net_worth.
         * @member {number} dire_net_worth
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.dire_net_worth = 0;

        /**
         * LiveMatch building_state.
         * @member {number} building_state
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.building_state = 0;

        /**
         * LiveMatch weekend_tourney_tournament_id.
         * @member {number} weekend_tourney_tournament_id
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.weekend_tourney_tournament_id = 0;

        /**
         * LiveMatch weekend_tourney_division.
         * @member {number} weekend_tourney_division
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.weekend_tourney_division = 0;

        /**
         * LiveMatch weekend_tourney_skill_level.
         * @member {number} weekend_tourney_skill_level
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.weekend_tourney_skill_level = 0;

        /**
         * LiveMatch weekend_tourney_bracket_round.
         * @member {number} weekend_tourney_bracket_round
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.weekend_tourney_bracket_round = 0;

        /**
         * LiveMatch activate_time.
         * @member {google.protobuf.ITimestamp|null|undefined} activate_time
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.activate_time = null;

        /**
         * LiveMatch deactivate_time.
         * @member {google.protobuf.ITimestamp|null|undefined} deactivate_time
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.deactivate_time = null;

        /**
         * LiveMatch last_update_time.
         * @member {google.protobuf.ITimestamp|null|undefined} last_update_time
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.last_update_time = null;

        /**
         * LiveMatch players.
         * @member {Array.<protocol.LiveMatch.IPlayer>} players
         * @memberof protocol.LiveMatch
         * @instance
         */
        LiveMatch.prototype.players = $util.emptyArray;

        /**
         * Creates a new LiveMatch instance using the specified properties.
         * @function create
         * @memberof protocol.LiveMatch
         * @static
         * @param {protocol.ILiveMatch=} [properties] Properties to set
         * @returns {protocol.LiveMatch} LiveMatch instance
         */
        LiveMatch.create = function create(properties) {
            return new LiveMatch(properties);
        };

        /**
         * Encodes the specified LiveMatch message. Does not implicitly {@link protocol.LiveMatch.verify|verify} messages.
         * @function encode
         * @memberof protocol.LiveMatch
         * @static
         * @param {protocol.ILiveMatch} message LiveMatch message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        LiveMatch.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.match_id != null && message.hasOwnProperty("match_id"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.match_id);
            if (message.server_steam_id != null && message.hasOwnProperty("server_steam_id"))
                writer.uint32(/* id 2, wireType 0 =*/16).uint64(message.server_steam_id);
            if (message.lobby_id != null && message.hasOwnProperty("lobby_id"))
                writer.uint32(/* id 3, wireType 0 =*/24).uint64(message.lobby_id);
            if (message.lobby_type != null && message.hasOwnProperty("lobby_type"))
                writer.uint32(/* id 4, wireType 0 =*/32).int32(message.lobby_type);
            if (message.league_id != null && message.hasOwnProperty("league_id"))
                writer.uint32(/* id 5, wireType 0 =*/40).uint64(message.league_id);
            if (message.series_id != null && message.hasOwnProperty("series_id"))
                writer.uint32(/* id 6, wireType 0 =*/48).uint64(message.series_id);
            if (message.game_mode != null && message.hasOwnProperty("game_mode"))
                writer.uint32(/* id 7, wireType 0 =*/56).int32(message.game_mode);
            if (message.game_state != null && message.hasOwnProperty("game_state"))
                writer.uint32(/* id 8, wireType 0 =*/64).int32(message.game_state);
            if (message.game_timestamp != null && message.hasOwnProperty("game_timestamp"))
                writer.uint32(/* id 9, wireType 0 =*/72).uint32(message.game_timestamp);
            if (message.game_time != null && message.hasOwnProperty("game_time"))
                writer.uint32(/* id 10, wireType 0 =*/80).int32(message.game_time);
            if (message.average_mmr != null && message.hasOwnProperty("average_mmr"))
                writer.uint32(/* id 11, wireType 0 =*/88).uint32(message.average_mmr);
            if (message.delay != null && message.hasOwnProperty("delay"))
                writer.uint32(/* id 12, wireType 0 =*/96).uint32(message.delay);
            if (message.spectators != null && message.hasOwnProperty("spectators"))
                writer.uint32(/* id 13, wireType 0 =*/104).uint32(message.spectators);
            if (message.sort_score != null && message.hasOwnProperty("sort_score"))
                writer.uint32(/* id 14, wireType 1 =*/113).double(message.sort_score);
            if (message.radiant_lead != null && message.hasOwnProperty("radiant_lead"))
                writer.uint32(/* id 15, wireType 0 =*/120).int32(message.radiant_lead);
            if (message.radiant_score != null && message.hasOwnProperty("radiant_score"))
                writer.uint32(/* id 16, wireType 0 =*/128).uint32(message.radiant_score);
            if (message.radiant_team_id != null && message.hasOwnProperty("radiant_team_id"))
                writer.uint32(/* id 17, wireType 0 =*/136).uint64(message.radiant_team_id);
            if (message.radiant_team_name != null && message.hasOwnProperty("radiant_team_name"))
                writer.uint32(/* id 18, wireType 2 =*/146).string(message.radiant_team_name);
            if (message.radiant_team_tag != null && message.hasOwnProperty("radiant_team_tag"))
                writer.uint32(/* id 19, wireType 2 =*/154).string(message.radiant_team_tag);
            if (message.radiant_team_logo != null && message.hasOwnProperty("radiant_team_logo"))
                writer.uint32(/* id 20, wireType 0 =*/160).uint64(message.radiant_team_logo);
            if (message.radiant_team_logo_url != null && message.hasOwnProperty("radiant_team_logo_url"))
                writer.uint32(/* id 21, wireType 2 =*/170).string(message.radiant_team_logo_url);
            if (message.radiant_net_worth != null && message.hasOwnProperty("radiant_net_worth"))
                writer.uint32(/* id 22, wireType 0 =*/176).uint32(message.radiant_net_worth);
            if (message.dire_score != null && message.hasOwnProperty("dire_score"))
                writer.uint32(/* id 23, wireType 0 =*/184).uint32(message.dire_score);
            if (message.dire_team_id != null && message.hasOwnProperty("dire_team_id"))
                writer.uint32(/* id 24, wireType 0 =*/192).uint64(message.dire_team_id);
            if (message.dire_team_name != null && message.hasOwnProperty("dire_team_name"))
                writer.uint32(/* id 25, wireType 2 =*/202).string(message.dire_team_name);
            if (message.dire_team_tag != null && message.hasOwnProperty("dire_team_tag"))
                writer.uint32(/* id 26, wireType 2 =*/210).string(message.dire_team_tag);
            if (message.dire_team_logo != null && message.hasOwnProperty("dire_team_logo"))
                writer.uint32(/* id 27, wireType 0 =*/216).uint64(message.dire_team_logo);
            if (message.dire_team_logo_url != null && message.hasOwnProperty("dire_team_logo_url"))
                writer.uint32(/* id 28, wireType 2 =*/226).string(message.dire_team_logo_url);
            if (message.dire_net_worth != null && message.hasOwnProperty("dire_net_worth"))
                writer.uint32(/* id 29, wireType 0 =*/232).uint32(message.dire_net_worth);
            if (message.building_state != null && message.hasOwnProperty("building_state"))
                writer.uint32(/* id 30, wireType 0 =*/240).uint32(message.building_state);
            if (message.weekend_tourney_tournament_id != null && message.hasOwnProperty("weekend_tourney_tournament_id"))
                writer.uint32(/* id 31, wireType 0 =*/248).uint32(message.weekend_tourney_tournament_id);
            if (message.weekend_tourney_division != null && message.hasOwnProperty("weekend_tourney_division"))
                writer.uint32(/* id 32, wireType 0 =*/256).uint32(message.weekend_tourney_division);
            if (message.weekend_tourney_skill_level != null && message.hasOwnProperty("weekend_tourney_skill_level"))
                writer.uint32(/* id 33, wireType 0 =*/264).uint32(message.weekend_tourney_skill_level);
            if (message.weekend_tourney_bracket_round != null && message.hasOwnProperty("weekend_tourney_bracket_round"))
                writer.uint32(/* id 34, wireType 0 =*/272).uint32(message.weekend_tourney_bracket_round);
            if (message.activate_time != null && message.hasOwnProperty("activate_time"))
                $root.google.protobuf.Timestamp.encode(message.activate_time, writer.uint32(/* id 35, wireType 2 =*/282).fork()).ldelim();
            if (message.deactivate_time != null && message.hasOwnProperty("deactivate_time"))
                $root.google.protobuf.Timestamp.encode(message.deactivate_time, writer.uint32(/* id 36, wireType 2 =*/290).fork()).ldelim();
            if (message.last_update_time != null && message.hasOwnProperty("last_update_time"))
                $root.google.protobuf.Timestamp.encode(message.last_update_time, writer.uint32(/* id 37, wireType 2 =*/298).fork()).ldelim();
            if (message.players != null && message.players.length)
                for (let i = 0; i < message.players.length; ++i)
                    $root.protocol.LiveMatch.Player.encode(message.players[i], writer.uint32(/* id 100, wireType 2 =*/802).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified LiveMatch message, length delimited. Does not implicitly {@link protocol.LiveMatch.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protocol.LiveMatch
         * @static
         * @param {protocol.ILiveMatch} message LiveMatch message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        LiveMatch.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a LiveMatch message from the specified reader or buffer.
         * @function decode
         * @memberof protocol.LiveMatch
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protocol.LiveMatch} LiveMatch
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        LiveMatch.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.protocol.LiveMatch();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.match_id = reader.uint64();
                    break;
                case 2:
                    message.server_steam_id = reader.uint64();
                    break;
                case 3:
                    message.lobby_id = reader.uint64();
                    break;
                case 4:
                    message.lobby_type = reader.int32();
                    break;
                case 5:
                    message.league_id = reader.uint64();
                    break;
                case 6:
                    message.series_id = reader.uint64();
                    break;
                case 7:
                    message.game_mode = reader.int32();
                    break;
                case 8:
                    message.game_state = reader.int32();
                    break;
                case 9:
                    message.game_timestamp = reader.uint32();
                    break;
                case 10:
                    message.game_time = reader.int32();
                    break;
                case 11:
                    message.average_mmr = reader.uint32();
                    break;
                case 12:
                    message.delay = reader.uint32();
                    break;
                case 13:
                    message.spectators = reader.uint32();
                    break;
                case 14:
                    message.sort_score = reader.double();
                    break;
                case 15:
                    message.radiant_lead = reader.int32();
                    break;
                case 16:
                    message.radiant_score = reader.uint32();
                    break;
                case 17:
                    message.radiant_team_id = reader.uint64();
                    break;
                case 18:
                    message.radiant_team_name = reader.string();
                    break;
                case 19:
                    message.radiant_team_tag = reader.string();
                    break;
                case 20:
                    message.radiant_team_logo = reader.uint64();
                    break;
                case 21:
                    message.radiant_team_logo_url = reader.string();
                    break;
                case 22:
                    message.radiant_net_worth = reader.uint32();
                    break;
                case 23:
                    message.dire_score = reader.uint32();
                    break;
                case 24:
                    message.dire_team_id = reader.uint64();
                    break;
                case 25:
                    message.dire_team_name = reader.string();
                    break;
                case 26:
                    message.dire_team_tag = reader.string();
                    break;
                case 27:
                    message.dire_team_logo = reader.uint64();
                    break;
                case 28:
                    message.dire_team_logo_url = reader.string();
                    break;
                case 29:
                    message.dire_net_worth = reader.uint32();
                    break;
                case 30:
                    message.building_state = reader.uint32();
                    break;
                case 31:
                    message.weekend_tourney_tournament_id = reader.uint32();
                    break;
                case 32:
                    message.weekend_tourney_division = reader.uint32();
                    break;
                case 33:
                    message.weekend_tourney_skill_level = reader.uint32();
                    break;
                case 34:
                    message.weekend_tourney_bracket_round = reader.uint32();
                    break;
                case 35:
                    message.activate_time = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                    break;
                case 36:
                    message.deactivate_time = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                    break;
                case 37:
                    message.last_update_time = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                    break;
                case 100:
                    if (!(message.players && message.players.length))
                        message.players = [];
                    message.players.push($root.protocol.LiveMatch.Player.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a LiveMatch message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protocol.LiveMatch
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protocol.LiveMatch} LiveMatch
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        LiveMatch.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a LiveMatch message.
         * @function verify
         * @memberof protocol.LiveMatch
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        LiveMatch.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.match_id != null && message.hasOwnProperty("match_id"))
                if (!$util.isInteger(message.match_id) && !(message.match_id && $util.isInteger(message.match_id.low) && $util.isInteger(message.match_id.high)))
                    return "match_id: integer|Long expected";
            if (message.server_steam_id != null && message.hasOwnProperty("server_steam_id"))
                if (!$util.isInteger(message.server_steam_id) && !(message.server_steam_id && $util.isInteger(message.server_steam_id.low) && $util.isInteger(message.server_steam_id.high)))
                    return "server_steam_id: integer|Long expected";
            if (message.lobby_id != null && message.hasOwnProperty("lobby_id"))
                if (!$util.isInteger(message.lobby_id) && !(message.lobby_id && $util.isInteger(message.lobby_id.low) && $util.isInteger(message.lobby_id.high)))
                    return "lobby_id: integer|Long expected";
            if (message.lobby_type != null && message.hasOwnProperty("lobby_type"))
                switch (message.lobby_type) {
                default:
                    return "lobby_type: enum value expected";
                case 0:
                case 1:
                case 4:
                case 5:
                case 6:
                case 7:
                case 8:
                case 9:
                case 10:
                case 11:
                case 12:
                    break;
                }
            if (message.league_id != null && message.hasOwnProperty("league_id"))
                if (!$util.isInteger(message.league_id) && !(message.league_id && $util.isInteger(message.league_id.low) && $util.isInteger(message.league_id.high)))
                    return "league_id: integer|Long expected";
            if (message.series_id != null && message.hasOwnProperty("series_id"))
                if (!$util.isInteger(message.series_id) && !(message.series_id && $util.isInteger(message.series_id.low) && $util.isInteger(message.series_id.high)))
                    return "series_id: integer|Long expected";
            if (message.game_mode != null && message.hasOwnProperty("game_mode"))
                switch (message.game_mode) {
                default:
                    return "game_mode: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                case 4:
                case 5:
                case 6:
                case 7:
                case 8:
                case 9:
                case 10:
                case 11:
                case 12:
                case 13:
                case 14:
                case 15:
                case 16:
                case 17:
                case 18:
                case 19:
                case 20:
                case 21:
                case 22:
                case 23:
                case 24:
                case 25:
                    break;
                }
            if (message.game_state != null && message.hasOwnProperty("game_state"))
                switch (message.game_state) {
                default:
                    return "game_state: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                case 4:
                case 5:
                case 6:
                case 7:
                case 8:
                case 9:
                case 10:
                case 11:
                    break;
                }
            if (message.game_timestamp != null && message.hasOwnProperty("game_timestamp"))
                if (!$util.isInteger(message.game_timestamp))
                    return "game_timestamp: integer expected";
            if (message.game_time != null && message.hasOwnProperty("game_time"))
                if (!$util.isInteger(message.game_time))
                    return "game_time: integer expected";
            if (message.average_mmr != null && message.hasOwnProperty("average_mmr"))
                if (!$util.isInteger(message.average_mmr))
                    return "average_mmr: integer expected";
            if (message.delay != null && message.hasOwnProperty("delay"))
                if (!$util.isInteger(message.delay))
                    return "delay: integer expected";
            if (message.spectators != null && message.hasOwnProperty("spectators"))
                if (!$util.isInteger(message.spectators))
                    return "spectators: integer expected";
            if (message.sort_score != null && message.hasOwnProperty("sort_score"))
                if (typeof message.sort_score !== "number")
                    return "sort_score: number expected";
            if (message.radiant_lead != null && message.hasOwnProperty("radiant_lead"))
                if (!$util.isInteger(message.radiant_lead))
                    return "radiant_lead: integer expected";
            if (message.radiant_score != null && message.hasOwnProperty("radiant_score"))
                if (!$util.isInteger(message.radiant_score))
                    return "radiant_score: integer expected";
            if (message.radiant_team_id != null && message.hasOwnProperty("radiant_team_id"))
                if (!$util.isInteger(message.radiant_team_id) && !(message.radiant_team_id && $util.isInteger(message.radiant_team_id.low) && $util.isInteger(message.radiant_team_id.high)))
                    return "radiant_team_id: integer|Long expected";
            if (message.radiant_team_name != null && message.hasOwnProperty("radiant_team_name"))
                if (!$util.isString(message.radiant_team_name))
                    return "radiant_team_name: string expected";
            if (message.radiant_team_tag != null && message.hasOwnProperty("radiant_team_tag"))
                if (!$util.isString(message.radiant_team_tag))
                    return "radiant_team_tag: string expected";
            if (message.radiant_team_logo != null && message.hasOwnProperty("radiant_team_logo"))
                if (!$util.isInteger(message.radiant_team_logo) && !(message.radiant_team_logo && $util.isInteger(message.radiant_team_logo.low) && $util.isInteger(message.radiant_team_logo.high)))
                    return "radiant_team_logo: integer|Long expected";
            if (message.radiant_team_logo_url != null && message.hasOwnProperty("radiant_team_logo_url"))
                if (!$util.isString(message.radiant_team_logo_url))
                    return "radiant_team_logo_url: string expected";
            if (message.radiant_net_worth != null && message.hasOwnProperty("radiant_net_worth"))
                if (!$util.isInteger(message.radiant_net_worth))
                    return "radiant_net_worth: integer expected";
            if (message.dire_score != null && message.hasOwnProperty("dire_score"))
                if (!$util.isInteger(message.dire_score))
                    return "dire_score: integer expected";
            if (message.dire_team_id != null && message.hasOwnProperty("dire_team_id"))
                if (!$util.isInteger(message.dire_team_id) && !(message.dire_team_id && $util.isInteger(message.dire_team_id.low) && $util.isInteger(message.dire_team_id.high)))
                    return "dire_team_id: integer|Long expected";
            if (message.dire_team_name != null && message.hasOwnProperty("dire_team_name"))
                if (!$util.isString(message.dire_team_name))
                    return "dire_team_name: string expected";
            if (message.dire_team_tag != null && message.hasOwnProperty("dire_team_tag"))
                if (!$util.isString(message.dire_team_tag))
                    return "dire_team_tag: string expected";
            if (message.dire_team_logo != null && message.hasOwnProperty("dire_team_logo"))
                if (!$util.isInteger(message.dire_team_logo) && !(message.dire_team_logo && $util.isInteger(message.dire_team_logo.low) && $util.isInteger(message.dire_team_logo.high)))
                    return "dire_team_logo: integer|Long expected";
            if (message.dire_team_logo_url != null && message.hasOwnProperty("dire_team_logo_url"))
                if (!$util.isString(message.dire_team_logo_url))
                    return "dire_team_logo_url: string expected";
            if (message.dire_net_worth != null && message.hasOwnProperty("dire_net_worth"))
                if (!$util.isInteger(message.dire_net_worth))
                    return "dire_net_worth: integer expected";
            if (message.building_state != null && message.hasOwnProperty("building_state"))
                if (!$util.isInteger(message.building_state))
                    return "building_state: integer expected";
            if (message.weekend_tourney_tournament_id != null && message.hasOwnProperty("weekend_tourney_tournament_id"))
                if (!$util.isInteger(message.weekend_tourney_tournament_id))
                    return "weekend_tourney_tournament_id: integer expected";
            if (message.weekend_tourney_division != null && message.hasOwnProperty("weekend_tourney_division"))
                if (!$util.isInteger(message.weekend_tourney_division))
                    return "weekend_tourney_division: integer expected";
            if (message.weekend_tourney_skill_level != null && message.hasOwnProperty("weekend_tourney_skill_level"))
                if (!$util.isInteger(message.weekend_tourney_skill_level))
                    return "weekend_tourney_skill_level: integer expected";
            if (message.weekend_tourney_bracket_round != null && message.hasOwnProperty("weekend_tourney_bracket_round"))
                if (!$util.isInteger(message.weekend_tourney_bracket_round))
                    return "weekend_tourney_bracket_round: integer expected";
            if (message.activate_time != null && message.hasOwnProperty("activate_time")) {
                let error = $root.google.protobuf.Timestamp.verify(message.activate_time);
                if (error)
                    return "activate_time." + error;
            }
            if (message.deactivate_time != null && message.hasOwnProperty("deactivate_time")) {
                let error = $root.google.protobuf.Timestamp.verify(message.deactivate_time);
                if (error)
                    return "deactivate_time." + error;
            }
            if (message.last_update_time != null && message.hasOwnProperty("last_update_time")) {
                let error = $root.google.protobuf.Timestamp.verify(message.last_update_time);
                if (error)
                    return "last_update_time." + error;
            }
            if (message.players != null && message.hasOwnProperty("players")) {
                if (!Array.isArray(message.players))
                    return "players: array expected";
                for (let i = 0; i < message.players.length; ++i) {
                    let error = $root.protocol.LiveMatch.Player.verify(message.players[i]);
                    if (error)
                        return "players." + error;
                }
            }
            return null;
        };

        /**
         * Creates a LiveMatch message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protocol.LiveMatch
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protocol.LiveMatch} LiveMatch
         */
        LiveMatch.fromObject = function fromObject(object) {
            if (object instanceof $root.protocol.LiveMatch)
                return object;
            let message = new $root.protocol.LiveMatch();
            if (object.match_id != null)
                if ($util.Long)
                    (message.match_id = $util.Long.fromValue(object.match_id)).unsigned = true;
                else if (typeof object.match_id === "string")
                    message.match_id = parseInt(object.match_id, 10);
                else if (typeof object.match_id === "number")
                    message.match_id = object.match_id;
                else if (typeof object.match_id === "object")
                    message.match_id = new $util.LongBits(object.match_id.low >>> 0, object.match_id.high >>> 0).toNumber(true);
            if (object.server_steam_id != null)
                if ($util.Long)
                    (message.server_steam_id = $util.Long.fromValue(object.server_steam_id)).unsigned = true;
                else if (typeof object.server_steam_id === "string")
                    message.server_steam_id = parseInt(object.server_steam_id, 10);
                else if (typeof object.server_steam_id === "number")
                    message.server_steam_id = object.server_steam_id;
                else if (typeof object.server_steam_id === "object")
                    message.server_steam_id = new $util.LongBits(object.server_steam_id.low >>> 0, object.server_steam_id.high >>> 0).toNumber(true);
            if (object.lobby_id != null)
                if ($util.Long)
                    (message.lobby_id = $util.Long.fromValue(object.lobby_id)).unsigned = true;
                else if (typeof object.lobby_id === "string")
                    message.lobby_id = parseInt(object.lobby_id, 10);
                else if (typeof object.lobby_id === "number")
                    message.lobby_id = object.lobby_id;
                else if (typeof object.lobby_id === "object")
                    message.lobby_id = new $util.LongBits(object.lobby_id.low >>> 0, object.lobby_id.high >>> 0).toNumber(true);
            switch (object.lobby_type) {
            case "LOBBY_TYPE_CASUAL_MATCH":
            case 0:
                message.lobby_type = 0;
                break;
            case "LOBBY_TYPE_PRACTICE":
            case 1:
                message.lobby_type = 1;
                break;
            case "LOBBY_TYPE_COOP_BOT_MATCH":
            case 4:
                message.lobby_type = 4;
                break;
            case "LOBBY_TYPE_LEGACY_TEAM_MATCH":
            case 5:
                message.lobby_type = 5;
                break;
            case "LOBBY_TYPE_LEGACY_SOLO_QUEUE_MATCH":
            case 6:
                message.lobby_type = 6;
                break;
            case "LOBBY_TYPE_COMPETITIVE_MATCH":
            case 7:
                message.lobby_type = 7;
                break;
            case "LOBBY_TYPE_CASUAL_1V1_MATCH":
            case 8:
                message.lobby_type = 8;
                break;
            case "LOBBY_TYPE_WEEKEND_TOURNEY":
            case 9:
                message.lobby_type = 9;
                break;
            case "LOBBY_TYPE_LOCAL_BOT_MATCH":
            case 10:
                message.lobby_type = 10;
                break;
            case "LOBBY_TYPE_SPECTATOR":
            case 11:
                message.lobby_type = 11;
                break;
            case "LOBBY_TYPE_EVENT_MATCH":
            case 12:
                message.lobby_type = 12;
                break;
            }
            if (object.league_id != null)
                if ($util.Long)
                    (message.league_id = $util.Long.fromValue(object.league_id)).unsigned = true;
                else if (typeof object.league_id === "string")
                    message.league_id = parseInt(object.league_id, 10);
                else if (typeof object.league_id === "number")
                    message.league_id = object.league_id;
                else if (typeof object.league_id === "object")
                    message.league_id = new $util.LongBits(object.league_id.low >>> 0, object.league_id.high >>> 0).toNumber(true);
            if (object.series_id != null)
                if ($util.Long)
                    (message.series_id = $util.Long.fromValue(object.series_id)).unsigned = true;
                else if (typeof object.series_id === "string")
                    message.series_id = parseInt(object.series_id, 10);
                else if (typeof object.series_id === "number")
                    message.series_id = object.series_id;
                else if (typeof object.series_id === "object")
                    message.series_id = new $util.LongBits(object.series_id.low >>> 0, object.series_id.high >>> 0).toNumber(true);
            switch (object.game_mode) {
            case "GAME_MODE_NONE":
            case 0:
                message.game_mode = 0;
                break;
            case "GAME_MODE_AP":
            case 1:
                message.game_mode = 1;
                break;
            case "GAME_MODE_CM":
            case 2:
                message.game_mode = 2;
                break;
            case "GAME_MODE_RD":
            case 3:
                message.game_mode = 3;
                break;
            case "GAME_MODE_SD":
            case 4:
                message.game_mode = 4;
                break;
            case "GAME_MODE_AR":
            case 5:
                message.game_mode = 5;
                break;
            case "GAME_MODE_INTRO":
            case 6:
                message.game_mode = 6;
                break;
            case "GAME_MODE_HW":
            case 7:
                message.game_mode = 7;
                break;
            case "GAME_MODE_REVERSE_CM":
            case 8:
                message.game_mode = 8;
                break;
            case "GAME_MODE_XMAS":
            case 9:
                message.game_mode = 9;
                break;
            case "GAME_MODE_TUTORIAL":
            case 10:
                message.game_mode = 10;
                break;
            case "GAME_MODE_MO":
            case 11:
                message.game_mode = 11;
                break;
            case "GAME_MODE_LP":
            case 12:
                message.game_mode = 12;
                break;
            case "GAME_MODE_POOL1":
            case 13:
                message.game_mode = 13;
                break;
            case "GAME_MODE_FH":
            case 14:
                message.game_mode = 14;
                break;
            case "GAME_MODE_CUSTOM":
            case 15:
                message.game_mode = 15;
                break;
            case "GAME_MODE_CD":
            case 16:
                message.game_mode = 16;
                break;
            case "GAME_MODE_BD":
            case 17:
                message.game_mode = 17;
                break;
            case "GAME_MODE_ABILITY_DRAFT":
            case 18:
                message.game_mode = 18;
                break;
            case "GAME_MODE_EVENT":
            case 19:
                message.game_mode = 19;
                break;
            case "GAME_MODE_ARDM":
            case 20:
                message.game_mode = 20;
                break;
            case "GAME_MODE_1V1_MID":
            case 21:
                message.game_mode = 21;
                break;
            case "GAME_MODE_ALL_DRAFT":
            case 22:
                message.game_mode = 22;
                break;
            case "GAME_MODE_TURBO":
            case 23:
                message.game_mode = 23;
                break;
            case "GAME_MODE_MUTATION":
            case 24:
                message.game_mode = 24;
                break;
            case "GAME_MODE_COACHES_CHALLENGE":
            case 25:
                message.game_mode = 25;
                break;
            }
            switch (object.game_state) {
            case "GAME_STATE_INIT":
            case 0:
                message.game_state = 0;
                break;
            case "GAME_STATE_WAIT_FOR_PLAYERS_TO_LOAD":
            case 1:
                message.game_state = 1;
                break;
            case "GAME_STATE_HERO_SELECTION":
            case 2:
                message.game_state = 2;
                break;
            case "GAME_STATE_STRATEGY_TIME":
            case 3:
                message.game_state = 3;
                break;
            case "GAME_STATE_PRE_GAME":
            case 4:
                message.game_state = 4;
                break;
            case "GAME_STATE_GAME_IN_PROGRESS":
            case 5:
                message.game_state = 5;
                break;
            case "GAME_STATE_POST_GAME":
            case 6:
                message.game_state = 6;
                break;
            case "GAME_STATE_DISCONNECT":
            case 7:
                message.game_state = 7;
                break;
            case "GAME_STATE_TEAM_SHOWCASE":
            case 8:
                message.game_state = 8;
                break;
            case "GAME_STATE_CUSTOM_GAME_SETUP":
            case 9:
                message.game_state = 9;
                break;
            case "GAME_STATE_WAIT_FOR_MAP_TO_LOAD":
            case 10:
                message.game_state = 10;
                break;
            case "GAME_STATE_LAST":
            case 11:
                message.game_state = 11;
                break;
            }
            if (object.game_timestamp != null)
                message.game_timestamp = object.game_timestamp >>> 0;
            if (object.game_time != null)
                message.game_time = object.game_time | 0;
            if (object.average_mmr != null)
                message.average_mmr = object.average_mmr >>> 0;
            if (object.delay != null)
                message.delay = object.delay >>> 0;
            if (object.spectators != null)
                message.spectators = object.spectators >>> 0;
            if (object.sort_score != null)
                message.sort_score = Number(object.sort_score);
            if (object.radiant_lead != null)
                message.radiant_lead = object.radiant_lead | 0;
            if (object.radiant_score != null)
                message.radiant_score = object.radiant_score >>> 0;
            if (object.radiant_team_id != null)
                if ($util.Long)
                    (message.radiant_team_id = $util.Long.fromValue(object.radiant_team_id)).unsigned = true;
                else if (typeof object.radiant_team_id === "string")
                    message.radiant_team_id = parseInt(object.radiant_team_id, 10);
                else if (typeof object.radiant_team_id === "number")
                    message.radiant_team_id = object.radiant_team_id;
                else if (typeof object.radiant_team_id === "object")
                    message.radiant_team_id = new $util.LongBits(object.radiant_team_id.low >>> 0, object.radiant_team_id.high >>> 0).toNumber(true);
            if (object.radiant_team_name != null)
                message.radiant_team_name = String(object.radiant_team_name);
            if (object.radiant_team_tag != null)
                message.radiant_team_tag = String(object.radiant_team_tag);
            if (object.radiant_team_logo != null)
                if ($util.Long)
                    (message.radiant_team_logo = $util.Long.fromValue(object.radiant_team_logo)).unsigned = true;
                else if (typeof object.radiant_team_logo === "string")
                    message.radiant_team_logo = parseInt(object.radiant_team_logo, 10);
                else if (typeof object.radiant_team_logo === "number")
                    message.radiant_team_logo = object.radiant_team_logo;
                else if (typeof object.radiant_team_logo === "object")
                    message.radiant_team_logo = new $util.LongBits(object.radiant_team_logo.low >>> 0, object.radiant_team_logo.high >>> 0).toNumber(true);
            if (object.radiant_team_logo_url != null)
                message.radiant_team_logo_url = String(object.radiant_team_logo_url);
            if (object.radiant_net_worth != null)
                message.radiant_net_worth = object.radiant_net_worth >>> 0;
            if (object.dire_score != null)
                message.dire_score = object.dire_score >>> 0;
            if (object.dire_team_id != null)
                if ($util.Long)
                    (message.dire_team_id = $util.Long.fromValue(object.dire_team_id)).unsigned = true;
                else if (typeof object.dire_team_id === "string")
                    message.dire_team_id = parseInt(object.dire_team_id, 10);
                else if (typeof object.dire_team_id === "number")
                    message.dire_team_id = object.dire_team_id;
                else if (typeof object.dire_team_id === "object")
                    message.dire_team_id = new $util.LongBits(object.dire_team_id.low >>> 0, object.dire_team_id.high >>> 0).toNumber(true);
            if (object.dire_team_name != null)
                message.dire_team_name = String(object.dire_team_name);
            if (object.dire_team_tag != null)
                message.dire_team_tag = String(object.dire_team_tag);
            if (object.dire_team_logo != null)
                if ($util.Long)
                    (message.dire_team_logo = $util.Long.fromValue(object.dire_team_logo)).unsigned = true;
                else if (typeof object.dire_team_logo === "string")
                    message.dire_team_logo = parseInt(object.dire_team_logo, 10);
                else if (typeof object.dire_team_logo === "number")
                    message.dire_team_logo = object.dire_team_logo;
                else if (typeof object.dire_team_logo === "object")
                    message.dire_team_logo = new $util.LongBits(object.dire_team_logo.low >>> 0, object.dire_team_logo.high >>> 0).toNumber(true);
            if (object.dire_team_logo_url != null)
                message.dire_team_logo_url = String(object.dire_team_logo_url);
            if (object.dire_net_worth != null)
                message.dire_net_worth = object.dire_net_worth >>> 0;
            if (object.building_state != null)
                message.building_state = object.building_state >>> 0;
            if (object.weekend_tourney_tournament_id != null)
                message.weekend_tourney_tournament_id = object.weekend_tourney_tournament_id >>> 0;
            if (object.weekend_tourney_division != null)
                message.weekend_tourney_division = object.weekend_tourney_division >>> 0;
            if (object.weekend_tourney_skill_level != null)
                message.weekend_tourney_skill_level = object.weekend_tourney_skill_level >>> 0;
            if (object.weekend_tourney_bracket_round != null)
                message.weekend_tourney_bracket_round = object.weekend_tourney_bracket_round >>> 0;
            if (object.activate_time != null) {
                if (typeof object.activate_time !== "object")
                    throw TypeError(".protocol.LiveMatch.activate_time: object expected");
                message.activate_time = $root.google.protobuf.Timestamp.fromObject(object.activate_time);
            }
            if (object.deactivate_time != null) {
                if (typeof object.deactivate_time !== "object")
                    throw TypeError(".protocol.LiveMatch.deactivate_time: object expected");
                message.deactivate_time = $root.google.protobuf.Timestamp.fromObject(object.deactivate_time);
            }
            if (object.last_update_time != null) {
                if (typeof object.last_update_time !== "object")
                    throw TypeError(".protocol.LiveMatch.last_update_time: object expected");
                message.last_update_time = $root.google.protobuf.Timestamp.fromObject(object.last_update_time);
            }
            if (object.players) {
                if (!Array.isArray(object.players))
                    throw TypeError(".protocol.LiveMatch.players: array expected");
                message.players = [];
                for (let i = 0; i < object.players.length; ++i) {
                    if (typeof object.players[i] !== "object")
                        throw TypeError(".protocol.LiveMatch.players: object expected");
                    message.players[i] = $root.protocol.LiveMatch.Player.fromObject(object.players[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a LiveMatch message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protocol.LiveMatch
         * @static
         * @param {protocol.LiveMatch} message LiveMatch
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        LiveMatch.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.arrays || options.defaults)
                object.players = [];
            if (options.defaults) {
                if ($util.Long) {
                    let long = new $util.Long(0, 0, true);
                    object.match_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.match_id = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    let long = new $util.Long(0, 0, true);
                    object.server_steam_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.server_steam_id = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    let long = new $util.Long(0, 0, true);
                    object.lobby_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.lobby_id = options.longs === String ? "0" : 0;
                object.lobby_type = options.enums === String ? "LOBBY_TYPE_CASUAL_MATCH" : 0;
                if ($util.Long) {
                    let long = new $util.Long(0, 0, true);
                    object.league_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.league_id = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    let long = new $util.Long(0, 0, true);
                    object.series_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.series_id = options.longs === String ? "0" : 0;
                object.game_mode = options.enums === String ? "GAME_MODE_NONE" : 0;
                object.game_state = options.enums === String ? "GAME_STATE_INIT" : 0;
                object.game_timestamp = 0;
                object.game_time = 0;
                object.average_mmr = 0;
                object.delay = 0;
                object.spectators = 0;
                object.sort_score = 0;
                object.radiant_lead = 0;
                object.radiant_score = 0;
                if ($util.Long) {
                    let long = new $util.Long(0, 0, true);
                    object.radiant_team_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.radiant_team_id = options.longs === String ? "0" : 0;
                object.radiant_team_name = "";
                object.radiant_team_tag = "";
                if ($util.Long) {
                    let long = new $util.Long(0, 0, true);
                    object.radiant_team_logo = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.radiant_team_logo = options.longs === String ? "0" : 0;
                object.radiant_team_logo_url = "";
                object.radiant_net_worth = 0;
                object.dire_score = 0;
                if ($util.Long) {
                    let long = new $util.Long(0, 0, true);
                    object.dire_team_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.dire_team_id = options.longs === String ? "0" : 0;
                object.dire_team_name = "";
                object.dire_team_tag = "";
                if ($util.Long) {
                    let long = new $util.Long(0, 0, true);
                    object.dire_team_logo = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.dire_team_logo = options.longs === String ? "0" : 0;
                object.dire_team_logo_url = "";
                object.dire_net_worth = 0;
                object.building_state = 0;
                object.weekend_tourney_tournament_id = 0;
                object.weekend_tourney_division = 0;
                object.weekend_tourney_skill_level = 0;
                object.weekend_tourney_bracket_round = 0;
                object.activate_time = null;
                object.deactivate_time = null;
                object.last_update_time = null;
            }
            if (message.match_id != null && message.hasOwnProperty("match_id"))
                if (typeof message.match_id === "number")
                    object.match_id = options.longs === String ? String(message.match_id) : message.match_id;
                else
                    object.match_id = options.longs === String ? $util.Long.prototype.toString.call(message.match_id) : options.longs === Number ? new $util.LongBits(message.match_id.low >>> 0, message.match_id.high >>> 0).toNumber(true) : message.match_id;
            if (message.server_steam_id != null && message.hasOwnProperty("server_steam_id"))
                if (typeof message.server_steam_id === "number")
                    object.server_steam_id = options.longs === String ? String(message.server_steam_id) : message.server_steam_id;
                else
                    object.server_steam_id = options.longs === String ? $util.Long.prototype.toString.call(message.server_steam_id) : options.longs === Number ? new $util.LongBits(message.server_steam_id.low >>> 0, message.server_steam_id.high >>> 0).toNumber(true) : message.server_steam_id;
            if (message.lobby_id != null && message.hasOwnProperty("lobby_id"))
                if (typeof message.lobby_id === "number")
                    object.lobby_id = options.longs === String ? String(message.lobby_id) : message.lobby_id;
                else
                    object.lobby_id = options.longs === String ? $util.Long.prototype.toString.call(message.lobby_id) : options.longs === Number ? new $util.LongBits(message.lobby_id.low >>> 0, message.lobby_id.high >>> 0).toNumber(true) : message.lobby_id;
            if (message.lobby_type != null && message.hasOwnProperty("lobby_type"))
                object.lobby_type = options.enums === String ? $root.protocol.LobbyType[message.lobby_type] : message.lobby_type;
            if (message.league_id != null && message.hasOwnProperty("league_id"))
                if (typeof message.league_id === "number")
                    object.league_id = options.longs === String ? String(message.league_id) : message.league_id;
                else
                    object.league_id = options.longs === String ? $util.Long.prototype.toString.call(message.league_id) : options.longs === Number ? new $util.LongBits(message.league_id.low >>> 0, message.league_id.high >>> 0).toNumber(true) : message.league_id;
            if (message.series_id != null && message.hasOwnProperty("series_id"))
                if (typeof message.series_id === "number")
                    object.series_id = options.longs === String ? String(message.series_id) : message.series_id;
                else
                    object.series_id = options.longs === String ? $util.Long.prototype.toString.call(message.series_id) : options.longs === Number ? new $util.LongBits(message.series_id.low >>> 0, message.series_id.high >>> 0).toNumber(true) : message.series_id;
            if (message.game_mode != null && message.hasOwnProperty("game_mode"))
                object.game_mode = options.enums === String ? $root.protocol.GameMode[message.game_mode] : message.game_mode;
            if (message.game_state != null && message.hasOwnProperty("game_state"))
                object.game_state = options.enums === String ? $root.protocol.GameState[message.game_state] : message.game_state;
            if (message.game_timestamp != null && message.hasOwnProperty("game_timestamp"))
                object.game_timestamp = message.game_timestamp;
            if (message.game_time != null && message.hasOwnProperty("game_time"))
                object.game_time = message.game_time;
            if (message.average_mmr != null && message.hasOwnProperty("average_mmr"))
                object.average_mmr = message.average_mmr;
            if (message.delay != null && message.hasOwnProperty("delay"))
                object.delay = message.delay;
            if (message.spectators != null && message.hasOwnProperty("spectators"))
                object.spectators = message.spectators;
            if (message.sort_score != null && message.hasOwnProperty("sort_score"))
                object.sort_score = options.json && !isFinite(message.sort_score) ? String(message.sort_score) : message.sort_score;
            if (message.radiant_lead != null && message.hasOwnProperty("radiant_lead"))
                object.radiant_lead = message.radiant_lead;
            if (message.radiant_score != null && message.hasOwnProperty("radiant_score"))
                object.radiant_score = message.radiant_score;
            if (message.radiant_team_id != null && message.hasOwnProperty("radiant_team_id"))
                if (typeof message.radiant_team_id === "number")
                    object.radiant_team_id = options.longs === String ? String(message.radiant_team_id) : message.radiant_team_id;
                else
                    object.radiant_team_id = options.longs === String ? $util.Long.prototype.toString.call(message.radiant_team_id) : options.longs === Number ? new $util.LongBits(message.radiant_team_id.low >>> 0, message.radiant_team_id.high >>> 0).toNumber(true) : message.radiant_team_id;
            if (message.radiant_team_name != null && message.hasOwnProperty("radiant_team_name"))
                object.radiant_team_name = message.radiant_team_name;
            if (message.radiant_team_tag != null && message.hasOwnProperty("radiant_team_tag"))
                object.radiant_team_tag = message.radiant_team_tag;
            if (message.radiant_team_logo != null && message.hasOwnProperty("radiant_team_logo"))
                if (typeof message.radiant_team_logo === "number")
                    object.radiant_team_logo = options.longs === String ? String(message.radiant_team_logo) : message.radiant_team_logo;
                else
                    object.radiant_team_logo = options.longs === String ? $util.Long.prototype.toString.call(message.radiant_team_logo) : options.longs === Number ? new $util.LongBits(message.radiant_team_logo.low >>> 0, message.radiant_team_logo.high >>> 0).toNumber(true) : message.radiant_team_logo;
            if (message.radiant_team_logo_url != null && message.hasOwnProperty("radiant_team_logo_url"))
                object.radiant_team_logo_url = message.radiant_team_logo_url;
            if (message.radiant_net_worth != null && message.hasOwnProperty("radiant_net_worth"))
                object.radiant_net_worth = message.radiant_net_worth;
            if (message.dire_score != null && message.hasOwnProperty("dire_score"))
                object.dire_score = message.dire_score;
            if (message.dire_team_id != null && message.hasOwnProperty("dire_team_id"))
                if (typeof message.dire_team_id === "number")
                    object.dire_team_id = options.longs === String ? String(message.dire_team_id) : message.dire_team_id;
                else
                    object.dire_team_id = options.longs === String ? $util.Long.prototype.toString.call(message.dire_team_id) : options.longs === Number ? new $util.LongBits(message.dire_team_id.low >>> 0, message.dire_team_id.high >>> 0).toNumber(true) : message.dire_team_id;
            if (message.dire_team_name != null && message.hasOwnProperty("dire_team_name"))
                object.dire_team_name = message.dire_team_name;
            if (message.dire_team_tag != null && message.hasOwnProperty("dire_team_tag"))
                object.dire_team_tag = message.dire_team_tag;
            if (message.dire_team_logo != null && message.hasOwnProperty("dire_team_logo"))
                if (typeof message.dire_team_logo === "number")
                    object.dire_team_logo = options.longs === String ? String(message.dire_team_logo) : message.dire_team_logo;
                else
                    object.dire_team_logo = options.longs === String ? $util.Long.prototype.toString.call(message.dire_team_logo) : options.longs === Number ? new $util.LongBits(message.dire_team_logo.low >>> 0, message.dire_team_logo.high >>> 0).toNumber(true) : message.dire_team_logo;
            if (message.dire_team_logo_url != null && message.hasOwnProperty("dire_team_logo_url"))
                object.dire_team_logo_url = message.dire_team_logo_url;
            if (message.dire_net_worth != null && message.hasOwnProperty("dire_net_worth"))
                object.dire_net_worth = message.dire_net_worth;
            if (message.building_state != null && message.hasOwnProperty("building_state"))
                object.building_state = message.building_state;
            if (message.weekend_tourney_tournament_id != null && message.hasOwnProperty("weekend_tourney_tournament_id"))
                object.weekend_tourney_tournament_id = message.weekend_tourney_tournament_id;
            if (message.weekend_tourney_division != null && message.hasOwnProperty("weekend_tourney_division"))
                object.weekend_tourney_division = message.weekend_tourney_division;
            if (message.weekend_tourney_skill_level != null && message.hasOwnProperty("weekend_tourney_skill_level"))
                object.weekend_tourney_skill_level = message.weekend_tourney_skill_level;
            if (message.weekend_tourney_bracket_round != null && message.hasOwnProperty("weekend_tourney_bracket_round"))
                object.weekend_tourney_bracket_round = message.weekend_tourney_bracket_round;
            if (message.activate_time != null && message.hasOwnProperty("activate_time"))
                object.activate_time = $root.google.protobuf.Timestamp.toObject(message.activate_time, options);
            if (message.deactivate_time != null && message.hasOwnProperty("deactivate_time"))
                object.deactivate_time = $root.google.protobuf.Timestamp.toObject(message.deactivate_time, options);
            if (message.last_update_time != null && message.hasOwnProperty("last_update_time"))
                object.last_update_time = $root.google.protobuf.Timestamp.toObject(message.last_update_time, options);
            if (message.players && message.players.length) {
                object.players = [];
                for (let j = 0; j < message.players.length; ++j)
                    object.players[j] = $root.protocol.LiveMatch.Player.toObject(message.players[j], options);
            }
            return object;
        };

        /**
         * Converts this LiveMatch to JSON.
         * @function toJSON
         * @memberof protocol.LiveMatch
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        LiveMatch.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        LiveMatch.Player = (function() {

            /**
             * Properties of a Player.
             * @memberof protocol.LiveMatch
             * @interface IPlayer
             * @property {number|null} [account_id] Player account_id
             * @property {string|null} [name] Player name
             * @property {string|null} [persona_name] Player persona_name
             * @property {string|null} [avatar_url] Player avatar_url
             * @property {string|null} [avatar_medium_url] Player avatar_medium_url
             * @property {string|null} [avatar_full_url] Player avatar_full_url
             * @property {boolean|null} [is_pro] Player is_pro
             * @property {Long|null} [hero_id] Player hero_id
             * @property {number|null} [player_slot] Player player_slot
             * @property {protocol.GameTeam|null} [team] Player team
             * @property {number|null} [level] Player level
             * @property {number|null} [kills] Player kills
             * @property {number|null} [deaths] Player deaths
             * @property {number|null} [assists] Player assists
             * @property {number|null} [denies] Player denies
             * @property {number|null} [last_hits] Player last_hits
             * @property {number|null} [gold] Player gold
             * @property {number|null} [net_worth] Player net_worth
             * @property {string|null} [label] Player label
             */

            /**
             * Constructs a new Player.
             * @memberof protocol.LiveMatch
             * @classdesc Represents a Player.
             * @implements IPlayer
             * @constructor
             * @param {protocol.LiveMatch.IPlayer=} [properties] Properties to set
             */
            function Player(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Player account_id.
             * @member {number} account_id
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.account_id = 0;

            /**
             * Player name.
             * @member {string} name
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.name = "";

            /**
             * Player persona_name.
             * @member {string} persona_name
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.persona_name = "";

            /**
             * Player avatar_url.
             * @member {string} avatar_url
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.avatar_url = "";

            /**
             * Player avatar_medium_url.
             * @member {string} avatar_medium_url
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.avatar_medium_url = "";

            /**
             * Player avatar_full_url.
             * @member {string} avatar_full_url
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.avatar_full_url = "";

            /**
             * Player is_pro.
             * @member {boolean} is_pro
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.is_pro = false;

            /**
             * Player hero_id.
             * @member {Long} hero_id
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.hero_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

            /**
             * Player player_slot.
             * @member {number} player_slot
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.player_slot = 0;

            /**
             * Player team.
             * @member {protocol.GameTeam} team
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.team = 0;

            /**
             * Player level.
             * @member {number} level
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.level = 0;

            /**
             * Player kills.
             * @member {number} kills
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.kills = 0;

            /**
             * Player deaths.
             * @member {number} deaths
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.deaths = 0;

            /**
             * Player assists.
             * @member {number} assists
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.assists = 0;

            /**
             * Player denies.
             * @member {number} denies
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.denies = 0;

            /**
             * Player last_hits.
             * @member {number} last_hits
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.last_hits = 0;

            /**
             * Player gold.
             * @member {number} gold
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.gold = 0;

            /**
             * Player net_worth.
             * @member {number} net_worth
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.net_worth = 0;

            /**
             * Player label.
             * @member {string} label
             * @memberof protocol.LiveMatch.Player
             * @instance
             */
            Player.prototype.label = "";

            /**
             * Creates a new Player instance using the specified properties.
             * @function create
             * @memberof protocol.LiveMatch.Player
             * @static
             * @param {protocol.LiveMatch.IPlayer=} [properties] Properties to set
             * @returns {protocol.LiveMatch.Player} Player instance
             */
            Player.create = function create(properties) {
                return new Player(properties);
            };

            /**
             * Encodes the specified Player message. Does not implicitly {@link protocol.LiveMatch.Player.verify|verify} messages.
             * @function encode
             * @memberof protocol.LiveMatch.Player
             * @static
             * @param {protocol.LiveMatch.IPlayer} message Player message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Player.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.account_id != null && message.hasOwnProperty("account_id"))
                    writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.account_id);
                if (message.name != null && message.hasOwnProperty("name"))
                    writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
                if (message.persona_name != null && message.hasOwnProperty("persona_name"))
                    writer.uint32(/* id 3, wireType 2 =*/26).string(message.persona_name);
                if (message.avatar_url != null && message.hasOwnProperty("avatar_url"))
                    writer.uint32(/* id 4, wireType 2 =*/34).string(message.avatar_url);
                if (message.avatar_medium_url != null && message.hasOwnProperty("avatar_medium_url"))
                    writer.uint32(/* id 5, wireType 2 =*/42).string(message.avatar_medium_url);
                if (message.avatar_full_url != null && message.hasOwnProperty("avatar_full_url"))
                    writer.uint32(/* id 6, wireType 2 =*/50).string(message.avatar_full_url);
                if (message.is_pro != null && message.hasOwnProperty("is_pro"))
                    writer.uint32(/* id 7, wireType 0 =*/56).bool(message.is_pro);
                if (message.hero_id != null && message.hasOwnProperty("hero_id"))
                    writer.uint32(/* id 8, wireType 0 =*/64).uint64(message.hero_id);
                if (message.player_slot != null && message.hasOwnProperty("player_slot"))
                    writer.uint32(/* id 9, wireType 0 =*/72).uint32(message.player_slot);
                if (message.team != null && message.hasOwnProperty("team"))
                    writer.uint32(/* id 10, wireType 0 =*/80).int32(message.team);
                if (message.level != null && message.hasOwnProperty("level"))
                    writer.uint32(/* id 11, wireType 0 =*/88).uint32(message.level);
                if (message.kills != null && message.hasOwnProperty("kills"))
                    writer.uint32(/* id 12, wireType 0 =*/96).uint32(message.kills);
                if (message.deaths != null && message.hasOwnProperty("deaths"))
                    writer.uint32(/* id 13, wireType 0 =*/104).uint32(message.deaths);
                if (message.assists != null && message.hasOwnProperty("assists"))
                    writer.uint32(/* id 14, wireType 0 =*/112).uint32(message.assists);
                if (message.denies != null && message.hasOwnProperty("denies"))
                    writer.uint32(/* id 15, wireType 0 =*/120).uint32(message.denies);
                if (message.last_hits != null && message.hasOwnProperty("last_hits"))
                    writer.uint32(/* id 16, wireType 0 =*/128).uint32(message.last_hits);
                if (message.gold != null && message.hasOwnProperty("gold"))
                    writer.uint32(/* id 17, wireType 0 =*/136).uint32(message.gold);
                if (message.net_worth != null && message.hasOwnProperty("net_worth"))
                    writer.uint32(/* id 18, wireType 0 =*/144).uint32(message.net_worth);
                if (message.label != null && message.hasOwnProperty("label"))
                    writer.uint32(/* id 19, wireType 2 =*/154).string(message.label);
                return writer;
            };

            /**
             * Encodes the specified Player message, length delimited. Does not implicitly {@link protocol.LiveMatch.Player.verify|verify} messages.
             * @function encodeDelimited
             * @memberof protocol.LiveMatch.Player
             * @static
             * @param {protocol.LiveMatch.IPlayer} message Player message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Player.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a Player message from the specified reader or buffer.
             * @function decode
             * @memberof protocol.LiveMatch.Player
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {protocol.LiveMatch.Player} Player
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Player.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.protocol.LiveMatch.Player();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.account_id = reader.uint32();
                        break;
                    case 2:
                        message.name = reader.string();
                        break;
                    case 3:
                        message.persona_name = reader.string();
                        break;
                    case 4:
                        message.avatar_url = reader.string();
                        break;
                    case 5:
                        message.avatar_medium_url = reader.string();
                        break;
                    case 6:
                        message.avatar_full_url = reader.string();
                        break;
                    case 7:
                        message.is_pro = reader.bool();
                        break;
                    case 8:
                        message.hero_id = reader.uint64();
                        break;
                    case 9:
                        message.player_slot = reader.uint32();
                        break;
                    case 10:
                        message.team = reader.int32();
                        break;
                    case 11:
                        message.level = reader.uint32();
                        break;
                    case 12:
                        message.kills = reader.uint32();
                        break;
                    case 13:
                        message.deaths = reader.uint32();
                        break;
                    case 14:
                        message.assists = reader.uint32();
                        break;
                    case 15:
                        message.denies = reader.uint32();
                        break;
                    case 16:
                        message.last_hits = reader.uint32();
                        break;
                    case 17:
                        message.gold = reader.uint32();
                        break;
                    case 18:
                        message.net_worth = reader.uint32();
                        break;
                    case 19:
                        message.label = reader.string();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a Player message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof protocol.LiveMatch.Player
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {protocol.LiveMatch.Player} Player
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Player.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a Player message.
             * @function verify
             * @memberof protocol.LiveMatch.Player
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Player.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.account_id != null && message.hasOwnProperty("account_id"))
                    if (!$util.isInteger(message.account_id))
                        return "account_id: integer expected";
                if (message.name != null && message.hasOwnProperty("name"))
                    if (!$util.isString(message.name))
                        return "name: string expected";
                if (message.persona_name != null && message.hasOwnProperty("persona_name"))
                    if (!$util.isString(message.persona_name))
                        return "persona_name: string expected";
                if (message.avatar_url != null && message.hasOwnProperty("avatar_url"))
                    if (!$util.isString(message.avatar_url))
                        return "avatar_url: string expected";
                if (message.avatar_medium_url != null && message.hasOwnProperty("avatar_medium_url"))
                    if (!$util.isString(message.avatar_medium_url))
                        return "avatar_medium_url: string expected";
                if (message.avatar_full_url != null && message.hasOwnProperty("avatar_full_url"))
                    if (!$util.isString(message.avatar_full_url))
                        return "avatar_full_url: string expected";
                if (message.is_pro != null && message.hasOwnProperty("is_pro"))
                    if (typeof message.is_pro !== "boolean")
                        return "is_pro: boolean expected";
                if (message.hero_id != null && message.hasOwnProperty("hero_id"))
                    if (!$util.isInteger(message.hero_id) && !(message.hero_id && $util.isInteger(message.hero_id.low) && $util.isInteger(message.hero_id.high)))
                        return "hero_id: integer|Long expected";
                if (message.player_slot != null && message.hasOwnProperty("player_slot"))
                    if (!$util.isInteger(message.player_slot))
                        return "player_slot: integer expected";
                if (message.team != null && message.hasOwnProperty("team"))
                    switch (message.team) {
                    default:
                        return "team: enum value expected";
                    case 0:
                    case 2:
                    case 3:
                    case 4:
                    case 5:
                    case 6:
                    case 7:
                    case 8:
                    case 9:
                    case 10:
                    case 11:
                    case 12:
                    case 13:
                        break;
                    }
                if (message.level != null && message.hasOwnProperty("level"))
                    if (!$util.isInteger(message.level))
                        return "level: integer expected";
                if (message.kills != null && message.hasOwnProperty("kills"))
                    if (!$util.isInteger(message.kills))
                        return "kills: integer expected";
                if (message.deaths != null && message.hasOwnProperty("deaths"))
                    if (!$util.isInteger(message.deaths))
                        return "deaths: integer expected";
                if (message.assists != null && message.hasOwnProperty("assists"))
                    if (!$util.isInteger(message.assists))
                        return "assists: integer expected";
                if (message.denies != null && message.hasOwnProperty("denies"))
                    if (!$util.isInteger(message.denies))
                        return "denies: integer expected";
                if (message.last_hits != null && message.hasOwnProperty("last_hits"))
                    if (!$util.isInteger(message.last_hits))
                        return "last_hits: integer expected";
                if (message.gold != null && message.hasOwnProperty("gold"))
                    if (!$util.isInteger(message.gold))
                        return "gold: integer expected";
                if (message.net_worth != null && message.hasOwnProperty("net_worth"))
                    if (!$util.isInteger(message.net_worth))
                        return "net_worth: integer expected";
                if (message.label != null && message.hasOwnProperty("label"))
                    if (!$util.isString(message.label))
                        return "label: string expected";
                return null;
            };

            /**
             * Creates a Player message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof protocol.LiveMatch.Player
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {protocol.LiveMatch.Player} Player
             */
            Player.fromObject = function fromObject(object) {
                if (object instanceof $root.protocol.LiveMatch.Player)
                    return object;
                let message = new $root.protocol.LiveMatch.Player();
                if (object.account_id != null)
                    message.account_id = object.account_id >>> 0;
                if (object.name != null)
                    message.name = String(object.name);
                if (object.persona_name != null)
                    message.persona_name = String(object.persona_name);
                if (object.avatar_url != null)
                    message.avatar_url = String(object.avatar_url);
                if (object.avatar_medium_url != null)
                    message.avatar_medium_url = String(object.avatar_medium_url);
                if (object.avatar_full_url != null)
                    message.avatar_full_url = String(object.avatar_full_url);
                if (object.is_pro != null)
                    message.is_pro = Boolean(object.is_pro);
                if (object.hero_id != null)
                    if ($util.Long)
                        (message.hero_id = $util.Long.fromValue(object.hero_id)).unsigned = true;
                    else if (typeof object.hero_id === "string")
                        message.hero_id = parseInt(object.hero_id, 10);
                    else if (typeof object.hero_id === "number")
                        message.hero_id = object.hero_id;
                    else if (typeof object.hero_id === "object")
                        message.hero_id = new $util.LongBits(object.hero_id.low >>> 0, object.hero_id.high >>> 0).toNumber(true);
                if (object.player_slot != null)
                    message.player_slot = object.player_slot >>> 0;
                switch (object.team) {
                case "GAME_TEAM_UNKNOWN":
                case 0:
                    message.team = 0;
                    break;
                case "GAME_TEAM_GOODGUYS":
                case 2:
                    message.team = 2;
                    break;
                case "GAME_TEAM_BADGUYS":
                case 3:
                    message.team = 3;
                    break;
                case "GAME_TEAM_NEUTRALS":
                case 4:
                    message.team = 4;
                    break;
                case "GAME_TEAM_NOTEAM":
                case 5:
                    message.team = 5;
                    break;
                case "GAME_TEAM_CUSTOM1":
                case 6:
                    message.team = 6;
                    break;
                case "GAME_TEAM_CUSTOM2":
                case 7:
                    message.team = 7;
                    break;
                case "GAME_TEAM_CUSTOM3":
                case 8:
                    message.team = 8;
                    break;
                case "GAME_TEAM_CUSTOM4":
                case 9:
                    message.team = 9;
                    break;
                case "GAME_TEAM_CUSTOM5":
                case 10:
                    message.team = 10;
                    break;
                case "GAME_TEAM_CUSTOM6":
                case 11:
                    message.team = 11;
                    break;
                case "GAME_TEAM_CUSTOM7":
                case 12:
                    message.team = 12;
                    break;
                case "GAME_TEAM_CUSTOM8":
                case 13:
                    message.team = 13;
                    break;
                }
                if (object.level != null)
                    message.level = object.level >>> 0;
                if (object.kills != null)
                    message.kills = object.kills >>> 0;
                if (object.deaths != null)
                    message.deaths = object.deaths >>> 0;
                if (object.assists != null)
                    message.assists = object.assists >>> 0;
                if (object.denies != null)
                    message.denies = object.denies >>> 0;
                if (object.last_hits != null)
                    message.last_hits = object.last_hits >>> 0;
                if (object.gold != null)
                    message.gold = object.gold >>> 0;
                if (object.net_worth != null)
                    message.net_worth = object.net_worth >>> 0;
                if (object.label != null)
                    message.label = String(object.label);
                return message;
            };

            /**
             * Creates a plain object from a Player message. Also converts values to other types if specified.
             * @function toObject
             * @memberof protocol.LiveMatch.Player
             * @static
             * @param {protocol.LiveMatch.Player} message Player
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Player.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.account_id = 0;
                    object.name = "";
                    object.persona_name = "";
                    object.avatar_url = "";
                    object.avatar_medium_url = "";
                    object.avatar_full_url = "";
                    object.is_pro = false;
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, true);
                        object.hero_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.hero_id = options.longs === String ? "0" : 0;
                    object.player_slot = 0;
                    object.team = options.enums === String ? "GAME_TEAM_UNKNOWN" : 0;
                    object.level = 0;
                    object.kills = 0;
                    object.deaths = 0;
                    object.assists = 0;
                    object.denies = 0;
                    object.last_hits = 0;
                    object.gold = 0;
                    object.net_worth = 0;
                    object.label = "";
                }
                if (message.account_id != null && message.hasOwnProperty("account_id"))
                    object.account_id = message.account_id;
                if (message.name != null && message.hasOwnProperty("name"))
                    object.name = message.name;
                if (message.persona_name != null && message.hasOwnProperty("persona_name"))
                    object.persona_name = message.persona_name;
                if (message.avatar_url != null && message.hasOwnProperty("avatar_url"))
                    object.avatar_url = message.avatar_url;
                if (message.avatar_medium_url != null && message.hasOwnProperty("avatar_medium_url"))
                    object.avatar_medium_url = message.avatar_medium_url;
                if (message.avatar_full_url != null && message.hasOwnProperty("avatar_full_url"))
                    object.avatar_full_url = message.avatar_full_url;
                if (message.is_pro != null && message.hasOwnProperty("is_pro"))
                    object.is_pro = message.is_pro;
                if (message.hero_id != null && message.hasOwnProperty("hero_id"))
                    if (typeof message.hero_id === "number")
                        object.hero_id = options.longs === String ? String(message.hero_id) : message.hero_id;
                    else
                        object.hero_id = options.longs === String ? $util.Long.prototype.toString.call(message.hero_id) : options.longs === Number ? new $util.LongBits(message.hero_id.low >>> 0, message.hero_id.high >>> 0).toNumber(true) : message.hero_id;
                if (message.player_slot != null && message.hasOwnProperty("player_slot"))
                    object.player_slot = message.player_slot;
                if (message.team != null && message.hasOwnProperty("team"))
                    object.team = options.enums === String ? $root.protocol.GameTeam[message.team] : message.team;
                if (message.level != null && message.hasOwnProperty("level"))
                    object.level = message.level;
                if (message.kills != null && message.hasOwnProperty("kills"))
                    object.kills = message.kills;
                if (message.deaths != null && message.hasOwnProperty("deaths"))
                    object.deaths = message.deaths;
                if (message.assists != null && message.hasOwnProperty("assists"))
                    object.assists = message.assists;
                if (message.denies != null && message.hasOwnProperty("denies"))
                    object.denies = message.denies;
                if (message.last_hits != null && message.hasOwnProperty("last_hits"))
                    object.last_hits = message.last_hits;
                if (message.gold != null && message.hasOwnProperty("gold"))
                    object.gold = message.gold;
                if (message.net_worth != null && message.hasOwnProperty("net_worth"))
                    object.net_worth = message.net_worth;
                if (message.label != null && message.hasOwnProperty("label"))
                    object.label = message.label;
                return object;
            };

            /**
             * Converts this Player to JSON.
             * @function toJSON
             * @memberof protocol.LiveMatch.Player
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Player.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return Player;
        })();

        return LiveMatch;
    })();

    /**
     * LobbyType enum.
     * @name protocol.LobbyType
     * @enum {string}
     * @property {number} LOBBY_TYPE_CASUAL_MATCH=0 LOBBY_TYPE_CASUAL_MATCH value
     * @property {number} LOBBY_TYPE_PRACTICE=1 LOBBY_TYPE_PRACTICE value
     * @property {number} LOBBY_TYPE_COOP_BOT_MATCH=4 LOBBY_TYPE_COOP_BOT_MATCH value
     * @property {number} LOBBY_TYPE_LEGACY_TEAM_MATCH=5 LOBBY_TYPE_LEGACY_TEAM_MATCH value
     * @property {number} LOBBY_TYPE_LEGACY_SOLO_QUEUE_MATCH=6 LOBBY_TYPE_LEGACY_SOLO_QUEUE_MATCH value
     * @property {number} LOBBY_TYPE_COMPETITIVE_MATCH=7 LOBBY_TYPE_COMPETITIVE_MATCH value
     * @property {number} LOBBY_TYPE_CASUAL_1V1_MATCH=8 LOBBY_TYPE_CASUAL_1V1_MATCH value
     * @property {number} LOBBY_TYPE_WEEKEND_TOURNEY=9 LOBBY_TYPE_WEEKEND_TOURNEY value
     * @property {number} LOBBY_TYPE_LOCAL_BOT_MATCH=10 LOBBY_TYPE_LOCAL_BOT_MATCH value
     * @property {number} LOBBY_TYPE_SPECTATOR=11 LOBBY_TYPE_SPECTATOR value
     * @property {number} LOBBY_TYPE_EVENT_MATCH=12 LOBBY_TYPE_EVENT_MATCH value
     */
    protocol.LobbyType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "LOBBY_TYPE_CASUAL_MATCH"] = 0;
        values[valuesById[1] = "LOBBY_TYPE_PRACTICE"] = 1;
        values[valuesById[4] = "LOBBY_TYPE_COOP_BOT_MATCH"] = 4;
        values[valuesById[5] = "LOBBY_TYPE_LEGACY_TEAM_MATCH"] = 5;
        values[valuesById[6] = "LOBBY_TYPE_LEGACY_SOLO_QUEUE_MATCH"] = 6;
        values[valuesById[7] = "LOBBY_TYPE_COMPETITIVE_MATCH"] = 7;
        values[valuesById[8] = "LOBBY_TYPE_CASUAL_1V1_MATCH"] = 8;
        values[valuesById[9] = "LOBBY_TYPE_WEEKEND_TOURNEY"] = 9;
        values[valuesById[10] = "LOBBY_TYPE_LOCAL_BOT_MATCH"] = 10;
        values[valuesById[11] = "LOBBY_TYPE_SPECTATOR"] = 11;
        values[valuesById[12] = "LOBBY_TYPE_EVENT_MATCH"] = 12;
        return values;
    })();

    /**
     * GameMode enum.
     * @name protocol.GameMode
     * @enum {string}
     * @property {number} GAME_MODE_NONE=0 GAME_MODE_NONE value
     * @property {number} GAME_MODE_AP=1 GAME_MODE_AP value
     * @property {number} GAME_MODE_CM=2 GAME_MODE_CM value
     * @property {number} GAME_MODE_RD=3 GAME_MODE_RD value
     * @property {number} GAME_MODE_SD=4 GAME_MODE_SD value
     * @property {number} GAME_MODE_AR=5 GAME_MODE_AR value
     * @property {number} GAME_MODE_INTRO=6 GAME_MODE_INTRO value
     * @property {number} GAME_MODE_HW=7 GAME_MODE_HW value
     * @property {number} GAME_MODE_REVERSE_CM=8 GAME_MODE_REVERSE_CM value
     * @property {number} GAME_MODE_XMAS=9 GAME_MODE_XMAS value
     * @property {number} GAME_MODE_TUTORIAL=10 GAME_MODE_TUTORIAL value
     * @property {number} GAME_MODE_MO=11 GAME_MODE_MO value
     * @property {number} GAME_MODE_LP=12 GAME_MODE_LP value
     * @property {number} GAME_MODE_POOL1=13 GAME_MODE_POOL1 value
     * @property {number} GAME_MODE_FH=14 GAME_MODE_FH value
     * @property {number} GAME_MODE_CUSTOM=15 GAME_MODE_CUSTOM value
     * @property {number} GAME_MODE_CD=16 GAME_MODE_CD value
     * @property {number} GAME_MODE_BD=17 GAME_MODE_BD value
     * @property {number} GAME_MODE_ABILITY_DRAFT=18 GAME_MODE_ABILITY_DRAFT value
     * @property {number} GAME_MODE_EVENT=19 GAME_MODE_EVENT value
     * @property {number} GAME_MODE_ARDM=20 GAME_MODE_ARDM value
     * @property {number} GAME_MODE_1V1_MID=21 GAME_MODE_1V1_MID value
     * @property {number} GAME_MODE_ALL_DRAFT=22 GAME_MODE_ALL_DRAFT value
     * @property {number} GAME_MODE_TURBO=23 GAME_MODE_TURBO value
     * @property {number} GAME_MODE_MUTATION=24 GAME_MODE_MUTATION value
     * @property {number} GAME_MODE_COACHES_CHALLENGE=25 GAME_MODE_COACHES_CHALLENGE value
     */
    protocol.GameMode = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "GAME_MODE_NONE"] = 0;
        values[valuesById[1] = "GAME_MODE_AP"] = 1;
        values[valuesById[2] = "GAME_MODE_CM"] = 2;
        values[valuesById[3] = "GAME_MODE_RD"] = 3;
        values[valuesById[4] = "GAME_MODE_SD"] = 4;
        values[valuesById[5] = "GAME_MODE_AR"] = 5;
        values[valuesById[6] = "GAME_MODE_INTRO"] = 6;
        values[valuesById[7] = "GAME_MODE_HW"] = 7;
        values[valuesById[8] = "GAME_MODE_REVERSE_CM"] = 8;
        values[valuesById[9] = "GAME_MODE_XMAS"] = 9;
        values[valuesById[10] = "GAME_MODE_TUTORIAL"] = 10;
        values[valuesById[11] = "GAME_MODE_MO"] = 11;
        values[valuesById[12] = "GAME_MODE_LP"] = 12;
        values[valuesById[13] = "GAME_MODE_POOL1"] = 13;
        values[valuesById[14] = "GAME_MODE_FH"] = 14;
        values[valuesById[15] = "GAME_MODE_CUSTOM"] = 15;
        values[valuesById[16] = "GAME_MODE_CD"] = 16;
        values[valuesById[17] = "GAME_MODE_BD"] = 17;
        values[valuesById[18] = "GAME_MODE_ABILITY_DRAFT"] = 18;
        values[valuesById[19] = "GAME_MODE_EVENT"] = 19;
        values[valuesById[20] = "GAME_MODE_ARDM"] = 20;
        values[valuesById[21] = "GAME_MODE_1V1_MID"] = 21;
        values[valuesById[22] = "GAME_MODE_ALL_DRAFT"] = 22;
        values[valuesById[23] = "GAME_MODE_TURBO"] = 23;
        values[valuesById[24] = "GAME_MODE_MUTATION"] = 24;
        values[valuesById[25] = "GAME_MODE_COACHES_CHALLENGE"] = 25;
        return values;
    })();

    /**
     * GameState enum.
     * @name protocol.GameState
     * @enum {string}
     * @property {number} GAME_STATE_INIT=0 GAME_STATE_INIT value
     * @property {number} GAME_STATE_WAIT_FOR_PLAYERS_TO_LOAD=1 GAME_STATE_WAIT_FOR_PLAYERS_TO_LOAD value
     * @property {number} GAME_STATE_HERO_SELECTION=2 GAME_STATE_HERO_SELECTION value
     * @property {number} GAME_STATE_STRATEGY_TIME=3 GAME_STATE_STRATEGY_TIME value
     * @property {number} GAME_STATE_PRE_GAME=4 GAME_STATE_PRE_GAME value
     * @property {number} GAME_STATE_GAME_IN_PROGRESS=5 GAME_STATE_GAME_IN_PROGRESS value
     * @property {number} GAME_STATE_POST_GAME=6 GAME_STATE_POST_GAME value
     * @property {number} GAME_STATE_DISCONNECT=7 GAME_STATE_DISCONNECT value
     * @property {number} GAME_STATE_TEAM_SHOWCASE=8 GAME_STATE_TEAM_SHOWCASE value
     * @property {number} GAME_STATE_CUSTOM_GAME_SETUP=9 GAME_STATE_CUSTOM_GAME_SETUP value
     * @property {number} GAME_STATE_WAIT_FOR_MAP_TO_LOAD=10 GAME_STATE_WAIT_FOR_MAP_TO_LOAD value
     * @property {number} GAME_STATE_LAST=11 GAME_STATE_LAST value
     */
    protocol.GameState = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "GAME_STATE_INIT"] = 0;
        values[valuesById[1] = "GAME_STATE_WAIT_FOR_PLAYERS_TO_LOAD"] = 1;
        values[valuesById[2] = "GAME_STATE_HERO_SELECTION"] = 2;
        values[valuesById[3] = "GAME_STATE_STRATEGY_TIME"] = 3;
        values[valuesById[4] = "GAME_STATE_PRE_GAME"] = 4;
        values[valuesById[5] = "GAME_STATE_GAME_IN_PROGRESS"] = 5;
        values[valuesById[6] = "GAME_STATE_POST_GAME"] = 6;
        values[valuesById[7] = "GAME_STATE_DISCONNECT"] = 7;
        values[valuesById[8] = "GAME_STATE_TEAM_SHOWCASE"] = 8;
        values[valuesById[9] = "GAME_STATE_CUSTOM_GAME_SETUP"] = 9;
        values[valuesById[10] = "GAME_STATE_WAIT_FOR_MAP_TO_LOAD"] = 10;
        values[valuesById[11] = "GAME_STATE_LAST"] = 11;
        return values;
    })();

    /**
     * GameTeam enum.
     * @name protocol.GameTeam
     * @enum {string}
     * @property {number} GAME_TEAM_UNKNOWN=0 GAME_TEAM_UNKNOWN value
     * @property {number} GAME_TEAM_GOODGUYS=2 GAME_TEAM_GOODGUYS value
     * @property {number} GAME_TEAM_BADGUYS=3 GAME_TEAM_BADGUYS value
     * @property {number} GAME_TEAM_NEUTRALS=4 GAME_TEAM_NEUTRALS value
     * @property {number} GAME_TEAM_NOTEAM=5 GAME_TEAM_NOTEAM value
     * @property {number} GAME_TEAM_CUSTOM1=6 GAME_TEAM_CUSTOM1 value
     * @property {number} GAME_TEAM_CUSTOM2=7 GAME_TEAM_CUSTOM2 value
     * @property {number} GAME_TEAM_CUSTOM3=8 GAME_TEAM_CUSTOM3 value
     * @property {number} GAME_TEAM_CUSTOM4=9 GAME_TEAM_CUSTOM4 value
     * @property {number} GAME_TEAM_CUSTOM5=10 GAME_TEAM_CUSTOM5 value
     * @property {number} GAME_TEAM_CUSTOM6=11 GAME_TEAM_CUSTOM6 value
     * @property {number} GAME_TEAM_CUSTOM7=12 GAME_TEAM_CUSTOM7 value
     * @property {number} GAME_TEAM_CUSTOM8=13 GAME_TEAM_CUSTOM8 value
     */
    protocol.GameTeam = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "GAME_TEAM_UNKNOWN"] = 0;
        values[valuesById[2] = "GAME_TEAM_GOODGUYS"] = 2;
        values[valuesById[3] = "GAME_TEAM_BADGUYS"] = 3;
        values[valuesById[4] = "GAME_TEAM_NEUTRALS"] = 4;
        values[valuesById[5] = "GAME_TEAM_NOTEAM"] = 5;
        values[valuesById[6] = "GAME_TEAM_CUSTOM1"] = 6;
        values[valuesById[7] = "GAME_TEAM_CUSTOM2"] = 7;
        values[valuesById[8] = "GAME_TEAM_CUSTOM3"] = 8;
        values[valuesById[9] = "GAME_TEAM_CUSTOM4"] = 9;
        values[valuesById[10] = "GAME_TEAM_CUSTOM5"] = 10;
        values[valuesById[11] = "GAME_TEAM_CUSTOM6"] = 11;
        values[valuesById[12] = "GAME_TEAM_CUSTOM7"] = 12;
        values[valuesById[13] = "GAME_TEAM_CUSTOM8"] = 13;
        return values;
    })();

    /**
     * BuildingType enum.
     * @name protocol.BuildingType
     * @enum {string}
     * @property {number} BUILDING_TYPE_TOWER=0 BUILDING_TYPE_TOWER value
     * @property {number} BUILDING_TYPE_BARRACKS=1 BUILDING_TYPE_BARRACKS value
     * @property {number} BUILDING_TYPE_ANCIENT=2 BUILDING_TYPE_ANCIENT value
     */
    protocol.BuildingType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "BUILDING_TYPE_TOWER"] = 0;
        values[valuesById[1] = "BUILDING_TYPE_BARRACKS"] = 1;
        values[valuesById[2] = "BUILDING_TYPE_ANCIENT"] = 2;
        return values;
    })();

    /**
     * FantasyRole enum.
     * @name protocol.FantasyRole
     * @enum {string}
     * @property {number} FANTASY_ROLE_UNDEFINED=0 FANTASY_ROLE_UNDEFINED value
     * @property {number} FANTASY_ROLE_CORE=1 FANTASY_ROLE_CORE value
     * @property {number} FANTASY_ROLE_SUPPORT=2 FANTASY_ROLE_SUPPORT value
     * @property {number} FANTASY_ROLE_OFFLANE=3 FANTASY_ROLE_OFFLANE value
     * @property {number} FANTASY_ROLE_MID=4 FANTASY_ROLE_MID value
     */
    protocol.FantasyRole = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "FANTASY_ROLE_UNDEFINED"] = 0;
        values[valuesById[1] = "FANTASY_ROLE_CORE"] = 1;
        values[valuesById[2] = "FANTASY_ROLE_SUPPORT"] = 2;
        values[valuesById[3] = "FANTASY_ROLE_OFFLANE"] = 3;
        values[valuesById[4] = "FANTASY_ROLE_MID"] = 4;
        return values;
    })();

    /**
     * LaneType enum.
     * @name protocol.LaneType
     * @enum {string}
     * @property {number} LANE_TYPE_UNKNOWN=0 LANE_TYPE_UNKNOWN value
     * @property {number} LANE_TYPE_SAFE=1 LANE_TYPE_SAFE value
     * @property {number} LANE_TYPE_OFF=2 LANE_TYPE_OFF value
     * @property {number} LANE_TYPE_MID=3 LANE_TYPE_MID value
     * @property {number} LANE_TYPE_JUNGLE=4 LANE_TYPE_JUNGLE value
     * @property {number} LANE_TYPE_ROAM=5 LANE_TYPE_ROAM value
     */
    protocol.LaneType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "LANE_TYPE_UNKNOWN"] = 0;
        values[valuesById[1] = "LANE_TYPE_SAFE"] = 1;
        values[valuesById[2] = "LANE_TYPE_OFF"] = 2;
        values[valuesById[3] = "LANE_TYPE_MID"] = 3;
        values[valuesById[4] = "LANE_TYPE_JUNGLE"] = 4;
        values[valuesById[5] = "LANE_TYPE_ROAM"] = 5;
        return values;
    })();

    /**
     * MatchOutcome enum.
     * @name protocol.MatchOutcome
     * @enum {string}
     * @property {number} MATCH_OUTCOME_UNKNOWN=0 MATCH_OUTCOME_UNKNOWN value
     * @property {number} MATCH_OUTCOME_RAD_VICTORY=2 MATCH_OUTCOME_RAD_VICTORY value
     * @property {number} MATCH_OUTCOME_DIRE_VICTORY=3 MATCH_OUTCOME_DIRE_VICTORY value
     * @property {number} MATCH_OUTCOME_NOT_SCORED_POOR_NETWORK_CONDITIONS=64 MATCH_OUTCOME_NOT_SCORED_POOR_NETWORK_CONDITIONS value
     * @property {number} MATCH_OUTCOME_NOT_SCORED_LEAVER=65 MATCH_OUTCOME_NOT_SCORED_LEAVER value
     * @property {number} MATCH_OUTCOME_NOT_SCORED_SERVER_CRASH=66 MATCH_OUTCOME_NOT_SCORED_SERVER_CRASH value
     * @property {number} MATCH_OUTCOME_NOT_SCORED_NEVER_STARTED=67 MATCH_OUTCOME_NOT_SCORED_NEVER_STARTED value
     * @property {number} MATCH_OUTCOME_NOT_SCORED_CANCELED=68 MATCH_OUTCOME_NOT_SCORED_CANCELED value
     */
    protocol.MatchOutcome = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "MATCH_OUTCOME_UNKNOWN"] = 0;
        values[valuesById[2] = "MATCH_OUTCOME_RAD_VICTORY"] = 2;
        values[valuesById[3] = "MATCH_OUTCOME_DIRE_VICTORY"] = 3;
        values[valuesById[64] = "MATCH_OUTCOME_NOT_SCORED_POOR_NETWORK_CONDITIONS"] = 64;
        values[valuesById[65] = "MATCH_OUTCOME_NOT_SCORED_LEAVER"] = 65;
        values[valuesById[66] = "MATCH_OUTCOME_NOT_SCORED_SERVER_CRASH"] = 66;
        values[valuesById[67] = "MATCH_OUTCOME_NOT_SCORED_NEVER_STARTED"] = 67;
        values[valuesById[68] = "MATCH_OUTCOME_NOT_SCORED_CANCELED"] = 68;
        return values;
    })();

    protocol.Search = (function() {

        /**
         * Properties of a Search.
         * @memberof protocol
         * @interface ISearch
         * @property {Array.<protocol.Search.IPlayer>|null} [players] Search players
         * @property {Array.<Long>|null} [hero_ids] Search hero_ids
         */

        /**
         * Constructs a new Search.
         * @memberof protocol
         * @classdesc Represents a Search.
         * @implements ISearch
         * @constructor
         * @param {protocol.ISearch=} [properties] Properties to set
         */
        function Search(properties) {
            this.players = [];
            this.hero_ids = [];
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Search players.
         * @member {Array.<protocol.Search.IPlayer>} players
         * @memberof protocol.Search
         * @instance
         */
        Search.prototype.players = $util.emptyArray;

        /**
         * Search hero_ids.
         * @member {Array.<Long>} hero_ids
         * @memberof protocol.Search
         * @instance
         */
        Search.prototype.hero_ids = $util.emptyArray;

        /**
         * Creates a new Search instance using the specified properties.
         * @function create
         * @memberof protocol.Search
         * @static
         * @param {protocol.ISearch=} [properties] Properties to set
         * @returns {protocol.Search} Search instance
         */
        Search.create = function create(properties) {
            return new Search(properties);
        };

        /**
         * Encodes the specified Search message. Does not implicitly {@link protocol.Search.verify|verify} messages.
         * @function encode
         * @memberof protocol.Search
         * @static
         * @param {protocol.ISearch} message Search message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Search.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.players != null && message.players.length)
                for (let i = 0; i < message.players.length; ++i)
                    $root.protocol.Search.Player.encode(message.players[i], writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            if (message.hero_ids != null && message.hero_ids.length) {
                writer.uint32(/* id 2, wireType 2 =*/18).fork();
                for (let i = 0; i < message.hero_ids.length; ++i)
                    writer.uint64(message.hero_ids[i]);
                writer.ldelim();
            }
            return writer;
        };

        /**
         * Encodes the specified Search message, length delimited. Does not implicitly {@link protocol.Search.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protocol.Search
         * @static
         * @param {protocol.ISearch} message Search message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Search.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Search message from the specified reader or buffer.
         * @function decode
         * @memberof protocol.Search
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protocol.Search} Search
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Search.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.protocol.Search();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    if (!(message.players && message.players.length))
                        message.players = [];
                    message.players.push($root.protocol.Search.Player.decode(reader, reader.uint32()));
                    break;
                case 2:
                    if (!(message.hero_ids && message.hero_ids.length))
                        message.hero_ids = [];
                    if ((tag & 7) === 2) {
                        let end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2)
                            message.hero_ids.push(reader.uint64());
                    } else
                        message.hero_ids.push(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Search message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protocol.Search
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protocol.Search} Search
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Search.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Search message.
         * @function verify
         * @memberof protocol.Search
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Search.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.players != null && message.hasOwnProperty("players")) {
                if (!Array.isArray(message.players))
                    return "players: array expected";
                for (let i = 0; i < message.players.length; ++i) {
                    let error = $root.protocol.Search.Player.verify(message.players[i]);
                    if (error)
                        return "players." + error;
                }
            }
            if (message.hero_ids != null && message.hasOwnProperty("hero_ids")) {
                if (!Array.isArray(message.hero_ids))
                    return "hero_ids: array expected";
                for (let i = 0; i < message.hero_ids.length; ++i)
                    if (!$util.isInteger(message.hero_ids[i]) && !(message.hero_ids[i] && $util.isInteger(message.hero_ids[i].low) && $util.isInteger(message.hero_ids[i].high)))
                        return "hero_ids: integer|Long[] expected";
            }
            return null;
        };

        /**
         * Creates a Search message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protocol.Search
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protocol.Search} Search
         */
        Search.fromObject = function fromObject(object) {
            if (object instanceof $root.protocol.Search)
                return object;
            let message = new $root.protocol.Search();
            if (object.players) {
                if (!Array.isArray(object.players))
                    throw TypeError(".protocol.Search.players: array expected");
                message.players = [];
                for (let i = 0; i < object.players.length; ++i) {
                    if (typeof object.players[i] !== "object")
                        throw TypeError(".protocol.Search.players: object expected");
                    message.players[i] = $root.protocol.Search.Player.fromObject(object.players[i]);
                }
            }
            if (object.hero_ids) {
                if (!Array.isArray(object.hero_ids))
                    throw TypeError(".protocol.Search.hero_ids: array expected");
                message.hero_ids = [];
                for (let i = 0; i < object.hero_ids.length; ++i)
                    if ($util.Long)
                        (message.hero_ids[i] = $util.Long.fromValue(object.hero_ids[i])).unsigned = true;
                    else if (typeof object.hero_ids[i] === "string")
                        message.hero_ids[i] = parseInt(object.hero_ids[i], 10);
                    else if (typeof object.hero_ids[i] === "number")
                        message.hero_ids[i] = object.hero_ids[i];
                    else if (typeof object.hero_ids[i] === "object")
                        message.hero_ids[i] = new $util.LongBits(object.hero_ids[i].low >>> 0, object.hero_ids[i].high >>> 0).toNumber(true);
            }
            return message;
        };

        /**
         * Creates a plain object from a Search message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protocol.Search
         * @static
         * @param {protocol.Search} message Search
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Search.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.arrays || options.defaults) {
                object.players = [];
                object.hero_ids = [];
            }
            if (message.players && message.players.length) {
                object.players = [];
                for (let j = 0; j < message.players.length; ++j)
                    object.players[j] = $root.protocol.Search.Player.toObject(message.players[j], options);
            }
            if (message.hero_ids && message.hero_ids.length) {
                object.hero_ids = [];
                for (let j = 0; j < message.hero_ids.length; ++j)
                    if (typeof message.hero_ids[j] === "number")
                        object.hero_ids[j] = options.longs === String ? String(message.hero_ids[j]) : message.hero_ids[j];
                    else
                        object.hero_ids[j] = options.longs === String ? $util.Long.prototype.toString.call(message.hero_ids[j]) : options.longs === Number ? new $util.LongBits(message.hero_ids[j].low >>> 0, message.hero_ids[j].high >>> 0).toNumber(true) : message.hero_ids[j];
            }
            return object;
        };

        /**
         * Converts this Search to JSON.
         * @function toJSON
         * @memberof protocol.Search
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Search.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        Search.Player = (function() {

            /**
             * Properties of a Player.
             * @memberof protocol.Search
             * @interface IPlayer
             * @property {number|null} [account_id] Player account_id
             * @property {string|null} [name] Player name
             * @property {string|null} [persona_name] Player persona_name
             * @property {string|null} [avatar_url] Player avatar_url
             * @property {string|null} [avatar_medium_url] Player avatar_medium_url
             * @property {string|null} [avatar_full_url] Player avatar_full_url
             * @property {boolean|null} [is_pro] Player is_pro
             */

            /**
             * Constructs a new Player.
             * @memberof protocol.Search
             * @classdesc Represents a Player.
             * @implements IPlayer
             * @constructor
             * @param {protocol.Search.IPlayer=} [properties] Properties to set
             */
            function Player(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Player account_id.
             * @member {number} account_id
             * @memberof protocol.Search.Player
             * @instance
             */
            Player.prototype.account_id = 0;

            /**
             * Player name.
             * @member {string} name
             * @memberof protocol.Search.Player
             * @instance
             */
            Player.prototype.name = "";

            /**
             * Player persona_name.
             * @member {string} persona_name
             * @memberof protocol.Search.Player
             * @instance
             */
            Player.prototype.persona_name = "";

            /**
             * Player avatar_url.
             * @member {string} avatar_url
             * @memberof protocol.Search.Player
             * @instance
             */
            Player.prototype.avatar_url = "";

            /**
             * Player avatar_medium_url.
             * @member {string} avatar_medium_url
             * @memberof protocol.Search.Player
             * @instance
             */
            Player.prototype.avatar_medium_url = "";

            /**
             * Player avatar_full_url.
             * @member {string} avatar_full_url
             * @memberof protocol.Search.Player
             * @instance
             */
            Player.prototype.avatar_full_url = "";

            /**
             * Player is_pro.
             * @member {boolean} is_pro
             * @memberof protocol.Search.Player
             * @instance
             */
            Player.prototype.is_pro = false;

            /**
             * Creates a new Player instance using the specified properties.
             * @function create
             * @memberof protocol.Search.Player
             * @static
             * @param {protocol.Search.IPlayer=} [properties] Properties to set
             * @returns {protocol.Search.Player} Player instance
             */
            Player.create = function create(properties) {
                return new Player(properties);
            };

            /**
             * Encodes the specified Player message. Does not implicitly {@link protocol.Search.Player.verify|verify} messages.
             * @function encode
             * @memberof protocol.Search.Player
             * @static
             * @param {protocol.Search.IPlayer} message Player message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Player.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.account_id != null && message.hasOwnProperty("account_id"))
                    writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.account_id);
                if (message.name != null && message.hasOwnProperty("name"))
                    writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
                if (message.persona_name != null && message.hasOwnProperty("persona_name"))
                    writer.uint32(/* id 3, wireType 2 =*/26).string(message.persona_name);
                if (message.avatar_url != null && message.hasOwnProperty("avatar_url"))
                    writer.uint32(/* id 4, wireType 2 =*/34).string(message.avatar_url);
                if (message.avatar_medium_url != null && message.hasOwnProperty("avatar_medium_url"))
                    writer.uint32(/* id 5, wireType 2 =*/42).string(message.avatar_medium_url);
                if (message.avatar_full_url != null && message.hasOwnProperty("avatar_full_url"))
                    writer.uint32(/* id 6, wireType 2 =*/50).string(message.avatar_full_url);
                if (message.is_pro != null && message.hasOwnProperty("is_pro"))
                    writer.uint32(/* id 7, wireType 0 =*/56).bool(message.is_pro);
                return writer;
            };

            /**
             * Encodes the specified Player message, length delimited. Does not implicitly {@link protocol.Search.Player.verify|verify} messages.
             * @function encodeDelimited
             * @memberof protocol.Search.Player
             * @static
             * @param {protocol.Search.IPlayer} message Player message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Player.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a Player message from the specified reader or buffer.
             * @function decode
             * @memberof protocol.Search.Player
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {protocol.Search.Player} Player
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Player.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.protocol.Search.Player();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.account_id = reader.uint32();
                        break;
                    case 2:
                        message.name = reader.string();
                        break;
                    case 3:
                        message.persona_name = reader.string();
                        break;
                    case 4:
                        message.avatar_url = reader.string();
                        break;
                    case 5:
                        message.avatar_medium_url = reader.string();
                        break;
                    case 6:
                        message.avatar_full_url = reader.string();
                        break;
                    case 7:
                        message.is_pro = reader.bool();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a Player message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof protocol.Search.Player
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {protocol.Search.Player} Player
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Player.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a Player message.
             * @function verify
             * @memberof protocol.Search.Player
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Player.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.account_id != null && message.hasOwnProperty("account_id"))
                    if (!$util.isInteger(message.account_id))
                        return "account_id: integer expected";
                if (message.name != null && message.hasOwnProperty("name"))
                    if (!$util.isString(message.name))
                        return "name: string expected";
                if (message.persona_name != null && message.hasOwnProperty("persona_name"))
                    if (!$util.isString(message.persona_name))
                        return "persona_name: string expected";
                if (message.avatar_url != null && message.hasOwnProperty("avatar_url"))
                    if (!$util.isString(message.avatar_url))
                        return "avatar_url: string expected";
                if (message.avatar_medium_url != null && message.hasOwnProperty("avatar_medium_url"))
                    if (!$util.isString(message.avatar_medium_url))
                        return "avatar_medium_url: string expected";
                if (message.avatar_full_url != null && message.hasOwnProperty("avatar_full_url"))
                    if (!$util.isString(message.avatar_full_url))
                        return "avatar_full_url: string expected";
                if (message.is_pro != null && message.hasOwnProperty("is_pro"))
                    if (typeof message.is_pro !== "boolean")
                        return "is_pro: boolean expected";
                return null;
            };

            /**
             * Creates a Player message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof protocol.Search.Player
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {protocol.Search.Player} Player
             */
            Player.fromObject = function fromObject(object) {
                if (object instanceof $root.protocol.Search.Player)
                    return object;
                let message = new $root.protocol.Search.Player();
                if (object.account_id != null)
                    message.account_id = object.account_id >>> 0;
                if (object.name != null)
                    message.name = String(object.name);
                if (object.persona_name != null)
                    message.persona_name = String(object.persona_name);
                if (object.avatar_url != null)
                    message.avatar_url = String(object.avatar_url);
                if (object.avatar_medium_url != null)
                    message.avatar_medium_url = String(object.avatar_medium_url);
                if (object.avatar_full_url != null)
                    message.avatar_full_url = String(object.avatar_full_url);
                if (object.is_pro != null)
                    message.is_pro = Boolean(object.is_pro);
                return message;
            };

            /**
             * Creates a plain object from a Player message. Also converts values to other types if specified.
             * @function toObject
             * @memberof protocol.Search.Player
             * @static
             * @param {protocol.Search.Player} message Player
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Player.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.account_id = 0;
                    object.name = "";
                    object.persona_name = "";
                    object.avatar_url = "";
                    object.avatar_medium_url = "";
                    object.avatar_full_url = "";
                    object.is_pro = false;
                }
                if (message.account_id != null && message.hasOwnProperty("account_id"))
                    object.account_id = message.account_id;
                if (message.name != null && message.hasOwnProperty("name"))
                    object.name = message.name;
                if (message.persona_name != null && message.hasOwnProperty("persona_name"))
                    object.persona_name = message.persona_name;
                if (message.avatar_url != null && message.hasOwnProperty("avatar_url"))
                    object.avatar_url = message.avatar_url;
                if (message.avatar_medium_url != null && message.hasOwnProperty("avatar_medium_url"))
                    object.avatar_medium_url = message.avatar_medium_url;
                if (message.avatar_full_url != null && message.hasOwnProperty("avatar_full_url"))
                    object.avatar_full_url = message.avatar_full_url;
                if (message.is_pro != null && message.hasOwnProperty("is_pro"))
                    object.is_pro = message.is_pro;
                return object;
            };

            /**
             * Converts this Player to JSON.
             * @function toJSON
             * @memberof protocol.Search.Player
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Player.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return Player;
        })();

        return Search;
    })();

    protocol.Player = (function() {

        /**
         * Properties of a Player.
         * @memberof protocol
         * @interface IPlayer
         * @property {number|null} [account_id] Player account_id
         * @property {string|null} [name] Player name
         * @property {string|null} [persona_name] Player persona_name
         * @property {string|null} [avatar_url] Player avatar_url
         * @property {string|null} [avatar_medium_url] Player avatar_medium_url
         * @property {string|null} [avatar_full_url] Player avatar_full_url
         * @property {boolean|null} [is_pro] Player is_pro
         * @property {protocol.Player.ITeam|null} [team] Player team
         * @property {Array.<protocol.Player.IMatch>|null} [matches] Player matches
         */

        /**
         * Constructs a new Player.
         * @memberof protocol
         * @classdesc Represents a Player.
         * @implements IPlayer
         * @constructor
         * @param {protocol.IPlayer=} [properties] Properties to set
         */
        function Player(properties) {
            this.matches = [];
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Player account_id.
         * @member {number} account_id
         * @memberof protocol.Player
         * @instance
         */
        Player.prototype.account_id = 0;

        /**
         * Player name.
         * @member {string} name
         * @memberof protocol.Player
         * @instance
         */
        Player.prototype.name = "";

        /**
         * Player persona_name.
         * @member {string} persona_name
         * @memberof protocol.Player
         * @instance
         */
        Player.prototype.persona_name = "";

        /**
         * Player avatar_url.
         * @member {string} avatar_url
         * @memberof protocol.Player
         * @instance
         */
        Player.prototype.avatar_url = "";

        /**
         * Player avatar_medium_url.
         * @member {string} avatar_medium_url
         * @memberof protocol.Player
         * @instance
         */
        Player.prototype.avatar_medium_url = "";

        /**
         * Player avatar_full_url.
         * @member {string} avatar_full_url
         * @memberof protocol.Player
         * @instance
         */
        Player.prototype.avatar_full_url = "";

        /**
         * Player is_pro.
         * @member {boolean} is_pro
         * @memberof protocol.Player
         * @instance
         */
        Player.prototype.is_pro = false;

        /**
         * Player team.
         * @member {protocol.Player.ITeam|null|undefined} team
         * @memberof protocol.Player
         * @instance
         */
        Player.prototype.team = null;

        /**
         * Player matches.
         * @member {Array.<protocol.Player.IMatch>} matches
         * @memberof protocol.Player
         * @instance
         */
        Player.prototype.matches = $util.emptyArray;

        /**
         * Creates a new Player instance using the specified properties.
         * @function create
         * @memberof protocol.Player
         * @static
         * @param {protocol.IPlayer=} [properties] Properties to set
         * @returns {protocol.Player} Player instance
         */
        Player.create = function create(properties) {
            return new Player(properties);
        };

        /**
         * Encodes the specified Player message. Does not implicitly {@link protocol.Player.verify|verify} messages.
         * @function encode
         * @memberof protocol.Player
         * @static
         * @param {protocol.IPlayer} message Player message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Player.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.account_id != null && message.hasOwnProperty("account_id"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.account_id);
            if (message.name != null && message.hasOwnProperty("name"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
            if (message.persona_name != null && message.hasOwnProperty("persona_name"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.persona_name);
            if (message.avatar_url != null && message.hasOwnProperty("avatar_url"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.avatar_url);
            if (message.avatar_medium_url != null && message.hasOwnProperty("avatar_medium_url"))
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.avatar_medium_url);
            if (message.avatar_full_url != null && message.hasOwnProperty("avatar_full_url"))
                writer.uint32(/* id 6, wireType 2 =*/50).string(message.avatar_full_url);
            if (message.is_pro != null && message.hasOwnProperty("is_pro"))
                writer.uint32(/* id 7, wireType 0 =*/56).bool(message.is_pro);
            if (message.team != null && message.hasOwnProperty("team"))
                $root.protocol.Player.Team.encode(message.team, writer.uint32(/* id 100, wireType 2 =*/802).fork()).ldelim();
            if (message.matches != null && message.matches.length)
                for (let i = 0; i < message.matches.length; ++i)
                    $root.protocol.Player.Match.encode(message.matches[i], writer.uint32(/* id 101, wireType 2 =*/810).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified Player message, length delimited. Does not implicitly {@link protocol.Player.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protocol.Player
         * @static
         * @param {protocol.IPlayer} message Player message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Player.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Player message from the specified reader or buffer.
         * @function decode
         * @memberof protocol.Player
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protocol.Player} Player
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Player.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.protocol.Player();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.account_id = reader.uint32();
                    break;
                case 2:
                    message.name = reader.string();
                    break;
                case 3:
                    message.persona_name = reader.string();
                    break;
                case 4:
                    message.avatar_url = reader.string();
                    break;
                case 5:
                    message.avatar_medium_url = reader.string();
                    break;
                case 6:
                    message.avatar_full_url = reader.string();
                    break;
                case 7:
                    message.is_pro = reader.bool();
                    break;
                case 100:
                    message.team = $root.protocol.Player.Team.decode(reader, reader.uint32());
                    break;
                case 101:
                    if (!(message.matches && message.matches.length))
                        message.matches = [];
                    message.matches.push($root.protocol.Player.Match.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Player message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protocol.Player
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protocol.Player} Player
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Player.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Player message.
         * @function verify
         * @memberof protocol.Player
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Player.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.account_id != null && message.hasOwnProperty("account_id"))
                if (!$util.isInteger(message.account_id))
                    return "account_id: integer expected";
            if (message.name != null && message.hasOwnProperty("name"))
                if (!$util.isString(message.name))
                    return "name: string expected";
            if (message.persona_name != null && message.hasOwnProperty("persona_name"))
                if (!$util.isString(message.persona_name))
                    return "persona_name: string expected";
            if (message.avatar_url != null && message.hasOwnProperty("avatar_url"))
                if (!$util.isString(message.avatar_url))
                    return "avatar_url: string expected";
            if (message.avatar_medium_url != null && message.hasOwnProperty("avatar_medium_url"))
                if (!$util.isString(message.avatar_medium_url))
                    return "avatar_medium_url: string expected";
            if (message.avatar_full_url != null && message.hasOwnProperty("avatar_full_url"))
                if (!$util.isString(message.avatar_full_url))
                    return "avatar_full_url: string expected";
            if (message.is_pro != null && message.hasOwnProperty("is_pro"))
                if (typeof message.is_pro !== "boolean")
                    return "is_pro: boolean expected";
            if (message.team != null && message.hasOwnProperty("team")) {
                let error = $root.protocol.Player.Team.verify(message.team);
                if (error)
                    return "team." + error;
            }
            if (message.matches != null && message.hasOwnProperty("matches")) {
                if (!Array.isArray(message.matches))
                    return "matches: array expected";
                for (let i = 0; i < message.matches.length; ++i) {
                    let error = $root.protocol.Player.Match.verify(message.matches[i]);
                    if (error)
                        return "matches." + error;
                }
            }
            return null;
        };

        /**
         * Creates a Player message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protocol.Player
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protocol.Player} Player
         */
        Player.fromObject = function fromObject(object) {
            if (object instanceof $root.protocol.Player)
                return object;
            let message = new $root.protocol.Player();
            if (object.account_id != null)
                message.account_id = object.account_id >>> 0;
            if (object.name != null)
                message.name = String(object.name);
            if (object.persona_name != null)
                message.persona_name = String(object.persona_name);
            if (object.avatar_url != null)
                message.avatar_url = String(object.avatar_url);
            if (object.avatar_medium_url != null)
                message.avatar_medium_url = String(object.avatar_medium_url);
            if (object.avatar_full_url != null)
                message.avatar_full_url = String(object.avatar_full_url);
            if (object.is_pro != null)
                message.is_pro = Boolean(object.is_pro);
            if (object.team != null) {
                if (typeof object.team !== "object")
                    throw TypeError(".protocol.Player.team: object expected");
                message.team = $root.protocol.Player.Team.fromObject(object.team);
            }
            if (object.matches) {
                if (!Array.isArray(object.matches))
                    throw TypeError(".protocol.Player.matches: array expected");
                message.matches = [];
                for (let i = 0; i < object.matches.length; ++i) {
                    if (typeof object.matches[i] !== "object")
                        throw TypeError(".protocol.Player.matches: object expected");
                    message.matches[i] = $root.protocol.Player.Match.fromObject(object.matches[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a Player message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protocol.Player
         * @static
         * @param {protocol.Player} message Player
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Player.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.arrays || options.defaults)
                object.matches = [];
            if (options.defaults) {
                object.account_id = 0;
                object.name = "";
                object.persona_name = "";
                object.avatar_url = "";
                object.avatar_medium_url = "";
                object.avatar_full_url = "";
                object.is_pro = false;
                object.team = null;
            }
            if (message.account_id != null && message.hasOwnProperty("account_id"))
                object.account_id = message.account_id;
            if (message.name != null && message.hasOwnProperty("name"))
                object.name = message.name;
            if (message.persona_name != null && message.hasOwnProperty("persona_name"))
                object.persona_name = message.persona_name;
            if (message.avatar_url != null && message.hasOwnProperty("avatar_url"))
                object.avatar_url = message.avatar_url;
            if (message.avatar_medium_url != null && message.hasOwnProperty("avatar_medium_url"))
                object.avatar_medium_url = message.avatar_medium_url;
            if (message.avatar_full_url != null && message.hasOwnProperty("avatar_full_url"))
                object.avatar_full_url = message.avatar_full_url;
            if (message.is_pro != null && message.hasOwnProperty("is_pro"))
                object.is_pro = message.is_pro;
            if (message.team != null && message.hasOwnProperty("team"))
                object.team = $root.protocol.Player.Team.toObject(message.team, options);
            if (message.matches && message.matches.length) {
                object.matches = [];
                for (let j = 0; j < message.matches.length; ++j)
                    object.matches[j] = $root.protocol.Player.Match.toObject(message.matches[j], options);
            }
            return object;
        };

        /**
         * Converts this Player to JSON.
         * @function toJSON
         * @memberof protocol.Player
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Player.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        Player.Match = (function() {

            /**
             * Properties of a Match.
             * @memberof protocol.Player
             * @interface IMatch
             * @property {Long|null} [match_id] Match match_id
             * @property {Long|null} [lobby_id] Match lobby_id
             * @property {protocol.LobbyType|null} [lobby_type] Match lobby_type
             * @property {Long|null} [league_id] Match league_id
             * @property {Long|null} [series_id] Match series_id
             * @property {protocol.GameMode|null} [game_mode] Match game_mode
             * @property {number|null} [average_mmr] Match average_mmr
             * @property {Long|null} [radiant_team_id] Match radiant_team_id
             * @property {string|null} [radiant_team_name] Match radiant_team_name
             * @property {string|null} [radiant_team_tag] Match radiant_team_tag
             * @property {Long|null} [radiant_team_logo] Match radiant_team_logo
             * @property {string|null} [radiant_team_logo_url] Match radiant_team_logo_url
             * @property {Long|null} [dire_team_id] Match dire_team_id
             * @property {string|null} [dire_team_name] Match dire_team_name
             * @property {string|null} [dire_team_tag] Match dire_team_tag
             * @property {Long|null} [dire_team_logo] Match dire_team_logo
             * @property {string|null} [dire_team_logo_url] Match dire_team_logo_url
             * @property {google.protobuf.ITimestamp|null} [activate_time] Match activate_time
             * @property {google.protobuf.ITimestamp|null} [deactivate_time] Match deactivate_time
             * @property {google.protobuf.ITimestamp|null} [last_update_time] Match last_update_time
             * @property {google.protobuf.ITimestamp|null} [start_time] Match start_time
             * @property {number|null} [series_type] Match series_type
             * @property {number|null} [series_game] Match series_game
             * @property {number|null} [duration] Match duration
             * @property {number|null} [radiant_score] Match radiant_score
             * @property {number|null} [dire_score] Match dire_score
             * @property {protocol.MatchOutcome|null} [outcome] Match outcome
             * @property {protocol.Player.Match.IPlayerDetails|null} [player_details] Match player_details
             */

            /**
             * Constructs a new Match.
             * @memberof protocol.Player
             * @classdesc Represents a Match.
             * @implements IMatch
             * @constructor
             * @param {protocol.Player.IMatch=} [properties] Properties to set
             */
            function Match(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Match match_id.
             * @member {Long} match_id
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.match_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

            /**
             * Match lobby_id.
             * @member {Long} lobby_id
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.lobby_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

            /**
             * Match lobby_type.
             * @member {protocol.LobbyType} lobby_type
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.lobby_type = 0;

            /**
             * Match league_id.
             * @member {Long} league_id
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.league_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

            /**
             * Match series_id.
             * @member {Long} series_id
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.series_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

            /**
             * Match game_mode.
             * @member {protocol.GameMode} game_mode
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.game_mode = 0;

            /**
             * Match average_mmr.
             * @member {number} average_mmr
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.average_mmr = 0;

            /**
             * Match radiant_team_id.
             * @member {Long} radiant_team_id
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.radiant_team_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

            /**
             * Match radiant_team_name.
             * @member {string} radiant_team_name
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.radiant_team_name = "";

            /**
             * Match radiant_team_tag.
             * @member {string} radiant_team_tag
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.radiant_team_tag = "";

            /**
             * Match radiant_team_logo.
             * @member {Long} radiant_team_logo
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.radiant_team_logo = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

            /**
             * Match radiant_team_logo_url.
             * @member {string} radiant_team_logo_url
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.radiant_team_logo_url = "";

            /**
             * Match dire_team_id.
             * @member {Long} dire_team_id
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.dire_team_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

            /**
             * Match dire_team_name.
             * @member {string} dire_team_name
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.dire_team_name = "";

            /**
             * Match dire_team_tag.
             * @member {string} dire_team_tag
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.dire_team_tag = "";

            /**
             * Match dire_team_logo.
             * @member {Long} dire_team_logo
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.dire_team_logo = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

            /**
             * Match dire_team_logo_url.
             * @member {string} dire_team_logo_url
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.dire_team_logo_url = "";

            /**
             * Match activate_time.
             * @member {google.protobuf.ITimestamp|null|undefined} activate_time
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.activate_time = null;

            /**
             * Match deactivate_time.
             * @member {google.protobuf.ITimestamp|null|undefined} deactivate_time
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.deactivate_time = null;

            /**
             * Match last_update_time.
             * @member {google.protobuf.ITimestamp|null|undefined} last_update_time
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.last_update_time = null;

            /**
             * Match start_time.
             * @member {google.protobuf.ITimestamp|null|undefined} start_time
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.start_time = null;

            /**
             * Match series_type.
             * @member {number} series_type
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.series_type = 0;

            /**
             * Match series_game.
             * @member {number} series_game
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.series_game = 0;

            /**
             * Match duration.
             * @member {number} duration
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.duration = 0;

            /**
             * Match radiant_score.
             * @member {number} radiant_score
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.radiant_score = 0;

            /**
             * Match dire_score.
             * @member {number} dire_score
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.dire_score = 0;

            /**
             * Match outcome.
             * @member {protocol.MatchOutcome} outcome
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.outcome = 0;

            /**
             * Match player_details.
             * @member {protocol.Player.Match.IPlayerDetails|null|undefined} player_details
             * @memberof protocol.Player.Match
             * @instance
             */
            Match.prototype.player_details = null;

            /**
             * Creates a new Match instance using the specified properties.
             * @function create
             * @memberof protocol.Player.Match
             * @static
             * @param {protocol.Player.IMatch=} [properties] Properties to set
             * @returns {protocol.Player.Match} Match instance
             */
            Match.create = function create(properties) {
                return new Match(properties);
            };

            /**
             * Encodes the specified Match message. Does not implicitly {@link protocol.Player.Match.verify|verify} messages.
             * @function encode
             * @memberof protocol.Player.Match
             * @static
             * @param {protocol.Player.IMatch} message Match message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Match.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.match_id != null && message.hasOwnProperty("match_id"))
                    writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.match_id);
                if (message.lobby_id != null && message.hasOwnProperty("lobby_id"))
                    writer.uint32(/* id 2, wireType 0 =*/16).uint64(message.lobby_id);
                if (message.lobby_type != null && message.hasOwnProperty("lobby_type"))
                    writer.uint32(/* id 3, wireType 0 =*/24).int32(message.lobby_type);
                if (message.league_id != null && message.hasOwnProperty("league_id"))
                    writer.uint32(/* id 4, wireType 0 =*/32).uint64(message.league_id);
                if (message.series_id != null && message.hasOwnProperty("series_id"))
                    writer.uint32(/* id 5, wireType 0 =*/40).uint64(message.series_id);
                if (message.game_mode != null && message.hasOwnProperty("game_mode"))
                    writer.uint32(/* id 6, wireType 0 =*/48).int32(message.game_mode);
                if (message.average_mmr != null && message.hasOwnProperty("average_mmr"))
                    writer.uint32(/* id 7, wireType 0 =*/56).uint32(message.average_mmr);
                if (message.radiant_team_id != null && message.hasOwnProperty("radiant_team_id"))
                    writer.uint32(/* id 8, wireType 0 =*/64).uint64(message.radiant_team_id);
                if (message.radiant_team_name != null && message.hasOwnProperty("radiant_team_name"))
                    writer.uint32(/* id 9, wireType 2 =*/74).string(message.radiant_team_name);
                if (message.radiant_team_tag != null && message.hasOwnProperty("radiant_team_tag"))
                    writer.uint32(/* id 10, wireType 2 =*/82).string(message.radiant_team_tag);
                if (message.radiant_team_logo != null && message.hasOwnProperty("radiant_team_logo"))
                    writer.uint32(/* id 11, wireType 0 =*/88).uint64(message.radiant_team_logo);
                if (message.radiant_team_logo_url != null && message.hasOwnProperty("radiant_team_logo_url"))
                    writer.uint32(/* id 12, wireType 2 =*/98).string(message.radiant_team_logo_url);
                if (message.dire_team_id != null && message.hasOwnProperty("dire_team_id"))
                    writer.uint32(/* id 13, wireType 0 =*/104).uint64(message.dire_team_id);
                if (message.dire_team_name != null && message.hasOwnProperty("dire_team_name"))
                    writer.uint32(/* id 14, wireType 2 =*/114).string(message.dire_team_name);
                if (message.dire_team_tag != null && message.hasOwnProperty("dire_team_tag"))
                    writer.uint32(/* id 15, wireType 2 =*/122).string(message.dire_team_tag);
                if (message.dire_team_logo != null && message.hasOwnProperty("dire_team_logo"))
                    writer.uint32(/* id 16, wireType 0 =*/128).uint64(message.dire_team_logo);
                if (message.dire_team_logo_url != null && message.hasOwnProperty("dire_team_logo_url"))
                    writer.uint32(/* id 17, wireType 2 =*/138).string(message.dire_team_logo_url);
                if (message.activate_time != null && message.hasOwnProperty("activate_time"))
                    $root.google.protobuf.Timestamp.encode(message.activate_time, writer.uint32(/* id 18, wireType 2 =*/146).fork()).ldelim();
                if (message.deactivate_time != null && message.hasOwnProperty("deactivate_time"))
                    $root.google.protobuf.Timestamp.encode(message.deactivate_time, writer.uint32(/* id 19, wireType 2 =*/154).fork()).ldelim();
                if (message.last_update_time != null && message.hasOwnProperty("last_update_time"))
                    $root.google.protobuf.Timestamp.encode(message.last_update_time, writer.uint32(/* id 20, wireType 2 =*/162).fork()).ldelim();
                if (message.start_time != null && message.hasOwnProperty("start_time"))
                    $root.google.protobuf.Timestamp.encode(message.start_time, writer.uint32(/* id 21, wireType 2 =*/170).fork()).ldelim();
                if (message.series_type != null && message.hasOwnProperty("series_type"))
                    writer.uint32(/* id 22, wireType 0 =*/176).uint32(message.series_type);
                if (message.series_game != null && message.hasOwnProperty("series_game"))
                    writer.uint32(/* id 23, wireType 0 =*/184).uint32(message.series_game);
                if (message.duration != null && message.hasOwnProperty("duration"))
                    writer.uint32(/* id 24, wireType 0 =*/192).uint32(message.duration);
                if (message.radiant_score != null && message.hasOwnProperty("radiant_score"))
                    writer.uint32(/* id 25, wireType 0 =*/200).uint32(message.radiant_score);
                if (message.dire_score != null && message.hasOwnProperty("dire_score"))
                    writer.uint32(/* id 26, wireType 0 =*/208).uint32(message.dire_score);
                if (message.outcome != null && message.hasOwnProperty("outcome"))
                    writer.uint32(/* id 27, wireType 0 =*/216).int32(message.outcome);
                if (message.player_details != null && message.hasOwnProperty("player_details"))
                    $root.protocol.Player.Match.PlayerDetails.encode(message.player_details, writer.uint32(/* id 100, wireType 2 =*/802).fork()).ldelim();
                return writer;
            };

            /**
             * Encodes the specified Match message, length delimited. Does not implicitly {@link protocol.Player.Match.verify|verify} messages.
             * @function encodeDelimited
             * @memberof protocol.Player.Match
             * @static
             * @param {protocol.Player.IMatch} message Match message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Match.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a Match message from the specified reader or buffer.
             * @function decode
             * @memberof protocol.Player.Match
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {protocol.Player.Match} Match
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Match.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.protocol.Player.Match();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.match_id = reader.uint64();
                        break;
                    case 2:
                        message.lobby_id = reader.uint64();
                        break;
                    case 3:
                        message.lobby_type = reader.int32();
                        break;
                    case 4:
                        message.league_id = reader.uint64();
                        break;
                    case 5:
                        message.series_id = reader.uint64();
                        break;
                    case 6:
                        message.game_mode = reader.int32();
                        break;
                    case 7:
                        message.average_mmr = reader.uint32();
                        break;
                    case 8:
                        message.radiant_team_id = reader.uint64();
                        break;
                    case 9:
                        message.radiant_team_name = reader.string();
                        break;
                    case 10:
                        message.radiant_team_tag = reader.string();
                        break;
                    case 11:
                        message.radiant_team_logo = reader.uint64();
                        break;
                    case 12:
                        message.radiant_team_logo_url = reader.string();
                        break;
                    case 13:
                        message.dire_team_id = reader.uint64();
                        break;
                    case 14:
                        message.dire_team_name = reader.string();
                        break;
                    case 15:
                        message.dire_team_tag = reader.string();
                        break;
                    case 16:
                        message.dire_team_logo = reader.uint64();
                        break;
                    case 17:
                        message.dire_team_logo_url = reader.string();
                        break;
                    case 18:
                        message.activate_time = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                        break;
                    case 19:
                        message.deactivate_time = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                        break;
                    case 20:
                        message.last_update_time = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                        break;
                    case 21:
                        message.start_time = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                        break;
                    case 22:
                        message.series_type = reader.uint32();
                        break;
                    case 23:
                        message.series_game = reader.uint32();
                        break;
                    case 24:
                        message.duration = reader.uint32();
                        break;
                    case 25:
                        message.radiant_score = reader.uint32();
                        break;
                    case 26:
                        message.dire_score = reader.uint32();
                        break;
                    case 27:
                        message.outcome = reader.int32();
                        break;
                    case 100:
                        message.player_details = $root.protocol.Player.Match.PlayerDetails.decode(reader, reader.uint32());
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a Match message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof protocol.Player.Match
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {protocol.Player.Match} Match
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Match.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a Match message.
             * @function verify
             * @memberof protocol.Player.Match
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Match.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.match_id != null && message.hasOwnProperty("match_id"))
                    if (!$util.isInteger(message.match_id) && !(message.match_id && $util.isInteger(message.match_id.low) && $util.isInteger(message.match_id.high)))
                        return "match_id: integer|Long expected";
                if (message.lobby_id != null && message.hasOwnProperty("lobby_id"))
                    if (!$util.isInteger(message.lobby_id) && !(message.lobby_id && $util.isInteger(message.lobby_id.low) && $util.isInteger(message.lobby_id.high)))
                        return "lobby_id: integer|Long expected";
                if (message.lobby_type != null && message.hasOwnProperty("lobby_type"))
                    switch (message.lobby_type) {
                    default:
                        return "lobby_type: enum value expected";
                    case 0:
                    case 1:
                    case 4:
                    case 5:
                    case 6:
                    case 7:
                    case 8:
                    case 9:
                    case 10:
                    case 11:
                    case 12:
                        break;
                    }
                if (message.league_id != null && message.hasOwnProperty("league_id"))
                    if (!$util.isInteger(message.league_id) && !(message.league_id && $util.isInteger(message.league_id.low) && $util.isInteger(message.league_id.high)))
                        return "league_id: integer|Long expected";
                if (message.series_id != null && message.hasOwnProperty("series_id"))
                    if (!$util.isInteger(message.series_id) && !(message.series_id && $util.isInteger(message.series_id.low) && $util.isInteger(message.series_id.high)))
                        return "series_id: integer|Long expected";
                if (message.game_mode != null && message.hasOwnProperty("game_mode"))
                    switch (message.game_mode) {
                    default:
                        return "game_mode: enum value expected";
                    case 0:
                    case 1:
                    case 2:
                    case 3:
                    case 4:
                    case 5:
                    case 6:
                    case 7:
                    case 8:
                    case 9:
                    case 10:
                    case 11:
                    case 12:
                    case 13:
                    case 14:
                    case 15:
                    case 16:
                    case 17:
                    case 18:
                    case 19:
                    case 20:
                    case 21:
                    case 22:
                    case 23:
                    case 24:
                    case 25:
                        break;
                    }
                if (message.average_mmr != null && message.hasOwnProperty("average_mmr"))
                    if (!$util.isInteger(message.average_mmr))
                        return "average_mmr: integer expected";
                if (message.radiant_team_id != null && message.hasOwnProperty("radiant_team_id"))
                    if (!$util.isInteger(message.radiant_team_id) && !(message.radiant_team_id && $util.isInteger(message.radiant_team_id.low) && $util.isInteger(message.radiant_team_id.high)))
                        return "radiant_team_id: integer|Long expected";
                if (message.radiant_team_name != null && message.hasOwnProperty("radiant_team_name"))
                    if (!$util.isString(message.radiant_team_name))
                        return "radiant_team_name: string expected";
                if (message.radiant_team_tag != null && message.hasOwnProperty("radiant_team_tag"))
                    if (!$util.isString(message.radiant_team_tag))
                        return "radiant_team_tag: string expected";
                if (message.radiant_team_logo != null && message.hasOwnProperty("radiant_team_logo"))
                    if (!$util.isInteger(message.radiant_team_logo) && !(message.radiant_team_logo && $util.isInteger(message.radiant_team_logo.low) && $util.isInteger(message.radiant_team_logo.high)))
                        return "radiant_team_logo: integer|Long expected";
                if (message.radiant_team_logo_url != null && message.hasOwnProperty("radiant_team_logo_url"))
                    if (!$util.isString(message.radiant_team_logo_url))
                        return "radiant_team_logo_url: string expected";
                if (message.dire_team_id != null && message.hasOwnProperty("dire_team_id"))
                    if (!$util.isInteger(message.dire_team_id) && !(message.dire_team_id && $util.isInteger(message.dire_team_id.low) && $util.isInteger(message.dire_team_id.high)))
                        return "dire_team_id: integer|Long expected";
                if (message.dire_team_name != null && message.hasOwnProperty("dire_team_name"))
                    if (!$util.isString(message.dire_team_name))
                        return "dire_team_name: string expected";
                if (message.dire_team_tag != null && message.hasOwnProperty("dire_team_tag"))
                    if (!$util.isString(message.dire_team_tag))
                        return "dire_team_tag: string expected";
                if (message.dire_team_logo != null && message.hasOwnProperty("dire_team_logo"))
                    if (!$util.isInteger(message.dire_team_logo) && !(message.dire_team_logo && $util.isInteger(message.dire_team_logo.low) && $util.isInteger(message.dire_team_logo.high)))
                        return "dire_team_logo: integer|Long expected";
                if (message.dire_team_logo_url != null && message.hasOwnProperty("dire_team_logo_url"))
                    if (!$util.isString(message.dire_team_logo_url))
                        return "dire_team_logo_url: string expected";
                if (message.activate_time != null && message.hasOwnProperty("activate_time")) {
                    let error = $root.google.protobuf.Timestamp.verify(message.activate_time);
                    if (error)
                        return "activate_time." + error;
                }
                if (message.deactivate_time != null && message.hasOwnProperty("deactivate_time")) {
                    let error = $root.google.protobuf.Timestamp.verify(message.deactivate_time);
                    if (error)
                        return "deactivate_time." + error;
                }
                if (message.last_update_time != null && message.hasOwnProperty("last_update_time")) {
                    let error = $root.google.protobuf.Timestamp.verify(message.last_update_time);
                    if (error)
                        return "last_update_time." + error;
                }
                if (message.start_time != null && message.hasOwnProperty("start_time")) {
                    let error = $root.google.protobuf.Timestamp.verify(message.start_time);
                    if (error)
                        return "start_time." + error;
                }
                if (message.series_type != null && message.hasOwnProperty("series_type"))
                    if (!$util.isInteger(message.series_type))
                        return "series_type: integer expected";
                if (message.series_game != null && message.hasOwnProperty("series_game"))
                    if (!$util.isInteger(message.series_game))
                        return "series_game: integer expected";
                if (message.duration != null && message.hasOwnProperty("duration"))
                    if (!$util.isInteger(message.duration))
                        return "duration: integer expected";
                if (message.radiant_score != null && message.hasOwnProperty("radiant_score"))
                    if (!$util.isInteger(message.radiant_score))
                        return "radiant_score: integer expected";
                if (message.dire_score != null && message.hasOwnProperty("dire_score"))
                    if (!$util.isInteger(message.dire_score))
                        return "dire_score: integer expected";
                if (message.outcome != null && message.hasOwnProperty("outcome"))
                    switch (message.outcome) {
                    default:
                        return "outcome: enum value expected";
                    case 0:
                    case 2:
                    case 3:
                    case 64:
                    case 65:
                    case 66:
                    case 67:
                    case 68:
                        break;
                    }
                if (message.player_details != null && message.hasOwnProperty("player_details")) {
                    let error = $root.protocol.Player.Match.PlayerDetails.verify(message.player_details);
                    if (error)
                        return "player_details." + error;
                }
                return null;
            };

            /**
             * Creates a Match message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof protocol.Player.Match
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {protocol.Player.Match} Match
             */
            Match.fromObject = function fromObject(object) {
                if (object instanceof $root.protocol.Player.Match)
                    return object;
                let message = new $root.protocol.Player.Match();
                if (object.match_id != null)
                    if ($util.Long)
                        (message.match_id = $util.Long.fromValue(object.match_id)).unsigned = true;
                    else if (typeof object.match_id === "string")
                        message.match_id = parseInt(object.match_id, 10);
                    else if (typeof object.match_id === "number")
                        message.match_id = object.match_id;
                    else if (typeof object.match_id === "object")
                        message.match_id = new $util.LongBits(object.match_id.low >>> 0, object.match_id.high >>> 0).toNumber(true);
                if (object.lobby_id != null)
                    if ($util.Long)
                        (message.lobby_id = $util.Long.fromValue(object.lobby_id)).unsigned = true;
                    else if (typeof object.lobby_id === "string")
                        message.lobby_id = parseInt(object.lobby_id, 10);
                    else if (typeof object.lobby_id === "number")
                        message.lobby_id = object.lobby_id;
                    else if (typeof object.lobby_id === "object")
                        message.lobby_id = new $util.LongBits(object.lobby_id.low >>> 0, object.lobby_id.high >>> 0).toNumber(true);
                switch (object.lobby_type) {
                case "LOBBY_TYPE_CASUAL_MATCH":
                case 0:
                    message.lobby_type = 0;
                    break;
                case "LOBBY_TYPE_PRACTICE":
                case 1:
                    message.lobby_type = 1;
                    break;
                case "LOBBY_TYPE_COOP_BOT_MATCH":
                case 4:
                    message.lobby_type = 4;
                    break;
                case "LOBBY_TYPE_LEGACY_TEAM_MATCH":
                case 5:
                    message.lobby_type = 5;
                    break;
                case "LOBBY_TYPE_LEGACY_SOLO_QUEUE_MATCH":
                case 6:
                    message.lobby_type = 6;
                    break;
                case "LOBBY_TYPE_COMPETITIVE_MATCH":
                case 7:
                    message.lobby_type = 7;
                    break;
                case "LOBBY_TYPE_CASUAL_1V1_MATCH":
                case 8:
                    message.lobby_type = 8;
                    break;
                case "LOBBY_TYPE_WEEKEND_TOURNEY":
                case 9:
                    message.lobby_type = 9;
                    break;
                case "LOBBY_TYPE_LOCAL_BOT_MATCH":
                case 10:
                    message.lobby_type = 10;
                    break;
                case "LOBBY_TYPE_SPECTATOR":
                case 11:
                    message.lobby_type = 11;
                    break;
                case "LOBBY_TYPE_EVENT_MATCH":
                case 12:
                    message.lobby_type = 12;
                    break;
                }
                if (object.league_id != null)
                    if ($util.Long)
                        (message.league_id = $util.Long.fromValue(object.league_id)).unsigned = true;
                    else if (typeof object.league_id === "string")
                        message.league_id = parseInt(object.league_id, 10);
                    else if (typeof object.league_id === "number")
                        message.league_id = object.league_id;
                    else if (typeof object.league_id === "object")
                        message.league_id = new $util.LongBits(object.league_id.low >>> 0, object.league_id.high >>> 0).toNumber(true);
                if (object.series_id != null)
                    if ($util.Long)
                        (message.series_id = $util.Long.fromValue(object.series_id)).unsigned = true;
                    else if (typeof object.series_id === "string")
                        message.series_id = parseInt(object.series_id, 10);
                    else if (typeof object.series_id === "number")
                        message.series_id = object.series_id;
                    else if (typeof object.series_id === "object")
                        message.series_id = new $util.LongBits(object.series_id.low >>> 0, object.series_id.high >>> 0).toNumber(true);
                switch (object.game_mode) {
                case "GAME_MODE_NONE":
                case 0:
                    message.game_mode = 0;
                    break;
                case "GAME_MODE_AP":
                case 1:
                    message.game_mode = 1;
                    break;
                case "GAME_MODE_CM":
                case 2:
                    message.game_mode = 2;
                    break;
                case "GAME_MODE_RD":
                case 3:
                    message.game_mode = 3;
                    break;
                case "GAME_MODE_SD":
                case 4:
                    message.game_mode = 4;
                    break;
                case "GAME_MODE_AR":
                case 5:
                    message.game_mode = 5;
                    break;
                case "GAME_MODE_INTRO":
                case 6:
                    message.game_mode = 6;
                    break;
                case "GAME_MODE_HW":
                case 7:
                    message.game_mode = 7;
                    break;
                case "GAME_MODE_REVERSE_CM":
                case 8:
                    message.game_mode = 8;
                    break;
                case "GAME_MODE_XMAS":
                case 9:
                    message.game_mode = 9;
                    break;
                case "GAME_MODE_TUTORIAL":
                case 10:
                    message.game_mode = 10;
                    break;
                case "GAME_MODE_MO":
                case 11:
                    message.game_mode = 11;
                    break;
                case "GAME_MODE_LP":
                case 12:
                    message.game_mode = 12;
                    break;
                case "GAME_MODE_POOL1":
                case 13:
                    message.game_mode = 13;
                    break;
                case "GAME_MODE_FH":
                case 14:
                    message.game_mode = 14;
                    break;
                case "GAME_MODE_CUSTOM":
                case 15:
                    message.game_mode = 15;
                    break;
                case "GAME_MODE_CD":
                case 16:
                    message.game_mode = 16;
                    break;
                case "GAME_MODE_BD":
                case 17:
                    message.game_mode = 17;
                    break;
                case "GAME_MODE_ABILITY_DRAFT":
                case 18:
                    message.game_mode = 18;
                    break;
                case "GAME_MODE_EVENT":
                case 19:
                    message.game_mode = 19;
                    break;
                case "GAME_MODE_ARDM":
                case 20:
                    message.game_mode = 20;
                    break;
                case "GAME_MODE_1V1_MID":
                case 21:
                    message.game_mode = 21;
                    break;
                case "GAME_MODE_ALL_DRAFT":
                case 22:
                    message.game_mode = 22;
                    break;
                case "GAME_MODE_TURBO":
                case 23:
                    message.game_mode = 23;
                    break;
                case "GAME_MODE_MUTATION":
                case 24:
                    message.game_mode = 24;
                    break;
                case "GAME_MODE_COACHES_CHALLENGE":
                case 25:
                    message.game_mode = 25;
                    break;
                }
                if (object.average_mmr != null)
                    message.average_mmr = object.average_mmr >>> 0;
                if (object.radiant_team_id != null)
                    if ($util.Long)
                        (message.radiant_team_id = $util.Long.fromValue(object.radiant_team_id)).unsigned = true;
                    else if (typeof object.radiant_team_id === "string")
                        message.radiant_team_id = parseInt(object.radiant_team_id, 10);
                    else if (typeof object.radiant_team_id === "number")
                        message.radiant_team_id = object.radiant_team_id;
                    else if (typeof object.radiant_team_id === "object")
                        message.radiant_team_id = new $util.LongBits(object.radiant_team_id.low >>> 0, object.radiant_team_id.high >>> 0).toNumber(true);
                if (object.radiant_team_name != null)
                    message.radiant_team_name = String(object.radiant_team_name);
                if (object.radiant_team_tag != null)
                    message.radiant_team_tag = String(object.radiant_team_tag);
                if (object.radiant_team_logo != null)
                    if ($util.Long)
                        (message.radiant_team_logo = $util.Long.fromValue(object.radiant_team_logo)).unsigned = true;
                    else if (typeof object.radiant_team_logo === "string")
                        message.radiant_team_logo = parseInt(object.radiant_team_logo, 10);
                    else if (typeof object.radiant_team_logo === "number")
                        message.radiant_team_logo = object.radiant_team_logo;
                    else if (typeof object.radiant_team_logo === "object")
                        message.radiant_team_logo = new $util.LongBits(object.radiant_team_logo.low >>> 0, object.radiant_team_logo.high >>> 0).toNumber(true);
                if (object.radiant_team_logo_url != null)
                    message.radiant_team_logo_url = String(object.radiant_team_logo_url);
                if (object.dire_team_id != null)
                    if ($util.Long)
                        (message.dire_team_id = $util.Long.fromValue(object.dire_team_id)).unsigned = true;
                    else if (typeof object.dire_team_id === "string")
                        message.dire_team_id = parseInt(object.dire_team_id, 10);
                    else if (typeof object.dire_team_id === "number")
                        message.dire_team_id = object.dire_team_id;
                    else if (typeof object.dire_team_id === "object")
                        message.dire_team_id = new $util.LongBits(object.dire_team_id.low >>> 0, object.dire_team_id.high >>> 0).toNumber(true);
                if (object.dire_team_name != null)
                    message.dire_team_name = String(object.dire_team_name);
                if (object.dire_team_tag != null)
                    message.dire_team_tag = String(object.dire_team_tag);
                if (object.dire_team_logo != null)
                    if ($util.Long)
                        (message.dire_team_logo = $util.Long.fromValue(object.dire_team_logo)).unsigned = true;
                    else if (typeof object.dire_team_logo === "string")
                        message.dire_team_logo = parseInt(object.dire_team_logo, 10);
                    else if (typeof object.dire_team_logo === "number")
                        message.dire_team_logo = object.dire_team_logo;
                    else if (typeof object.dire_team_logo === "object")
                        message.dire_team_logo = new $util.LongBits(object.dire_team_logo.low >>> 0, object.dire_team_logo.high >>> 0).toNumber(true);
                if (object.dire_team_logo_url != null)
                    message.dire_team_logo_url = String(object.dire_team_logo_url);
                if (object.activate_time != null) {
                    if (typeof object.activate_time !== "object")
                        throw TypeError(".protocol.Player.Match.activate_time: object expected");
                    message.activate_time = $root.google.protobuf.Timestamp.fromObject(object.activate_time);
                }
                if (object.deactivate_time != null) {
                    if (typeof object.deactivate_time !== "object")
                        throw TypeError(".protocol.Player.Match.deactivate_time: object expected");
                    message.deactivate_time = $root.google.protobuf.Timestamp.fromObject(object.deactivate_time);
                }
                if (object.last_update_time != null) {
                    if (typeof object.last_update_time !== "object")
                        throw TypeError(".protocol.Player.Match.last_update_time: object expected");
                    message.last_update_time = $root.google.protobuf.Timestamp.fromObject(object.last_update_time);
                }
                if (object.start_time != null) {
                    if (typeof object.start_time !== "object")
                        throw TypeError(".protocol.Player.Match.start_time: object expected");
                    message.start_time = $root.google.protobuf.Timestamp.fromObject(object.start_time);
                }
                if (object.series_type != null)
                    message.series_type = object.series_type >>> 0;
                if (object.series_game != null)
                    message.series_game = object.series_game >>> 0;
                if (object.duration != null)
                    message.duration = object.duration >>> 0;
                if (object.radiant_score != null)
                    message.radiant_score = object.radiant_score >>> 0;
                if (object.dire_score != null)
                    message.dire_score = object.dire_score >>> 0;
                switch (object.outcome) {
                case "MATCH_OUTCOME_UNKNOWN":
                case 0:
                    message.outcome = 0;
                    break;
                case "MATCH_OUTCOME_RAD_VICTORY":
                case 2:
                    message.outcome = 2;
                    break;
                case "MATCH_OUTCOME_DIRE_VICTORY":
                case 3:
                    message.outcome = 3;
                    break;
                case "MATCH_OUTCOME_NOT_SCORED_POOR_NETWORK_CONDITIONS":
                case 64:
                    message.outcome = 64;
                    break;
                case "MATCH_OUTCOME_NOT_SCORED_LEAVER":
                case 65:
                    message.outcome = 65;
                    break;
                case "MATCH_OUTCOME_NOT_SCORED_SERVER_CRASH":
                case 66:
                    message.outcome = 66;
                    break;
                case "MATCH_OUTCOME_NOT_SCORED_NEVER_STARTED":
                case 67:
                    message.outcome = 67;
                    break;
                case "MATCH_OUTCOME_NOT_SCORED_CANCELED":
                case 68:
                    message.outcome = 68;
                    break;
                }
                if (object.player_details != null) {
                    if (typeof object.player_details !== "object")
                        throw TypeError(".protocol.Player.Match.player_details: object expected");
                    message.player_details = $root.protocol.Player.Match.PlayerDetails.fromObject(object.player_details);
                }
                return message;
            };

            /**
             * Creates a plain object from a Match message. Also converts values to other types if specified.
             * @function toObject
             * @memberof protocol.Player.Match
             * @static
             * @param {protocol.Player.Match} message Match
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Match.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, true);
                        object.match_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.match_id = options.longs === String ? "0" : 0;
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, true);
                        object.lobby_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.lobby_id = options.longs === String ? "0" : 0;
                    object.lobby_type = options.enums === String ? "LOBBY_TYPE_CASUAL_MATCH" : 0;
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, true);
                        object.league_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.league_id = options.longs === String ? "0" : 0;
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, true);
                        object.series_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.series_id = options.longs === String ? "0" : 0;
                    object.game_mode = options.enums === String ? "GAME_MODE_NONE" : 0;
                    object.average_mmr = 0;
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, true);
                        object.radiant_team_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.radiant_team_id = options.longs === String ? "0" : 0;
                    object.radiant_team_name = "";
                    object.radiant_team_tag = "";
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, true);
                        object.radiant_team_logo = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.radiant_team_logo = options.longs === String ? "0" : 0;
                    object.radiant_team_logo_url = "";
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, true);
                        object.dire_team_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.dire_team_id = options.longs === String ? "0" : 0;
                    object.dire_team_name = "";
                    object.dire_team_tag = "";
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, true);
                        object.dire_team_logo = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.dire_team_logo = options.longs === String ? "0" : 0;
                    object.dire_team_logo_url = "";
                    object.activate_time = null;
                    object.deactivate_time = null;
                    object.last_update_time = null;
                    object.start_time = null;
                    object.series_type = 0;
                    object.series_game = 0;
                    object.duration = 0;
                    object.radiant_score = 0;
                    object.dire_score = 0;
                    object.outcome = options.enums === String ? "MATCH_OUTCOME_UNKNOWN" : 0;
                    object.player_details = null;
                }
                if (message.match_id != null && message.hasOwnProperty("match_id"))
                    if (typeof message.match_id === "number")
                        object.match_id = options.longs === String ? String(message.match_id) : message.match_id;
                    else
                        object.match_id = options.longs === String ? $util.Long.prototype.toString.call(message.match_id) : options.longs === Number ? new $util.LongBits(message.match_id.low >>> 0, message.match_id.high >>> 0).toNumber(true) : message.match_id;
                if (message.lobby_id != null && message.hasOwnProperty("lobby_id"))
                    if (typeof message.lobby_id === "number")
                        object.lobby_id = options.longs === String ? String(message.lobby_id) : message.lobby_id;
                    else
                        object.lobby_id = options.longs === String ? $util.Long.prototype.toString.call(message.lobby_id) : options.longs === Number ? new $util.LongBits(message.lobby_id.low >>> 0, message.lobby_id.high >>> 0).toNumber(true) : message.lobby_id;
                if (message.lobby_type != null && message.hasOwnProperty("lobby_type"))
                    object.lobby_type = options.enums === String ? $root.protocol.LobbyType[message.lobby_type] : message.lobby_type;
                if (message.league_id != null && message.hasOwnProperty("league_id"))
                    if (typeof message.league_id === "number")
                        object.league_id = options.longs === String ? String(message.league_id) : message.league_id;
                    else
                        object.league_id = options.longs === String ? $util.Long.prototype.toString.call(message.league_id) : options.longs === Number ? new $util.LongBits(message.league_id.low >>> 0, message.league_id.high >>> 0).toNumber(true) : message.league_id;
                if (message.series_id != null && message.hasOwnProperty("series_id"))
                    if (typeof message.series_id === "number")
                        object.series_id = options.longs === String ? String(message.series_id) : message.series_id;
                    else
                        object.series_id = options.longs === String ? $util.Long.prototype.toString.call(message.series_id) : options.longs === Number ? new $util.LongBits(message.series_id.low >>> 0, message.series_id.high >>> 0).toNumber(true) : message.series_id;
                if (message.game_mode != null && message.hasOwnProperty("game_mode"))
                    object.game_mode = options.enums === String ? $root.protocol.GameMode[message.game_mode] : message.game_mode;
                if (message.average_mmr != null && message.hasOwnProperty("average_mmr"))
                    object.average_mmr = message.average_mmr;
                if (message.radiant_team_id != null && message.hasOwnProperty("radiant_team_id"))
                    if (typeof message.radiant_team_id === "number")
                        object.radiant_team_id = options.longs === String ? String(message.radiant_team_id) : message.radiant_team_id;
                    else
                        object.radiant_team_id = options.longs === String ? $util.Long.prototype.toString.call(message.radiant_team_id) : options.longs === Number ? new $util.LongBits(message.radiant_team_id.low >>> 0, message.radiant_team_id.high >>> 0).toNumber(true) : message.radiant_team_id;
                if (message.radiant_team_name != null && message.hasOwnProperty("radiant_team_name"))
                    object.radiant_team_name = message.radiant_team_name;
                if (message.radiant_team_tag != null && message.hasOwnProperty("radiant_team_tag"))
                    object.radiant_team_tag = message.radiant_team_tag;
                if (message.radiant_team_logo != null && message.hasOwnProperty("radiant_team_logo"))
                    if (typeof message.radiant_team_logo === "number")
                        object.radiant_team_logo = options.longs === String ? String(message.radiant_team_logo) : message.radiant_team_logo;
                    else
                        object.radiant_team_logo = options.longs === String ? $util.Long.prototype.toString.call(message.radiant_team_logo) : options.longs === Number ? new $util.LongBits(message.radiant_team_logo.low >>> 0, message.radiant_team_logo.high >>> 0).toNumber(true) : message.radiant_team_logo;
                if (message.radiant_team_logo_url != null && message.hasOwnProperty("radiant_team_logo_url"))
                    object.radiant_team_logo_url = message.radiant_team_logo_url;
                if (message.dire_team_id != null && message.hasOwnProperty("dire_team_id"))
                    if (typeof message.dire_team_id === "number")
                        object.dire_team_id = options.longs === String ? String(message.dire_team_id) : message.dire_team_id;
                    else
                        object.dire_team_id = options.longs === String ? $util.Long.prototype.toString.call(message.dire_team_id) : options.longs === Number ? new $util.LongBits(message.dire_team_id.low >>> 0, message.dire_team_id.high >>> 0).toNumber(true) : message.dire_team_id;
                if (message.dire_team_name != null && message.hasOwnProperty("dire_team_name"))
                    object.dire_team_name = message.dire_team_name;
                if (message.dire_team_tag != null && message.hasOwnProperty("dire_team_tag"))
                    object.dire_team_tag = message.dire_team_tag;
                if (message.dire_team_logo != null && message.hasOwnProperty("dire_team_logo"))
                    if (typeof message.dire_team_logo === "number")
                        object.dire_team_logo = options.longs === String ? String(message.dire_team_logo) : message.dire_team_logo;
                    else
                        object.dire_team_logo = options.longs === String ? $util.Long.prototype.toString.call(message.dire_team_logo) : options.longs === Number ? new $util.LongBits(message.dire_team_logo.low >>> 0, message.dire_team_logo.high >>> 0).toNumber(true) : message.dire_team_logo;
                if (message.dire_team_logo_url != null && message.hasOwnProperty("dire_team_logo_url"))
                    object.dire_team_logo_url = message.dire_team_logo_url;
                if (message.activate_time != null && message.hasOwnProperty("activate_time"))
                    object.activate_time = $root.google.protobuf.Timestamp.toObject(message.activate_time, options);
                if (message.deactivate_time != null && message.hasOwnProperty("deactivate_time"))
                    object.deactivate_time = $root.google.protobuf.Timestamp.toObject(message.deactivate_time, options);
                if (message.last_update_time != null && message.hasOwnProperty("last_update_time"))
                    object.last_update_time = $root.google.protobuf.Timestamp.toObject(message.last_update_time, options);
                if (message.start_time != null && message.hasOwnProperty("start_time"))
                    object.start_time = $root.google.protobuf.Timestamp.toObject(message.start_time, options);
                if (message.series_type != null && message.hasOwnProperty("series_type"))
                    object.series_type = message.series_type;
                if (message.series_game != null && message.hasOwnProperty("series_game"))
                    object.series_game = message.series_game;
                if (message.duration != null && message.hasOwnProperty("duration"))
                    object.duration = message.duration;
                if (message.radiant_score != null && message.hasOwnProperty("radiant_score"))
                    object.radiant_score = message.radiant_score;
                if (message.dire_score != null && message.hasOwnProperty("dire_score"))
                    object.dire_score = message.dire_score;
                if (message.outcome != null && message.hasOwnProperty("outcome"))
                    object.outcome = options.enums === String ? $root.protocol.MatchOutcome[message.outcome] : message.outcome;
                if (message.player_details != null && message.hasOwnProperty("player_details"))
                    object.player_details = $root.protocol.Player.Match.PlayerDetails.toObject(message.player_details, options);
                return object;
            };

            /**
             * Converts this Match to JSON.
             * @function toJSON
             * @memberof protocol.Player.Match
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Match.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            Match.PlayerDetails = (function() {

                /**
                 * Properties of a PlayerDetails.
                 * @memberof protocol.Player.Match
                 * @interface IPlayerDetails
                 * @property {Long|null} [hero_id] PlayerDetails hero_id
                 * @property {number|null} [player_slot] PlayerDetails player_slot
                 * @property {string|null} [pro_name] PlayerDetails pro_name
                 * @property {number|null} [kills] PlayerDetails kills
                 * @property {number|null} [deaths] PlayerDetails deaths
                 * @property {number|null} [assists] PlayerDetails assists
                 * @property {Array.<Long>|null} [items] PlayerDetails items
                 */

                /**
                 * Constructs a new PlayerDetails.
                 * @memberof protocol.Player.Match
                 * @classdesc Represents a PlayerDetails.
                 * @implements IPlayerDetails
                 * @constructor
                 * @param {protocol.Player.Match.IPlayerDetails=} [properties] Properties to set
                 */
                function PlayerDetails(properties) {
                    this.items = [];
                    if (properties)
                        for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * PlayerDetails hero_id.
                 * @member {Long} hero_id
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @instance
                 */
                PlayerDetails.prototype.hero_id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

                /**
                 * PlayerDetails player_slot.
                 * @member {number} player_slot
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @instance
                 */
                PlayerDetails.prototype.player_slot = 0;

                /**
                 * PlayerDetails pro_name.
                 * @member {string} pro_name
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @instance
                 */
                PlayerDetails.prototype.pro_name = "";

                /**
                 * PlayerDetails kills.
                 * @member {number} kills
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @instance
                 */
                PlayerDetails.prototype.kills = 0;

                /**
                 * PlayerDetails deaths.
                 * @member {number} deaths
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @instance
                 */
                PlayerDetails.prototype.deaths = 0;

                /**
                 * PlayerDetails assists.
                 * @member {number} assists
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @instance
                 */
                PlayerDetails.prototype.assists = 0;

                /**
                 * PlayerDetails items.
                 * @member {Array.<Long>} items
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @instance
                 */
                PlayerDetails.prototype.items = $util.emptyArray;

                /**
                 * Creates a new PlayerDetails instance using the specified properties.
                 * @function create
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @static
                 * @param {protocol.Player.Match.IPlayerDetails=} [properties] Properties to set
                 * @returns {protocol.Player.Match.PlayerDetails} PlayerDetails instance
                 */
                PlayerDetails.create = function create(properties) {
                    return new PlayerDetails(properties);
                };

                /**
                 * Encodes the specified PlayerDetails message. Does not implicitly {@link protocol.Player.Match.PlayerDetails.verify|verify} messages.
                 * @function encode
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @static
                 * @param {protocol.Player.Match.IPlayerDetails} message PlayerDetails message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                PlayerDetails.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.hero_id != null && message.hasOwnProperty("hero_id"))
                        writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.hero_id);
                    if (message.player_slot != null && message.hasOwnProperty("player_slot"))
                        writer.uint32(/* id 2, wireType 0 =*/16).uint32(message.player_slot);
                    if (message.pro_name != null && message.hasOwnProperty("pro_name"))
                        writer.uint32(/* id 3, wireType 2 =*/26).string(message.pro_name);
                    if (message.kills != null && message.hasOwnProperty("kills"))
                        writer.uint32(/* id 4, wireType 0 =*/32).uint32(message.kills);
                    if (message.deaths != null && message.hasOwnProperty("deaths"))
                        writer.uint32(/* id 5, wireType 0 =*/40).uint32(message.deaths);
                    if (message.assists != null && message.hasOwnProperty("assists"))
                        writer.uint32(/* id 6, wireType 0 =*/48).uint32(message.assists);
                    if (message.items != null && message.items.length) {
                        writer.uint32(/* id 7, wireType 2 =*/58).fork();
                        for (let i = 0; i < message.items.length; ++i)
                            writer.int64(message.items[i]);
                        writer.ldelim();
                    }
                    return writer;
                };

                /**
                 * Encodes the specified PlayerDetails message, length delimited. Does not implicitly {@link protocol.Player.Match.PlayerDetails.verify|verify} messages.
                 * @function encodeDelimited
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @static
                 * @param {protocol.Player.Match.IPlayerDetails} message PlayerDetails message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                PlayerDetails.encodeDelimited = function encodeDelimited(message, writer) {
                    return this.encode(message, writer).ldelim();
                };

                /**
                 * Decodes a PlayerDetails message from the specified reader or buffer.
                 * @function decode
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {protocol.Player.Match.PlayerDetails} PlayerDetails
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                PlayerDetails.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    let end = length === undefined ? reader.len : reader.pos + length, message = new $root.protocol.Player.Match.PlayerDetails();
                    while (reader.pos < end) {
                        let tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.hero_id = reader.uint64();
                            break;
                        case 2:
                            message.player_slot = reader.uint32();
                            break;
                        case 3:
                            message.pro_name = reader.string();
                            break;
                        case 4:
                            message.kills = reader.uint32();
                            break;
                        case 5:
                            message.deaths = reader.uint32();
                            break;
                        case 6:
                            message.assists = reader.uint32();
                            break;
                        case 7:
                            if (!(message.items && message.items.length))
                                message.items = [];
                            if ((tag & 7) === 2) {
                                let end2 = reader.uint32() + reader.pos;
                                while (reader.pos < end2)
                                    message.items.push(reader.int64());
                            } else
                                message.items.push(reader.int64());
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Decodes a PlayerDetails message from the specified reader or buffer, length delimited.
                 * @function decodeDelimited
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @returns {protocol.Player.Match.PlayerDetails} PlayerDetails
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                PlayerDetails.decodeDelimited = function decodeDelimited(reader) {
                    if (!(reader instanceof $Reader))
                        reader = new $Reader(reader);
                    return this.decode(reader, reader.uint32());
                };

                /**
                 * Verifies a PlayerDetails message.
                 * @function verify
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @static
                 * @param {Object.<string,*>} message Plain object to verify
                 * @returns {string|null} `null` if valid, otherwise the reason why it is not
                 */
                PlayerDetails.verify = function verify(message) {
                    if (typeof message !== "object" || message === null)
                        return "object expected";
                    if (message.hero_id != null && message.hasOwnProperty("hero_id"))
                        if (!$util.isInteger(message.hero_id) && !(message.hero_id && $util.isInteger(message.hero_id.low) && $util.isInteger(message.hero_id.high)))
                            return "hero_id: integer|Long expected";
                    if (message.player_slot != null && message.hasOwnProperty("player_slot"))
                        if (!$util.isInteger(message.player_slot))
                            return "player_slot: integer expected";
                    if (message.pro_name != null && message.hasOwnProperty("pro_name"))
                        if (!$util.isString(message.pro_name))
                            return "pro_name: string expected";
                    if (message.kills != null && message.hasOwnProperty("kills"))
                        if (!$util.isInteger(message.kills))
                            return "kills: integer expected";
                    if (message.deaths != null && message.hasOwnProperty("deaths"))
                        if (!$util.isInteger(message.deaths))
                            return "deaths: integer expected";
                    if (message.assists != null && message.hasOwnProperty("assists"))
                        if (!$util.isInteger(message.assists))
                            return "assists: integer expected";
                    if (message.items != null && message.hasOwnProperty("items")) {
                        if (!Array.isArray(message.items))
                            return "items: array expected";
                        for (let i = 0; i < message.items.length; ++i)
                            if (!$util.isInteger(message.items[i]) && !(message.items[i] && $util.isInteger(message.items[i].low) && $util.isInteger(message.items[i].high)))
                                return "items: integer|Long[] expected";
                    }
                    return null;
                };

                /**
                 * Creates a PlayerDetails message from a plain object. Also converts values to their respective internal types.
                 * @function fromObject
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @static
                 * @param {Object.<string,*>} object Plain object
                 * @returns {protocol.Player.Match.PlayerDetails} PlayerDetails
                 */
                PlayerDetails.fromObject = function fromObject(object) {
                    if (object instanceof $root.protocol.Player.Match.PlayerDetails)
                        return object;
                    let message = new $root.protocol.Player.Match.PlayerDetails();
                    if (object.hero_id != null)
                        if ($util.Long)
                            (message.hero_id = $util.Long.fromValue(object.hero_id)).unsigned = true;
                        else if (typeof object.hero_id === "string")
                            message.hero_id = parseInt(object.hero_id, 10);
                        else if (typeof object.hero_id === "number")
                            message.hero_id = object.hero_id;
                        else if (typeof object.hero_id === "object")
                            message.hero_id = new $util.LongBits(object.hero_id.low >>> 0, object.hero_id.high >>> 0).toNumber(true);
                    if (object.player_slot != null)
                        message.player_slot = object.player_slot >>> 0;
                    if (object.pro_name != null)
                        message.pro_name = String(object.pro_name);
                    if (object.kills != null)
                        message.kills = object.kills >>> 0;
                    if (object.deaths != null)
                        message.deaths = object.deaths >>> 0;
                    if (object.assists != null)
                        message.assists = object.assists >>> 0;
                    if (object.items) {
                        if (!Array.isArray(object.items))
                            throw TypeError(".protocol.Player.Match.PlayerDetails.items: array expected");
                        message.items = [];
                        for (let i = 0; i < object.items.length; ++i)
                            if ($util.Long)
                                (message.items[i] = $util.Long.fromValue(object.items[i])).unsigned = false;
                            else if (typeof object.items[i] === "string")
                                message.items[i] = parseInt(object.items[i], 10);
                            else if (typeof object.items[i] === "number")
                                message.items[i] = object.items[i];
                            else if (typeof object.items[i] === "object")
                                message.items[i] = new $util.LongBits(object.items[i].low >>> 0, object.items[i].high >>> 0).toNumber();
                    }
                    return message;
                };

                /**
                 * Creates a plain object from a PlayerDetails message. Also converts values to other types if specified.
                 * @function toObject
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @static
                 * @param {protocol.Player.Match.PlayerDetails} message PlayerDetails
                 * @param {$protobuf.IConversionOptions} [options] Conversion options
                 * @returns {Object.<string,*>} Plain object
                 */
                PlayerDetails.toObject = function toObject(message, options) {
                    if (!options)
                        options = {};
                    let object = {};
                    if (options.arrays || options.defaults)
                        object.items = [];
                    if (options.defaults) {
                        if ($util.Long) {
                            let long = new $util.Long(0, 0, true);
                            object.hero_id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                        } else
                            object.hero_id = options.longs === String ? "0" : 0;
                        object.player_slot = 0;
                        object.pro_name = "";
                        object.kills = 0;
                        object.deaths = 0;
                        object.assists = 0;
                    }
                    if (message.hero_id != null && message.hasOwnProperty("hero_id"))
                        if (typeof message.hero_id === "number")
                            object.hero_id = options.longs === String ? String(message.hero_id) : message.hero_id;
                        else
                            object.hero_id = options.longs === String ? $util.Long.prototype.toString.call(message.hero_id) : options.longs === Number ? new $util.LongBits(message.hero_id.low >>> 0, message.hero_id.high >>> 0).toNumber(true) : message.hero_id;
                    if (message.player_slot != null && message.hasOwnProperty("player_slot"))
                        object.player_slot = message.player_slot;
                    if (message.pro_name != null && message.hasOwnProperty("pro_name"))
                        object.pro_name = message.pro_name;
                    if (message.kills != null && message.hasOwnProperty("kills"))
                        object.kills = message.kills;
                    if (message.deaths != null && message.hasOwnProperty("deaths"))
                        object.deaths = message.deaths;
                    if (message.assists != null && message.hasOwnProperty("assists"))
                        object.assists = message.assists;
                    if (message.items && message.items.length) {
                        object.items = [];
                        for (let j = 0; j < message.items.length; ++j)
                            if (typeof message.items[j] === "number")
                                object.items[j] = options.longs === String ? String(message.items[j]) : message.items[j];
                            else
                                object.items[j] = options.longs === String ? $util.Long.prototype.toString.call(message.items[j]) : options.longs === Number ? new $util.LongBits(message.items[j].low >>> 0, message.items[j].high >>> 0).toNumber() : message.items[j];
                    }
                    return object;
                };

                /**
                 * Converts this PlayerDetails to JSON.
                 * @function toJSON
                 * @memberof protocol.Player.Match.PlayerDetails
                 * @instance
                 * @returns {Object.<string,*>} JSON object
                 */
                PlayerDetails.prototype.toJSON = function toJSON() {
                    return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                };

                return PlayerDetails;
            })();

            return Match;
        })();

        Player.Team = (function() {

            /**
             * Properties of a Team.
             * @memberof protocol.Player
             * @interface ITeam
             * @property {Long|null} [id] Team id
             * @property {string|null} [name] Team name
             * @property {string|null} [tag] Team tag
             * @property {string|null} [logo_url] Team logo_url
             */

            /**
             * Constructs a new Team.
             * @memberof protocol.Player
             * @classdesc Represents a Team.
             * @implements ITeam
             * @constructor
             * @param {protocol.Player.ITeam=} [properties] Properties to set
             */
            function Team(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Team id.
             * @member {Long} id
             * @memberof protocol.Player.Team
             * @instance
             */
            Team.prototype.id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

            /**
             * Team name.
             * @member {string} name
             * @memberof protocol.Player.Team
             * @instance
             */
            Team.prototype.name = "";

            /**
             * Team tag.
             * @member {string} tag
             * @memberof protocol.Player.Team
             * @instance
             */
            Team.prototype.tag = "";

            /**
             * Team logo_url.
             * @member {string} logo_url
             * @memberof protocol.Player.Team
             * @instance
             */
            Team.prototype.logo_url = "";

            /**
             * Creates a new Team instance using the specified properties.
             * @function create
             * @memberof protocol.Player.Team
             * @static
             * @param {protocol.Player.ITeam=} [properties] Properties to set
             * @returns {protocol.Player.Team} Team instance
             */
            Team.create = function create(properties) {
                return new Team(properties);
            };

            /**
             * Encodes the specified Team message. Does not implicitly {@link protocol.Player.Team.verify|verify} messages.
             * @function encode
             * @memberof protocol.Player.Team
             * @static
             * @param {protocol.Player.ITeam} message Team message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Team.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.id != null && message.hasOwnProperty("id"))
                    writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.id);
                if (message.name != null && message.hasOwnProperty("name"))
                    writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
                if (message.tag != null && message.hasOwnProperty("tag"))
                    writer.uint32(/* id 3, wireType 2 =*/26).string(message.tag);
                if (message.logo_url != null && message.hasOwnProperty("logo_url"))
                    writer.uint32(/* id 4, wireType 2 =*/34).string(message.logo_url);
                return writer;
            };

            /**
             * Encodes the specified Team message, length delimited. Does not implicitly {@link protocol.Player.Team.verify|verify} messages.
             * @function encodeDelimited
             * @memberof protocol.Player.Team
             * @static
             * @param {protocol.Player.ITeam} message Team message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Team.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a Team message from the specified reader or buffer.
             * @function decode
             * @memberof protocol.Player.Team
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {protocol.Player.Team} Team
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Team.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.protocol.Player.Team();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.id = reader.uint64();
                        break;
                    case 2:
                        message.name = reader.string();
                        break;
                    case 3:
                        message.tag = reader.string();
                        break;
                    case 4:
                        message.logo_url = reader.string();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a Team message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof protocol.Player.Team
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {protocol.Player.Team} Team
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Team.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a Team message.
             * @function verify
             * @memberof protocol.Player.Team
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Team.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.id != null && message.hasOwnProperty("id"))
                    if (!$util.isInteger(message.id) && !(message.id && $util.isInteger(message.id.low) && $util.isInteger(message.id.high)))
                        return "id: integer|Long expected";
                if (message.name != null && message.hasOwnProperty("name"))
                    if (!$util.isString(message.name))
                        return "name: string expected";
                if (message.tag != null && message.hasOwnProperty("tag"))
                    if (!$util.isString(message.tag))
                        return "tag: string expected";
                if (message.logo_url != null && message.hasOwnProperty("logo_url"))
                    if (!$util.isString(message.logo_url))
                        return "logo_url: string expected";
                return null;
            };

            /**
             * Creates a Team message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof protocol.Player.Team
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {protocol.Player.Team} Team
             */
            Team.fromObject = function fromObject(object) {
                if (object instanceof $root.protocol.Player.Team)
                    return object;
                let message = new $root.protocol.Player.Team();
                if (object.id != null)
                    if ($util.Long)
                        (message.id = $util.Long.fromValue(object.id)).unsigned = true;
                    else if (typeof object.id === "string")
                        message.id = parseInt(object.id, 10);
                    else if (typeof object.id === "number")
                        message.id = object.id;
                    else if (typeof object.id === "object")
                        message.id = new $util.LongBits(object.id.low >>> 0, object.id.high >>> 0).toNumber(true);
                if (object.name != null)
                    message.name = String(object.name);
                if (object.tag != null)
                    message.tag = String(object.tag);
                if (object.logo_url != null)
                    message.logo_url = String(object.logo_url);
                return message;
            };

            /**
             * Creates a plain object from a Team message. Also converts values to other types if specified.
             * @function toObject
             * @memberof protocol.Player.Team
             * @static
             * @param {protocol.Player.Team} message Team
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Team.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, true);
                        object.id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.id = options.longs === String ? "0" : 0;
                    object.name = "";
                    object.tag = "";
                    object.logo_url = "";
                }
                if (message.id != null && message.hasOwnProperty("id"))
                    if (typeof message.id === "number")
                        object.id = options.longs === String ? String(message.id) : message.id;
                    else
                        object.id = options.longs === String ? $util.Long.prototype.toString.call(message.id) : options.longs === Number ? new $util.LongBits(message.id.low >>> 0, message.id.high >>> 0).toNumber(true) : message.id;
                if (message.name != null && message.hasOwnProperty("name"))
                    object.name = message.name;
                if (message.tag != null && message.hasOwnProperty("tag"))
                    object.tag = message.tag;
                if (message.logo_url != null && message.hasOwnProperty("logo_url"))
                    object.logo_url = message.logo_url;
                return object;
            };

            /**
             * Converts this Team to JSON.
             * @function toJSON
             * @memberof protocol.Player.Team
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Team.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return Team;
        })();

        return Player;
    })();

    /**
     * CollectionOp enum.
     * @name protocol.CollectionOp
     * @enum {string}
     * @property {number} REPLACE=0 REPLACE value
     * @property {number} ADD=1 ADD value
     * @property {number} UPDATE=2 UPDATE value
     * @property {number} REMOVE=3 REMOVE value
     */
    protocol.CollectionOp = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "REPLACE"] = 0;
        values[valuesById[1] = "ADD"] = 1;
        values[valuesById[2] = "UPDATE"] = 2;
        values[valuesById[3] = "REMOVE"] = 3;
        return values;
    })();

    protocol.Hero = (function() {

        /**
         * Properties of a Hero.
         * @memberof protocol
         * @interface IHero
         * @property {Long|null} [id] Hero id
         * @property {string|null} [name] Hero name
         * @property {string|null} [localized_name] Hero localized_name
         * @property {string|null} [image_full_url] Hero image_full_url
         * @property {string|null} [image_large_url] Hero image_large_url
         * @property {string|null} [image_small_url] Hero image_small_url
         * @property {string|null} [image_portrait_url] Hero image_portrait_url
         */

        /**
         * Constructs a new Hero.
         * @memberof protocol
         * @classdesc Represents a Hero.
         * @implements IHero
         * @constructor
         * @param {protocol.IHero=} [properties] Properties to set
         */
        function Hero(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Hero id.
         * @member {Long} id
         * @memberof protocol.Hero
         * @instance
         */
        Hero.prototype.id = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * Hero name.
         * @member {string} name
         * @memberof protocol.Hero
         * @instance
         */
        Hero.prototype.name = "";

        /**
         * Hero localized_name.
         * @member {string} localized_name
         * @memberof protocol.Hero
         * @instance
         */
        Hero.prototype.localized_name = "";

        /**
         * Hero image_full_url.
         * @member {string} image_full_url
         * @memberof protocol.Hero
         * @instance
         */
        Hero.prototype.image_full_url = "";

        /**
         * Hero image_large_url.
         * @member {string} image_large_url
         * @memberof protocol.Hero
         * @instance
         */
        Hero.prototype.image_large_url = "";

        /**
         * Hero image_small_url.
         * @member {string} image_small_url
         * @memberof protocol.Hero
         * @instance
         */
        Hero.prototype.image_small_url = "";

        /**
         * Hero image_portrait_url.
         * @member {string} image_portrait_url
         * @memberof protocol.Hero
         * @instance
         */
        Hero.prototype.image_portrait_url = "";

        /**
         * Creates a new Hero instance using the specified properties.
         * @function create
         * @memberof protocol.Hero
         * @static
         * @param {protocol.IHero=} [properties] Properties to set
         * @returns {protocol.Hero} Hero instance
         */
        Hero.create = function create(properties) {
            return new Hero(properties);
        };

        /**
         * Encodes the specified Hero message. Does not implicitly {@link protocol.Hero.verify|verify} messages.
         * @function encode
         * @memberof protocol.Hero
         * @static
         * @param {protocol.IHero} message Hero message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Hero.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.id != null && message.hasOwnProperty("id"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.id);
            if (message.name != null && message.hasOwnProperty("name"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
            if (message.localized_name != null && message.hasOwnProperty("localized_name"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.localized_name);
            if (message.image_full_url != null && message.hasOwnProperty("image_full_url"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.image_full_url);
            if (message.image_large_url != null && message.hasOwnProperty("image_large_url"))
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.image_large_url);
            if (message.image_small_url != null && message.hasOwnProperty("image_small_url"))
                writer.uint32(/* id 6, wireType 2 =*/50).string(message.image_small_url);
            if (message.image_portrait_url != null && message.hasOwnProperty("image_portrait_url"))
                writer.uint32(/* id 7, wireType 2 =*/58).string(message.image_portrait_url);
            return writer;
        };

        /**
         * Encodes the specified Hero message, length delimited. Does not implicitly {@link protocol.Hero.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protocol.Hero
         * @static
         * @param {protocol.IHero} message Hero message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Hero.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Hero message from the specified reader or buffer.
         * @function decode
         * @memberof protocol.Hero
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protocol.Hero} Hero
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Hero.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.protocol.Hero();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.id = reader.uint64();
                    break;
                case 2:
                    message.name = reader.string();
                    break;
                case 3:
                    message.localized_name = reader.string();
                    break;
                case 4:
                    message.image_full_url = reader.string();
                    break;
                case 5:
                    message.image_large_url = reader.string();
                    break;
                case 6:
                    message.image_small_url = reader.string();
                    break;
                case 7:
                    message.image_portrait_url = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Hero message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protocol.Hero
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protocol.Hero} Hero
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Hero.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Hero message.
         * @function verify
         * @memberof protocol.Hero
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Hero.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isInteger(message.id) && !(message.id && $util.isInteger(message.id.low) && $util.isInteger(message.id.high)))
                    return "id: integer|Long expected";
            if (message.name != null && message.hasOwnProperty("name"))
                if (!$util.isString(message.name))
                    return "name: string expected";
            if (message.localized_name != null && message.hasOwnProperty("localized_name"))
                if (!$util.isString(message.localized_name))
                    return "localized_name: string expected";
            if (message.image_full_url != null && message.hasOwnProperty("image_full_url"))
                if (!$util.isString(message.image_full_url))
                    return "image_full_url: string expected";
            if (message.image_large_url != null && message.hasOwnProperty("image_large_url"))
                if (!$util.isString(message.image_large_url))
                    return "image_large_url: string expected";
            if (message.image_small_url != null && message.hasOwnProperty("image_small_url"))
                if (!$util.isString(message.image_small_url))
                    return "image_small_url: string expected";
            if (message.image_portrait_url != null && message.hasOwnProperty("image_portrait_url"))
                if (!$util.isString(message.image_portrait_url))
                    return "image_portrait_url: string expected";
            return null;
        };

        /**
         * Creates a Hero message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protocol.Hero
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protocol.Hero} Hero
         */
        Hero.fromObject = function fromObject(object) {
            if (object instanceof $root.protocol.Hero)
                return object;
            let message = new $root.protocol.Hero();
            if (object.id != null)
                if ($util.Long)
                    (message.id = $util.Long.fromValue(object.id)).unsigned = true;
                else if (typeof object.id === "string")
                    message.id = parseInt(object.id, 10);
                else if (typeof object.id === "number")
                    message.id = object.id;
                else if (typeof object.id === "object")
                    message.id = new $util.LongBits(object.id.low >>> 0, object.id.high >>> 0).toNumber(true);
            if (object.name != null)
                message.name = String(object.name);
            if (object.localized_name != null)
                message.localized_name = String(object.localized_name);
            if (object.image_full_url != null)
                message.image_full_url = String(object.image_full_url);
            if (object.image_large_url != null)
                message.image_large_url = String(object.image_large_url);
            if (object.image_small_url != null)
                message.image_small_url = String(object.image_small_url);
            if (object.image_portrait_url != null)
                message.image_portrait_url = String(object.image_portrait_url);
            return message;
        };

        /**
         * Creates a plain object from a Hero message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protocol.Hero
         * @static
         * @param {protocol.Hero} message Hero
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Hero.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                if ($util.Long) {
                    let long = new $util.Long(0, 0, true);
                    object.id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.id = options.longs === String ? "0" : 0;
                object.name = "";
                object.localized_name = "";
                object.image_full_url = "";
                object.image_large_url = "";
                object.image_small_url = "";
                object.image_portrait_url = "";
            }
            if (message.id != null && message.hasOwnProperty("id"))
                if (typeof message.id === "number")
                    object.id = options.longs === String ? String(message.id) : message.id;
                else
                    object.id = options.longs === String ? $util.Long.prototype.toString.call(message.id) : options.longs === Number ? new $util.LongBits(message.id.low >>> 0, message.id.high >>> 0).toNumber(true) : message.id;
            if (message.name != null && message.hasOwnProperty("name"))
                object.name = message.name;
            if (message.localized_name != null && message.hasOwnProperty("localized_name"))
                object.localized_name = message.localized_name;
            if (message.image_full_url != null && message.hasOwnProperty("image_full_url"))
                object.image_full_url = message.image_full_url;
            if (message.image_large_url != null && message.hasOwnProperty("image_large_url"))
                object.image_large_url = message.image_large_url;
            if (message.image_small_url != null && message.hasOwnProperty("image_small_url"))
                object.image_small_url = message.image_small_url;
            if (message.image_portrait_url != null && message.hasOwnProperty("image_portrait_url"))
                object.image_portrait_url = message.image_portrait_url;
            return object;
        };

        /**
         * Converts this Hero to JSON.
         * @function toJSON
         * @memberof protocol.Hero
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Hero.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Hero;
    })();

    protocol.LiveMatches = (function() {

        /**
         * Properties of a LiveMatches.
         * @memberof protocol
         * @interface ILiveMatches
         * @property {Array.<protocol.ILiveMatch>|null} [matches] LiveMatches matches
         */

        /**
         * Constructs a new LiveMatches.
         * @memberof protocol
         * @classdesc Represents a LiveMatches.
         * @implements ILiveMatches
         * @constructor
         * @param {protocol.ILiveMatches=} [properties] Properties to set
         */
        function LiveMatches(properties) {
            this.matches = [];
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * LiveMatches matches.
         * @member {Array.<protocol.ILiveMatch>} matches
         * @memberof protocol.LiveMatches
         * @instance
         */
        LiveMatches.prototype.matches = $util.emptyArray;

        /**
         * Creates a new LiveMatches instance using the specified properties.
         * @function create
         * @memberof protocol.LiveMatches
         * @static
         * @param {protocol.ILiveMatches=} [properties] Properties to set
         * @returns {protocol.LiveMatches} LiveMatches instance
         */
        LiveMatches.create = function create(properties) {
            return new LiveMatches(properties);
        };

        /**
         * Encodes the specified LiveMatches message. Does not implicitly {@link protocol.LiveMatches.verify|verify} messages.
         * @function encode
         * @memberof protocol.LiveMatches
         * @static
         * @param {protocol.ILiveMatches} message LiveMatches message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        LiveMatches.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.matches != null && message.matches.length)
                for (let i = 0; i < message.matches.length; ++i)
                    $root.protocol.LiveMatch.encode(message.matches[i], writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified LiveMatches message, length delimited. Does not implicitly {@link protocol.LiveMatches.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protocol.LiveMatches
         * @static
         * @param {protocol.ILiveMatches} message LiveMatches message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        LiveMatches.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a LiveMatches message from the specified reader or buffer.
         * @function decode
         * @memberof protocol.LiveMatches
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protocol.LiveMatches} LiveMatches
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        LiveMatches.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.protocol.LiveMatches();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    if (!(message.matches && message.matches.length))
                        message.matches = [];
                    message.matches.push($root.protocol.LiveMatch.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a LiveMatches message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protocol.LiveMatches
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protocol.LiveMatches} LiveMatches
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        LiveMatches.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a LiveMatches message.
         * @function verify
         * @memberof protocol.LiveMatches
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        LiveMatches.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.matches != null && message.hasOwnProperty("matches")) {
                if (!Array.isArray(message.matches))
                    return "matches: array expected";
                for (let i = 0; i < message.matches.length; ++i) {
                    let error = $root.protocol.LiveMatch.verify(message.matches[i]);
                    if (error)
                        return "matches." + error;
                }
            }
            return null;
        };

        /**
         * Creates a LiveMatches message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protocol.LiveMatches
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protocol.LiveMatches} LiveMatches
         */
        LiveMatches.fromObject = function fromObject(object) {
            if (object instanceof $root.protocol.LiveMatches)
                return object;
            let message = new $root.protocol.LiveMatches();
            if (object.matches) {
                if (!Array.isArray(object.matches))
                    throw TypeError(".protocol.LiveMatches.matches: array expected");
                message.matches = [];
                for (let i = 0; i < object.matches.length; ++i) {
                    if (typeof object.matches[i] !== "object")
                        throw TypeError(".protocol.LiveMatches.matches: object expected");
                    message.matches[i] = $root.protocol.LiveMatch.fromObject(object.matches[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a LiveMatches message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protocol.LiveMatches
         * @static
         * @param {protocol.LiveMatches} message LiveMatches
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        LiveMatches.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.arrays || options.defaults)
                object.matches = [];
            if (message.matches && message.matches.length) {
                object.matches = [];
                for (let j = 0; j < message.matches.length; ++j)
                    object.matches[j] = $root.protocol.LiveMatch.toObject(message.matches[j], options);
            }
            return object;
        };

        /**
         * Converts this LiveMatches to JSON.
         * @function toJSON
         * @memberof protocol.LiveMatches
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        LiveMatches.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return LiveMatches;
    })();

    protocol.LiveMatchesChange = (function() {

        /**
         * Properties of a LiveMatchesChange.
         * @memberof protocol
         * @interface ILiveMatchesChange
         * @property {protocol.CollectionOp|null} [op] LiveMatchesChange op
         * @property {protocol.ILiveMatches|null} [change] LiveMatchesChange change
         */

        /**
         * Constructs a new LiveMatchesChange.
         * @memberof protocol
         * @classdesc Represents a LiveMatchesChange.
         * @implements ILiveMatchesChange
         * @constructor
         * @param {protocol.ILiveMatchesChange=} [properties] Properties to set
         */
        function LiveMatchesChange(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * LiveMatchesChange op.
         * @member {protocol.CollectionOp} op
         * @memberof protocol.LiveMatchesChange
         * @instance
         */
        LiveMatchesChange.prototype.op = 0;

        /**
         * LiveMatchesChange change.
         * @member {protocol.ILiveMatches|null|undefined} change
         * @memberof protocol.LiveMatchesChange
         * @instance
         */
        LiveMatchesChange.prototype.change = null;

        /**
         * Creates a new LiveMatchesChange instance using the specified properties.
         * @function create
         * @memberof protocol.LiveMatchesChange
         * @static
         * @param {protocol.ILiveMatchesChange=} [properties] Properties to set
         * @returns {protocol.LiveMatchesChange} LiveMatchesChange instance
         */
        LiveMatchesChange.create = function create(properties) {
            return new LiveMatchesChange(properties);
        };

        /**
         * Encodes the specified LiveMatchesChange message. Does not implicitly {@link protocol.LiveMatchesChange.verify|verify} messages.
         * @function encode
         * @memberof protocol.LiveMatchesChange
         * @static
         * @param {protocol.ILiveMatchesChange} message LiveMatchesChange message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        LiveMatchesChange.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.op != null && message.hasOwnProperty("op"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.op);
            if (message.change != null && message.hasOwnProperty("change"))
                $root.protocol.LiveMatches.encode(message.change, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified LiveMatchesChange message, length delimited. Does not implicitly {@link protocol.LiveMatchesChange.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protocol.LiveMatchesChange
         * @static
         * @param {protocol.ILiveMatchesChange} message LiveMatchesChange message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        LiveMatchesChange.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a LiveMatchesChange message from the specified reader or buffer.
         * @function decode
         * @memberof protocol.LiveMatchesChange
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protocol.LiveMatchesChange} LiveMatchesChange
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        LiveMatchesChange.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.protocol.LiveMatchesChange();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.op = reader.int32();
                    break;
                case 2:
                    message.change = $root.protocol.LiveMatches.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a LiveMatchesChange message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protocol.LiveMatchesChange
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protocol.LiveMatchesChange} LiveMatchesChange
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        LiveMatchesChange.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a LiveMatchesChange message.
         * @function verify
         * @memberof protocol.LiveMatchesChange
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        LiveMatchesChange.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.op != null && message.hasOwnProperty("op"))
                switch (message.op) {
                default:
                    return "op: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                    break;
                }
            if (message.change != null && message.hasOwnProperty("change")) {
                let error = $root.protocol.LiveMatches.verify(message.change);
                if (error)
                    return "change." + error;
            }
            return null;
        };

        /**
         * Creates a LiveMatchesChange message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protocol.LiveMatchesChange
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protocol.LiveMatchesChange} LiveMatchesChange
         */
        LiveMatchesChange.fromObject = function fromObject(object) {
            if (object instanceof $root.protocol.LiveMatchesChange)
                return object;
            let message = new $root.protocol.LiveMatchesChange();
            switch (object.op) {
            case "REPLACE":
            case 0:
                message.op = 0;
                break;
            case "ADD":
            case 1:
                message.op = 1;
                break;
            case "UPDATE":
            case 2:
                message.op = 2;
                break;
            case "REMOVE":
            case 3:
                message.op = 3;
                break;
            }
            if (object.change != null) {
                if (typeof object.change !== "object")
                    throw TypeError(".protocol.LiveMatchesChange.change: object expected");
                message.change = $root.protocol.LiveMatches.fromObject(object.change);
            }
            return message;
        };

        /**
         * Creates a plain object from a LiveMatchesChange message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protocol.LiveMatchesChange
         * @static
         * @param {protocol.LiveMatchesChange} message LiveMatchesChange
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        LiveMatchesChange.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.op = options.enums === String ? "REPLACE" : 0;
                object.change = null;
            }
            if (message.op != null && message.hasOwnProperty("op"))
                object.op = options.enums === String ? $root.protocol.CollectionOp[message.op] : message.op;
            if (message.change != null && message.hasOwnProperty("change"))
                object.change = $root.protocol.LiveMatches.toObject(message.change, options);
            return object;
        };

        /**
         * Converts this LiveMatchesChange to JSON.
         * @function toJSON
         * @memberof protocol.LiveMatchesChange
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        LiveMatchesChange.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return LiveMatchesChange;
    })();

    return protocol;
})();

export const google = $root.google = (() => {

    /**
     * Namespace google.
     * @exports google
     * @namespace
     */
    const google = {};

    google.protobuf = (function() {

        /**
         * Namespace protobuf.
         * @memberof google
         * @namespace
         */
        const protobuf = {};

        protobuf.Timestamp = (function() {

            /**
             * Properties of a Timestamp.
             * @memberof google.protobuf
             * @interface ITimestamp
             * @property {Long|null} [seconds] Timestamp seconds
             * @property {number|null} [nanos] Timestamp nanos
             */

            /**
             * Constructs a new Timestamp.
             * @memberof google.protobuf
             * @classdesc Represents a Timestamp.
             * @implements ITimestamp
             * @constructor
             * @param {google.protobuf.ITimestamp=} [properties] Properties to set
             */
            function Timestamp(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Timestamp seconds.
             * @member {Long} seconds
             * @memberof google.protobuf.Timestamp
             * @instance
             */
            Timestamp.prototype.seconds = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

            /**
             * Timestamp nanos.
             * @member {number} nanos
             * @memberof google.protobuf.Timestamp
             * @instance
             */
            Timestamp.prototype.nanos = 0;

            /**
             * Creates a new Timestamp instance using the specified properties.
             * @function create
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.ITimestamp=} [properties] Properties to set
             * @returns {google.protobuf.Timestamp} Timestamp instance
             */
            Timestamp.create = function create(properties) {
                return new Timestamp(properties);
            };

            /**
             * Encodes the specified Timestamp message. Does not implicitly {@link google.protobuf.Timestamp.verify|verify} messages.
             * @function encode
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.ITimestamp} message Timestamp message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Timestamp.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.seconds != null && message.hasOwnProperty("seconds"))
                    writer.uint32(/* id 1, wireType 0 =*/8).int64(message.seconds);
                if (message.nanos != null && message.hasOwnProperty("nanos"))
                    writer.uint32(/* id 2, wireType 0 =*/16).int32(message.nanos);
                return writer;
            };

            /**
             * Encodes the specified Timestamp message, length delimited. Does not implicitly {@link google.protobuf.Timestamp.verify|verify} messages.
             * @function encodeDelimited
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.ITimestamp} message Timestamp message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Timestamp.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a Timestamp message from the specified reader or buffer.
             * @function decode
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {google.protobuf.Timestamp} Timestamp
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Timestamp.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.google.protobuf.Timestamp();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.seconds = reader.int64();
                        break;
                    case 2:
                        message.nanos = reader.int32();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a Timestamp message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {google.protobuf.Timestamp} Timestamp
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Timestamp.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a Timestamp message.
             * @function verify
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Timestamp.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.seconds != null && message.hasOwnProperty("seconds"))
                    if (!$util.isInteger(message.seconds) && !(message.seconds && $util.isInteger(message.seconds.low) && $util.isInteger(message.seconds.high)))
                        return "seconds: integer|Long expected";
                if (message.nanos != null && message.hasOwnProperty("nanos"))
                    if (!$util.isInteger(message.nanos))
                        return "nanos: integer expected";
                return null;
            };

            /**
             * Creates a Timestamp message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {google.protobuf.Timestamp} Timestamp
             */
            Timestamp.fromObject = function fromObject(object) {
                if (object instanceof $root.google.protobuf.Timestamp)
                    return object;
                let message = new $root.google.protobuf.Timestamp();
                if (object.seconds != null)
                    if ($util.Long)
                        (message.seconds = $util.Long.fromValue(object.seconds)).unsigned = false;
                    else if (typeof object.seconds === "string")
                        message.seconds = parseInt(object.seconds, 10);
                    else if (typeof object.seconds === "number")
                        message.seconds = object.seconds;
                    else if (typeof object.seconds === "object")
                        message.seconds = new $util.LongBits(object.seconds.low >>> 0, object.seconds.high >>> 0).toNumber();
                if (object.nanos != null)
                    message.nanos = object.nanos | 0;
                return message;
            };

            /**
             * Creates a plain object from a Timestamp message. Also converts values to other types if specified.
             * @function toObject
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.Timestamp} message Timestamp
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Timestamp.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, false);
                        object.seconds = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.seconds = options.longs === String ? "0" : 0;
                    object.nanos = 0;
                }
                if (message.seconds != null && message.hasOwnProperty("seconds"))
                    if (typeof message.seconds === "number")
                        object.seconds = options.longs === String ? String(message.seconds) : message.seconds;
                    else
                        object.seconds = options.longs === String ? $util.Long.prototype.toString.call(message.seconds) : options.longs === Number ? new $util.LongBits(message.seconds.low >>> 0, message.seconds.high >>> 0).toNumber() : message.seconds;
                if (message.nanos != null && message.hasOwnProperty("nanos"))
                    object.nanos = message.nanos;
                return object;
            };

            /**
             * Converts this Timestamp to JSON.
             * @function toJSON
             * @memberof google.protobuf.Timestamp
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Timestamp.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return Timestamp;
        })();

        return protobuf;
    })();

    return google;
})();

export { $root as default };
