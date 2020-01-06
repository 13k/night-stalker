import Vue from "vue";
import _ from "lodash";

import pb from "@/protocol/proto";
import { getLiveMatches } from "@/protocol/api";
import { handleLiveMatchesChange } from "@/protocol/ws";

const log = Vue.log({ context: { location: "store/liveMatches" } });

const state = {
  all: []
};

const getters = {};

const actions = {
  watch({ commit, rootState }) {
    log.debug("<watch>", rootState.ws);

    rootState.ws.addEventListener("message", ev => {
      const liveMatchesChange = handleLiveMatchesChange(rootState, ev);

      log.debug("<watch:message>", liveMatchesChange);

      let mutation;

      switch (liveMatchesChange.op) {
        case pb.protocol.LiveMatchesChange.Op.REPLACE:
          mutation = "setLiveMatches";
          break;
        case pb.protocol.LiveMatchesChange.Op.UPDATE:
          mutation = "updateLiveMatches";
          break;
        default:
          log.error(
            "<watch:message> unknown LiveMatchesChange.op:",
            liveMatchesChange.op
          );
          break;
      }

      commit(mutation, liveMatchesChange.change);
    });
  },

  fetch({ commit, rootState }) {
    log.debug("<fetch>");

    getLiveMatches(rootState).then(liveMatches => {
      log.debug("<fetch:response>", liveMatches);
      commit("setLiveMatches", liveMatches);
    });
  }
};

const mutations = {
  setLiveMatches(state, { matches }) {
    state.all = matches;
  },
  updateLiveMatches(state, { matches }) {
    _.each(matches, match => {
      let idx = _.findIndex(state.all, { match_id: match.match_id });
      let delCount = 1;

      if (idx < 0) {
        idx = _.sortedIndexBy(state.all, match, m => -m.sort_score);
        delCount = 0;
      }

      state.all.splice(idx, delCount, match);
    });
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
