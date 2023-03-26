package store

import (
	film "github.com/Ropho/Cinema/internal/model/film"
	user "github.com/Ropho/Cinema/internal/model/user"
)

type UserRepository interface {
	Create(*user.User) error
	FindByEmail(string) (*user.User, error)
	FindById(id int) (*user.User, error)
}

type FilmRepository interface {
	// Create(*film.Film) error
	FindById(id int) (*film.Film, error)
	FindByName(name string) (*film.Film, error)
	CountAllRows() (int, error)
	GetRandomFilms(num int) ([]film.Film, error)
	GetNewFilms(num int) ([]film.Film, error)
}
