import _ from "lodash";

import pb from "@/protocol/proto";
import { preprocessMatches } from "@/protocol/preprocess";
import { transformLiveMatchesChange } from "@/protocol/transform";

export function handleLiveMatchesChange(state, ev) {
  if (!_.isPlainObject(ev.body)) {
    throw new TypeError("received message with non-object body", ev.body);
  }

  preprocessMatches(ev.body.change.matches);

  const liveMatchesChange = pb.protocol.LiveMatchesChange.fromObject(ev.body);

  transformLiveMatchesChange(liveMatchesChange, state);

  return liveMatchesChange;
}
