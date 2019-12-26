<template>
  <div class="live-match">
    <h3 class="header">
      {{ match.match_id }}<br />
      <span class="watch-command monospace">
        watch_server {{ match.server_steam_id }}
      </span>
    </h3>

    <div class="info" v-if="hasInfo">
      <p v-if="hasMMR">MMR: {{ match.average_mmr }}</p>

      <p class="teams" v-if="hasTeamTags">
        <img
          class="team-logo"
          v-if="match.radiant_team_logo_url"
          :src="match.radiant_team_logo_url"
          :title="match.radiant_team_name"
        />

        <span class="tags">
          {{ match.radiant_team_tag }} vs {{ match.dire_team_tag }}
        </span>

        <img
          class="team-logo"
          v-if="match.dire_team_logo_url"
          :src="match.dire_team_logo_url"
          :title="match.dire_team_name"
        />
      </p>
    </div>

    <div class="players">
      <live-match-player
        v-for="player in match.players"
        :key="player.account_id"
        :player="player"
        :side="teamSides[player.team]"
      />
    </div>
  </div>
</template>

<script>
import _ from "lodash";

import filters from "@/components/filters";
import LiveMatchPlayer from "@/components/LiveMatchPlayer.vue";

export default {
  name: "live-match",
  filters,
  components: {
    LiveMatchPlayer
  },
  props: {
    match: Object
  },
  computed: {
    hasInfo() {
      return this.hasTeamTags || this.hasMMR;
    },
    hasMMR() {
      return this.match.average_mmr > 0;
    },
    hasTeamTags() {
      return this.match.radiant_team_tag && this.match.dire_team_tag;
    },
    teamSides() {
      return _.chain(this.match.players)
        .map("team")
        .sortBy()
        .sortedUniq()
        .take(2)
        .zipObject(["left", "right"])
        .value();
    }
  }
};
</script>

<style lang="scss" scoped>
.live-match {
  margin: 20px;
  width: 300px;
  border: 1px solid #000;
  border-radius: 4px;
  box-shadow: 1px 1px #444;
  background-color: #000d;
  color: #ddd;
}

.header {
  font-size: 18px;
  margin: 6px 0;
  padding-bottom: 4px;
  border-bottom: 1px solid #bbb;
  text-align: center;

  .watch-command {
    font-size: small;
    font-weight: normal;
    font-style: italic;
    text-align: center;
  }
}

.info {
  font-size: 14px;
  text-align: center;
  padding-top: 4px;
  padding-bottom: 4px;
  border-bottom: 1px solid #bbb;

  > p {
    margin: 4px 0;
  }

  .teams {
    display: flex;
    align-items: center;

    .tags {
      flex-grow: 2;
    }

    .team-logo {
      max-height: 48px;
      margin-left: 4px;
      margin-right: 4px;
    }
  }
}

.players {
  padding: 8px;

  ul {
    padding-left: 0;
  }

  li {
    margin: 2px 0 2px 0;
    list-style-type: none;
  }
}
</style>
