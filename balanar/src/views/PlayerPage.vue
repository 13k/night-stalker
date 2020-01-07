<template>
  <div class="player-page">
    <router-link :to="{ name: 'home' }">
      <v-tooltip bottom>
        <template v-slot:activator="{ on }">
          <v-icon v-on="on">
            mdi-arrow-left-bold-circle-outline
          </v-icon>
        </template>
        <span>Home</span>
      </v-tooltip>
    </router-link>

    <v-skeleton-loader
      v-if="loading"
      type="list-item-avatar-three-line"
    />

    <v-alert
      v-if="error"
      type="error"
    >
      {{ error }}
    </v-alert>

    <transition name="slide">
      <v-container
        v-if="player"
        :key="player.account_id"
        class="content"
      >
        <v-row>
          <v-col cols="12">
            <section class="profile">
              <v-row justify="start">
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
                    class="player-avatar"
                    max-width="64"
                    max-height="64"
                    contain
                  />

                  <span class="display-1 mx-4 d-flex d-sm-none">{{ player.name }}</span>

                  <span class="display-2 mx-4 d-none d-sm-flex d-lg-none">{{ player.name }}</span>

                  <span class="display-3 mx-4 d-none d-lg-flex">{{ player.name }}</span>

                  <span class="order-first order-sm-last mr-2 mr-sm-0 d-inline-flex flex-column flex-sm-row flex-md-row flex-lg-row flex-xl-row">
                    <CommunitySiteBtn
                      site="opendota"
                      :href="player | opendotaPlayerURL"
                      :alt="`View ${player.name} on OpenDota`"
                      target="_blank"
                      width="20"
                      height="20"
                    />

                    <CommunitySiteBtn
                      site="dotabuff"
                      :href="player | dotabuffPlayerURL"
                      :alt="`View ${player.name} on Dotabuff`"
                      target="_blank"
                      width="20"
                      height="20"
                    />

                    <CommunitySiteBtn
                      site="stratz"
                      :href="player | stratzPlayerURL"
                      :alt="`View ${player.name} on Stratz`"
                      target="_blank"
                      width="20"
                      height="20"
                    />

                    <CommunitySiteBtn
                      v-if="player.is_pro"
                      site="datdota"
                      :href="player | datdotaPlayerURL"
                      :alt="`View ${player.name} on DatDota`"
                      target="_blank"
                      width="20"
                      height="20"
                    />
                  </span>
                </v-col>

                <v-spacer />

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
                  <div class="d-flex">
                    <div class="d-inline-flex flex-column justify-center align-center">
                      <v-img
                        v-if="player.team.logo_url"
                        :src="player.team.logo_url"
                        :title="player.team.name"
                        max-width="64"
                        max-height="64"
                        contain
                      />

                      <span class="text-center">{{ player.team.name }}</span>
                    </div>

                    <div class="d-flex flex-column ml-3 justify-center">
                      <CommunitySiteBtn
                        site="opendota"
                        :href="player.team | opendotaTeamURL"
                        :alt="`View ${player.team.name} on OpenDota`"
                        target="_blank"
                        width="20"
                        height="20"
                      />

                      <CommunitySiteBtn
                        site="dotabuff"
                        :href="player.team | dotabuffTeamURL"
                        :alt="`View ${player.team.name} on Dotabuff`"
                        target="_blank"
                        width="20"
                        height="20"
                      />

                      <CommunitySiteBtn
                        site="datdota"
                        :href="player.team | datdotaTeamURL"
                        :alt="`View ${player.team.name} on DatDota`"
                        target="_blank"
                        width="20"
                        height="20"
                      />
                    </div>
                  </div>
                </v-col>
              </v-row>
            </section>
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12">
            <section class="history">
              <PlayerMatches :matches="player.matches" />
            </section>
          </v-col>
        </v-row>
      </v-container>
    </transition>
  </div>
</template>

<script>
import { getPlayer } from "@/protocol/api";
import filters from "@/components/filters";
import CommunitySiteBtn from "@/components/CommunitySiteBtn.vue";
import PlayerMatches from "@/components/PlayerMatches.vue";

export default {
  name: "PlayerPage",

  components: {
    CommunitySiteBtn,
    PlayerMatches,
  },

  filters,

  data() {
    return {
      loading: false,
      player: null,
      error: null,
    };
  },

  watch: {
    $route: "fetchData",
  },

  created() {
    this.fetchData();
  },

  methods: {
    fetchData() {
      this.error = this.player = null;
      this.loading = true;

      getPlayer(this.$store.state, this.$route.params.accountId)
        .then(player => {
          this.player = player;
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

.profile {
  .player-avatar {
    box-shadow: 2px 2px 4px #000;
  }
}
</style>
