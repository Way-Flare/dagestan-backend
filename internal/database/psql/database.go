package psql

type Authorization interface {
}

type Place interface {
}

type Repository struct {
	Authorization
	Place
}

func NewRepository() *Repository {
	return &Repository{}
}
