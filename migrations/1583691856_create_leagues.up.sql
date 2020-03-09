CREATE TABLE leagues (
	id bigint PRIMARY KEY,
	name varchar(255) NOT NULL,
	tier bigint NOT NULL,
	region bigint NOT NULL,
	status bigint NOT NULL,
	total_prize_pool bigint,
  last_activity_at timestamp with time zone,
  start_at timestamp with time zone,
  finish_at timestamp with time zone,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

CREATE INDEX idx_leagues_deleted_at ON leagues (deleted_at);
