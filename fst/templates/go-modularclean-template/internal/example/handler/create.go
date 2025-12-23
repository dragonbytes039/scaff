package handler

import (
	"net/http"
)

func (h Handler) CreateExample(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := h.Service.Create(ctx, struct{ Name string }{"body"})

	if err != nil {
		w.Write([]byte{})
		return
	}

	w.Write([]byte{})
}
