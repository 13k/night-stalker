<template>
  <v-expansion-panel>
    <v-expansion-panel-header class="justify-start">
      <HeroImage
        :hero="hero"
        version="icon"
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
          target="_blank"
          :max-width="communitySiteIconSize"
          :max-height="communitySiteIconSize"
          class="ml-1"
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
                  <v-list-item-title>{{ match.average_mmr }}</v-list-item-title>
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
                  <v-list-item-title>{{ playerSlot | teamSideName }}</v-list-item-title>
                  <v-list-item-subtitle>Side</v-list-item-subtitle>
                </v-list-item-content>
              </v-list-item>

              <v-list-item>
                <v-list-item-icon>
                  <v-icon>mdi-sword</v-icon>
                </v-list-item-icon>

                <v-list-item-content>
                  <v-list-item-title>
                    {{ match.player_details.kills }}/{{ match.player_details.deaths }}/{{ match.player_details.assists }}
                  </v-list-item-title>
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
import {
  l10n,
  colonDuration,
  teamSideName,
  opendotaMatchURL,
  dotabuffMatchURL,
  stratzMatchURL,
  datdotaMatchURL,
} from "@/components/filters";

import * as t from "@/protocol/transform";
import CommunitySiteBtn from "@/components/CommunitySiteBtn.vue";
import HeroImage from "@/components/HeroImage.vue";

export default {
  name: "PlayerMatch",

  components: {
    CommunitySiteBtn,
    HeroImage,
  },

  filters: {
    l10n,
    colonDuration,
    teamSideName,
  },

  props: {
    match: {
      type: Object,
      required: true,
    },
  },

  computed: {
    hero() {
      return t.get(this.match, "hero");
    },
    playerSlot() {
      return t.get(this.match, "slot");
    },
    date() {
      return t.get(this.match, "start_time") || t.get(this.match, "activate_time");
    },
    outcome() {
      const outcome = t.get(this.match, "outcome");

      if (!outcome) {
        return null;
      }

      return {
        text: outcome.playerVictory ? "Victory" : "Defeat",
        color: outcome.playerVictory ? "green" : "red",
        icon: outcome.playerVictory ? "mdi-trophy" : "mdi-bomb",
      };
    },
    communitySiteIconSize() {
      return this.$vuetify.breakpoint.xsOnly ? 22 : 28;
    },
    communitySitesMatch() {
      const sites = [
        {
          site: "opendota",
          url: opendotaMatchURL(this.match),
          text: "View match on OpenDota",
        },
        {
          site: "dotabuff",
          url: dotabuffMatchURL(this.match),
          text: "View match on Dotabuff",
        },
        {
          site: "stratz",
          url: stratzMatchURL(this.match),
          text: "View match on Stratz",
        },
      ];

      if (!this.match.league_id.isZero()) {
        sites.push({
          site: "datdota",
          url: datdotaMatchURL(this.match),
          text: "View match on DatDota",
        });
      }

      return sites;
    },
  },
};
</script>
