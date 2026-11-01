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

type ProjectRoute struct {
	Service *services.ProjectService
}

type CreateProjectRequest struct {
	Name string `json:"name"`
}

type ProjectResponse struct {
	Project repos.Project `json:"project"`
}

// Initialize the repo and the services
func NewProjectRoute(db *sql.DB) *ProjectRoute {
	projectService := services.NewProjectService(db)

	return &ProjectRoute{
		Service: projectService,
	}
}

// GET
func (route *ProjectRoute) Get(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")

	proj, err := route.Service.GetProject(
		r.Context(),
		projectId,
	)

	if err != nil {
		// No projects found
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
		}

		log.Printf("Error (ProjectService-Get-Error) %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	if err != nil {
		log.Fatal("Error while fetching the project")
	}

	project := ProjectResponse{
		Project: proj,
	}

	json.NewEncoder(w).Encode(project)
}

// POST
func (route *ProjectRoute) Post(w http.ResponseWriter, r *http.Request) {
	var request CreateProjectRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Printf("Error: %v", err)
		log.Fatal("Error while converting the project request from JSON to object")
	}

	proj, err := route.Service.CreateProject(
		r.Context(),
		request.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "No project was created", 500)
		}

		log.Printf("Error (ProjectService-POST-Error): %v", err)
		http.Error(w, "Internal server error", 500)
	}

	log.Printf("Project created (id: %s,name: %s)", proj.Id, proj.Name)

	response := ProjectResponse{
		Project: proj,
	}

	json.NewEncoder(w).Encode(response)

}
