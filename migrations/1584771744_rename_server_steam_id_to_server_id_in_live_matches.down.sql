ALTER TABLE live_matches
  ADD COLUMN server_steam_id bigint;

UPDATE live_matches
  SET server_steam_id = server_id;

ALTER TABLE live_matches
  ALTER COLUMN server_steam_id SET NOT NULL,
  DROP COLUMN server_id;

