import { normalizePlayerName } from "@/util";

export function heroParam(hero) {
  return `${hero.id}-${hero.name.replace(/^npc_dota_hero_/, "")}`;
}

export function heroRoute(hero) {
  return {
    name: "heroes.show",
    params: { id: heroParam(hero) },
  };
}

export function playerParam(player) {
  return `${player.account_id}-${normalizePlayerName(player.name)}`;
}

export function playerRoute(player) {
  return {
    name: "players.show",
    params: { id: playerParam(player) },
  };
}
