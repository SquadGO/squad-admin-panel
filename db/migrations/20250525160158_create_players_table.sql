-- migrate:up

CREATE TABLE players (
  player_id SERIAL PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  eos_id VARCHAR(32) UNIQUE NOT NULL,
  steam_id VARCHAR(17) UNIQUE NOT NULL,
  first_seen TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  last_seen TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);

-- migrate:down

DROP TABLE players