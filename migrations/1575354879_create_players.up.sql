CREATE TABLE players (
  id bigserial PRIMARY KEY,
  account_id bigint NOT NULL,
  steam_id bigint,
  name character varying(255),
  persona_name character varying(255),
  avatar_url text,
  avatar_medium_url text,
  avatar_full_url text,
  profile_url text,
  country_code character varying(255),
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

CREATE UNIQUE INDEX uix_players_account_id ON players USING btree (account_id);
CREATE INDEX idx_players_deleted_at ON players USING btree (deleted_at);
