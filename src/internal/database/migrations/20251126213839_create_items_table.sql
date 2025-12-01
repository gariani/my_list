-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items (
    id UUID PRIMARY KEY,
    list_id UUID REFERENCES lists(id),
    user_id UUID REFERENCES users(id),
    type_id INT REFERENCES item_types(id),
    title TEXT,
    content TEXT,
    url TEXT,
    thumbnail TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
