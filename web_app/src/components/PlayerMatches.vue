<template>
  <v-list dense>
    <v-list-item
      v-for="match in matches"
      :key="match.match_id"
      link
      ripple
      :href="`https://www.opendota.com/matches/${match.match_id}`"
      target="_blank"
    >
      <v-list-item-icon>
        <hero-image :hero="match.hero" version="icon" />
      </v-list-item-icon>

      <v-list-item-content>
        <!-- <span>{{ `View match ${match.match_id} on OpenDota` }}</span> -->
        <v-list-item-title v-text="match.match_id" v-on="on">
        </v-list-item-title>
      </v-list-item-content>
    </v-list-item>
  </v-list>
</template>

<script>
import HeroImage from "@/components/HeroImage.vue";

export default {
  name: "player-matches",
  components: {
    HeroImage
  },
  props: {
    matches: Array
  }
};
</script>
