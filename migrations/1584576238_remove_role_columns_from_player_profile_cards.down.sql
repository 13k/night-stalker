ALTER TABLE player_profile_cards
  ADD COLUMN background_def_index bigint,
  ADD COLUMN rank_tier_core bigint,
  ADD COLUMN rank_tier_core_score bigint,
  ADD COLUMN rank_tier_support bigint,
  ADD COLUMN rank_tier_support_score bigint,
  ADD COLUMN leaderboard_rank_core bigint,
  ADD COLUMN leaderboard_rank_support bigint,
  ADD COLUMN plus_original_start_date timestamp with time zone;

UPDATE player_profile_cards
  SET plus_original_start_date = plus_start_at;

ALTER TABLE player_profile_cards
  DROP COLUMN plus_start_at;

