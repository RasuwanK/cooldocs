package services

import (
	"context"
	"cooldocs/internal/repos"
	"database/sql"
	"log"
)

type DatabasesService struct {
	repo *repos.DatabasesRepository
}

func NewDatabasesService(db *sql.DB) *DatabasesService {
	return &DatabasesService{
		repo: repos.NewDatbasesRepository(db),
	}
}

// Handling the document read
func (sr *DatabasesService) GetDatabase(ctx context.Context, projectId string, databaseId string) (*repos.Database, error) {
	database, err := sr.repo.Get(ctx, databaseId, projectId)

	if err != nil {
		log.Printf("Error occured while getting database, (Error: %v)", err)
		return nil, err
	}

	return database, nil
}

func (sr *DatabasesService) CreateDatabase(ctx context.Context, projectId string, databaseName string) (*repos.Database, error) {
	database, err := sr.repo.Create(
		ctx,
		projectId,
		databaseName,
	)

	if err != nil {
		log.Printf("Error occured while creating the database, (Error: %v)", err)
		return nil, err
	}

	return database, nil
}
