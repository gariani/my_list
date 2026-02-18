-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS item_tags (
    item_id UUID REFERENCES items(id),
    tag_id UUID REFERENCES tags(id),
    PRIMARY KEY (item_id, tag_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS item_tags;
-- +goose StatementEnd
