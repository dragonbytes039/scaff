package main

import (
	handler "example/api-hexagonal/cmd/http/handlers"
	"example/api-hexagonal/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

//struct server
//funcion para inicializar el server
// funcion para cargar rutas

func main() {

	r := mux.NewRouter()

	s := services.NewServices("insert Db o repository")

	h := handler.NewHandler(s)

	r.HandleFunc("/get", h.GetUserHandler)

	port := ":8080"
	http.ListenAndServe(port, r)
}
