ALTER TABLE live_match_players
  ADD COLUMN match_id bigint;

UPDATE live_match_players
  SET match_id = live_matches.match_id
  FROM live_matches
  WHERE live_matches.id = live_match_players.live_match_id;

ALTER TABLE live_match_players
  ALTER COLUMN match_id SET NOT NULL;

CREATE UNIQUE INDEX uix_live_match_players_match_id_account_id ON live_match_players (match_id, account_id);
