-- name: CreateTag :one
INSERT INTO tags (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: ListTags :many
SELECT *
FROM tags
ORDER BY name;

-- name: GetTag :one
SELECT *
FROM tags
WHERE id = $1;

-- name: DeleteTag :exec
DELETE FROM tags
WHERE id = $1;
