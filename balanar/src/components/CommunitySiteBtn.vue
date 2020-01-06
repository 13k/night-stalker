<template>
  <v-btn :icon="icon" :small="small" :href="href" :title="altText" :target="target">
    <v-img
      contain
      :src="source"
      :alt="altText"
      :title="altText"
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
export default {
  name: "community-site-btn",
  props: {
    site: String,
    alt: String,
    icon: {
      type: Boolean,
      default: true,
    },
    small: {
      type: Boolean,
      default: true,
    },
    href: String,
    target: String,
    width: [Number, String],
    height: [Number, String],
    maxWidth: [Number, String],
    maxHeight: [Number, String],
    minWidth: [Number, String],
    minHeight: [Number, String],
  },
  computed: {
    source() {
      return require(`@/assets/sites/${this.site}.png`);
    },
    altText() {
      if (this.alt) {
        return this.alt;
      }

      return `View on ${this.site}`;
    },
  },
};
</script>
