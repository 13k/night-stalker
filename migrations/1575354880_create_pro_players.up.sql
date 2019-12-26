CREATE TABLE pro_players (
  id bigserial PRIMARY KEY,
  account_id bigint NOT NULL,
  team_id bigint NOT NULL,
  fantasy_role bigint,
  is_locked boolean,
  locked_until timestamp with time zone,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

CREATE UNIQUE INDEX uix_pro_players_account_id ON pro_players USING btree (account_id);
CREATE INDEX idx_pro_players_team_id ON pro_players USING btree (team_id);
CREATE INDEX idx_pro_players_deleted_at ON pro_players USING btree (deleted_at);
