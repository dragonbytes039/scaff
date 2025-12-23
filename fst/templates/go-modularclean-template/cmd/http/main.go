package main

import (
	"example/modular-clean/cmd/http/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// assembly example
	routes.RegisterExampleRoutes(r)

	port := ":8080"
	http.ListenAndServe(port, r)
}
