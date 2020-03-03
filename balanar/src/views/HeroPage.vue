<template>
  <div>
    <v-row v-if="loading">
      <v-col
        cols="12"
        lg="8"
        offset-lg="2"
      >
        <v-skeleton-loader
          type="card-avatar"
          class="mb-3"
        />
        <v-skeleton-loader
          type="list-item-avatar-three-line"
          class="mb-3"
        />
        <v-skeleton-loader
          type="list-item-avatar-three-line"
          class="mb-3"
        />
      </v-col>
    </v-row>

    <v-row v-if="error">
      <v-col
        cols="12"
        lg="8"
        offset-lg="2"
      >
        <v-alert type="error">
          {{ error }}
        </v-alert>
      </v-col>
    </v-row>

    <transition name="slide">
      <section
        v-if="hero"
        :key="hero.id.toString()"
        class="content"
      >
        <v-row>
          <v-col
            cols="12"
            lg="8"
            offset-lg="2"
          >
            <v-row>
              <v-col
                cols="8"
                sm="9"
                md="10"
                lg="10"
                xl="10"
                class="d-flex align-end"
              >
                <HeroImage
                  :hero="hero"
                  :max-width="avatarWidth"
                  :max-height="avatarHeight"
                  class="avatar"
                  orientation="portrait"
                />

                <span class="title mx-4 d-flex d-sm-none">{{ hero.localized_name }}</span>
                <span class="display-2 mx-4 d-none d-sm-flex d-lg-none">{{ hero.localized_name }}</span>
                <span class="display-3 mx-4 d-none d-lg-flex">{{ hero.localized_name }}</span>

                <v-menu
                  open-on-click
                  bottom
                  offset-y
                >
                  <template v-slot:activator="{ on }">
                    <v-btn
                      icon
                      color="accent"
                      v-on="on"
                    >
                      <v-icon large>
                        mdi-menu-down
                      </v-icon>
                    </v-btn>
                  </template>

                  <v-list
                    dense
                    nav
                  >
                    <v-list-item
                      v-for="link in communitySitesHero"
                      :key="link.site"
                      :href="link.url"
                      target="_blank"
                    >
                      <v-list-item-icon>
                        <CommunitySiteBtn
                          :site="link.site"
                          :width="communitySiteIconSize"
                          :height="communitySiteIconSize"
                        />
                      </v-list-item-icon>

                      <v-list-item-title>{{ link.text }}</v-list-item-title>

                      <v-list-item-action>
                        <v-icon>mdi-open-in-new</v-icon>
                      </v-list-item-action>
                    </v-list-item>
                  </v-list>
                </v-menu>
              </v-col>
            </v-row>
          </v-col>
        </v-row>

        <v-row>
          <v-col
            cols="12"
            lg="8"
            offset-lg="2"
          >
            <HeroMatches
              :hero="hero"
              :matches="matches"
              :known-players="knownPlayers"
            />
          </v-col>
        </v-row>
      </section>
    </transition>
  </div>
</template>

<script>
import _ from "lodash";

import * as $f from "@/components/filters";
import { fetchHeroMatches } from "@/protocol/api";
import CommunitySiteBtn from "@/components/CommunitySiteBtn.vue";
import HeroImage from "@/components/HeroImage.vue";
import HeroMatches from "@/components/HeroMatches.vue";

export default {
  name: "HeroPage",

  components: {
    CommunitySiteBtn,
    HeroImage,
    HeroMatches,
  },

  data() {
    return {
      loading: false,
      heroMatches: null,
      error: null,
    };
  },

  computed: {
    hero() {
      return _.get(this.heroMatches, "hero");
    },
    matches() {
      return _.get(this.heroMatches, "matches");
    },
    knownPlayers() {
      return _.get(this.heroMatches, "known_players");
    },
    avatarWidth() {
      return this.$vuetify.breakpoint.xsOnly ? 40 : 80;
    },
    avatarHeight() {
      return this.$vuetify.breakpoint.xsOnly ? 53 : 106;
    },
    communitySiteIconSize() {
      return this.$vuetify.breakpoint.xsOnly ? 16 : 24;
    },
    communitySitesHero() {
      if (this.hero == null) {
        return [];
      }

      return [
        {
          site: "opendota",
          url: $f.opendotaHeroURL(this.hero),
          text: "View hero on OpenDota",
        },
        {
          site: "dotabuff",
          url: $f.dotabuffHeroURL(this.hero),
          text: "View hero on Dotabuff",
        },
        {
          site: "stratz",
          url: $f.stratzHeroURL(this.hero),
          text: "View hero on Stratz",
        },
      ];
    },
  },

  watch: {
    $route: "fetchData",
  },

  created() {
    this.fetchData();
  },

  methods: {
    fetchData() {
      this.error = this.heroMatches = null;
      this.loading = true;

      fetchHeroMatches(this.$store.state, parseInt(this.$route.params.id))
        .then(heroMatches => {
          this.heroMatches = heroMatches;
        })
        .catch(err => {
          this.error = err.toString();
        })
        .finally(() => {
          this.loading = false;
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.content {
  transition: all 0.35s ease;
}

.slide-enter {
  opacity: 0;
  transform: translate(30px, 0);
}

.slide-leave-active {
  opacity: 0;
  transform: translate(-30px, 0);
}

.avatar {
  box-shadow: 2px 2px 4px #000;
}
</style>
