<template>
  <v-img
    contain
    :src="source"
    :lazy-src="lazySource"
    :alt="altText"
    :width="width"
    :height="height"
    :max-width="maxWidth"
    :max-height="maxHeight"
    :min-width="minWidth"
    :min-height="minHeight"
  />
</template>

<script>
import { get } from "lodash/object";
import { isString } from "lodash/lang";

export default {
  name: "hero-image",
  props: {
    hero: Object,
    version: {
      type: String,
      default: "portrait"
    },
    placeholder: {
      type: Boolean,
      default: true
    },
    alt: {
      type: [Boolean, String],
      default: true
    },
    altPlaceholder: String,
    width: [Number, String],
    height: [Number, String],
    maxWidth: [Number, String],
    maxHeight: [Number, String],
    minWidth: [Number, String],
    minHeight: [Number, String]
  },
  computed: {
    source() {
      if (!this.hero) {
        return this.placeholder ? this.lazySource : "";
      }

      switch (this.version) {
        case "icon":
          return require(`@/assets/heroes/icons/${this.hero.name}_png.png`);
        case "selection":
          return require(`@/assets/heroes/selection/${this.hero.name}_png.png`);
        default:
          return require(`@/assets/heroes/${this.hero.name}_png.png`);
      }
    },
    lazySource() {
      switch (this.version) {
        case "icon":
          return "";
        case "selection":
          return require("@/assets/heroes/default_vert.png");
        default:
          return require("@/assets/heroes/default_full.png");
      }
    },
    altText() {
      if (isString(this.alt)) {
        return this.alt;
      }

      if (this.alt) {
        return get(this.hero, "localized_name", this.altPlaceholder);
      }

      return "";
    }
  }
};
</script>
