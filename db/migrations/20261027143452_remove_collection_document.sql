-- +goose Up
ALTER TABLE documents DROP COLUMN collection;

-- +goose Down
ALTER TABLE document ADD COLUMN collection TEXT NOT NULL;
