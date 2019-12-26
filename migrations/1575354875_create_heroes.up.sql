CREATE TABLE heroes (
  id bigserial PRIMARY KEY,
  name character varying(255) NOT NULL,
  localized_name character varying(255) NOT NULL,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

CREATE UNIQUE INDEX uix_heroes_name ON heroes USING btree (name);
CREATE INDEX idx_heroes_deleted_at ON heroes USING btree (deleted_at);
