CREATE TABLE steam_servers (
  id bigserial PRIMARY KEY,
  address character varying(255) NOT NULL,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

CREATE UNIQUE INDEX uix_steam_servers_address ON steam_servers USING btree (address);
CREATE INDEX idx_steam_servers_deleted_at ON steam_servers USING btree (deleted_at);
