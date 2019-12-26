CREATE TABLE live_match_stats_players (
  id bigserial PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  live_match_stats_id bigint NOT NULL,
  hero_id bigint,
  account_id bigint,
  player_slot bigint,
  name character varying(255),
  game_team bigint,
  level bigint,
  kills bigint,
  deaths bigint,
  assists bigint,
  denies bigint,
  last_hits bigint,
  gold bigint,
  pos_x numeric,
  pos_y numeric,
  net_worth bigint,
  abilities bigint[],
  items bigint[]
);

CREATE INDEX idx_live_match_stats_players_live_match_stats_id ON live_match_stats_players USING btree (live_match_stats_id);
CREATE INDEX idx_live_match_stats_players_account_id ON live_match_stats_players USING btree (account_id);
CREATE INDEX idx_live_match_stats_players_hero_id ON live_match_stats_players USING btree (hero_id);
CREATE INDEX idx_live_match_stats_players_deleted_at ON live_match_stats_players USING btree (deleted_at);
