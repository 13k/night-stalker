<template>
  <v-list
    link
    dense
  >
    <v-subheader
      v-if="teamLabel"
      class="mb-3"
    >
      <v-img
        v-if="team.logo_url"
        :src="team.logo_url"
        :title="team.name"
        max-width="64"
        class="mr-3"
        contain
      />

      <span class="caption">{{ teamLabel }}</span>

      <v-divider class="ml-3" />
    </v-subheader>

    <v-list-item
      v-for="player in team.players"
      :key="player.account_id"
      :to="player | playerRoute"
    >
      <LiveMatchPlayer
        :match="match"
        :team="team"
        :player="player"
      />
    </v-list-item>
  </v-list>
</template>

<script>
import pb from "@/protocol/proto";
import { playerRoute } from "@/router/helpers";
import LiveMatchPlayer from "@/components/LiveMatchPlayer.vue";

export default {
  name: "LiveMatchTeam",

  components: {
    LiveMatchPlayer,
  },

  filters: {
    playerRoute,
  },

  props: {
    match: {
      type: pb.protocol.LiveMatch,
      required: true,
    },
    team: {
      type: Object,
      required: true,
    },
  },

  computed: {
    teamLabel() {
      return this.team.tag || this.team.name;
    },
  },
};
</script>
