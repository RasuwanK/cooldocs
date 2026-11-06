package repos

import (
	"context"
	"database/sql"
	"log"
)

// Modeling project data
type Database struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ProjectId string `json:"projectId"`
}

// Modeling the project repository
type DatabasesRepository struct {
	db *sql.DB
}

// Acts as the contructor for project repo
func NewDatbasesRepository(db *sql.DB) *DatabasesRepository {
	return &DatabasesRepository{
		db: db,
	}
}

// Fetch the name of a single project
func (repo *DatabasesRepository) Get(ctx context.Context, projectId string, databaseId string) (*Database, error) {
	var name string

	err := repo.db.QueryRowContext(
		ctx,
		"SELECT name FROM databases WHERE id = $1",
		databaseId,
	).Scan(&name)

	if err != nil {
		log.Printf("Error while executing query to select the database (Error: %v)", err)
		return nil, err
	}

	database := Database{
		Id:        databaseId,
		Name:      name,
		ProjectId: projectId,
	}

	return &database, nil
}

func (repo *DatabasesRepository) Create(ctx context.Context, projectId string, name string) (*Database, error) {
	var database Database

	err := repo.db.QueryRowContext(
		ctx,
		"INSERT INTO databases(name, project_id) VALUES ($1, $2) RETURNING id, name, project_id",
		name,
		projectId,
	).Scan(&database.Id, &database.Name, &database.ProjectId)

	if err != nil {
		log.Printf("Error while executing query to insert a database (Error: %v)", err)
		return nil, err
	}

	return &database, nil
}
