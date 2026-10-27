-- +goose Up
CREATE TABLE projects (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE documents (
    project_id UUID NOT NULL REFERENCES projects(id),
    collection TEXT NOT NULL,
    document_id TEXT NOT NULL,
    data JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (project_id, collection, document_id)
);

CREATE INDEX documents_data_gin_idx
    ON documents USING GIN (data);


-- +goose Down
DROP TABLE projects;
DROP TABLE documents;
DROP INDEX documents.documents_data_gin_idx;
