CREATE TABLE live_match_stats_teams (
  id bigserial PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  live_match_stats_id bigint NOT NULL,
  game_team bigint,
  team_id bigint,
  name character varying(255),
  tag character varying(255),
  logo_id bigint,
  logo_url text,
  score bigint,
  net_worth bigint
);

CREATE INDEX idx_live_match_stats_teams_live_match_stats_id ON live_match_stats_teams USING btree (live_match_stats_id);
CREATE INDEX idx_live_match_stats_teams_team_id ON live_match_stats_teams USING btree (team_id);
CREATE INDEX idx_live_match_stats_teams_deleted_at ON live_match_stats_teams USING btree (deleted_at);
