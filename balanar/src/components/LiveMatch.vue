<template>
  <v-card height="100%">
    <v-card-title>
      {{ match.match_id }}

      <div
        v-if="hasMMR"
        class="grey--text ml-3 subtitle-2"
      >
        {{ match.average_mmr }} MMR
      </div>
    </v-card-title>

    <v-card-subtitle>
      <kbd>
        {{ watchCommand }}

        <ClipboardBtn
          :content="watchCommand"
          :success="onClipboardSuccess"
          :error="onClipboardError"
        />
      </kbd>

      <v-list
        v-if="match.game_time > 0"
        dense
        disabled
      >
        <v-list-item
          :two-line="match.delay > 0"
          dense
        >
          <v-list-item-icon class="mr-3">
            <v-icon>mdi-clock-outline</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ match.game_time | colonDuration }}</v-list-item-title>

            <v-list-item-subtitle v-if="match.delay > 0">
              {{ match.delay | humanDuration({verbose: true}) }} delay
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>

        <v-list-item dense>
          <v-list-item-icon class="mr-3">
            <v-icon>mdi-scoreboard-outline</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            {{ match.radiant_score }} - {{ match.dire_score }}
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-card-subtitle>

    <v-divider class="mx-4" />

    <v-container>
      <v-row
        v-for="team in teams"
        :key="team.number"
      >
        <v-col
          :cols="team | playersColWidth"
          :class="team | playersColClasses"
          order="1"
          class="players-col"
        >
          <v-list
            link
            dense
          >
            <v-list-item
              v-for="player in team.players"
              :key="player.account_id"
              :to="player | playerRoute"
              class="player"
            >
              <LiveMatchPlayer
                :team="team"
                :player="player"
              />
            </v-list-item>
          </v-list>
        </v-col>

        <v-col
          v-if="team.tag || team.name"
          :order="team.number % 2 === 0 ? 1 : 0"
          cols="3"
          align-self="center"
          class="d-flex flex-column justify-center align-center"
        >
          <img
            v-if="team.logo_url"
            :src="team.logo_url"
            :title="team.name"
            class="team-logo"
          >

          <span class="team-name caption">
            {{ team.tag || team.name }}
          </span>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script>
import * as $t from "@/protocol/transform";
import pb from "@/protocol/proto";
import { playerRoute } from "@/router/helpers";
import { colonDuration, humanDuration } from "@/filters";
import ClipboardBtn from "@/components/ClipboardBtn.vue";
import LiveMatchPlayer from "@/components/LiveMatchPlayer.vue";

export default {
  name: "LiveMatch",

  components: {
    ClipboardBtn,
    LiveMatchPlayer,
  },

  filters: {
    colonDuration,
    humanDuration,
    playerRoute,
    playersColWidth(team) {
      return team.tag || team.name ? 9 : 12;
    },
    playersColClasses(team) {
      if (!(team.tag || team.name)) {
        return {};
      }

      return {
        left: team.number % 2 === 0,
        right: team.number % 2 !== 0,
      };
    },
  },

  props: {
    match: {
      type: pb.protocol.LiveMatch,
      required: true,
    },
  },

  computed: {
    teams() {
      return $t.get(this.match, "teams");
    },
    hasMMR() {
      return this.match.average_mmr > 0;
    },
    watchCommand() {
      return `watch_server ${this.match.server_steam_id}`;
    },
  },

  methods: {
    onClipboardSuccess() {
      this.$store.commit("liveMatches/showClipboardNotification", {
        type: "success",
        text: "Command copied to clipboard",
      });
    },
    onClipboardError() {
      this.$store.commit("liveMatches/showClipboardNotification", {
        type: "error",
        text: "Failed to copy command to clipboard",
      });
    },
  },
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
