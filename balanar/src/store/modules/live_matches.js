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
      handleLiveMatchesChange(rootState, ev).then(msg => {
        log.debug("<watch:message>", {
          op: pb.ns.protocol.CollectionOp[msg.op],
          matchIDs: _.map(msg.change.matches, m => m.match_id.toString()),
        });

        let mutation;

        switch (msg.op) {
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
            log.error("<watch:message> unknown LiveMatchesChange.op:", msg.op);
            break;
        }

        commit(mutation, msg.change);
      });
    });
  },

  fetch({ commit, rootState }) {
    log.debug("<fetch>");

    fetchLiveMatches(rootState).then(msg => {
      log.debug("<fetch:response>", msg);
      commit("setLiveMatches", msg);
    });
  },
};

function findMatchIndex(matches, match) {
  return _.findIndex(matches, m => m.match_id.equals(match.match_id));
}

function upsertLiveMatch(matches, match) {
  const lenBefore = matches.length;
  let idx = findMatchIndex(matches, match);
  let delCount = 1;

  if (idx < 0) {
    idx = _.sortedIndexBy(matches, match, m => -m.sort_score);
    delCount = 0;
  }

  matches.splice(idx, delCount, match);

  log.debug("upsertLiveMatch", {
    matchID: match.match_id.toString(),
    before: lenBefore,
    after: matches.length,
  });
}

function removeLiveMatch(matches, match) {
  const lenBefore = matches.length;
  const idx = findMatchIndex(matches, match);

  if (idx >= 0) {
    matches.splice(idx, 1);
  }

  log.debug("removeLiveMatch", {
    matchID: match.match_id.toString(),
    before: lenBefore,
    after: matches.length,
  });
}

const mutations = {
  setLiveMatches(state, { matches }) {
    state.all = matches;
  },
  addLiveMatches(state, { matches }) {
    _.each(matches, match => upsertLiveMatch(state.all, match));
  },
  updateLiveMatches(state, { matches }) {
    _.each(matches, match => upsertLiveMatch(state.all, match));
  },
  removeLiveMatches(state, { matches }) {
    _.each(matches, match => removeLiveMatch(state.all, match));
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
