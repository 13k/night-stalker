import _ from "lodash";

import { transformProperty } from "./object";
import { transformMatchHistory } from "./match_common";

export const transformHero = hero => hero;

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
