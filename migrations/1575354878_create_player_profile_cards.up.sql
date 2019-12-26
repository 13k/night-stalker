CREATE TABLE player_profile_cards (
  id bigserial PRIMARY KEY,
  account_id bigint NOT NULL,
  background_def_index bigint,
  badge_points bigint,
  event_points bigint,
  event_id bigint,
  rank_tier bigint,
  leaderboard_rank bigint,
  rank_tier_score bigint,
  previous_rank_tier bigint,
  rank_tier_mmr_type bigint,
  rank_tier_core bigint,
  rank_tier_core_score bigint,
  leaderboard_rank_core bigint,
  rank_tier_support bigint,
  rank_tier_support_score bigint,
  leaderboard_rank_support bigint,
  is_plus_subscriber boolean,
  plus_original_start_date timestamp with time zone,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

CREATE UNIQUE INDEX uix_player_profile_cards_account_id ON player_profile_cards USING btree (account_id);
CREATE INDEX idx_player_profile_cards_deleted_at ON player_profile_cards USING btree (deleted_at);
