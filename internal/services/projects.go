package services

import (
	"context"
	"cooldocs/internal/repos"
	"database/sql"
	"log"
)

type ProjectService struct {
	repo *repos.ProjectRepository
}

func NewProjectService(db *sql.DB) *ProjectService {
	return &ProjectService{
		repo: repos.NewProjectRepository(db),
	}
}

// Handling the document read
func (sr *ProjectService) GetProject(ctx context.Context, projectId string) (*repos.Project, error) {
	project, err := sr.repo.Get(ctx, projectId)

	if err != nil {
		log.Printf("Error occured while getting project, (Error: %v)", err)
		return nil, err
	}

	return project, nil
}

func (sr *ProjectService) CreateProject(ctx context.Context, projectName string) (*repos.Project, error) {
	project, err := sr.repo.Create(
		ctx,
		projectName,
	)

	if err != nil {
		log.Printf("Error occured while creating the project, (Error: %v)", err)
		return nil, err
	}

	return project, nil
}
