<template>
  <v-card hover :color="cardColor" @click="toggle" height="100%">
    <v-card-title>
      {{ match.match_id }}

      <div v-if="hasMMR" class="grey--text ml-3 subtitle-2">
        {{ match.average_mmr }} MMR
      </div>
    </v-card-title>

    <v-card-subtitle>
      <kbd>watch_server {{ match.server_steam_id }}</kbd>
    </v-card-subtitle>

    <v-divider class="mx-4"></v-divider>

    <v-container>
      <v-row v-for="team in match.teams" :key="team.number">
        <v-col
          order="1"
          class="players-col"
          :cols="team | playersColWidth"
          :class="team | playersColClasses"
        >
          <v-list dense link>
            <v-list-item
              dense
              class="player"
              v-for="player in team.players"
              :key="player.account_id"
              :to="{
                name: 'players.show',
                params: { accountId: player.account_id }
              }"
            >
              <live-match-player :team="team" :player="player" />
            </v-list-item>
          </v-list>
        </v-col>

        <v-col
          v-if="team.tag || team.name"
          cols="3"
          align-self="center"
          class="d-flex flex-column justify-center align-center"
          :order="team.number % 2 === 0 ? 1 : 0"
        >
          <img
            class="team-logo"
            v-if="team.logo_url"
            :src="team.logo_url"
            :title="team.name"
          />

          <span class="team-name caption">
            {{ team.tag || team.name }}
          </span>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script>
import filters from "@/components/filters";
import LiveMatchPlayer from "@/components/LiveMatchPlayer.vue";

export default {
  name: "live-match",
  filters: {
    ...filters,
    playersColWidth(team) {
      return team.tag || team.name ? 9 : 12;
    },
    playersColClasses(team) {
      if (!(team.tag || team.name)) {
        return {};
      }

      return {
        left: team.number % 2 === 0,
        right: team.number % 2 !== 0
      };
    }
  },
  components: {
    LiveMatchPlayer
  },
  props: {
    match: Object,
    active: Boolean,
    toggle: Function
  },
  computed: {
    cardColor() {
      return this.active ? "primary" : "";
    },
    hasMMR() {
      return this.match.average_mmr > 0;
    }
  }
};
</script>

<style lang="scss" scoped>
.player {
  padding: 0;
}

.players-col {
  &.left {
    border-right: 1px solid #333;
  }

  &.right {
    border-left: 1px solid #333;
  }
}

.team-logo {
  max-width: 48px;
  max-height: 48px;
}

.team-name {
  color: black;
  text-align: center;
  line-height: 1em;
  margin-top: 8px;
}
</style>
