CREATE TABLE live_matches (
  id bigserial PRIMARY KEY,
  match_id bigint NOT NULL,
  server_steam_id bigint NOT NULL,
  lobby_id bigint NOT NULL,
  lobby_type bigint,
  league_id bigint,
  series_id bigint,
  game_mode bigint,
  average_mmr bigint,
  radiant_lead integer,
  radiant_team_id bigint,
  radiant_team_name character varying(255),
  radiant_team_logo bigint,
  radiant_score bigint,
  dire_team_id bigint,
  dire_team_name character varying(255),
  dire_team_logo bigint,
  dire_score bigint,
  delay bigint,
  activate_time timestamp with time zone,
  deactivate_time timestamp with time zone,
  last_update_time timestamp with time zone,
  game_time integer,
  spectators bigint,
  sort_score numeric,
  building_state bigint,
  weekend_tourney_tournament_id bigint,
  weekend_tourney_division bigint,
  weekend_tourney_skill_level bigint,
  weekend_tourney_bracket_round bigint,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

CREATE UNIQUE INDEX uix_live_matches_match_id ON live_matches USING btree (match_id);
CREATE UNIQUE INDEX uix_live_matches_lobby_id ON live_matches USING btree (lobby_id);
CREATE INDEX idx_live_matches_deleted_at ON live_matches USING btree (deleted_at);
