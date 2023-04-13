package postgrsql

type Authorization interface {
}

type Lineup interface {
}

type Repository struct {
	Authorization
	Lineup
}

func NewRepository() *Repository {
	return &Repository{}
}
