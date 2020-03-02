import _ from "lodash";

import pb from "@/protocol/proto";

export const MATCH_TIMESTAMP_FIELDS = [
  "activate_time",
  "deactivate_time",
  "last_update_time",
  "start_time",
  "created_at",
  "updated_at",
];

// https://github.com/protobufjs/protobuf.js/issues/677
// https://github.com/protobufjs/protobuf.js/issues/893
// https://github.com/protobufjs/protobuf.js/pull/1258
export function preprocessMatch(match) {
  _.each(MATCH_TIMESTAMP_FIELDS, field => {
    const value = match[field];

    if (_.isString(value)) {
      var dt = Date.parse(value);

      if (isNaN(dt)) {
        throw TypeError(`${field}: invalid timestamp: ${value}`);
      }

      match[field] = new pb.google.protobuf.Timestamp({
        seconds: Math.floor(dt / 1000),
        nanos: (dt % 1000) * 1000000,
      });
    }
  });

  return match;
}

export function preprocessMatches(matches) {
  return _.map(matches, preprocessMatch);
}
