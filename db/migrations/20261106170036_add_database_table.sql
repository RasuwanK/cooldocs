-- +goose Up
SELECT 'up SQL query';

CREATE TABLE databases (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  project_id UUID NOT NULL,
  name TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT now(),

  UNIQUE(id, project_id)
);

ALTER TABLE collections ADD COLUMN database_id UUID NOT NULL;

ALTER TABLE collections ADD CONSTRAINT fk_databases_collections FOREIGN KEY(database_id) REFERENCES databases(id);
ALTER TABLE databases ADD CONSTRAINT fk_databases_projects FOREIGN KEY(project_id) REFERENCES projects(id);

-- +goose Down
SELECT 'down SQL query';

ALTER TABLE collections DROP CONSTRAINT fk_databases_collections; 
ALTER TABLE databases DROP CONSTRAINT fk_databases_projects; 

ALTER TABLE collections DROP COLUMN database_id;

DROP TABLE databases;
