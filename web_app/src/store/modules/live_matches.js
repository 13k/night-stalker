import Vue from "vue";
import _ from "lodash";

import api from "@/api";
import { protocol as enums } from "@/protocol/enums_pb";

const log = Vue.log({ context: { location: "store/liveMatches" } });

const TEAM_ATTRIBUTE_NAME_TPL = _.template("<%= side %>_team_<%= attr %>");
const TEAM_ATTRIBUTES = ["id", "name", "tag", "logo", "logo_url"];
const TEAM_SIDES = {
  [enums.GameTeam.GAME_TEAM_GOODGUYS]: "radiant",
  [enums.GameTeam.GAME_TEAM_BADGUYS]: "dire"
};

const state = {
  all: []
};

const getters = {};

function getTeamAttributes(match, side) {
  return _.transform(
    TEAM_ATTRIBUTES,
    (attrs, attr) => {
      const attrName = TEAM_ATTRIBUTE_NAME_TPL({ side, attr });
      attrs[attr] = match[attrName];
    },
    {}
  );
}

function createMatchTeam(match, number, players) {
  const side = TEAM_SIDES[number];
  return _.assign({ number, players, side }, getTeamAttributes(match, side));
}

function transformMatches(matches, { heroes }) {
  return _.map(matches || [], match => {
    match.players = _.chain(match.players)
      .map(player => {
        player.hero = _.get(heroes, ["byId", player.hero_id]);
        return player;
      })
      .sortBy("player_slot")
      .value();

    match.teams = _.chain(match.players)
      .groupBy("team")
      .toPairs()
      .sortBy(([number]) => number)
      .transform((teams, [number, players]) => {
        teams[number] = createMatchTeam(match, _.toInteger(number), players);
      }, {})
      .value();

    return match;
  });
}

const actions = {
  watch({ commit, rootState }) {
    log.debug("<watch>", rootState.ws);

    rootState.ws.addEventListener("message", ev => {
      log.debug("<watch:message>", ev);
      commit("setLiveMatches", transformMatches(ev.body, rootState));
    });
  },

  fetch({ commit, rootState }) {
    log.debug("<fetch>");

    api.getLiveMatches().then(matches => {
      log.debug("<fetch:response>", matches);
      commit("setLiveMatches", transformMatches(matches, rootState));
    });
  }
};

const mutations = {
  setLiveMatches(state, matches) {
    state.all = matches;
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
