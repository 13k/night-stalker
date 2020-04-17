ALTER TABLE players
  ADD COLUMN team_id bigint,
  ADD COLUMN fantasy_role bigint,
  ADD COLUMN is_locked boolean,
  ADD COLUMN locked_until timestamp with time zone;

UPDATE players
  SET (team_id, fantasy_role, is_locked, locked_until) = (
    SELECT team_id, fantasy_role, is_locked, locked_until
    FROM pro_players
    WHERE players.account_id = pro_players.account_id
  );

CREATE INDEX idx_players_team_id ON players (team_id);
