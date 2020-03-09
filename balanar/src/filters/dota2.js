import _ from "lodash";

const TEAM_SIDES = {
  0: "Radiant",
  1: "Dire",
};

export function heroName(hero, fallback) {
  return _.get(hero, "localized_name", fallback);
}

export function teamSideName(slot) {
  return TEAM_SIDES[slot.team];
}
