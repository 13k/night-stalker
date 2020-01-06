import { isPlainObject } from "lodash/lang";

import pb from "@/protocol/proto";
import { preprocessMatches } from "@/protocol/preprocess";
import { transformMatches } from "@/protocol/transform";

export function handleLiveMatchesChange(state, ev) {
  if (!isPlainObject(ev.body)) {
    throw new TypeError("received message with non-object body", ev.body);
  }

  preprocessMatches(ev.body.change.matches);

  const liveMatchesChange = pb.protocol.LiveMatchesChange.fromObject(ev.body);

  transformMatches(liveMatchesChange.change.matches, state);

  return liveMatchesChange;
}
