import Vue from "vue";
import ky from "ky";
import _ from "lodash";
import { compile } from "path-to-regexp";

const log = Vue.log({ context: { location: "api" } });

const DEBUG = process.env.NODE_ENV !== "production";

const CLIENT_OPTIONS = {
  timeout: false,
  retry: 0,
};

const API_URL = "/api/v1";

const ROUTES = {
  heroes: {
    index: compile("heroes"),
    matches: compile("heroes/:id/matches"),
  },
  live_matches: {
    index: compile("live_matches"),
  },
  players: {
    matches: compile("players/:accountId/matches"),
  },
  search: {
    index: compile("search"),
  },
};

const debugResponse = (request, _options, response) => {
  log.debug(request.method, request.url, "->", response.status, response.statusText);
};

const beforeRequest = [];
const afterResponse = DEBUG ? [debugResponse] : [];

class API {
  constructor(baseURL) {
    this.client = ky.create(
      Object.assign({}, CLIENT_OPTIONS, {
        prefixUrl: baseURL,
        hooks: {
          beforeRequest,
          afterResponse,
        },
      })
    );
  }

  request(method, route, options) {
    const routeParams = _.get(route, "params", {});
    const toPath = _.get(ROUTES, route.name);
    const path = toPath(routeParams);

    return this.client[method].call(this.client, path, options).json();
  }

  get(route, options) {
    return this.request("get", route, options);
  }

  search(query) {
    const route = { name: "search.index" };
    return this.get(route, { searchParams: { q: query } });
  }

  liveMatches() {
    const route = { name: "live_matches.index" };
    return this.get(route);
  }

  heroes() {
    const route = { name: "heroes.index" };
    return this.get(route);
  }

  playerMatches(accountId) {
    const route = { name: "players.matches", params: { accountId } };
    return this.get(route);
  }

  heroMatches(id) {
    const route = { name: "heroes.matches", params: { id } };
    return this.get(route);
  }
}

export default new API(API_URL);
