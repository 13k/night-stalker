<template>
  <v-app id="night-stalker">
    <v-navigation-drawer
      v-model="drawer"
      app
      clipped
      :mini-variant="miniDrawer"
    >
      <v-list
        dense
        nav
      >
        <v-list-item>
          <v-list-item-action v-if="miniDrawer">
            <v-btn
              icon
              title="Expand"
              @click.stop="toggleMiniDrawer"
            >
              <v-icon>mdi-chevron-right</v-icon>
            </v-btn>
          </v-list-item-action>

          <v-spacer />

          <v-list-item-action>
            <v-btn
              icon
              title="Collapse"
              @click.stop="toggleMiniDrawer"
            >
              <v-icon>mdi-chevron-left</v-icon>
            </v-btn>
          </v-list-item-action>
        </v-list-item>

        <v-divider />

        <v-list-item
          :to="{ name: 'home' }"
          title="Home"
          exact
        >
          <v-list-item-action>
            <v-icon>mdi-home</v-icon>
          </v-list-item-action>

          <v-list-item-content>
            <v-list-item-title>Home</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-list-item
          link
          title="Match history"
        >
          <v-list-item-action>
            <v-icon>mdi-history</v-icon>
          </v-list-item-action>

          <v-list-item-content>
            <v-list-item-title>History</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-subheader class="mt-4 grey--text text--darken-1 text-center">
          LIVE STALKED
        </v-subheader>

        <v-list>
          <v-list-item
            v-for="(player, i) in followed"
            :key="player.account_id"
            link
          >
            <v-list-item-avatar>
              <img :src="`https://randomuser.me/api/portraits/men/${i}.jpg`">
            </v-list-item-avatar>
            <v-list-item-title v-text="player.name" />
          </v-list-item>
        </v-list>
      </v-list>

      <template v-slot:append>
        <v-divider />

        <v-list>
          <v-list-item>
            <v-list-item-action>
              <ThemeToggle />
            </v-list-item-action>
          </v-list-item>
        </v-list>
      </template>
    </v-navigation-drawer>

    <v-app-bar
      app
      clipped-left
      dark
      color="primary"
    >
      <v-app-bar-nav-icon
        class="hidden-lg-and-up"
        @click="drawer = !drawer"
      />

      <v-toolbar-title>
        <div class="d-flex justify-left">
          <HeroImage
            :hero="balanar"
            orientation="icon"
            width="28"
            height="28"
            class="mx-4"
            :alt="appName"
          />

          <span class="title app-title grey--text text--darken-4">
            {{ appName }}
          </span>
        </div>
      </v-toolbar-title>

      <v-spacer />

      <v-col
        cols="6"
        lg="4"
        xl="4"
      >
        <AppSearch />
      </v-col>
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

import AppSearch from "@/components/AppSearch.vue";
import HeroImage from "@/components/HeroImage.vue";
import ThemeToggle from "@/components/ThemeToggle.vue";

export default {
  name: "App",

  components: {
    AppSearch,
    HeroImage,
    ThemeToggle,
  },

  data: () => ({
    appName: process.env.VUE_APP_NAME,
    drawer: null,
    miniDrawer: false,
    followed: [{ name: "13k", account_id: 13, picture: 28 }],
  }),

  computed: mapState({
    balanar: state => state.heroes.byName.npc_dota_hero_night_stalker,
  }),

  watch: {
    miniDrawer(val) {
      localStorage.setItem("balanar.drawer.mini", val);
    },
  },

  created() {
    this.$vuetify.theme.dark = localStorage.getItem("balanar.theme.dark") === "true";
    this.miniDrawer = localStorage.getItem("balanar.drawer.mini") === "true";

    this.$store.dispatch("heroes/fetch");
    this.$store.dispatch("liveMatches/watch");
    this.$store.dispatch("liveMatches/fetch");
  },

  methods: {
    toggleMiniDrawer() {
      this.miniDrawer = !this.miniDrawer;
    },
  },
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
