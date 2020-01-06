import { isPlainObject } from "lodash/lang";

import api from "@/api";
import pb from "@/protocol/proto";
import { preprocessMatches } from "@/protocol/preprocess";
import { transformPlayer } from "@/protocol/transform";
import { transformMatches } from "@/protocol/transform";

export function getPlayer(state, accountId) {
  return api.getPlayer(accountId).then(res => {
    preprocessMatches(res.matches);

    const player = pb.protocol.Player.fromObject(res);

    transformPlayer(player, state);

    return player;
  });
}

export function getLiveMatches(state) {
  return api.getLiveMatches().then(res => {
    if (!isPlainObject(res)) {
      throw new TypeError("received non-object response", res);
    }

    preprocessMatches(res.matches);

    const liveMatches = pb.protocol.LiveMatches.fromObject(res);

    transformMatches(liveMatches.matches, state);

    return liveMatches;
  });
}
