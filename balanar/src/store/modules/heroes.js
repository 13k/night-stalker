import _ from "lodash";
import Vue from "vue";

import { fetchHeroes } from "@/protocol/api";

const log = Vue.log({ context: { location: "store/heroes" } });

const state = {
  byId: {},
  byName: {},
};

const getters = {};

const actions = {
  fetch({ commit }) {
    log.debug("<fetch>");

    fetchHeroes().then(heroes => {
      log.debug("<fetch> received response", heroes);
      commit("setHeroes", heroes || []);
    });
  },
};

const mutations = {
  setHeroes(state, heroes) {
    state.byId = _.keyBy(heroes, "id");
    state.byName = _.keyBy(heroes, "name");
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
