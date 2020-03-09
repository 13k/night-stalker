import _ from "lodash";

import pb from "@/protocol/proto";
import { PlayerSlot } from "@/dota2/flags";

import { set, transformProperty } from "./object";
import { transformMatchTimestamps } from "./match_common";

const TEAM_ATTRIBUTE_NAME_TPL = _.template("<%= side %>_team_<%= attr %>");
const TEAM_ATTRIBUTES = ["id", "name", "tag", "logo", "logo_url"];
const TEAM_SIDES = {
  [pb.protocol.GameTeam.GAME_TEAM_GOODGUYS]: "radiant",
  [pb.protocol.GameTeam.GAME_TEAM_BADGUYS]: "dire",
};

function createLiveMatchTeamAttributes(liveMatch, side) {
  return _.transform(
    TEAM_ATTRIBUTES,
    (attrs, attr) => {
      attrs[attr] = liveMatch[TEAM_ATTRIBUTE_NAME_TPL({ side, attr })];
    },
    {}
  );
}

function createLiveMatchTeam({ liveMatch, number, players }) {
  number = _.toInteger(number);
  const side = TEAM_SIDES[number];
  const attrs = side != null ? createLiveMatchTeamAttributes(liveMatch, side) : {};

  return {
    number,
    side,
    players,
    ...attrs,
  };
}

export function transformLiveMatch(liveMatch, state) {
  transformMatchTimestamps(liveMatch);

  _.each(liveMatch.players, player => transformLiveMatchPlayer(player, state));

  liveMatch.players = _.sortBy(liveMatch.players, "player_slot");

  const teams = _.chain(liveMatch.players)
    .groupBy("team")
    .transform((teams, players, number) => {
      teams[number] = createLiveMatchTeam({ liveMatch, number, players });
    })
    .value();

  set(liveMatch, "teams", teams);

  return liveMatch;
}

export function transformLiveMatchPlayer(player, { heroes }) {
  transformProperty(player, "hero_id", "hero", heroId => _.get(heroes, ["byId", heroId]));
  transformProperty(player, "player_slot", "slot", slot => new PlayerSlot(slot));

  return player;
}

export function transformLiveMatches(liveMatches, state) {
  _.each(liveMatches.matches, liveMatch => transformLiveMatch(liveMatch, state));

  return liveMatches;
}

export function transformLiveMatchesChange(liveMatchesChange, state) {
  transformLiveMatches(liveMatchesChange.change, state);

  return liveMatchesChange;
}
