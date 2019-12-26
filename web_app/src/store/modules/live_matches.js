import Vue from "vue";
import { sortBy } from "lodash/collection";
import { get } from "lodash/object";

import api from "@/api";

const log = Vue.log({ context: { location: "store/liveMatches" } });

const state = {
  all: []
};

const getters = {};

function transformMatches(matches, { heroes }) {
  matches = matches || [];

  matches.forEach(match => {
    match.players = match.players || [];

    match.players.forEach(player => {
      player.hero = get(heroes, ["byId", player.hero_id]);
    });

    match.players = sortBy(match.players, "player_slot");
  });

  return matches;
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
