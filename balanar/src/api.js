import Vue from "vue";
import ky from "ky";
import { compile } from "path-to-regexp";
import { get } from "lodash/object";

const log = Vue.log({ context: { location: "api" } });

const DEBUG = process.env.NODE_ENV !== "production";
const API_URL = "/api/v1";

const ROUTES = {
  heroes: {
    index: compile("heroes"),
  },
  live_matches: {
    index: compile("live_matches"),
  },
  players: {
    show: compile("players/:accountId"),
  },
};

const debugResponse = (request, _options, response) => {
  log.debug(request.method, request.url, "->", response.status, response.statusText);
};

const beforeRequest = [];
const afterResponse = DEBUG ? [debugResponse] : [];

class API {
  constructor(baseURL) {
    this.client = ky.create({
      prefixUrl: baseURL,
      hooks: {
        beforeRequest,
        afterResponse,
      },
    });
  }

  request(method, route, options) {
    const routeParams = get(route, "params", {});
    const toPath = get(ROUTES, route.name);
    const path = toPath(routeParams);

    return this.client[method].call(this.client, path, options).json();
  }

  get(route, options) {
    return this.request("get", route, options);
  }

  getLiveMatches() {
    const route = { name: "live_matches.index" };
    return this.get(route);
  }

  getHeroes() {
    const route = { name: "heroes.index" };
    return this.get(route);
  }

  getPlayer(accountId) {
    const route = { name: "players.show", params: { accountId } };
    return this.get(route);
  }
}

export default new API(API_URL);
