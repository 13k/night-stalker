ALTER TABLE live_matches
  ADD COLUMN server_id bigint;

UPDATE live_matches
  SET server_id = server_steam_id;

ALTER TABLE live_matches
  ALTER COLUMN server_id SET NOT NULL,
  DROP COLUMN server_steam_id;
