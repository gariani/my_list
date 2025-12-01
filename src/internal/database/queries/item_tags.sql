-- name: AddTagToItem :exec
INSERT INTO item_tags (item_id, tag_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: RemoveTagFromItem :exec
DELETE FROM item_tags
WHERE item_id = $1 AND tag_id = $2;

-- name: ListItemTags :many
SELECT t.*
FROM item_tags it
JOIN tags t ON it.tag_id = t.id
WHERE it.item_id = $1;

-- name: ListItemsByTag :many
SELECT i.*
FROM item_tags it
JOIN items i ON it.item_id = i.id
WHERE it.tag_id = $1
ORDER BY i.created_at DESC;
