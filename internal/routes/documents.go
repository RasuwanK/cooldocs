package routes

import (
	"cooldocs/internal/repos"
	"cooldocs/internal/services"
	"database/sql"
	"net/http"
)

type DocumentsRoute struct {
	Service *services.DocumentsService
}

type CreateDocumentRequest struct {
	Name string `json:"name"`
}

type DocumentsResponse struct {
	Database repos.Database `json:"project"`
}

// Initialize the repo and the services
func NewDocumentsRoute(db *sql.DB) *DatabasesRoute {
	databasesService := services.NewDatabasesService(db)

	return &DatabasesRoute{
		Service: databasesService,
	}
}

// GET
func (route *DocumentsRoute) Get(w http.ResponseWriter, r *http.Request) {

}

// POST
func (route *DocumentsRoute) Post(w http.ResponseWriter, r *http.Request) {
}
