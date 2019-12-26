CREATE TABLE live_match_stats_picksbans (
  id bigserial PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  live_match_stats_id bigint NOT NULL,
  hero_id bigint,
  game_team bigint,
  is_ban boolean
);

CREATE INDEX idx_live_match_stats_picksbans_live_match_stats_id ON live_match_stats_picksbans USING btree (live_match_stats_id);
CREATE INDEX idx_live_match_stats_picksbans_hero_id ON live_match_stats_picksbans USING btree (hero_id);
CREATE INDEX idx_live_match_stats_picksbans_deleted_at ON live_match_stats_picksbans USING btree (deleted_at);
