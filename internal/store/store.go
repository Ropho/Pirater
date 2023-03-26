package store

type Store interface {
	User() UserRepository
	Film() FilmRepository
}
