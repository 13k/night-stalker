import _ from "lodash";

import { transformProperty } from "./object";

export function transformSearch(search, state) {
  transformProperty(search, "hero_ids", "heroes", heroIds => {
    return _.map(heroIds, id => state.heroes.byId[id]);
  });

  return search;
}
