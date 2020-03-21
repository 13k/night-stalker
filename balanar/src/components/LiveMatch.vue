<template>
  <v-card
    class="d-flex flex-column"
    height="100%"
  >
    <template v-if="league">
      <v-img
        :src="league.id | leagueImageURL"
        :aspect-ratio="512/200"
        min-height="148"
        height="148"
        max-height="148"
        class="white--text align-end"
      >
        <v-card-title>
          <span class="title--shadow">{{ match.match_id }}</span>
        </v-card-title>
      </v-img>
    </template>

    <template v-else>
      <v-card-title>{{ match.match_id }}</v-card-title>
    </template>

    <v-card-subtitle
      v-if="league"
      class="d-flex"
    >
      <v-img
        :src="leagueIcon"
        max-width="16"
        class="mr-2"
        title="Tournament"
        contain
      />

      {{ league.name }}
    </v-card-subtitle>

    <v-card-subtitle
      v-else-if="match.average_mmr > 0"
      class="d-flex"
    >
      <v-img
        :src="mmrIcon"
        max-width="16"
        class="mr-2"
        title="Ranked match"
        contain
      />

      {{ match.average_mmr }} MMR
    </v-card-subtitle>

    <v-card-subtitle
      v-else-if="match.weekend_tourney_tournament_id > 0"
      class="d-flex"
    >
      <v-img
        :src="battlecupIcon"
        max-width="16"
        class="mr-2"
        title="Battlecup"
        contain
      />

      Battlecup
    </v-card-subtitle>

    <LiveMatchTeam
      v-for="team in teams"
      :key="team.number"
      :match="match"
      :team="team"
    />

    <v-spacer />

    <v-card-actions>
      <v-spacer />

      <ClipboardBtn
        :content="watchCommand"
        :success="onClipboardSuccess"
        :error="onClipboardError"
        title="Copy command to clipboard"
        icon
      >
        <v-icon>mdi-console</v-icon>
      </ClipboardBtn>

      <v-btn
        :title="expand ? 'Hide details' : 'Show details'"
        icon
        @click="expand = !expand"
      >
        <v-icon>{{ expand ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
      </v-btn>
    </v-card-actions>

    <v-expand-transition>
      <v-list
        v-show="expand"
        color="primary"
        two-line
        dense
        dark
      >
        <v-list-item>
          <v-list-item-icon>
            <v-icon>mdi-console-line</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ watchCommand }}</v-list-item-title>
            <v-list-item-subtitle>Console command</v-list-item-subtitle>
          </v-list-item-content>

          <v-list-item-action>
            <ClipboardBtn
              :content="watchCommand"
              :success="onClipboardSuccess"
              :error="onClipboardError"
              title="Copy command to clipboard"
              icon
              small
            />
          </v-list-item-action>
        </v-list-item>

        <v-list-item
          v-if="match.game_time > 0"
          :two-line="match.delay > 0"
        >
          <v-list-item-icon>
            <v-icon>mdi-clock-outline</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ match.game_time | colonDuration }}</v-list-item-title>

            <v-list-item-subtitle v-if="match.delay > 0">
              {{ match.delay | humanDuration({verbose: true}) }} delay
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>

        <v-list-item v-if="match.game_time > 0">
          <v-list-item-icon>
            <v-icon>mdi-scoreboard-outline</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ match.radiant_score }} - {{ match.dire_score }}</v-list-item-title>
            <v-list-item-subtitle>Score</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-expand-transition>
  </v-card>
</template>

<script>
import * as $t from "@/protocol/transform";
import pb from "@/protocol/proto";
import { image } from "@/assets/helpers";
import { colonDuration, humanDuration, leagueImageURL } from "@/filters";
import ClipboardBtn from "@/components/ClipboardBtn.vue";
import LiveMatchTeam from "@/components/LiveMatchTeam.vue";

export default {
  name: "LiveMatch",

  components: {
    ClipboardBtn,
    LiveMatchTeam,
  },

  filters: {
    colonDuration,
    humanDuration,
    leagueImageURL,
  },

  props: {
    match: {
      type: pb.ns.protocol.LiveMatch,
      required: true,
    },
  },

  data: () => ({
    expand: false,
    mmrIcon: image("match_making/icon_mmr_medium.png"),
    leagueIcon: image("leagues/icon_league.png"),
    battlecupIcon: image("battlecup/battlecup_icon.png"),
  }),

  computed: {
    teams() {
      return $t.get(this.match, "teams");
    },
    league() {
      return $t.get(this.match, "league");
    },
    watchCommand() {
      return `watch_server ${this.match.server_id}`;
    },
  },

  methods: {
    onClipboardSuccess() {
      this.$store.commit("snackbar/show", {
        type: "success",
        text: "Command copied to clipboard",
      });
    },
    onClipboardError() {
      this.$store.commit("snackbar/show", {
        type: "error",
        text: "Failed to copy command to clipboard",
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.title--shadow {
  text-shadow: 1px 3px 3px rgba(0, 0, 0, 0.8);
}
</style>
