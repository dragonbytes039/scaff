package routes

import (
	"example/modular-clean/internal/example"

	"github.com/gorilla/mux"
)

func RegisterExampleRoutes(r *mux.Router) {

	handler := example.Assembly()

	exampleRouter := r.PathPrefix("/example").Subrouter()

	exampleRouter.HandleFunc("/", handler.CreateExample)

}
