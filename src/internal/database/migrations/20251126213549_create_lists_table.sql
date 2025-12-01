-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS lists (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS lists;
-- +goose StatementEnd
