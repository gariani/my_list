-- Add these to your queries.sql file

-- name: GetItemWithDetails :one
SELECT
    i.*,
    it.name as type_name,
    it.description as type_description
FROM items i
LEFT JOIN item_types it ON i.type_id = it.id
WHERE i.id = $1;

-- name: GetAllItemsByListWithDetails :many
SELECT
    i.*,
    it.name as type_name,
    it.description as type_description
FROM items i
LEFT JOIN item_types it ON i.type_id = it.id
JOIN lists l ON i.list_id = l.id
WHERE l.user_id = $1 AND l.id = $2
ORDER BY i.created_at DESC;

-- name: GetItemsWithTypeByList :many
SELECT
    i.*,
    it.id as type_id,
    it.name as type_name,
    it.description as type_description
FROM items i
LEFT JOIN item_types it ON i.type_id = it.id
JOIN lists l ON i.list_id = l.id
WHERE l.user_id = $1 AND l.id = $2
ORDER BY i.created_at DESC;

-- name: GetTagsByItems :many
SELECT
    it.item_id,
    t.id as tag_id,
    t.name as tag_name
FROM item_tags it
JOIN tags t ON it.tag_id = t.id
WHERE it.item_id = ANY($1::uuid[]);