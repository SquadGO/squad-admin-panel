-- name: GetPlayer :one
SELECT * FROM players
WHERE player_id = $1
LIMIT 1;

-- name: GetPlayerBySteamID :one
SELECT * FROM players
WHERE steam_id = $1
LIMIT 1;

-- name: GetPlayerByEosID :one
SELECT * FROM players
WHERE eos_id = $1
LIMIT 1;

-- name: InsertPlayer :one
INSERT INTO players (name, eos_id, steam_id)
VALUES ($1, $2, $3)
RETURNING player_id;

-- name: UpdatePlayerName :exec
UPDATE players
SET name = $1
WHERE player_id = $1;