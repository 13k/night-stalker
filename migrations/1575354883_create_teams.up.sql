CREATE TABLE teams (
  id bigserial PRIMARY KEY,
  name character varying(255) NOT NULL,
  tag character varying(255) NOT NULL,
  rating numeric,
  wins bigint,
  losses bigint,
  last_match_time timestamp with time zone,
  logo_url text,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

CREATE INDEX idx_teams_deleted_at ON teams USING btree (deleted_at);
