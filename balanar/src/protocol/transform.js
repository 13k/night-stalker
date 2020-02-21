import _ from "lodash";

import pb from "@/protocol/proto";
import { PlayerSlot } from "@/dota2/flags";
import { MATCH_TIMESTAMP_FIELDS } from "@/protocol/preprocess";

const TRANSFORM_KEY = "$t";

const TEAM_ATTRIBUTE_NAME_TPL = _.template("<%= side %>_team_<%= attr %>");
const TEAM_ATTRIBUTES = ["id", "name", "tag", "logo", "logo_url"];
const TEAM_SIDES = {
  [pb.protocol.GameTeam.GAME_TEAM_GOODGUYS]: "radiant",
  [pb.protocol.GameTeam.GAME_TEAM_BADGUYS]: "dire",
};

export function propertyPath(key) {
  let path = [TRANSFORM_KEY];

  if (_.isArray(key)) {
    path = _.concat(path, key);
  } else if (_.isString(key)) {
    path = _.concat(path, _.split(key, "."));
  } else {
    path.push(key);
  }

  return path;
}

export function property(key) {
  return _.property(propertyPath(key));
}

export function propertyMatches(pattern) {
  return _.matches({ [TRANSFORM_KEY]: pattern });
}

export function get(object, key, defaultValue) {
  return _.get(object, propertyPath(key), defaultValue);
}

export function bindGet(key) {
  return _.chain(get)
    .partialRight(key)
    .unary()
    .value();
}

export function set(object, key, value) {
  _.set(object, propertyPath(key), value);
  return object;
}

export function transformMatchTimestamps(match) {
  _.each(MATCH_TIMESTAMP_FIELDS, field => {
    const ts = match[field];

    if (ts instanceof pb.google.protobuf.Timestamp) {
      const date = new Date(Math.floor(ts.seconds * 1000 + ts.nanos / 1000000));
      match = set(match, field, date);
    }
  });

  return match;
}

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

function createMatchTeam({ number, match, players }) {
  number = _.toInteger(number);
  const side = TEAM_SIDES[number];
  return _.assign({ number, players, side }, getTeamAttributes(match, side));
}

export function transformAssignHero(object, { heroes }, heroId) {
  return set(object, "hero", _.get(heroes, ["byId", heroId]));
}

export function transformAssignPlayerSlot(object, slot) {
  return set(object, "slot", new PlayerSlot(slot));
}

export function transformMatches(matches, state) {
  return _.map(matches, match => {
    match = transformMatchTimestamps(match);

    const players = _.chain(match.players)
      .map(player => {
        player = transformAssignHero(player, state, player.hero_id);
        player = transformAssignPlayerSlot(player, player.player_slot);
        return player;
      })
      .sortBy("player_slot")
      .value();

    match = set(match, "players", players);

    const teams = _.chain(players)
      .groupBy("team")
      .toPairs()
      .sortBy(([number]) => number)
      .transform((teams, [number, players]) => {
        teams[number] = createMatchTeam({ number, match, players });
      }, {})
      .value();

    match = set(match, "teams", teams);

    return match;
  });
}

export function transformPlayer(player, state) {
  const matches = _.map(player.matches, match => {
    match = transformMatchTimestamps(match);
    match = transformPlayerMatchPlayerDetails(match, state);
    match = transformPlayerMatchOutcome(match);
    return match;
  });

  return set(player, "matches", matches);
}

function transformPlayerMatchPlayerDetails(match, state) {
  if (!match.player_details) {
    return;
  }

  match = transformAssignHero(match, state, match.player_details.hero_id);
  match = transformAssignPlayerSlot(match, match.player_details.player_slot);

  return match;
}

function transformPlayerMatchOutcome(match) {
  const eOutcome = match.outcome || pb.protocol.MatchOutcome.MATCH_OUTCOME_UNKNOWN;

  if (eOutcome === pb.protocol.MatchOutcome.MATCH_OUTCOME_UNKNOWN) {
    return match;
  }

  const slot = get(match, "slot");

  if (!slot) {
    return match;
  }

  const outcome = {
    radiantVictory: eOutcome === pb.protocol.MatchOutcome.MATCH_OUTCOME_RAD_VICTORY,
    direVictory: eOutcome === pb.protocol.MatchOutcome.MATCH_OUTCOME_DIRE_VICTORY,
  };

  outcome.playerVictory =
    (slot.isRadiant && outcome.radiantVictory) || (slot.isDire && outcome.direVictory);

  return set(match, "outcome", outcome);
}
