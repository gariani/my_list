-- +goose Up
-- +goose StatementBegin
ALTER TABLE items ADD COLUMN category TEXT;
ALTER TABLE items ADD COLUMN tags TEXT[];
ALTER TABLE items ADD COLUMN summary TEXT;
ALTER TABLE items ADD COLUMN embedding FLOAT8[];

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE items DROP COLUMN IS EXISTS category;
ALTER TABLE items DROP COLUMN IS EXISTS tags;
ALTER TABLE items DROP COLUMN IS EXISTS summary;
ALTER TABLE items DROP COLUMN IS EXISTS embedding;
-- +goose StatementEnd
