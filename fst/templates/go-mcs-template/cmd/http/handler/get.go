package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Obteniendo usuario ID: " + vars["id"]))
}
