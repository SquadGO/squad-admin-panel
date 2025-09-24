-- name: GetMaps :many
SELECT * FROM maps;

-- name: InsertMap :one
INSERT INTO maps (server_id, created_at, end_at, map_name, winner_name, winner_team_id, winner_tickets, loser_name, loser_team_id, loser_tickets)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING map_id;

-- name: UpdateMap :exec
UPDATE maps
SET
  end_at = $2,
  winner_name = $3,
  winner_team_id = $4,
  winner_tickets = $5,
  loser_name = $6,
  loser_team_id = $7,
  loser_tickets = $8
WHERE map_id = $1;
