package projects

import (
	"context"
	"log"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Handling the document read
func (sr *Service) GetProject(ctx context.Context, projectId string) (Project, error) {
	project, err := sr.repo.Get(ctx, projectId)

	if err != nil {
		log.Printf("Error occured while getting project, (Error: %v)", err)
		return project, err
	}

	return project, nil
}
