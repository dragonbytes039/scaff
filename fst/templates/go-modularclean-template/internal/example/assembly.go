package example

import (
	"example/modular-clean/internal/example/handler"
	"example/modular-clean/internal/example/ports"
	"example/modular-clean/internal/example/repository/mongo"
	"example/modular-clean/internal/example/services"
)

func Assembly() ports.IExampleHandler {
	repo := mongo.NewExampleRepo("Insert DB")
	service := services.NewExampleService(repo)

	Handler := handler.NewExampleHandler(service)

	return Handler
}
