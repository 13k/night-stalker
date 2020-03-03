<template>
  <v-btn
    :title="title"
    icon
    @click="$vuetify.theme.dark = !$vuetify.theme.dark"
  >
    <v-icon>mdi-brightness-{{ $vuetify.theme.dark ? '7' : '4' }}</v-icon>
  </v-btn>
</template>

<script>
export default {
  name: "ThemeToggle",

  computed: {
    title() {
      return this.$vuetify.theme.dark ? "Daytime" : "Nighttime";
    },
  },

  watch: {
    "$vuetify.theme.dark"(val) {
      localStorage.setItem("balanar.theme.dark", val);
    },
  },
};
</script>
