-- name: GetServer :one
SELECT * FROM servers
WHERE server_id = $1
LIMIT 1;

-- name: GetServers :many
SELECT * FROM servers;
