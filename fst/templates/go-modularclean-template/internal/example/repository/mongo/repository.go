package mongo

import "example/modular-clean/internal/example/ports"

type Repository struct {
	Db string
}

func NewExampleRepo(db string) ports.IExampleMongo {
	return Repository{Db: db}
}
