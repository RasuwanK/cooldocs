package routes

import (
	"cooldocs/internal/repos"
	"cooldocs/internal/services"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type DatabasesRoute struct {
	Service *services.DatabasesService
}

type CreateDatabasesRequest struct {
	Name string `json:"name"`
}

type DatabasesResponse struct {
	Database repos.Database `json:"project"`
}

// Initialize the repo and the services
func NewDatabasesRoute(db *sql.DB) *DatabasesRoute {
	databasesService := services.NewDatabasesService(db)

	return &DatabasesRoute{
		Service: databasesService,
	}
}

// GET
func (route *DatabasesRoute) Get(w http.ResponseWriter, r *http.Request) {
	databaseId := r.PathValue("databaseId")
	projectId := r.PathValue("projectId")

	database, err := route.Service.GetDatabase(
		r.Context(),
		databaseId,
		projectId,
	)

	if err != nil {
		// No projects found
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
		}

		log.Printf("Error (DatabasesService-Get-Error) %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	if err != nil {
		log.Fatal("Error while fetching the project")
	}

	project := DatabasesResponse{
		Database: *database,
	}

	json.NewEncoder(w).Encode(project)
}

// POST
func (route *DatabasesRoute) Post(w http.ResponseWriter, r *http.Request) {
	var request CreateDatabasesRequest
	projectId := r.PathValue("projectId")

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Printf("Error: %v", err)
		log.Fatal("Error while converting the database request from JSON to object")
	}

	database, err := route.Service.CreateDatabase(
		r.Context(),
		projectId,
		request.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "No database was created", 500)
		}

		log.Printf("Error (DatbaseService-POST-Error): %v", err)
		http.Error(w, "Internal server error", 500)
	}

	log.Printf("Database created (id: %s,name: %s)", database.Id, database.Name)

	response := DatabasesResponse{
		Database: *database,
	}

	json.NewEncoder(w).Encode(response)

}
