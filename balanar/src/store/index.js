import Vue from "vue";
import Vuex from "vuex";
import createLogger from "vuex/dist/logger";

import WS from "@/ws";
import * as modules from "./modules";

Vue.use(Vuex);

const DEBUG = process.env.NODE_ENV !== "production";
const WS_URL = `${DEBUG ? "ws" : "wss"}://${window.location.host}/ws`;

const plugins = DEBUG ? [createLogger()] : [];
const ws = new WS(WS_URL);

ws.connect();

export default new Vuex.Store({
  strict: DEBUG,
  state: { ws },
  mutations: {},
  actions: {},
  modules,
  plugins,
});
