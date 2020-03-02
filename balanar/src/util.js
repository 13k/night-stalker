import _ from "lodash";

export function normalizePlayerName(name) {
  return _.chain((name || "").normalize())
    .deburr()
    .toLower()
    .replace(/\W/g, "-")
    .replace(/-+/g, "-")
    .replace(/(^-|-$)/, "")
    .value();
}
