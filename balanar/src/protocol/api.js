import api from "@/api";
import pb from "@/protocol/proto";
import * as $t from "@/protocol/transform";
import { decode } from "@/protocol/decode";
import { prefetchMatchesLeagues } from "@/protocol/prefetch";

const {
  ns: {
    protocol: { Heroes, Leagues, LiveMatches, PlayerMatches, HeroMatches, Search },
  },
} = pb;

export const fetchHeroes = () =>
  api
    .heroes()
    .then(res => decode(Heroes, res))
    .then($t.transformHeroes);

export const fetchLeagues = id =>
  api
    .leagues(id)
    .then(res => decode(Leagues, res))
    .then($t.transformLeagues);

export const fetchLiveMatches = state =>
  api
    .liveMatches()
    .then(res => decode(LiveMatches, res))
    .then(msg => $t.transformLiveMatches(msg, state))
    .then(msg => prefetchMatchesLeagues(msg.matches).then(() => msg));

export const fetchPlayerMatches = (state, accountId) =>
  api
    .playerMatches(accountId)
    .then(res => decode(PlayerMatches, res))
    .then(msg => $t.transformPlayerMatches(msg, state));

export const fetchHeroMatches = (state, heroId) =>
  api
    .heroMatches(heroId)
    .then(res => decode(HeroMatches, res))
    .then(msg => $t.transformHeroMatches(msg, state));

export const search = (state, query) =>
  api
    .search(query)
    .then(res => decode(Search, res))
    .then(msg => $t.transformSearch(msg, state));
