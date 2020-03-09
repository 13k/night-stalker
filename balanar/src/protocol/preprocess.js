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

export const LEAGUE_TIMESTAMP_FIELDS = [
  "last_activity_at",
  "start_at",
  "finish_at",
  "created_at",
  "updated_at",
];

// https://github.com/protobufjs/protobuf.js/issues/677
// https://github.com/protobufjs/protobuf.js/issues/893
// https://github.com/protobufjs/protobuf.js/pull/1258
export function preprocessTimestamps(object, fields) {
  _.each(fields, field => {
    const value = object[field];

    if (_.isString(value)) {
      var dt = Date.parse(value);

      if (isNaN(dt)) {
        throw TypeError(`${field}: invalid timestamp: ${value}`);
      }

      object[field] = new pb.google.protobuf.Timestamp({
        seconds: Math.floor(dt / 1000),
        nanos: (dt % 1000) * 1000000,
      });
    }
  });

  return object;
}

export function preprocessMatch(match) {
  return preprocessTimestamps(match, MATCH_TIMESTAMP_FIELDS);
}

export function preprocessMatches(matches) {
  return _.map(matches, preprocessMatch);
}

export function preprocessLeague(league) {
  return preprocessTimestamps(league, LEAGUE_TIMESTAMP_FIELDS);
}

export function preprocessLeagues(leagues) {
  return _.map(leagues, preprocessLeague);
}
