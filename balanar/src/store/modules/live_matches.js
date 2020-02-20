import Vue from "vue";
import _ from "lodash";

import pb from "@/protocol/proto";
import { getLiveMatches } from "@/protocol/api";
import { handleLiveMatchesChange } from "@/protocol/ws";

const log = Vue.log({ context: { location: "store/liveMatches" } });

const state = {
  all: [],
  clipboardNotification: {
    show: false,
    type: "success",
    text: "",
  },
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
        case pb.protocol.CollectionOp.REPLACE:
          mutation = "setLiveMatches";
          break;
        case pb.protocol.CollectionOp.ADD:
          mutation = "addLiveMatches";
          break;
        case pb.protocol.CollectionOp.UPDATE:
          mutation = "updateLiveMatches";
          break;
        case pb.protocol.CollectionOp.REMOVE:
          mutation = "removeLiveMatches";
          break;
        default:
          log.error("<watch:message> unknown LiveMatchesChange.op:", liveMatchesChange.op);
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
  showClipboardNotification(state, { type, text }) {
    state.clipboardNotification.show = true;
    state.clipboardNotification.type = type;
    state.clipboardNotification.text = text;
  },
  hideClipboardNotification(state) {
    state.clipboardNotification.show = false;
    state.clipboardNotification.type = "success";
    state.clipboardNotification.text = "";
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
