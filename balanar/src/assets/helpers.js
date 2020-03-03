import _ from "lodash";

export function image(path) {
  return require(`@/assets/images/${path}`);
}

export const HERO_IMAGE_ORIENTATIONS = ["portrait", "landscape", "icon"];
export const HERO_IMAGE_SIZES = ["regular", "large"];
export const DEFAULT_HERO_IMAGE_ORIENTATION = "landscape";
export const DEFAULT_HERO_IMAGE_SIZE = "regular";

const HERO_PLACEHOLDER_IMAGES = {
  "portrait.regular": "heroes/default_portrait.png",
  "portrait.large": "heroes/default_portrait.png",
  "landscape.regular": "heroes/default_landscape.png",
  "landscape.large": "heroes/default_landscape.png",
};

const HERO_IMAGES = {
  "icon": hero => `heroes/vpk/icons/${hero.name}_png.png`,
  "portrait.regular": hero => `heroes/vpk/selection/${hero.name}_png.png`,
  "portrait.large": hero => `heroes/cdn/${hero.name}_vert.jpg`,
  "landscape.regular": hero => `heroes/vpk/${hero.name}_png.png`,
  "landscape.large": hero => `heroes/cdn/${hero.name}_full.png`,
};

function heroImageVersion({
  orientation = DEFAULT_HERO_IMAGE_ORIENTATION,
  size = DEFAULT_HERO_IMAGE_SIZE,
} = {}) {
  if (orientation === "icon" || orientation.indexOf(".") >= 0) {
    return orientation;
  }

  return `${orientation}.${size}`;
}

export function heroPlaceholderImage(options = {}) {
  const version = heroImageVersion(options);
  const path = HERO_PLACEHOLDER_IMAGES[version];

  return path && image(path);
}

export function heroImage(hero, options = {}) {
  options = _.assign({ placeholder: true }, options);

  if (hero == null) {
    if (options.placeholder === true) {
      return heroPlaceholderImage(options);
    } else if (options.placeholder) {
      return image(options.placeholder);
    }
  }

  const version = heroImageVersion(options);
  const pathFn = HERO_IMAGES[version];

  return pathFn && image(pathFn(hero));
}
