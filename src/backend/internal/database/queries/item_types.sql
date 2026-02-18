-- name: CreateItemType :one
INSERT INTO item_types (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: ListItemTypes :many
SELECT *
FROM item_types
ORDER BY id;

-- name: GetItemType :one
SELECT *
FROM item_types
WHERE id = $1;

-- name: DeleteItemType :exec
DELETE FROM item_types
WHERE id = $1;
