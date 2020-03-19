ALTER TABLE player_profile_cards
  DROP COLUMN background_def_index,
  DROP COLUMN rank_tier_core,
  DROP COLUMN rank_tier_core_score,
  DROP COLUMN rank_tier_support,
  DROP COLUMN rank_tier_support_score,
  DROP COLUMN leaderboard_rank_core,
  DROP COLUMN leaderboard_rank_support,
  ADD COLUMN plus_start_at timestamp with time zone;

UPDATE player_profile_cards
  SET plus_start_at = plus_original_start_date;

ALTER TABLE player_profile_cards
  DROP COLUMN plus_original_start_date;
