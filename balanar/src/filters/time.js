import _ from "lodash";
import prettyMs from "pretty-ms";

const PRETTY_DURATION_DEFAULTS = { unit: "seconds" };

export function prettyDuration(duration, { unit, ...options } = PRETTY_DURATION_DEFAULTS) {
  unit = unit || "seconds";

  switch (unit) {
    case "nanoseconds":
      duration *= 10 ** -6;
      break;
    case "microseconds":
      duration *= 10 ** -3;
      break;
    case "milliseconds":
      break;
    case "seconds":
      duration *= 10 ** 3;
      break;
  }

  return prettyMs(duration, options);
}

export function humanDuration(duration, options = PRETTY_DURATION_DEFAULTS) {
  return prettyDuration(duration, options);
}

export function colonDuration(duration, options = PRETTY_DURATION_DEFAULTS) {
  return prettyDuration(duration, _.assign({ colonNotation: true }, options));
}
