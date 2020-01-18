CREATE TABLE matches (
  id bigint PRIMARY KEY,
  league_id bigint,
  series_type bigint,
  series_game integer,
  game_mode bigint,
  start_time timestamp with time zone,
  duration integer,
  outcome bigint,
  radiant_team_id bigint,
  radiant_team_name character varying(255),
  radiant_team_logo bigint,
  radiant_team_logo_url text,
  radiant_score integer,
  dire_team_id bigint,
  dire_team_name character varying(255),
  dire_team_logo bigint,
  dire_team_logo_url text,
  dire_score integer,
  weekend_tourney_tournament_id bigint,
  weekend_tourney_season_trophy_id bigint,
  weekend_tourney_division bigint,
  weekend_tourney_skill_level bigint,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

CREATE INDEX idx_matches_league_id ON matches USING btree (league_id);
CREATE INDEX idx_matches_start_time ON matches USING btree (start_time);
CREATE INDEX idx_matches_deleted_at ON matches USING btree (deleted_at);
