import Vue from "vue";
import _ from "lodash";
import ky from "ky";
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
  leagues: {
    index: compile("leagues"),
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

const MEDIA_TYPES = {
  json: "application/json",
  protobuf: "application/protobuf",
};

const RESPONSE_BODY_METHODS = {
  json: "json",
  protobuf: "arrayBuffer",
};

const debugResponse = (request, _options, response) => {
  log.debug(request.method, request.url, "->", response.status, response.statusText);
};

const beforeRequest = [];
const afterResponse = DEBUG ? [debugResponse] : [];

class API {
  constructor(baseURL) {
    this.client = ky.create({
      ...CLIENT_OPTIONS,
      prefixUrl: baseURL,
      hooks: {
        beforeRequest,
        afterResponse,
      },
    });
  }

  request(method, route, { type = "protobuf", ...options } = {}) {
    if (!_.has(MEDIA_TYPES, type)) {
      throw new TypeError(`invalid resource type '${type}'`);
    }

    const headers = { accept: MEDIA_TYPES[type] };
    const bodyMethod = RESPONSE_BODY_METHODS[type];
    const routeParams = _.get(route, "params", {});
    const toPath = _.get(ROUTES, route.name);
    const path = toPath(routeParams);
    const res = this.client(path, { method, headers, ...options });

    return res[bodyMethod].call(res);
  }

  get(route, options = {}) {
    return this.request("get", route, options);
  }

  heroes() {
    return this.get({ name: "heroes.index" });
  }

  heroMatches(id) {
    return this.get({ name: "heroes.matches", params: { id } });
  }

  leagues(id) {
    const searchParams = new URLSearchParams();

    if (_.isString(id)) {
      searchParams.set("id", id);
    } else if (_.isArray(id)) {
      _.each(id, id => searchParams.append("id", id));
    }

    return this.get({ name: "leagues.index" }, { searchParams });
  }

  liveMatches() {
    return this.get({ name: "live_matches.index" });
  }

  playerMatches(accountId) {
    return this.get({ name: "players.matches", params: { accountId } });
  }

  search(query) {
    return this.get({ name: "search.index" }, { searchParams: { q: query } });
  }
}

export default new API(API_URL);
