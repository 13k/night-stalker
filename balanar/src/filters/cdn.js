import _ from "lodash";

import pb from "@/protocol/proto";

const {
  protocol: { CDNLeagueImageVersion },
} = pb;

const CDN_URL = "http://cdn.dota2.com";
const CDN_DOTA2_URL = `${CDN_URL}/apps/dota2`;

export function leagueImageURL(
  leagueId,
  version = CDNLeagueImageVersion.CDN_LEAGUE_IMAGE_VERSION_BANNER
) {
  if (_.isString(version)) {
    version = CDNLeagueImageVersion[`CDN_LEAGUE_IMAGE_VERSION_${version.toUpperCase()}`];
  }

  return `${CDN_DOTA2_URL}/images/leagues/${leagueId}/images/image_${version}.png`;
}
