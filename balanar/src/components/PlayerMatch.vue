
<template>
  <v-expansion-panel>
    <v-expansion-panel-header class="justify-start">
      <HeroImage
        :hero="match.hero"
        version="icon"
        max-width="32"
        max-height="32"
        class="mr-6"
      />

      <div class="d-flex flex-column">
        <span>{{ match.match_id }}</span>
        <span class="caption">{{ match.activate_time.toLocaleString() }}</span>
      </div>

      <v-spacer />

      <div class="flex-grow-0 d-flex flex-column flex-sm-row mr-4">
        <CommunitySiteBtn
          v-for="link in communitySitesMatch"
          :key="link.site"
          :site="link.site"
          :href="link.url"
          target="_blank"
          :max-width="communitySiteIconSize"
          :max-height="communitySiteIconSize"
          class="ml-1"
        />
      </div>
    </v-expansion-panel-header>

    <v-expansion-panel-content />
  </v-expansion-panel>
</template>

<script>
import filters from "@/components/filters";
import CommunitySiteBtn from "@/components/CommunitySiteBtn.vue";
import HeroImage from "@/components/HeroImage.vue";

export default {
  name: "PlayerMatch",

  components: {
    CommunitySiteBtn,
    HeroImage,
  },

  props: {
    match: {
      type: Object,
      required: true,
    },
  },

  computed: {
    communitySiteIconSize() {
      return this.$vuetify.breakpoint.xsOnly ? 22 : 28;
    },
    communitySitesMatch() {
      const sites = [
        {
          site: "opendota",
          url: filters.opendotaMatchURL(this.match),
          text: "View match on OpenDota",
        },
        {
          site: "dotabuff",
          url: filters.dotabuffMatchURL(this.match),
          text: "View match on Dotabuff",
        },
        {
          site: "stratz",
          url: filters.stratzMatchURL(this.match),
          text: "View match on Stratz",
        },
      ];

      if (!this.match.league_id.isZero()) {
        sites.push({
          site: "datdota",
          url: filters.datdotaMatchURL(this.match),
          text: "View match on DatDota",
        });
      }

      return sites;
    },
  },
};
</script>
