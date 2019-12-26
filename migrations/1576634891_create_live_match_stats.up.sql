CREATE TABLE live_match_stats (
  id bigserial PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  match_id bigint NOT NULL,
  server_steam_id bigint NOT NULL,
  league_id bigint,
  league_node_id bigint,
  game_timestamp bigint,
  game_time integer,
  game_mode bigint,
  game_state bigint,
  steam_broadcaster_account_ids bigint[],
  delta_frame boolean,
	graph_gold integer[],
	graph_xp integer[],
	graph_kill integer[],
	graph_tower integer[],
	graph_rax integer[]
);

CREATE INDEX idx_live_match_stats_match_id ON live_match_stats USING btree (match_id);
CREATE INDEX idx_live_match_stats_server_steam_id ON live_match_stats USING btree (server_steam_id);
CREATE INDEX idx_live_match_stats_deleted_at ON live_match_stats USING btree (deleted_at);
