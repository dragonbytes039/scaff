package ports

import (
	"context"
)

type IExample interface {
	Create(ctx context.Context, param struct{ Name string }) error
}
