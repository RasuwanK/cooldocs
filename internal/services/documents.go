package services

import (
	"cooldocs/internal/repos"
	"database/sql"
	"path/filepath"
)

type DocumentsService struct {
	repo *repos.DocumentsRespository
}

func NewDocumentService(db *sql.DB) *DocumentsService {
	return &DocumentsService{
		repo: repos.NewDocumentsRespository(db),
	}
}

func (sr *DocumentsService) Get(
	projectId string,
	databaseId string,
	documentPath string,
) {
	p := filepath.Dir(documentPath)
}
