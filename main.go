package main

import (
	"context"
	"cooldocs/internal/database"
	"cooldocs/internal/httpapi"
	"cooldocs/internal/projects"
	"encoding/json"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Checking the health of the server
func healthHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: "CoolDocs API is running",
		Status:  "Ok",
	}
	json.NewEncoder(w).Encode(response)

}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found; using system environment")
	}

	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		log.Fatal("No DB URL has configured")
	}

	// To enable concurrency
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	db, err := database.Open(ctx, databaseUrl)

	cancel()

	if err != nil {
		log.Printf("%v", err)
		log.Fatal("Error while connecting to the DB")
	}

	// projects
	projectRepository := projects.NewRepository(db)
	projectService := projects.NewService(projectRepository)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandle)

	// Document read Handler
	mux.HandleFunc("GET /v1/projects/{projectId}/{documentPath...}",
		func(w http.ResponseWriter, r *http.Request) {
			proj, err := projectService.GetProject(r.Context(), r.PathValue("projectId"))

			if err != nil {
				log.Fatal("Error while taking the project")
			}

			log.Printf("%s", proj.Name)

			response := Response{
				Message: "Done!",
				Status:  "200",
			}

			json.NewEncoder(w).Encode(response)
		})

	server := &http.Server{
		Addr:    ":8080",
		Handler: httpapi.CORS(mux),
	}

	log.Println("CoolDocs server is listening on 8080")

	log.Fatal(server.ListenAndServe())
}
