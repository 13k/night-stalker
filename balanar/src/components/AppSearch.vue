<template>
  <v-autocomplete
    ref="search"
    v-model="select"
    :error="error"
    :filter="() => true"
    :items="results"
    :label="label"
    :loading="loading"
    :search-input.sync="query"
    append-icon="mdi-magnify"
    clearable
    dense
    flat
    hide-details
    return-object
    single-line
    solo-inverted
    @keyup.esc="clearSearch"
  >
    <template
      v-if="error"
      v-slot:no-data
    >
      <v-list-item>
        <v-list-item-content>
          <v-alert type="error">
            {{ error }}
          </v-alert>
        </v-list-item-content>
      </v-list-item>
    </template>

    <template v-slot:item="{ item }">
      <v-list-item-avatar :color="item.type === 'player' ? 'accent' : 'secondary'">
        <v-img
          v-if="item.type === 'player'"
          :src="item.player.avatar_medium_url"
          :max-width="avatarSize"
          :max-height="avatarSize"
          contain
        />

        <HeroImage
          v-if="item.type === 'hero'"
          :hero="item.hero"
          :width="avatarSize"
          :height="avatarSize"
          orientation="icon"
        />
      </v-list-item-avatar>

      <v-list-item-content>
        <v-list-item-title v-text="item.text" />
      </v-list-item-content>
    </template>
  </v-autocomplete>
</template>

<script>
import _ from "lodash";

import * as $t from "@/protocol/transform";
import { search } from "@/protocol/api";
import { heroRoute, playerRoute } from "@/router/helpers";
import HeroImage from "@/components/HeroImage.vue";

export default {
  name: "AppSearch",

  components: {
    HeroImage,
  },

  data: () => ({
    query: null,
    select: null,
    error: null,
    results: [],
    loading: false,
  }),

  computed: {
    label() {
      return this.$vuetify.breakpoint.xsOnly ? "Search..." : 'Search ("Ctrl+Enter" to focus)';
    },
    avatarSize() {
      return this.$vuetify.breakpoint.xsOnly ? 18 : 32;
    },
  },

  watch: {
    query(val) {
      if (val && val.length > 1) {
        this.search(val);
      }
    },
    select(item) {
      if (item == null) {
        return;
      }

      let route;

      switch (item.type) {
        case "player":
          route = playerRoute(item.player);
          break;
        case "hero":
          route = heroRoute(item.hero);
          break;
      }

      if (route) {
        this.$router.push(route);
      }
    },
  },

  mounted() {
    document.addEventListener("keyup", ev => {
      ev = ev || window.event;

      if (ev.target !== this.$refs.search.$refs.input && ev.key === "Enter" && ev.ctrlKey) {
        ev.preventDefault();
        this.focusSearch();
      }
    });
  },

  methods: {
    focusSearch() {
      this.$refs.search.focus();
    },
    clearSearch() {
      this.query = null;
      this.select = null;
      this.error = null;
      this.results = [];
      this.$refs.search.reset();
      this.$refs.search.blur();
    },
    search: _.debounce(function(query) {
      this.loading = true;

      search(this.$store.state, query)
        .then(res => {
          const players = res.players || [];
          const heroes = $t.get(res, "heroes", []);
          let results = [];

          if (players.length > 0) {
            results.push({ header: "Players" });

            results = results.concat(
              players.map(player => {
                return {
                  type: "player",
                  text: player.name,
                  value: `player-${player.account_id}`,
                  player,
                };
              })
            );
          }

          if (heroes.length > 0) {
            if (results.length > 0) {
              results.push({ divider: true });
            }

            results.push({ header: "Heroes" });

            results = results.concat(
              heroes.map(hero => {
                return {
                  type: "hero",
                  text: hero.localized_name,
                  value: `hero-${hero.id}`,
                  hero,
                };
              })
            );
          }

          this.results = results;
        })
        .catch(err => {
          this.results = [];
          this.error = err.toString();
        })
        .finally(() => {
          this.loading = false;
        });
    }, 500),
  },
};
</script>
