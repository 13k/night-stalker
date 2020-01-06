<template>
  <v-card class="mx-auto" outlined>
    <v-toolbar flat dark color="secondary">
      <v-container>
        <v-row>
          <v-spacer />

          <v-col cols="12" md="2">
            <v-select
              :items="sortValues"
              v-model="sortBy"
              dense
              hide-details
              :class="$vuetify.breakpoint.mdAndUp ? '' : 'mb-6'"
            />
          </v-col>

          <v-col cols="12" md="2">
            <v-text-field
              v-model="searchQuery"
              append-icon="mdi-magnify"
              clearable
              color="white"
              hide-details
              label="Search ..."
              single-line
              type="search"
            />
          </v-col>
        </v-row>
      </v-container>
    </v-toolbar>

    <v-list two-line dense>
      <v-list-item v-for="match in filteredMatches" :key="match.match_id">
        <v-list-item-icon>
          <hero-image
            :hero="match.hero"
            version="icon"
            width="32"
            height="32"
          />
        </v-list-item-icon>

        <v-list-item-content>
          <v-list-item-title v-text="match.match_id" />
          <v-list-item-subtitle v-text="match.activate_time.toLocaleString()" />
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
  </v-card>
</template>

<script>
import _ from "lodash";

import heroAliases from "@/dota2/hero_aliases.json";
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
  filters,
  data: () => ({
    searchQuery: null,
    sortBy: "time:desc",
    sortValues: [
      { text: "Newest", value: "time:desc" },
      { text: "Oldest", value: "time:asc" },
      { text: "Hero", value: "hero:asc" },
      { text: "Hero (desc)", value: "hero:desc" }
    ],
    filteredMatches: []
  }),
  created() {
    this.filterMatches();
  },
  watch: {
    searchQuery() {
      this.filterMatches();
    },
    sortBy() {
      this.filterMatches();
    }
  },
  computed: {
    tokenizedHeroNames() {
      return _.chain(this.matches)
        .filter("hero")
        .transform((tokenized, { hero: { id, name, localized_name } }) => {
          if (tokenized[id]) return;

          tokenized[id] = _.chain(localized_name)
            .words()
            .map(_.toLower)
            .concat(
              _.chain(name)
                .replace(/^npc_dota_hero_/, "")
                .words()
                .value()
            )
            .sortBy()
            .sortedUniq()
            .value();
        }, {})
        .value();
    }
  },
  methods: {
    filterMatches: _.throttle(function() {
      const query = _.toLower(this.searchQuery);

      let matches = this.matches;

      if (query.length > 1) {
        const heroNamesByAlias = _.chain(heroAliases)
          .toPairs()
          .filter(([, aliases]) => {
            return _.sortedIndexOf(aliases, query) >= 0;
          })
          .map(([name]) => name)
          .sortBy()
          .value();

        matches = _.filter(matches, ({ hero }) => {
          if (!hero) {
            return false;
          }

          return (
            _.sortedIndexOf(this.tokenizedHeroNames[hero.id], query) >= 0 ||
            _.sortedIndexOf(heroNamesByAlias, hero.name) >= 0
          );
        });
      }

      switch (this.sortBy) {
        case "time:asc":
          matches = _.sortBy(matches, "activate_time");
          break;
        case "time:desc":
          matches = _.sortBy(matches, m => -m.activate_time.getTime());
          break;
        case "hero:asc":
          matches = _.sortBy(matches, ["hero", "localized_name"]);
          break;
        case "hero:desc":
          // matches.sort()
          matches = _.sortBy(matches, m => -m.hero_id);
          break;
        default:
          this.$log.error("Invalid player matches sorting:", this.sortBy);
      }

      this.filteredMatches = matches;
    }, 500)
  }
};
</script>
