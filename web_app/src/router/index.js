import Vue from "vue";
import VueRouter from "vue-router";

import Home from "@/views/Home.vue";
import PlayerPage from "@/views/PlayerPage.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: Home
  },
  {
    path: "/about",
    name: "about",
    component: () => import(/* webpackChunkName: "about" */ "@/views/About.vue")
  },
  {
    path: "/players/:accountId",
    name: "players.show",
    component: PlayerPage
  }
];

export default new VueRouter({
  routes
});
