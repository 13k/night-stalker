import _ from "lodash";

import pb from "@/protocol/proto";
import { PlayerSlot } from "@/dota2/flags";

import { get, set, transformProperty } from "./object";
import { transformMatchTimestamps } from "./match_common";

export function transformMatch(match, state) {
  transformMatchTimestamps(match);

  transformProperty(match, "outcome", outcome => {
    const radiantVictory = outcome === pb.ns.protocol.MatchOutcome.MATCH_OUTCOME_RAD_VICTORY;
    const direVictory = outcome === pb.ns.protocol.MatchOutcome.MATCH_OUTCOME_DIRE_VICTORY;
    return { radiantVictory, direVictory };
  });

  _.each(match.players, player => transformMatchPlayer(match, player, state));

  return match;
}

export function transformMatchPlayer(match, player, { heroes }) {
  const slot = new PlayerSlot(player.player_slot);
  const outcome = get(match, "outcome");
  const victory =
    (slot.isRadiant && outcome.radiantVictory) || (slot.isDire && outcome.direVictory);

  set(player, "slot", slot);
  set(player, "victory", victory);
  transformProperty(player, "hero_id", "hero", heroId => _.get(heroes, ["byId", heroId]));

  return player;
}
