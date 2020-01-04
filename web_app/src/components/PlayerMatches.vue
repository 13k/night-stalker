<template>
  <v-list dense>
    <v-list-item v-for="match in matches" :key="match.match_id">
      <v-list-item-icon>
        <hero-image :hero="match.hero" version="icon" width="32" height="32" />
      </v-list-item-icon>

      <v-list-item-content>
        <v-list-item-title v-text="match.match_id" />
      </v-list-item-content>

      <v-list-item-icon>
        <community-site-btn
          site="opendota"
          :href="match | opendotaMatchURL"
          :alt="`View match ${match.match_id} on OpenDota`"
          target="_blank"
          width="28"
          height="28"
        />
      </v-list-item-icon>

      <v-list-item-icon class="ml-1">
        <community-site-btn
          site="dotabuff"
          :href="match | dotabuffMatchURL"
          :alt="`View match ${match.match_id} on Dotabuff`"
          target="_blank"
          width="28"
          height="28"
        />
      </v-list-item-icon>

      <v-list-item-icon class="ml-1">
        <community-site-btn
          site="stratz"
          :href="match | stratzMatchURL"
          :alt="`View match ${match.match_id} on Stratz`"
          target="_blank"
          width="28"
          height="28"
        />
      </v-list-item-icon>

      <v-list-item-icon v-if="match.league_id" class="ml-1">
        <community-site-btn
          site="datdota"
          :href="match | datdotaMatchURL"
          :alt="`View match ${match.match_id} on DatDota`"
          target="_blank"
          width="28"
          height="28"
        />
      </v-list-item-icon>

      <v-list-item-icon>
        <v-icon>mdi-star</v-icon>
      </v-list-item-icon>
    </v-list-item>
  </v-list>
</template>

<script>
import filters from "@/components/filters";
import CommunitySiteBtn from "@/components/CommunitySiteBtn.vue";
import HeroImage from "@/components/HeroImage.vue";

export default {
  name: "player-matches",
  components: {
    CommunitySiteBtn,
    HeroImage
  },
  props: {
    matches: Array
  },
  filters
};
</script>
