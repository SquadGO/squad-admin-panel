-- migrate:up

CREATE TABLE bans (
  ban_id BIGSERIAL PRIMARY KEY,
  user_id INT REFERENCES users(user_id) ON DELETE NO ACTION NOT NULL,
  player_id INT REFERENCES players(player_id) ON DELETE NO ACTION NOT NULL,
  reason VARCHAR(255) NOT NULL,
  note TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  end_at TIMESTAMPTZ
)

-- migrate:down

DROP TABLE BANS