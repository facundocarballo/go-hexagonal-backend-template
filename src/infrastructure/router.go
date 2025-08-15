package infrastructure

import (
	controllers "github.com/facundocarballo/go-hexagonal-backend-template/src/infrastructure/controllers/health"
	"github.com/gorilla/mux"
)

type Message struct {
	Message string `json:"message"`
}

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", controllers.HealthController).Methods("GET")

	return router
}
