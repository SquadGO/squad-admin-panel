-- migrate:up

CREATE TYPE chat_type AS ENUM (
  'ChatAll',
  'ChatTeam',
  'ChatSquad',
  'ChatAdmin'
);

CREATE TABLE logs (
    log_id BIGSERIAL PRIMARY KEY,
    log_type VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    server_id INT REFERENCES servers(server_id) ON DELETE NO ACTION NULL,
    user_id INT REFERENCES users(user_id) ON DELETE NO ACTION NULL,
    player_id INT REFERENCES players(player_id) ON DELETE NO ACTION NULL,
    victim_id INT REFERENCES players(player_id) ON DELETE NO ACTION NULL,
    attacker_id INT REFERENCES players(player_id) ON DELETE NO ACTION NULL,
    player_ip VARCHAR NULL,
    squad_id VARCHAR NULL,
    squad_name TEXT NULL,
    team_id VARCHAR NULL,
    chat_type chat_type NULL,
    is_teamkill BOOLEAN NULL,
    message TEXT NULL,
    weapon TEXT NULL,
    map TEXT NULL
);


-- migrate:down

DROP TABLE logs
