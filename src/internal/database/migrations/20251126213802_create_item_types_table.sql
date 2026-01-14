-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS item_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS item_types;
-- +goose StatementEnd
