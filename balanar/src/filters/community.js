const OPENDOTA_URL = "https://www.opendota.com";
const DOTABUFF_URL = "https://www.dotabuff.com";
const STRATZ_URL = "https://www.stratz.com";
const DATDOTA_URL = "https://datdota.com";

export function opendotaMatchURL(match) {
  return `${OPENDOTA_URL}/matches/${match.match_id}`;
}

export function opendotaPlayerURL(player) {
  return `${OPENDOTA_URL}/players/${player.account_id}`;
}

export function opendotaTeamURL(team) {
  return `${OPENDOTA_URL}/teams/${team.id}`;
}

export function opendotaHeroURL(hero) {
  return `${OPENDOTA_URL}/heroes/${hero.id}`;
}

export function dotabuffMatchURL(match) {
  return `${DOTABUFF_URL}/matches/${match.match_id}`;
}

export function dotabuffPlayerURL(player) {
  return `${DOTABUFF_URL}/players/${player.account_id}`;
}

export function dotabuffTeamURL(team) {
  return `${DOTABUFF_URL}/esports/teams/${team.id}`;
}

export function dotabuffHeroURL(hero) {
  return `${DOTABUFF_URL}/heroes/${hero.name}`;
}

export function stratzMatchURL(match) {
  return `${STRATZ_URL}/matches/${match.match_id}`;
}

export function stratzPlayerURL(player) {
  return `${STRATZ_URL}/players/${player.account_id}`;
}

export function stratzHeroURL(hero) {
  return `${STRATZ_URL}/heroes/${hero.id}`;
}

export function datdotaMatchURL(match) {
  return `${DATDOTA_URL}/matches/${match.match_id}`;
}

export function datdotaPlayerURL(player) {
  return `${DATDOTA_URL}/players/${player.account_id}`;
}

export function datdotaTeamURL(team) {
  return `${DATDOTA_URL}/teams/${team.id}`;
}
