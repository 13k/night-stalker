import "@/logger";
import "@/protocol/configure";

import Vue from "vue";

import vuetify from "@/plugins/vuetify";
import store from "@/store";
import router from "@/router";
import App from "@/App.vue";

import "@mdi/font/css/materialdesignicons.css";
import "typeface-roboto";
import "typeface-roboto-mono";

Vue.config.productionTip = false;
Vue.config.performance = process.env.NODE_ENV === "development";

new Vue({
  store,
  router,
  vuetify,
  render: h => h(App),
}).$mount("#app");
