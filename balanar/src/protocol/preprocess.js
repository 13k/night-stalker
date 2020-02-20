import { each } from "lodash/collection";
import { isString } from "lodash/lang";

import pb from "@/protocol/proto";

export const MATCH_TIMESTAMP_FIELDS = [
  "start_time",
  "activate_time",
  "deactivate_time",
  "last_update_time",
];

// https://github.com/protobufjs/protobuf.js/issues/677
// https://github.com/protobufjs/protobuf.js/issues/893
export function preprocessMatches(matches) {
  each(matches, match => {
    each(MATCH_TIMESTAMP_FIELDS, field => {
      const value = match[field];

      if (isString(value)) {
        var dt = Date.parse(value);

        if (isNaN(dt)) {
          throw TypeError(`.protocol.LiveMatch.${field}: invalid timestamp: ${value}`);
        }

        match[field] = new pb.google.protobuf.Timestamp({
          seconds: Math.floor(dt / 1000),
          nanos: (dt % 1000) * 1000000,
        });
      }
    });
  });
}
