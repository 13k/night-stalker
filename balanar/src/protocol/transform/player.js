import _ from "lodash";

import { transformProperty } from "./object";
import { transformMatchHistory } from "./match_common";

export const transformPlayer = player => player;

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
