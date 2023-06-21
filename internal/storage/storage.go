package storage

type Storage struct {
	tt int
}

func NewStorage() *Storage {
	return &Storage{tt: 10}
}
