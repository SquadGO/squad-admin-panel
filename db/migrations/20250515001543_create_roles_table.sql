-- migrate:up

CREATE TABLE roles (
  role_id SERIAL PRIMARY KEY,
  name VARCHAR UNIQUE NOT NULL,
  ban BOOLEAN NOT NULL,
  kick BOOLEAN NOT NULL,
  warn BOOLEAN NOT NULL,
  disband BOOLEAN NOT NULL,
  remove_player_squad BOOLEAN NOT NULL,
  change_map BOOLEAN NOT NULL,
  change_next_map BOOLEAN NOT NULL,
  end_match BOOLEAN NOT NULL,
  broadcast BOOLEAN NOT NULL,
  force_player BOOLEAN NOT NULL,
  flags BOOLEAN NOT NULL,
  change_role BOOLEAN NOT NULL
);

INSERT INTO roles (name, ban, kick, warn, disband, remove_player_squad, change_map, change_next_map, end_match, broadcast, force_player, flags, change_role)
VALUES ('Гость', false, false, false, false, false, false, false, false, false, false, false, false);

INSERT INTO roles (name, ban, kick, warn, disband, remove_player_squad, change_map, change_next_map, end_match, broadcast, force_player, flags, change_role)
VALUES ('Админ', true, true, true, true, true, true, true, true, true, true, true, true);

-- migrate:down

DROP TABLE roles;