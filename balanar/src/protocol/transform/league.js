import _ from "lodash";

export const transformLeague = league => league;

export function transformLeagues(leagues) {
  _.each(leagues.leagues, transformLeague);
  return leagues;
}
