<template>
  <div class="player-page">
    <router-link :to="{ name: 'home' }">Home</router-link>

    <div v-if="loading" class="loading">Loading...</div>

    <div v-if="error" class="error">
      {{ error }}
    </div>

    <transition name="slide">
      <section v-if="player" class="content" :key="player.account_id">
        <header class="profile">
          <div class="persona">
            <img
              v-if="player.avatar_medium_url"
              class="player-avatar"
              :src="player.avatar_medium_url"
            />
            <span class="player-name">{{ player.name }}</span>
            <a
              class="player-opendota"
              :href="odPlayerURL"
              :title="`View ${player.name} on OpenDota`"
              target="_blank"
            >
              [opendota]
            </a>
          </div>

          <div v-if="player.team" class="team">
            <img
              v-if="player.team.logo_url"
              class="team-logo"
              :src="player.team.logo_url"
              :title="player.team.name"
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
          </div>
        </header>

        <player-matches class="history" :matches="player.matches" />
      </section>
    </transition>
  </div>
</template>

<script>
import { get } from "lodash/object";

import api from "@/api";
import PlayerMatches from "@/components/PlayerMatches.vue";

const transformPlayer = (player, { heroes }) => {
  player.matches = player.matches || [];

  player.matches.forEach(match => {
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
          this.loading = false;
          this.player = transformPlayer(player, this.$store.state);
        })
        .catch(err => {
          this.error = err.toString();
        });
    }
  }
};
</script>

<style lang="scss" scoped>
.loading {
  position: absolute;
  top: 10px;
  right: 10px;
}

.error {
  color: red;
}

.content {
  transition: all 0.35s ease;
  position: absolute;
}

.slide-enter {
  opacity: 0;
  transform: translate(30px, 0);
}

.slide-leave-active {
  opacity: 0;
  transform: translate(-30px, 0);
}

.content {
  width: 800px;
  margin: 2em;
}

.profile {
  display: flex;
  flex-direction: row;
  padding: 0 16px;
  border-bottom: 1px solid #444;

  .persona {
    display: flex;
    flex-direction: row;
    align-items: center;
    flex-grow: 2;
  }

  .player-avatar {
    max-width: 64px;
    margin-right: 8px;
    box-shadow: 2px 2px 4px #000;
  }

  .player-name {
    align-self: flex-end;
    font-weight: 700;
    font-size: 42px;
    text-shadow: 1px 1px 1px #000;
  }

  .player-opendota {
    align-self: flex-end;
    margin-left: 1em;
    margin-bottom: 4px;
  }

  .team {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-right: 8px;
  }

  .team-logo {
    max-width: 64px;
  }

  .team-name {
    font-size: 16px;
    font-weight: bold;
    text-decoration: none;
  }
}
</style>
