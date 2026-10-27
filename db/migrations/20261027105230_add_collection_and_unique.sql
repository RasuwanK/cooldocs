-- +goose Up
ALTER TABLE documents DROP CONSTRAINT documents_pkey;
ALTER TABLE documents ADD COLUMN id UUID PRIMARY KEY; 
ALTER TABLE documents ADD COLUMN collection_id UUID NOT NULL; 
ALTER TABLE documents ADD CONSTRAINT unique_collection_document_id UNIQUE (collection_id, document_id);
ALTER TABLE documents ADD CONSTRAINT unique_document_project_id UNIQUE (id, project_id);
ALTER TABLE documents ADD FOREIGN KEY(collection_id, project_id) REFERENCES collections(id, project_id);

-- Add this after both tables exist.
ALTER TABLE collections
ADD CONSTRAINT collections_parent_document_fk
FOREIGN KEY (parent_document_id, project_id)
REFERENCES documents (id, project_id);

-- Root collections must be unique within a project.
CREATE UNIQUE INDEX collections_root_name_idx
ON collections (project_id, name)
WHERE parent_document_id IS NULL;

-- Subcollection names must be unique under a parent document.
CREATE UNIQUE INDEX collections_child_name_idx
ON collections (parent_document_id, name)
WHERE parent_document_id IS NOT NULL;

CREATE INDEX documents_project_collection_idx
ON documents (project_id, collection_id);

-- +goose Down
DROP TABLE documents;
DROP TABLE collections;
