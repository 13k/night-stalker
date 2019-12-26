CREATE TABLE followed_players (
  id bigserial PRIMARY KEY,
  account_id bigint NOT NULL,
  label character varying(255) NOT NULL,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

CREATE INDEX idx_followed_players_deleted_at ON followed_players USING btree (deleted_at);
CREATE UNIQUE INDEX uix_followed_players_account_id ON followed_players USING btree (account_id);
