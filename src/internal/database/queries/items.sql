-- name: CreateItem :one
INSERT INTO items (
    list_id, type_id,title, content, url, thumbnail
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetItem :one
SELECT *
FROM items
WHERE id = $1;

-- name: GetAllItemsByList :many
SELECT i.*
FROM items i
JOIN lists l on i.list_id = l.id
WHERE l.user_id = $1
  AND l.id = $2;

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
DELETE FROM items i
WHERE id = $1;

