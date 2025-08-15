package main

import (
	"net/http"

	"github.com/facundocarballo/go-hexagonal-backend-template/scr/infrastructure"
)

func main() {
	router := infrastructure.Router()

	http.ListenAndServe(":8080", router)
}
