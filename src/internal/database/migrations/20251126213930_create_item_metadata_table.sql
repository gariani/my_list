-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS item_metadata (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    item_id UUID REFERENCES items(id),
    key TEXT NOT NULL,
    value TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS item_metadata;
-- +goose StatementEnd
