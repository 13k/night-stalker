ALTER TABLE live_match_stats_players
  ADD COLUMN match_id bigint;

UPDATE live_match_stats_players
  SET match_id = live_match_stats.match_id
  FROM live_match_stats
  WHERE live_match_stats.id = live_match_stats_players.live_match_stats_id;

ALTER TABLE live_match_stats_players
  ALTER COLUMN match_id SET NOT NULL;

CREATE INDEX idx_live_match_stats_players_match_id ON live_match_stats_players (match_id);
