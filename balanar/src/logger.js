import Vue from "vue";
import VueLog from "@dreipol/vue-log";
import { presets } from "@dreipol/vue-log/src/presets";

const LEVEL = process.env.NODE_ENV === "production" ? "warn" : "debug";
const LEVEL_INDEX = presets.levels.findIndex(l => l.name === LEVEL);

const options = {
  filter({ config, level }) {
    const logIndex = config.levels.findIndex(l => l.name === level.name);
    return logIndex >= LEVEL_INDEX;
  }
};

Vue.use(VueLog, options);
