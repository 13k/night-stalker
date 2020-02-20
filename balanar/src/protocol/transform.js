import _ from "lodash";

import pb from "@/protocol/proto";
import { PlayerSlot } from "@/dota2/flags";
import { MATCH_TIMESTAMP_FIELDS } from "@/protocol/preprocess";

const TEAM_ATTRIBUTE_NAME_TPL = _.template("<%= side %>_team_<%= attr %>");
const TEAM_ATTRIBUTES = ["id", "name", "tag", "logo", "logo_url"];
const TEAM_SIDES = {
  [pb.protocol.GameTeam.GAME_TEAM_GOODGUYS]: "radiant",
  [pb.protocol.GameTeam.GAME_TEAM_BADGUYS]: "dire",
};

const PLAYER_MATCH_DETAILS_CUSTOM_TYPE_FIELDS = {
  player_slot: {
    type: PlayerSlot,
    fieldName: "slot",
  },
};

function getTeamAttributes(match, side) {
  return _.transform(
    TEAM_ATTRIBUTES,
    (attrs, attr) => {
      const attrName = TEAM_ATTRIBUTE_NAME_TPL({ side, attr });
      attrs[attr] = match[attrName];
    },
    {}
  );
}

function createMatchTeam(match, number, players) {
  const side = TEAM_SIDES[number];
  return _.assign({ number, players, side }, getTeamAttributes(match, side));
}

export function transformMatchTimestamps(match) {
  _.each(MATCH_TIMESTAMP_FIELDS, field => {
    const ts = match[field];

    if (ts instanceof pb.google.protobuf.Timestamp) {
      match[field] = new Date(Math.floor(ts.seconds * 1000 + ts.nanos / 1000000));
    }
  });
}

export function transformMatches(matches, { heroes }) {
  return _.map(matches, match => {
    transformMatchTimestamps(match);

    match.players = _.chain(match.players)
      .map(player => {
        player.hero = _.get(heroes, ["byId", player.hero_id]);
        return player;
      })
      .sortBy("player_slot")
      .value();

    match.teams = _.chain(match.players)
      .groupBy("team")
      .toPairs()
      .sortBy(([number]) => number)
      .transform((teams, [number, players]) => {
        teams[number] = createMatchTeam(match, _.toInteger(number), players);
      }, {})
      .value();

    return match;
  });
}

export function transformPlayer(player, state) {
  player.matches = _.map(player.matches, match => {
    transformMatchTimestamps(match);
    transformPlayerMatchPlayerDetails(match.player_details, state);
    transformPlayerMatchVictory(match);
    return match;
  });

  return player;
}

function transformPlayerMatchVictory(match) {
  const slot = match.player_details.slot;
  match.radiantVictory = match.outcome === pb.protocol.MatchOutcome.MATCH_OUTCOME_RAD_VICTORY;
  match.direVictory = match.outcome === pb.protocol.MatchOutcome.MATCH_OUTCOME_DIRE_VICTORY;
  match.unknownOutcome = !match.radiantVictory && !match.direVictory;
  match.playerVictory =
    (slot.isRadiant && match.radiantVictory) || (slot.isDire && match.direVictory);
}

function transformPlayerMatchPlayerDetails(details, { heroes }) {
  if (!details) {
    return;
  }

  details.hero = _.get(heroes, ["byId", details.hero_id]);

  _.each(PLAYER_MATCH_DETAILS_CUSTOM_TYPE_FIELDS, (spec, field) => {
    details[spec.fieldName] = new spec.type(details[field]);
  });
}
