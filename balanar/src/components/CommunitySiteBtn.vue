<template>
  <v-btn
    :icon="icon"
    :small="small"
    :href="href"
    :title="alt"
    :target="target"
  >
    <v-img
      contain
      :src="source"
      :alt="alt"
      :title="alt"
      :width="width"
      :height="height"
      :max-width="maxWidth"
      :max-height="maxHeight"
      :min-width="minWidth"
      :min-height="minHeight"
    />
  </v-btn>
</template>

<script>
import * as $a from "@/assets/helpers";

export default {
  name: "CommunitySiteBtn",

  props: {
    site: {
      type: String,
      required: true,
    },
    alt: {
      type: String,
      default() {
        return `View on ${this.site}`;
      },
    },
    icon: {
      type: Boolean,
      default: true,
    },
    small: {
      type: Boolean,
      default: true,
    },
    href: {
      type: String,
      default: null,
    },
    target: {
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
      return $a.image(`sites/${this.site}.png`);
    },
  },
};
</script>
