-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN pass_hash TEXT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN IF EXISTS pass_hash;
-- +goose StatementEnd
