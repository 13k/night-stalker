import _ from "lodash";

import pb from "@/protocol/proto";
import { PlayerSlot } from "@/dota2/flags";
import { MATCH_TIMESTAMP_FIELDS } from "@/protocol/preprocess";

const TRANSFORM_KEY = "$t";
const TRANSFORM_KEY_PLACEHOLDER = "$t";

const TEAM_ATTRIBUTE_NAME_TPL = _.template("<%= side %>_team_<%= attr %>");
const TEAM_ATTRIBUTES = ["id", "name", "tag", "logo", "logo_url"];
const TEAM_SIDES = {
  [pb.protocol.GameTeam.GAME_TEAM_GOODGUYS]: "radiant",
  [pb.protocol.GameTeam.GAME_TEAM_BADGUYS]: "dire",
};

export function propertyPath(path) {
  if (_.isString(path)) {
    path = _.split(path, ".");
  } else if (!_.isArray(path)) {
    throw new TypeError(`invalid path type ${path} (${typeof path})`);
  }

  if (path.length === 0) {
    throw new Error("empty path");
  }

  if (path[0] !== TRANSFORM_KEY_PLACEHOLDER) {
    path.unshift(TRANSFORM_KEY);
  }

  return _.map(path, seg => (seg === TRANSFORM_KEY_PLACEHOLDER ? TRANSFORM_KEY : seg));
}

export function property(path) {
  return _.property(propertyPath(path));
}

export function propertyMatches(path, value) {
  if (_.isPlainObject(path)) {
    return _.matches({ [TRANSFORM_KEY]: path });
  }

  return _.matchesProperty(propertyPath(path), value);
}

export function get(object, path, defaultValue) {
  return _.get(object, propertyPath(path), defaultValue);
}

export function bindGet(path) {
  return _.chain(get)
    .partialRight(path)
    .unary()
    .value();
}

export function set(object, path, value) {
  _.set(object, propertyPath(path), value);

  return object;
}

export function transformProperty(object, path, newPath, transformation) {
  if (_.isFunction(newPath)) {
    transformation = newPath;
    newPath = path;
  }

  return set(object, newPath, transformation(_.get(object, path)));
}

function transformMatchTimestamps(match) {
  _.each(MATCH_TIMESTAMP_FIELDS, field => {
    const ts = match[field];

    if (ts instanceof pb.google.protobuf.Timestamp) {
      const date = new Date(Math.floor(ts.seconds * 1000 + ts.nanos / 1000000));
      set(match, field, date);
    }
  });

  return match;
}

function createliveMatchTeamAttributes(liveMatch, side) {
  return _.transform(
    TEAM_ATTRIBUTES,
    (attrs, attr) => {
      const attrName = TEAM_ATTRIBUTE_NAME_TPL({ side, attr });
      attrs[attr] = liveMatch[attrName];
    },
    {}
  );
}

function createLiveMatchTeam({ number, liveMatch, players }) {
  number = _.toInteger(number);
  const side = TEAM_SIDES[number];
  return _.assign({ number, players, side }, createliveMatchTeamAttributes(liveMatch, side));
}

export const transformHero = hero => hero;
export const transformLeague = league => league;
export const transformPlayer = player => player;

export function transformLiveMatch(liveMatch, state) {
  transformMatchTimestamps(liveMatch);

  _.each(liveMatch.players, player => transformLiveMatchPlayer(player, state));

  liveMatch.players = _.sortBy(liveMatch.players, "player_slot");

  const teams = _.chain(liveMatch.players)
    .groupBy("team")
    .toPairs()
    .sortBy(([number]) => number)
    .transform((teams, [number, players]) => {
      teams[number] = createLiveMatchTeam({ number, liveMatch, players });
    }, {})
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

export function transformMatch(match, state) {
  transformMatchTimestamps(match);

  transformProperty(match, "outcome", outcome => {
    const radiantVictory = outcome === pb.protocol.MatchOutcome.MATCH_OUTCOME_RAD_VICTORY;
    const direVictory = outcome === pb.protocol.MatchOutcome.MATCH_OUTCOME_DIRE_VICTORY;
    return { radiantVictory, direVictory };
  });

  _.each(match.players, player => transformMatchPlayer(match, player, state));

  return match;
}

export function transformMatchPlayer(match, player, { heroes }) {
  const slot = new PlayerSlot(player.player_slot);
  const outcome = get(match, "outcome");
  const victory =
    (slot.isRadiant && outcome.radiantVictory) || (slot.isDire && outcome.direVictory);

  set(player, "slot", slot);
  set(player, "victory", victory);
  transformProperty(player, "hero_id", "hero", heroId => _.get(heroes, ["byId", heroId]));

  return player;
}

function transformMatchHistory(history, state) {
  _.each(history.matches, match => transformMatch(match, state));
  _.each(history.known_players, player => transformPlayer(player, state));

  return history;
}

export function transformPlayerMatches(playerMatches, state) {
  transformPlayer(playerMatches.player);
  transformMatchHistory(playerMatches, state);

  _.each(playerMatches.matches, match => {
    transformProperty(match, "players", "poi", players => {
      return _.find(players, { account_id: playerMatches.player.account_id });
    });
  });

  return playerMatches;
}

export function transformHeroMatches(heroMatches, state) {
  transformHero(heroMatches.hero);
  transformMatchHistory(heroMatches, state);

  _.each(heroMatches.matches, match => {
    transformProperty(match, "players", "poi", players => {
      return _.find(players, { hero_id: heroMatches.hero.id });
    });
  });

  return heroMatches;
}

export function transformSearch(search, state) {
  transformProperty(search, "hero_ids", "heroes", heroIds => {
    return _.map(heroIds, id => state.heroes.byId[id]);
  });

  return search;
}
