ALTER TABLE live_match_stats
  ADD COLUMN live_match_id bigint;

UPDATE live_match_stats
  SET live_match_id = live_matches.id
  FROM live_matches
  WHERE live_match_stats.match_id = live_matches.match_id;

ALTER TABLE live_match_stats
  ALTER COLUMN live_match_id SET NOT NULL;

CREATE INDEX idx_live_match_stats_live_match_id ON live_match_stats (live_match_id);
