package cmd

import (
	"context"
	"cooldocs/internal/database"
	"cooldocs/internal/httpapi"
	"cooldocs/internal/routes"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type CreateProjectRequest struct {
	Name string `json:"name"`
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
	dbCtx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	db, err := database.Open(dbCtx, databaseUrl)

	cancel()

	if err != nil {
		log.Printf("%v", err)
		log.Fatal("Error while connecting to the DB")
	}

	// routes
	projectRoute := routes.NewProjectRoute(db)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", routes.Health)

	// Document read Handler
	// mux.HandleFunc("GET /v1/projects/{projectId}/{documentPath...}")

	server := &http.Server{
		Addr:    ":8080",
		Handler: httpapi.CORS(mux),
	}

	mux.HandleFunc("GET /v1/projects/{projectId}", projectRoute.Get)

	mux.HandleFunc("POST /v1/projects", projectRoute.Post)

	log.Println("CoolDocs server is listening on 8080")

	log.Fatal(server.ListenAndServe())
}
