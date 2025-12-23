package services

import (
	"context"
	"example/api-hexagonal/internal/models"
	"fmt"
)

func (s *Services) CreateExample(ctx context.Context, usuario *models.Example) error {

	fmt.Printf(s.db)

	return nil
}
