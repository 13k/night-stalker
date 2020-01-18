CREATE TABLE match_players (
  id bigserial PRIMARY KEY,
  match_id bigint NOT NULL,
  account_id bigint NOT NULL,
  hero_id bigint,
  player_slot integer,
  pro_name character varying(255),
  kills integer,
  deaths integer,
  assists integer,
  items bigint[],
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

CREATE UNIQUE INDEX uix_match_players_match_id_account_id ON match_players USING btree (match_id, account_id);
CREATE INDEX idx_match_players_deleted_at ON match_players USING btree (deleted_at);
