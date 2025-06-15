-- name: InsertLog :one
INSERT INTO logs (log_type, server_id, user_id, player_id, victim_id, attacker_id, player_ip, squad_id, squad_name, team_id, chat_type, is_teamkill, message, map)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
RETURNING log_id;