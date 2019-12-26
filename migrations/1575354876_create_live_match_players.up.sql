CREATE TABLE live_match_players (
  id bigserial PRIMARY KEY,
  live_match_id bigint NOT NULL,
  account_id bigint NOT NULL,
  hero_id bigint,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

CREATE UNIQUE INDEX uix_live_match_players_live_match_id_account_id ON live_match_players USING btree (live_match_id, account_id);
CREATE INDEX idx_live_match_players_deleted_at ON live_match_players USING btree (deleted_at);
