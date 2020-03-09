import _ from "lodash";

export function l10n(value) {
  if (_.isDate(value)) {
    return value.toLocaleString();
  }

  return value;
}
