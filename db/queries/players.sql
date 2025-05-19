-- name: GetPlayer :one
SELECT * FROM players
WHERE player_id = $1 LIMIT 1;

-- name: InsertPlayer :one
INSERT INTO players (name, eos_id, steam_id, last_seen)
VALUES ($1, $2, $3, $4)
RETURNING player_id;

-- name: UpdatePlayerName :exec
UPDATE players
SET name = $1
WHERE player_id = $1;