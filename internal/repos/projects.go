package repos

import (
	"context"
	"database/sql"
	"log"
)

// Modeling project data
type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Modeling the project repository
type ProjectRepository struct {
	db *sql.DB
}

// Acts as the contructor for project repo
func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

// Fetch the name of a single project
func (repo *ProjectRepository) Get(ctx context.Context, projectId string) (Project, error) {
	var name string

	err := repo.db.QueryRowContext(
		ctx,
		"SELECT name FROM projects WHERE id = $1",
		projectId,
	).Scan(&name)

	if err != nil {
		log.Printf("Error while executing query to select the project (Error: %v)", err)
		return Project{}, nil
	}

	project := Project{
		Id:   projectId,
		Name: name,
	}

	return project, nil
}

func (repo *ProjectRepository) Create(ctx context.Context, name string) (Project, error) {
	var project Project

	err := repo.db.QueryRowContext(
		ctx,
		"INSERT INTO projects(name) VALUES ($1) RETURNING id, name",
		name,
	).Scan(&project.Name, &project.Id)

	if err != nil {
		log.Printf("Error while executing query to insert a project (Error: %v)", err)
		return Project{}, err
	}

	return project, nil
}
