package controllers

import (
	"encoding/json"
	"net/http"

	usecases "github.com/facundocarballo/go-hexagonal-backend-template/src/domain/usecases/health"
	repositories "github.com/facundocarballo/go-hexagonal-backend-template/src/infrastructure/repositories/counter"
)

func HealthController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	repository := repositories.CounterMemory{}
	output := usecases.HealthUseCase(repository)

	response := HealthResponse{
		Message: "Ok",
		Author:  "Facundo Carballo",
		Counter: output.Value,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
