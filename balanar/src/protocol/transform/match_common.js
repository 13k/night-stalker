import _ from "lodash";

import pb from "@/protocol/proto";

import { set } from "./object";
import { transformMatch } from "./match";
import { transformPlayer } from "./player";

export const MATCH_TIMESTAMP_FIELDS = [
  "activate_time",
  "deactivate_time",
  "last_update_time",
  "start_time",
  "created_at",
  "updated_at",
];

export function transformMatchTimestamps(match) {
  _.each(MATCH_TIMESTAMP_FIELDS, field => {
    const ts = match[field];

    if (ts instanceof pb.google.protobuf.Timestamp) {
      const date = new Date(Math.floor(ts.seconds * 1000 + ts.nanos / 1000000));
      set(match, field, date);
    }
  });

  return match;
}

export function transformMatchHistory(history, state) {
  _.each(history.matches, match => transformMatch(match, state));
  _.each(history.known_players, player => transformPlayer(player, state));

  return history;
}
