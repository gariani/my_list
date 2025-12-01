-- name: CreateItem :one
INSERT INTO items (
    id, list_id, user_id, type_id,
    title, content, url, thumbnail
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetItem :one
SELECT *
FROM items
WHERE id = $1;

-- name: ListItemsByList :many
SELECT *
FROM items
WHERE list_id = $1
ORDER BY created_at DESC;

-- name: UpdateItem :one
UPDATE items
SET
    title = COALESCE($2, title),
    content = COALESCE($3, content),
    url = COALESCE($4, url),
    thumbnail = COALESCE($5, thumbnail),
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1;
