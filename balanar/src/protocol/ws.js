import pb from "@/protocol/proto";
import { decode } from "@/protocol/decode";
import { prefetchMatchesLeagues } from "@/protocol/prefetch";
import { transformLiveMatchesChange } from "@/protocol/transform/live_match";

export const handleLiveMatchesChange = (state, ev) =>
  decode(pb.ns.protocol.LiveMatchesChange, ev.data)
    .then(msg => transformLiveMatchesChange(msg, state))
    .then(msg => {
      if (msg.op === pb.ns.protocol.CollectionOp.COLLECTION_OP_REMOVE) {
        return msg;
      }

      return prefetchMatchesLeagues(msg.change.matches).then(() => msg);
    });
