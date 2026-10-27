-- +goose Up
CREATE TABLE collections (
    id UUID PRIMARY KEY,
    project_id UUID NOT NULL REFERENCES projects(id),
    name TEXT NOT NULL,
    parent_document_id UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (id, project_id)
);

-- +goose Down
DROP TABLE collections;
