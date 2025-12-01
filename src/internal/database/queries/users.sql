-- name: CreateUser :one
INSERT INTO users (id, email, pass_hash)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY created_at DESC;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
