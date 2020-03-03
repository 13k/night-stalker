<template>
  <v-expansion-panel>
    <v-expansion-panel-header class="justify-start">
      <HeroImage
        :hero="hero"
        orientation="icon"
        max-width="32"
        max-height="32"
        class="mr-6"
      />

      <div class="d-flex flex-column">
        <div class="d-flex align-center">
          <v-icon
            v-if="outcome"
            :color="outcome.color"
            :title="outcome.text"
            class="mr-2"
            small
          >
            {{ outcome.icon }}
          </v-icon>

          <span class="subtitle-2">{{ match.match_id }}</span>
        </div>

        <span class="caption">{{ date | l10n }}</span>
      </div>

      <v-spacer />

      <div class="flex-grow-0 d-flex flex-column flex-sm-row mr-4">
        <CommunitySiteBtn
          v-for="link in communitySitesMatch"
          :key="link.site"
          :site="link.site"
          :href="link.url"
          :max-width="communitySiteIconSize"
          :max-height="communitySiteIconSize"
          class="ml-1"
          target="_blank"
        />
      </div>
    </v-expansion-panel-header>

    <v-expansion-panel-content>
      <v-container>
        <v-row>
          <v-col
            cols="12"
            md="6"
          >
            <v-list two-line>
              <v-list-item>
                <v-list-item-icon>
                  <v-icon>mdi-shield-star-outline</v-icon>
                </v-list-item-icon>

                <v-list-item-content>
                  <v-list-item-title>{{ avgMMR }}</v-list-item-title>
                  <v-list-item-subtitle>Average MMR</v-list-item-subtitle>
                </v-list-item-content>
              </v-list-item>

              <v-list-item>
                <v-list-item-icon>
                  <v-icon>mdi-clock-outline</v-icon>
                </v-list-item-icon>

                <v-list-item-content>
                  <v-list-item-title>{{ match.duration | colonDuration }}</v-list-item-title>
                  <v-list-item-subtitle>Duration</v-list-item-subtitle>
                </v-list-item-content>
              </v-list-item>

              <v-list-item>
                <v-list-item-icon>
                  <v-icon>mdi-scoreboard-outline</v-icon>
                </v-list-item-icon>

                <v-list-item-content>
                  <v-list-item-title>{{ match.radiant_score }} - {{ match.dire_score }}</v-list-item-title>
                  <v-list-item-subtitle>Score</v-list-item-subtitle>
                </v-list-item-content>
              </v-list-item>
            </v-list>
          </v-col>

          <v-col
            cols="12"
            md="6"
          >
            <v-list two-line>
              <v-list-item>
                <v-list-item-icon>
                  <v-icon>mdi-checkerboard</v-icon>
                </v-list-item-icon>

                <v-list-item-content>
                  <v-list-item-title>{{ teamSideName }}</v-list-item-title>
                  <v-list-item-subtitle>Side</v-list-item-subtitle>
                </v-list-item-content>
              </v-list-item>

              <v-list-item>
                <v-list-item-icon>
                  <v-icon>mdi-sword</v-icon>
                </v-list-item-icon>

                <v-list-item-content>
                  <v-list-item-title>{{ kda.text }}</v-list-item-title>
                  <v-list-item-subtitle>KDA</v-list-item-subtitle>
                </v-list-item-content>
              </v-list-item>
            </v-list>
          </v-col>
        </v-row>
      </v-container>
    </v-expansion-panel-content>
  </v-expansion-panel>
</template>

<script>
import _ from "lodash";

import * as $t from "@/protocol/transform";
import * as $f from "@/components/filters";
import pb from "@/protocol/proto";
import CommunitySiteBtn from "@/components/CommunitySiteBtn.vue";
import HeroImage from "@/components/HeroImage.vue";

export default {
  name: "PlayerMatch",

  components: {
    CommunitySiteBtn,
    HeroImage,
  },

  filters: _.pick($f, "colonDuration", "l10n"),

  props: {
    player: {
      type: pb.protocol.Player,
      required: true,
    },
    match: {
      type: pb.protocol.Match,
      required: true,
    },
    knownPlayers: {
      type: Array,
      default: () => [],
      validator: v => _.every(v, i => i instanceof pb.protocol.Player),
    },
  },

  computed: {
    matchPlayer() {
      return $t.get(this.match, "poi");
    },
    hero() {
      return $t.get(this.matchPlayer, "hero");
    },
    avgMMR() {
      return this.match.average_mmr !== 0 ? this.match.average_mmr : "N/A";
    },
    teamSideName() {
      const slot = $t.get(this.matchPlayer, "slot");
      return slot ? $f.teamSideName(slot) : "N/A";
    },
    date() {
      return $t.get(this.match, "start_time") || $t.get(this.match, "activate_time");
    },
    outcome() {
      if (this.matchPlayer == null) {
        return null;
      }

      const victory = $t.get(this.matchPlayer, "victory");

      return {
        text: victory ? "Victory" : "Defeat",
        color: victory ? "green" : "red",
        icon: victory ? "mdi-trophy" : "mdi-bomb",
      };
    },
    kda() {
      const { kills = 0, deaths = 0, assists = 0 } = this.matchPlayer || {};
      const kda = {
        kills,
        deaths,
        assists,
        text: "N/A",
      };

      if (kda.kills > 0 || kda.deaths > 0 || kda.assists > 0) {
        kda.text = `${kda.kills}/${kda.deaths}/${kda.assists}`;
      }

      return kda;
    },
    communitySiteIconSize() {
      return this.$vuetify.breakpoint.xsOnly ? 22 : 28;
    },
    communitySitesMatch() {
      const sites = [
        {
          site: "opendota",
          url: $f.opendotaMatchURL(this.match),
          text: "View match on OpenDota",
        },
        {
          site: "dotabuff",
          url: $f.dotabuffMatchURL(this.match),
          text: "View match on Dotabuff",
        },
        {
          site: "stratz",
          url: $f.stratzMatchURL(this.match),
          text: "View match on Stratz",
        },
      ];

      if (!this.match.league_id.isZero()) {
        sites.push({
          site: "datdota",
          url: $f.datdotaMatchURL(this.match),
          text: "View match on DatDota",
        });
      }

      return sites;
    },
  },
};
</script>
