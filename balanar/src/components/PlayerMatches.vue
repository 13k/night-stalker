<template>
  <v-card
    class="mx-auto"
    outlined
  >
    <v-data-iterator
      :items="filteredMatches"
      :items-per-page="itemsPerPage"
      :page="page"
      :search="filterByHeroName"
      :custom-filter="filterMatches"
      :sort-by="sortBy"
      :sort-desc="sortDesc"
      :custom-sort="sortMatches"
      hide-default-footer
    >
      <template v-slot:header>
        <v-toolbar
          color="secondary"
          height="auto"
          dark
        >
          <v-container>
            <v-row>
              <v-col
                cols="12"
                md="3"
              >
                <v-autocomplete
                  v-model="filterByHeroName"
                  :items="heroes"
                  :class="$vuetify.breakpoint.mdAndUp ? '' : 'mb-3'"
                  item-text="localized_name"
                  item-value="name"
                  label="Filter by hero"
                  clearable
                  dense
                  hide-details
                  single-line
                >
                  <template v-slot:item="{ item }">
                    <v-list-item-avatar>
                      <HeroImage
                        :hero="item"
                        orientation="icon"
                        width="28"
                        height="28"
                      />
                    </v-list-item-avatar>

                    <v-list-item-content>
                      <v-list-item-title v-text="item.localized_name" />
                    </v-list-item-content>
                  </template>
                </v-autocomplete>
              </v-col>

              <v-col
                cols="12"
                md="3"
                class="d-flex"
              >
                <v-select
                  v-model="sortBy"
                  :items="sortValues"
                  class="mr-1"
                  dense
                  hide-details
                />

                <v-btn
                  :title="sortDesc ? 'Sort ascending' : 'Sort descending'"
                  icon
                  small
                  @click.stop="sortDesc = !sortDesc"
                >
                  <v-icon>{{ sortDesc ? "mdi-sort-ascending" : "mdi-sort-descending" }}</v-icon>
                </v-btn>
              </v-col>

              <v-col
                cols="12"
                md="3"
              >
                <v-switch
                  v-model="onlyWins"
                  label="Victories only"
                  color="white"
                  dense
                />
              </v-col>
            </v-row>
          </v-container>
        </v-toolbar>
      </template>

      <template v-slot="{ items }">
        <v-expansion-panels hover>
          <PlayerMatch
            v-for="match in items"
            :key="match.match_id.toString()"
            :player="player"
            :match="match"
          />
        </v-expansion-panels>
      </template>

      <template v-slot:footer="{ pagination }">
        <v-row
          class="mt-3 mb-3"
          align="center"
          justify="center"
        >
          <v-btn
            :disabled="pagination.page === 1"
            class="mr-1"
            dark
            icon
            small
            @click="prevPage"
          >
            <v-icon>mdi-chevron-left</v-icon>
          </v-btn>

          <span>
            {{ pagination.page }} / {{ pagination.pageCount }}
          </span>

          <v-btn
            :disabled="pagination.page === pagination.pageCount"
            class="ml-1"
            dark
            icon
            small
            @click="nextPage"
          >
            <v-icon>mdi-chevron-right</v-icon>
          </v-btn>
        </v-row>
      </template>
    </v-data-iterator>
  </v-card>
</template>

<script>
import _ from "lodash";

import * as $t from "@/protocol/transform";
import pb from "@/protocol/proto";
import HeroImage from "@/components/HeroImage.vue";
import PlayerMatch from "@/components/PlayerMatch.vue";

export default {
  name: "PlayerMatches",

  components: {
    HeroImage,
    PlayerMatch,
  },

  props: {
    player: {
      type: pb.ns.protocol.Player,
      required: true,
    },
    matches: {
      type: Array,
      default: () => [],
      validator: v => _.every(v, i => i instanceof pb.ns.protocol.Match),
    },
    knownPlayers: {
      type: Array,
      default: () => [],
      validator: v => _.every(v, i => i instanceof pb.ns.protocol.Player),
    },
    itemsPerPage: {
      type: Number,
      default: 15,
    },
  },

  data: () => ({
    page: 1,
    filterByHeroName: null,
    sortBy: "time",
    sortDesc: true,
    sortValues: [
      { text: "Date", value: "time" },
      { text: "Hero", value: "hero" },
      { text: "MMR", value: "mmr" },
    ],
    onlyWins: false,
  }),

  computed: {
    filteredMatches() {
      let matches = this.matches;

      if (this.onlyWins) {
        matches = _.filter(matches, $t.bindGet("poi.$t.victory"));
      }

      return matches;
    },
    heroes() {
      return _.chain(this.filteredMatches)
        .map($t.bindGet("poi.$t.hero"))
        .filter("id")
        .uniqBy("id")
        .sortBy("localized_name")
        .value();
    },
  },

  methods: {
    nextPage() {
      this.page += 1;
    },
    prevPage() {
      this.page -= 1;
    },
    filterMatches(matches, search) {
      if (_.isEmpty(search)) {
        return matches;
      }

      return _.filter(matches, $t.propertyMatches("poi.$t.hero.name", search));
    },
    sortMatches(matches, sortBy, sortDesc) {
      sortBy = _.get(sortBy, "[0]", "time");
      sortDesc = _.get(sortDesc, "[0]", true);

      switch (sortBy) {
        case "time":
          matches = _.orderBy(matches, $t.property("activate_time"), sortDesc ? "desc" : "asc");
          break;
        case "hero":
          matches = _.orderBy(
            matches,
            $t.property("hero.localized_name"),
            sortDesc ? "desc" : "asc"
          );

          break;
        case "mmr":
          matches = _.orderBy(matches, "average_mmr", sortDesc ? "desc" : "asc");
          break;
        default:
          this.$log.error("Invalid player matches sorting:", sortBy);
      }

      return matches;
    },
  },
};
</script>
