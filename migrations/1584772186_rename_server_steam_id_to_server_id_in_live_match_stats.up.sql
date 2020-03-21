ALTER TABLE live_match_stats
  ADD COLUMN server_id bigint;

UPDATE live_match_stats
  SET server_id = server_steam_id;

ALTER TABLE live_match_stats
  ALTER COLUMN server_id SET NOT NULL,
  DROP COLUMN server_steam_id;

