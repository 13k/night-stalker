import { assign, get } from "lodash/object";
import { isEmpty, isDate } from "lodash/lang";
import prettyMs from "pretty-ms";

const PRETTY_DURATION_DEFAULTS = { unit: "seconds" };

const TEAM_SIDES = {
  0: "Radiant",
  1: "Dire",
};

export function l10n(value) {
  if (isDate(value)) {
    return value.toLocaleString();
  }

  return value;
}

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
  return prettyDuration(duration, assign({ colonNotation: true }, options));
}

export function heroPlaceholderImageURL(version) {
  switch (version) {
    case "portrait":
      return require("@/assets/heroes/default_vert.png");
    default:
      return require("@/assets/heroes/default_full.png");
  }
}

// available versions: full (full), large (lg), small (sb), portrait (vert)
export function heroImageURL(hero, options) {
  options = assign({ version: "full", placeholder: true }, options);
  let url = get(hero, `image_${options.version}_url`);

  if (isEmpty(url)) {
    if (options.placeholder === true) {
      url = heroPlaceholderImageURL(options.version);
    } else if (options.placeholder) {
      url = options.placeholder;
    }
  }

  return url;
}

export function heroName(hero, fallback) {
  return get(hero, "localized_name", fallback);
}

export function opendotaMatchURL(match) {
  return `https://www.opendota.com/matches/${match.match_id}`;
}

export function opendotaPlayerURL(player) {
  return `https://www.opendota.com/players/${player.account_id}`;
}

export function opendotaTeamURL(team) {
  return `https://www.opendota.com/teams/${team.id}`;
}

export function dotabuffMatchURL(match) {
  return `https://www.dotabuff.com/matches/${match.match_id}`;
}

export function dotabuffPlayerURL(player) {
  return `https://www.dotabuff.com/players/${player.account_id}`;
}

export function dotabuffTeamURL(team) {
  return `https://www.dotabuff.com/esports/teams/${team.id}`;
}

export function stratzMatchURL(match) {
  return `https://www.stratz.com/matches/${match.match_id}`;
}

export function stratzPlayerURL(player) {
  return `https://www.stratz.com/players/${player.account_id}`;
}

export function datdotaMatchURL(match) {
  return `https://www.datdota.com/matches/${match.match_id}`;
}

export function datdotaPlayerURL(player) {
  return `https://datdota.com/players/${player.account_id}`;
}

export function datdotaTeamURL(team) {
  return `https://www.datdota.com/teams/${team.id}`;
}

export function teamSideName(slot) {
  return TEAM_SIDES[slot.teamNumber];
}
