package mongo

import (
	"context"
	"example/modular-clean/internal/example/entities"
)

func (s Repository) Create(ctx context.Context, param struct{ Name string }) entities.Example {

	return entities.Example{Name: param.Name}
}
