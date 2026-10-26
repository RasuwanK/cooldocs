-- +goose Up
CREATE TABLE folder (
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  name TEXT NOT NULL,
  created_At TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE folder; 
