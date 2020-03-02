import _ from "lodash";

import api from "@/api";
import pb from "@/protocol/proto";
import { preprocessMatches } from "@/protocol/preprocess";
import {
  transformHero,
  transformHeroMatches,
  transformLiveMatches,
  transformPlayerMatches,
  transformSearch,
} from "@/protocol/transform";

export function fetchHeroes() {
  return api.heroes().then(res => {
    if (!_.isArray(res)) {
      throw new TypeError("received non-array response", res);
    }

    return res.map(attrs => {
      const hero = pb.protocol.Hero.fromObject(attrs);
      return transformHero(hero);
    });
  });
}

export function fetchLiveMatches(state) {
  return api.liveMatches().then(res => {
    if (!_.isPlainObject(res)) {
      throw new TypeError("received non-object response", res);
    }

    preprocessMatches(res.matches);

    const liveMatches = pb.protocol.LiveMatches.fromObject(res);

    transformLiveMatches(liveMatches, state);

    return liveMatches;
  });
}

export function fetchPlayerMatches(state, accountId) {
  return api.playerMatches(accountId).then(res => {
    if (!_.isPlainObject(res)) {
      throw new TypeError("received non-object response", res);
    }

    preprocessMatches(res.matches);

    const playerMatches = pb.protocol.PlayerMatches.fromObject(res);

    transformPlayerMatches(playerMatches, state);

    return playerMatches;
  });
}

export function fetchHeroMatches(state, heroId) {
  return api.heroMatches(heroId).then(res => {
    if (!_.isPlainObject(res)) {
      throw new TypeError("received non-object response", res);
    }

    preprocessMatches(res.matches);

    const heroMatches = pb.protocol.HeroMatches.fromObject(res);

    transformHeroMatches(heroMatches, state);

    return heroMatches;
  });
}

export function search(state, query) {
  return api.search(query).then(res => {
    if (!_.isPlainObject(res)) {
      throw new TypeError("received non-object response", res);
    }

    const search = pb.protocol.Search.fromObject(res);

    transformSearch(search, state);

    return search;
  });
}
