package services

type Services struct {
	db string
}

func NewServices(db string) *Services {
	return &Services{db: db}
}
