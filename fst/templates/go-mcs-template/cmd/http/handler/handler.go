package handler

import "example/api-hexagonal/internal/services"

type Handler struct {
	services *services.Services
}

func NewHandler(s *services.Services) *Handler {
	return &Handler{services: s}
}
