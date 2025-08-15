package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Message string `json:"message"`
}

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := Message{
			Message: "OK",
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}).Methods("GET")

	return router
}
