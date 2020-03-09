import _ from "lodash";

import { fetchLeagues } from "@/protocol/api";
import { set } from "@/protocol/transform";

export function prefetchMatchesLeagues(matches) {
  const leagueIds = _.chain(matches)
    .map(match => match.league_id)
    .filter(id => id && !id.isZero())
    .value();

  if (_.isEmpty(leagueIds)) {
    return Promise.resolve(matches);
  }

  return fetchLeagues(leagueIds).then(leagues => {
    const leaguesById = _.keyBy(leagues, "id");

    _.each(matches, match => {
      set(match, "league", leaguesById[match.league_id]);
    });

    return matches;
  });
}
