package repos

import (
	"database/sql"
	"path/filepath"
)

type Document struct {
	ProjectId    string `json:"projectId"`
	CollectionId string `json:"collectionId"`
	DocumentId   string `json:"documentId"`
	Data         string `json:"data"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type DocumentsRespository struct {
	db *sql.DB
}

func NewDocumentsRespository(db *sql.DB) *DocumentsRespository {
	return &DocumentsRespository{
		db: db,
	}
}

func (repo *DocumentsRespository) Get(
	projectId string,
	databaseId string,
	documentPath ,
) (Database, error) {

}

func (repo *DocumentsRespository) Create(
	projectId string,
	databaseId string,
	documentPath string,
) (Database, error) {

}
