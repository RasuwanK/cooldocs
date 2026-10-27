package projects

import (
	"context"
	"database/sql"
	"log"
)

// Modeling project data
type Project struct {
	Id   string
	Name string
}

// Modeling the project repository
type Repository struct {
	db *sql.DB
}

// Acts as the contructor for project repo
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Fetch the name of a single project
func (repo *Repository) Get(ctx context.Context, projectId string) (Project, error) {
	var name string

	err := repo.db.QueryRowContext(
		ctx,
		"SELECT name FROM projects WHERE id = $1",
		projectId,
	).Scan(&name)

	if err != nil {
		log.Printf("Error while executing query to select the project (Error: %v)", err)
		return Project{}, err
	}

	project := Project{
		Id:   projectId,
		Name: name,
	}

	return project, nil
}
