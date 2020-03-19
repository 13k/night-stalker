import _ from "lodash";

import api from "@/api";
import pb from "@/protocol/proto";
import * as $t from "@/protocol/transform";
import { prefetchMatchesLeagues } from "@/protocol/prefetch";
import { preprocessMatches, preprocessLeagues } from "@/protocol/preprocess";

const {
  ns: {
    protocol: { Hero, League, LiveMatches, PlayerMatches, HeroMatches, Search },
  },
} = pb;

export function fetchHeroes() {
  return api.heroes().then(res => {
    if (!_.isArray(res)) {
      throw new TypeError("received non-array response", res);
    }

    return res.map(attrs => $t.transformHero(Hero.fromObject(attrs)));
  });
}

export function fetchLeagues(id) {
  return api.leagues(id).then(res => {
    if (!_.isArray(res)) {
      throw new TypeError("received non-array response", res);
    }

    preprocessLeagues(res);

    return res.map(attrs => $t.transformLeague(League.fromObject(attrs)));
  });
}

export function fetchLiveMatches(state) {
  return api.liveMatches().then(res => {
    if (!_.isPlainObject(res)) {
      throw new TypeError("received non-object response", res);
    }

    preprocessMatches(res.matches);

    const liveMatches = LiveMatches.fromObject(res);

    $t.transformLiveMatches(liveMatches, state);

    return prefetchMatchesLeagues(liveMatches.matches).then(() => liveMatches);
  });
}

export function fetchPlayerMatches(state, accountId) {
  return api.playerMatches(accountId).then(res => {
    if (!_.isPlainObject(res)) {
      throw new TypeError("received non-object response", res);
    }

    preprocessMatches(res.matches);

    const playerMatches = PlayerMatches.fromObject(res);

    $t.transformPlayerMatches(playerMatches, state);

    return playerMatches;
  });
}

export function fetchHeroMatches(state, heroId) {
  return api.heroMatches(heroId).then(res => {
    if (!_.isPlainObject(res)) {
      throw new TypeError("received non-object response", res);
    }

    preprocessMatches(res.matches);

    const heroMatches = HeroMatches.fromObject(res);

    $t.transformHeroMatches(heroMatches, state);

    return heroMatches;
  });
}

export function search(state, query) {
  return api.search(query).then(res => {
    if (!_.isPlainObject(res)) {
      throw new TypeError("received non-object response", res);
    }

    const search = Search.fromObject(res);

    $t.transformSearch(search, state);

    return search;
  });
}
