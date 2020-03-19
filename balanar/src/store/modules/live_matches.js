import Vue from "vue";
import _ from "lodash";

import pb from "@/protocol/proto";
import { fetchLiveMatches } from "@/protocol/api";
import { handleLiveMatchesChange } from "@/protocol/ws";

const log = Vue.log({ context: { location: "store/liveMatches" } });

const state = {
  all: [],
};

const getters = {};

const actions = {
  watch({ commit, rootState }) {
    log.debug("<watch>", rootState.ws);

    rootState.ws.addEventListener("message", ev => {
      handleLiveMatchesChange(rootState, ev).then(liveMatchesChange => {
        log.debug("<watch:message>", liveMatchesChange);

        let mutation;

        switch (liveMatchesChange.op) {
          case pb.ns.protocol.CollectionOp.COLLECTION_OP_REPLACE:
            mutation = "setLiveMatches";
            break;
          case pb.ns.protocol.CollectionOp.COLLECTION_OP_ADD:
            mutation = "addLiveMatches";
            break;
          case pb.ns.protocol.CollectionOp.COLLECTION_OP_UPDATE:
            mutation = "updateLiveMatches";
            break;
          case pb.ns.protocol.CollectionOp.COLLECTION_OP_REMOVE:
            mutation = "removeLiveMatches";
            break;
          default:
            log.error("<watch:message> unknown LiveMatchesChange.op:", liveMatchesChange.op);
            break;
        }

        commit(mutation, liveMatchesChange.change);
      });
    });
  },

  fetch({ commit, rootState }) {
    log.debug("<fetch>");

    fetchLiveMatches(rootState).then(liveMatches => {
      log.debug("<fetch:response>", liveMatches);
      commit("setLiveMatches", liveMatches);
    });
  },
};

function upsertLiveMatch(matches, match) {
  let idx = _.findIndex(matches, { match_id: match.match_id });
  let delCount = 1;

  if (idx < 0) {
    idx = _.sortedIndexBy(matches, match, m => -m.sort_score);
    delCount = 0;
  }

  matches.splice(idx, delCount, match);
}

const mutations = {
  setLiveMatches(state, { matches }) {
    state.all = matches;
  },
  addLiveMatches(state, { matches }) {
    _.each(matches, match => {
      upsertLiveMatch(state.all, match);
    });
  },
  updateLiveMatches(state, { matches }) {
    _.each(matches, match => {
      upsertLiveMatch(state.all, match);
    });
  },
  removeLiveMatches(state, { matches }) {
    _.each(matches, match => {
      let idx = _.findIndex(state.all, { match_id: match.match_id });

      if (idx >= 0) {
        state.all.splice(idx, 1);
      }
    });
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
