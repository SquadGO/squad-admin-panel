-- migrate:up

CREATE TABLE maps (
  map_id BIGSERIAL PRIMARY KEY,
  server_id INT REFERENCES servers(server_id) ON DELETE NO ACTION NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  end_at TIMESTAMPTZ,
  map_name TEXT NOT NULL,
  winner_name TEXT,
  winner_team_id INT,
  winner_tickets INT,
  loser_name TEXT,
  loser_team_id INT,
  loser_tickets INT
);

-- migrate:down

DROP TABLE maps;
