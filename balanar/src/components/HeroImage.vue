<template>
  <v-img
    contain
    :src="source"
    :lazy-src="lazySource"
    :alt="altText"
    :title="altText"
    :width="width"
    :height="height"
    :max-width="maxWidth"
    :max-height="maxHeight"
    :min-width="minWidth"
    :min-height="minHeight"
  />
</template>

<script>
import _ from "lodash";

import * as $a from "@/assets/helpers";
import pb from "@/protocol/proto";

export default {
  name: "HeroImage",

  props: {
    hero: {
      type: pb.protocol.Hero,
      default: null,
    },
    orientation: {
      type: String,
      default: $a.DEFAULT_HERO_IMAGE_ORIENTATION,
      validator: value => $a.HERO_IMAGE_ORIENTATIONS.indexOf(value) >= 0,
    },
    size: {
      type: String,
      default: $a.DEFAULT_HERO_IMAGE_SIZE,
      validator: value => $a.HERO_IMAGE_SIZES.indexOf(value) >= 0,
    },
    placeholder: {
      type: Boolean,
      default: true,
    },
    alt: {
      type: [Boolean, String],
      default: true,
    },
    altPlaceholder: {
      type: String,
      default: null,
    },
    width: {
      type: [Number, String],
      default: null,
    },
    height: {
      type: [Number, String],
      default: null,
    },
    maxWidth: {
      type: [Number, String],
      default: null,
    },
    maxHeight: {
      type: [Number, String],
      default: null,
    },
    minWidth: {
      type: [Number, String],
      default: null,
    },
    minHeight: {
      type: [Number, String],
      default: null,
    },
  },

  computed: {
    source() {
      return $a.heroImage(this.hero, {
        orientation: this.orientation,
        size: this.size,
        placeholder: this.placeholder,
      });
    },
    lazySource() {
      return $a.heroPlaceholderImage({
        orientation: this.orientation,
        size: this.size,
      });
    },
    altText() {
      if (_.isString(this.alt)) {
        return this.alt;
      }

      if (this.alt) {
        return _.get(this.hero, "localized_name", this.altPlaceholder);
      }

      return null;
    },
  },
};
</script>
