package handler

import (
	"example/modular-clean/internal/example/ports"
)

type Handler struct {
	Service ports.IExample
}

func NewExampleHandler(sv ports.IExample) ports.IExampleHandler {
	return Handler{Service: sv}
}
