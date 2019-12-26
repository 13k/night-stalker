import { assign, get } from "lodash/object";
import { isEmpty } from "lodash/lang";
import prettyMs from "pretty-ms";

function humanDuration(duration, inputPrecision) {
  inputPrecision = inputPrecision || "seconds";

  if (inputPrecision === "seconds") {
    duration *= 1000;
  }

  return prettyMs(duration);
}

function heroPlaceholderImageURL(version) {
  switch (version) {
    case "portrait":
      return require("@/assets/heroes/default_vert.png");
    default:
      return require("@/assets/heroes/default_full.png");
  }
}

// available versions: full (full), large (lg), small (sb), portrait (vert)
function heroImageURL(hero, options) {
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

function heroName(hero, fallback) {
  return get(hero, "localized_name", fallback);
}

export default {
  humanDuration,
  heroImageURL,
  heroName
};
