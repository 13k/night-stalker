const state = {
  show: false,
  type: "info",
  text: null,
  color: null,
  timeout: 1500,
};

const getters = {};

const actions = {};

const mutations = {
  show(state, { text, type = "info", color = null, timeout = 1500 }) {
    state.show = true;
    state.type = type;
    state.text = text;
    state.color = color;
    state.timeout = timeout;
  },
  hide(state) {
    state.show = false;
    state.type = "info";
    state.text = null;
    state.color = null;
    state.timeout = 1500;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
