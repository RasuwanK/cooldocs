package routes

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: "CoolDocs API is running",
		Status:  "Ok",
	}

	json.NewEncoder(w).Encode(response)

}
