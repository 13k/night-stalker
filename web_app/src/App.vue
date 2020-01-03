<template>
  <v-app id="night-stalker">
    <v-navigation-drawer v-model="drawer" app clipped>
      <v-list dense>
        <v-list-item link>
          <v-list-item-action>
            <v-icon>mdi-history</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>
              History
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-subheader class="mt-4 grey--text text--darken-1">
          LIVE STALKED
        </v-subheader>

        <v-list>
          <v-list-item
            v-for="(player, i) in followed"
            :key="player.account_id"
            link
          >
            <v-list-item-avatar>
              <img
                :src="`https://randomuser.me/api/portraits/men/${i}.jpg`"
                alt=""
              />
            </v-list-item-avatar>
            <v-list-item-title v-text="player.name" />
          </v-list-item>
        </v-list>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar
      app
      clipped-left
      dark
      dense
      color="primary"
      :extension-height="expandedSearch ? 80 : 0"
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />

      <v-btn icon class="mx-2">
        <hero-image :hero="balanar" version="icon" width="28" height="28" />
      </v-btn>

      <v-toolbar-title class="mr-4 align-center">
        <router-link
          :to="{ name: 'home' }"
          class="title app-title grey--text text--darken-4"
        >
          Balanar
        </router-link>
      </v-toolbar-title>

      <v-spacer />

      <v-col cols="6" lg="4" xl="4" v-if="!isXSmall">
        <v-text-field
          single-line
          hide-details
          v-model="query"
          color="white"
          placeholder="Search..."
          append-icon="mdi-magnify"
          ref="search"
          v-on:keyup.esc="clearSearch"
        />
      </v-col>

      <v-btn icon v-if="isXSmall" @click.stop="toggleSearch">
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <template v-slot:extension>
        <v-expand-transition>
          <v-text-field
            clearable
            single-line
            hide-details
            v-model="query"
            color="white"
            placeholder="Search..."
            ref="expandableSearch"
            v-show="expandedSearch"
            @click:clear="toggleSearch"
          />
        </v-expand-transition>
      </template>
    </v-app-bar>

    <v-content>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import { mapState } from "vuex";
import HeroImage from "@/components/HeroImage.vue";

export default {
  name: "App",

  components: {
    HeroImage
  },

  data: () => ({
    drawer: null,
    focusSearch: false,
    query: null,
    followed: [{ name: "13k", account_id: 13, picture: 28 }]
  }),

  created() {
    this.$vuetify.theme.dark = true;
    this.$store.dispatch("heroes/fetch");
    this.$store.dispatch("liveMatches/watch");
    this.$store.dispatch("liveMatches/fetch");
  },

  computed: {
    ...mapState({
      balanar: state => state.heroes.byName["npc_dota_hero_night_stalker"]
    }),
    isXSmall() {
      return this.$vuetify.breakpoint.name === "xs";
    },
    expandedSearch() {
      return this.isXSmall && this.focusSearch;
    }
  },

  methods: {
    toggleSearch() {
      this.focusSearch = !this.focusSearch;

      if (this.$refs.expandableSearch && this.focusSearch) {
        setTimeout(() => {
          this.$refs.expandableSearch.focus();
        }, 1);
      }
    },
    clearSearch() {
      if (this.$refs.search) {
        this.$refs.search.reset();
      }

      if (this.$refs.expandableSearch) {
        this.$refs.expandableSearch.reset();
        this.$refs.expandableSearch.blur();
        this.toggleSearch();
      }
    }
  }
};
</script>

<style lang="scss">
a {
  text-decoration: none;
}

.mono {
  font-family: "Roboto Mono", monospace !important;
}

.app-title {
  text-shadow: 1px 0 #666;
}
</style>
