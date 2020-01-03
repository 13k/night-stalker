<template>
  <div class="player-page">
    <router-link :to="{ name: 'home' }">
      <v-tooltip bottom>
        <template v-slot:activator="{ on }">
          <v-icon v-on="on">mdi-arrow-left-bold-circle-outline</v-icon>
        </template>
        <span>Home</span>
      </v-tooltip>
    </router-link>

    <v-skeleton-loader v-if="loading" type="list-item-avatar-three-line" />

    <v-alert v-if="error" type="error">
      {{ error }}
    </v-alert>

    <transition name="slide">
      <v-container v-if="player" class="content" :key="player.account_id">
        <v-row>
          <v-col cols="12">
            <section class="profile">
              <v-row>
                <v-col>
                  <v-img
                    v-if="player.avatar_medium_url"
                    :src="player.avatar_medium_url"
                    class="player-avatar mr-4"
                    max-width="64"
                  />
                </v-col>

                <v-col>
                  <span class="player-name display-3">{{ player.name }}</span>
                </v-col>

                <v-col>
                  <a
                    class="player-opendota"
                    :href="odPlayerURL"
                    :title="`View ${player.name} on OpenDota`"
                    target="_blank"
                  >
                    [opendota]
                  </a>
                </v-col>
              </v-row>

              <v-row v-if="player.team" class="team">
                <v-col>
                  <v-img
                    v-if="player.team.logo_url"
                    class="team-logo"
                    max-width="64"
                    :src="player.team.logo_url"
                    :alt="player.team.name"
                  />

                  <a
                    v-if="odTeamURL"
                    class="team-name"
                    :href="odTeamURL"
                    :title="`View ${player.team.name} on OpenDota`"
                    target="_blank"
                  >
                    {{ player.team.tag }}
                  </a>
                </v-col>
              </v-row>
            </section>
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12">
            <section class="history">
              <player-matches :matches="player.matches" />
            </section>
          </v-col>
        </v-row>
      </v-container>
    </transition>
  </div>
</template>

<script>
import { each } from "lodash/collection";
import { get } from "lodash/object";

import api from "@/api";
import PlayerMatches from "@/components/PlayerMatches.vue";

const transformPlayer = (player, { heroes }) => {
  player.matches = each(player.matches || [], match => {
    match.hero = get(heroes, ["byId", match.hero_id]);
  });

  return player;
};

export default {
  name: "player-page",
  components: {
    PlayerMatches
  },
  data() {
    return {
      loading: false,
      player: null,
      error: null
    };
  },
  created() {
    this.fetchData();
  },
  watch: {
    $route: "fetchData"
  },
  computed: {
    odPlayerURL() {
      return `https://www.opendota.com/players/${this.player.account_id}`;
    },
    odTeamURL() {
      if (!this.player.team) {
        return null;
      }

      return `https://www.opendota.com/teams/${this.player.team.id}`;
    }
  },
  methods: {
    fetchData() {
      this.error = this.player = null;
      this.loading = true;

      api
        .getPlayer(this.$route.params.accountId)
        .then(player => {
          this.player = transformPlayer(player, this.$store.state);
        })
        .catch(err => {
          this.error = err.toString();
        })
        .finally(() => {
          this.loading = false;
        });
    }
  }
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

  .player-name {
    text-shadow: 1px 1px 1px #000;
  }

  .player-opendota {
    align-self: flex-end;
    margin-left: 1em;
    margin-bottom: 4px;
  }

  .team-name {
    font-size: 16px;
    font-weight: bold;
    text-decoration: none;
  }
}
</style>
