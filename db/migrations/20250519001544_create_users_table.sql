-- migrate:up

CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  steam_id VARCHAR(17) UNIQUE NOT NULL,
  name VARCHAR(255) NOT NULL,
  avatar TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);

-- migrate:down

DROP TABLE users