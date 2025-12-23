package services

import "example/modular-clean/internal/example/ports"

type Example struct {
	Mongo ports.IExampleMongo
}

func NewExampleService(mongo ports.IExampleMongo) ports.IExample {
	return Example{Mongo: mongo}
}
