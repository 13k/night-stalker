<template>
  <v-expansion-panel>
    <v-expansion-panel-header>
      <v-row>
        <v-col
          cols="12"
          sm="4"
          md="2"
          lg="2"
          xl="2"
        >
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

            <span class="subtitle-2">
              {{ match.match_id }}

              <ClipboardBtn
                :content="match.match_id.toString()"
                :success="onClipboardSuccess"
                :error="onClipboardError"
                icon
                small
              />
            </span>
          </div>

          <span class="caption">{{ date | l10n }}</span>
        </v-col>

        <v-col
          cols="12"
          sm="4"
          md="2"
          lg="2"
          xl="2"
        >
          <span class="subtitle-1">{{ player && player.name }}</span>
        </v-col>

        <v-spacer />

        <v-col
          cols="12"
          sm="4"
          md="2"
          lg="2"
          xl="2"
        >
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
        </v-col>
      </v-row>
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
import * as $f from "@/filters";
import pb from "@/protocol/proto";
import ClipboardBtn from "@/components/ClipboardBtn.vue";
import CommunitySiteBtn from "@/components/CommunitySiteBtn.vue";

export default {
  name: "HeroMatch",

  components: {
    ClipboardBtn,
    CommunitySiteBtn,
  },

  filters: _.pick($f, "colonDuration", "l10n"),

  props: {
    hero: {
      type: pb.ns.protocol.Hero,
      required: true,
    },
    match: {
      type: pb.ns.protocol.Match,
      required: true,
    },
    knownPlayers: {
      type: Array,
      default: () => [],
      validator: v => _.every(v, i => i instanceof pb.ns.protocol.Player),
    },
  },

  computed: {
    matchPlayer() {
      return $t.get(this.match, "poi");
    },
    player() {
      return (
        this.matchPlayer && _.find(this.knownPlayers, { account_id: this.matchPlayer.account_id })
      );
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
    avatarSize() {
      return this.$vuetify.breakpoint.xsOnly ? 18 : 32;
    },
  },

  methods: {
    onClipboardSuccess() {
      this.$store.commit("snackbar/show", {
        type: "success",
        text: `MatchID ${this.match.match_id} copied to clipboard`,
      });
    },
    onClipboardError() {
      this.$store.commit("snackbar/show", {
        type: "error",
        text: "Failed to copy MatchID to clipboard",
      });
    },
  },
};
</script>
