-- name: AddMetadata :one
INSERT INTO item_metadata (id, item_id, key, value)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetItemMetadata :many
SELECT *
FROM item_metadata
WHERE item_id = $1;

-- name: DeleteMetadataByItem :exec
DELETE FROM item_metadata
WHERE item_id = $1;
