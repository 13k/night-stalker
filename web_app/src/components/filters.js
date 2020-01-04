import { assign, get } from "lodash/object";
import { isEmpty } from "lodash/lang";
import prettyMs from "pretty-ms";

export default {
  humanDuration(duration, inputPrecision) {
    inputPrecision = inputPrecision || "seconds";

    if (inputPrecision === "seconds") {
      duration *= 1000;
    }

    return prettyMs(duration);
  },

  heroPlaceholderImageURL(version) {
    switch (version) {
      case "portrait":
        return require("@/assets/heroes/default_vert.png");
      default:
        return require("@/assets/heroes/default_full.png");
    }
  },

  // available versions: full (full), large (lg), small (sb), portrait (vert)
  heroImageURL(hero, options) {
    options = assign({ version: "full", placeholder: true }, options);
    let url = get(hero, `image_${options.version}_url`);

    if (isEmpty(url)) {
      if (options.placeholder === true) {
        url = this.heroPlaceholderImageURL(options.version);
      } else if (options.placeholder) {
        url = options.placeholder;
      }
    }

    return url;
  },

  heroName(hero, fallback) {
    return get(hero, "localized_name", fallback);
  },

  opendotaMatchURL(match) {
    return `https://www.opendota.com/matches/${match.match_id}`;
  },

  opendotaPlayerURL(player) {
    return `https://www.opendota.com/players/${player.account_id}`;
  },

  opendotaTeamURL(team) {
    return `https://www.opendota.com/teams/${team.id}`;
  },

  dotabuffMatchURL(match) {
    return `https://www.dotabuff.com/matches/${match.match_id}`;
  },

  dotabuffPlayerURL(player) {
    return `https://www.dotabuff.com/players/${player.account_id}`;
  },

  dotabuffTeamURL(team) {
    return `https://www.dotabuff.com/esports/teams/${team.id}`;
  },

  stratzMatchURL(match) {
    return `https://www.stratz.com/matches/${match.match_id}`;
  },

  stratzPlayerURL(player) {
    return `https://www.stratz.com/players/${player.account_id}`;
  },

  datdotaMatchURL(match) {
    return `https://www.datdota.com/matches/${match.match_id}`;
  },

  datdotaPlayerURL(player) {
    return `https://datdota.com/players/${player.account_id}`;
  },

  datdotaTeamURL(team) {
    return `https://www.datdota.com/teams/${team.id}`;
  }
};
