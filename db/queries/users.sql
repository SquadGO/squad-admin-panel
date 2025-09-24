-- name: GetUserBySteamID :one
SELECT * FROM users
WHERE steam_id = $1
LIMIT 1;

-- name: InsertUser :one
INSERT INTO users (steam_id, name, avatar)
VALUES ($1, $2, $3)
RETURNING user_id;

-- name: UpdateUser :exec
UPDATE users
SET
  steam_id = $2,
  name = $3,
  avatar = $4
WHERE user_id = $1;
