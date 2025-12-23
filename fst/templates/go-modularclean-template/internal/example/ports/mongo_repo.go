package ports

import (
	"context"
	"example/modular-clean/internal/example/entities"
)

type IExampleMongo interface {
	Create(ctx context.Context, param struct{ Name string }) entities.Example
}
