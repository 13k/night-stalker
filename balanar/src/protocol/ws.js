import _ from "lodash";

import pb from "@/protocol/proto";
import * as $t from "@/protocol/transform";
import { preprocessMatches } from "@/protocol/preprocess";
import { prefetchMatchesLeagues } from "@/protocol/prefetch";

export function handleLiveMatchesChange(state, ev) {
  if (!_.isPlainObject(ev.body)) {
    throw new TypeError("received message with non-object body", ev.body);
  }

  preprocessMatches(ev.body.change.matches);

  const liveMatchesChange = pb.ns.protocol.LiveMatchesChange.fromObject(ev.body);

  $t.transformLiveMatchesChange(liveMatchesChange, state);

  if (liveMatchesChange.op === pb.ns.protocol.CollectionOp.COLLECTION_OP_REMOVE) {
    return Promise.resolve(liveMatchesChange);
  }

  return prefetchMatchesLeagues(liveMatchesChange.change.matches).then(() => liveMatchesChange);
}
