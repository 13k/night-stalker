export function heroParam(hero) {
  return `${hero.id}-${hero.slug}`;
}

export function heroRoute(hero) {
  return {
    name: "heroes.show",
    params: { id: heroParam(hero) },
  };
}

export function playerParam(player) {
  return `${player.account_id}-${player.slug}`;
}

export function playerRoute(player) {
  return {
    name: "players.show",
    params: { accountId: playerParam(player) },
  };
}
