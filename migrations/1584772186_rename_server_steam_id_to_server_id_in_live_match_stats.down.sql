ALTER TABLE live_match_stats
  ADD COLUMN server_steam_id bigint;

UPDATE live_match_stats
  SET server_steam_id = server_id;

ALTER TABLE live_match_stats
  ALTER COLUMN server_steam_id SET NOT NULL,
  DROP COLUMN server_id;


