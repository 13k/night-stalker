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
        v-if="player"
        :key="player.account_id"
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
                <v-img
                  v-if="player.avatar_medium_url"
                  :src="player.avatar_medium_url"
                  :max-width="playerAvatarSize"
                  :max-height="playerAvatarSize"
                  class="player-avatar"
                  contain
                />

                <span class="title mx-4 d-flex d-sm-none">{{ player.name }}</span>
                <span class="display-2 mx-4 d-none d-sm-flex d-lg-none">{{ player.name }}</span>
                <span class="display-3 mx-4 d-none d-lg-flex">{{ player.name }}</span>

                <v-menu
                  open-on-click
                  bottom
                  offset-y
                >
                  <template v-slot:activator="{ on }">
                    <v-btn
                      color="accent"
                      icon
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
                      v-for="link in communitySitesPlayer"
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

              <v-col
                v-if="player.team"
                cols="4"
                sm="3"
                md="2"
                lg="2"
                xl="2"
                align-self="end"
                class="d-flex justify-end"
              >
                <div class="d-flex align-end">
                  <div class="d-inline-flex flex-column justify-center align-center">
                    <v-img
                      v-if="player.team.logo_url"
                      :src="player.team.logo_url"
                      :title="player.team.name"
                      :max-width="teamLogoSize"
                      :max-height="teamLogoSize"
                      contain
                    />

                    <span class="text-center caption d-flex d-sm-none">{{ player.team.name }}</span>
                    <span class="text-center body-2 d-none d-sm-flex d-lg-none">{{ player.team.name }}</span>
                    <span class="text-center subtitle-1 d-none d-lg-flex">{{ player.team.name }}</span>
                  </div>

                  <v-menu
                    open-on-click
                    left
                    offset-x
                  >
                    <template v-slot:activator="{ on }">
                      <v-btn
                        color="accent"
                        icon
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
                        v-for="link in communitySitesTeam"
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
                </div>
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
            <PlayerMatches
              :player="player"
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

import * as $f from "@/filters";
import { fetchPlayerMatches } from "@/protocol/api";
import CommunitySiteBtn from "@/components/CommunitySiteBtn.vue";
import PlayerMatches from "@/components/PlayerMatches.vue";

export default {
  name: "PlayerPage",

  components: {
    CommunitySiteBtn,
    PlayerMatches,
  },

  data() {
    return {
      loading: false,
      playerMatches: null,
      error: null,
    };
  },

  computed: {
    player() {
      return _.get(this.playerMatches, "player");
    },
    matches() {
      return _.get(this.playerMatches, "matches");
    },
    knownPlayers() {
      return _.get(this.playerMatches, "known_players");
    },
    playerAvatarSize() {
      return this.$vuetify.breakpoint.xsOnly ? 32 : 64;
    },
    teamLogoSize() {
      return this.$vuetify.breakpoint.xsOnly ? 32 : 64;
    },
    communitySiteIconSize() {
      return this.$vuetify.breakpoint.xsOnly ? 16 : 24;
    },
    communitySitesPlayer() {
      if (this.player == null) {
        return [];
      }

      const sites = [
        {
          site: "opendota",
          url: $f.opendotaPlayerURL(this.player),
          text: "View player on OpenDota",
        },
        {
          site: "dotabuff",
          url: $f.dotabuffPlayerURL(this.player),
          text: "View player on Dotabuff",
        },
        {
          site: "stratz",
          url: $f.stratzPlayerURL(this.player),
          text: "View player on Stratz",
        },
      ];

      if (this.player.is_pro) {
        sites.push({
          site: "datdota",
          url: $f.datdotaPlayerURL(this.player),
          text: "View player on DatDota",
        });
      }

      return sites;
    },
    communitySitesTeam() {
      if (this.player.team == null) {
        return [];
      }

      return [
        {
          site: "opendota",
          url: $f.opendotaTeamURL(this.player.team),
          text: "View team on OpenDota",
        },
        {
          site: "dotabuff",
          url: $f.dotabuffTeamURL(this.player.team),
          text: "View team on Dotabuff",
        },
        {
          site: "datdota",
          url: $f.datdotaTeamURL(this.player.team),
          text: "View team on DatDota",
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
      this.error = this.playerMatches = null;
      this.loading = true;

      fetchPlayerMatches(this.$store.state, parseInt(this.$route.params.accountId))
        .then(playerMatches => {
          this.playerMatches = playerMatches;
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

.player-avatar {
  box-shadow: 2px 2px 4px #000;
}
</style>
