-- name: CreateList :one
INSERT INTO lists (id, user_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetList :one
SELECT *
FROM lists
WHERE id = $1;

-- name: ListUserLists :many
SELECT *
FROM lists
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: DeleteList :exec
DELETE FROM lists
WHERE id = $1;
