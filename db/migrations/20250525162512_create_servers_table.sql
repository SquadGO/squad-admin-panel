-- migrate:up

CREATE TABLE servers (
    server_id SERIAL PRIMARY KEY,
    "user" TEXT NOT NULL,
    host TEXT NOT NULL,
    password TEXT NOT NULL,
    rcon_password TEXT NOT NULL,
    rcon_port TEXT UNIQUE NOT NULL,
    admin_log_file_path TEXT NOT NULL,
    game_log_file_path TEXT UNIQUE NOT NULL
);

-- migrate:down

DROP TABLE servers;