CREATE TABLE live_match_stats_buildings (
  id bigserial PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  live_match_stats_id bigint NOT NULL,
  game_team bigint,
  heading numeric,
  type bigint,
  lane bigint,
  tier bigint,
  pos_x numeric,
  pos_y numeric,
  destroyed boolean
);

CREATE INDEX idx_live_match_stats_buildings_live_match_stats_id ON live_match_stats_buildings USING btree (live_match_stats_id);
CREATE INDEX idx_live_match_stats_buildings_deleted_at ON live_match_stats_buildings USING btree (deleted_at);
