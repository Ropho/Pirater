package store

import (
	"net/http"

	"github.com/gorilla/sessions"

	film "github.com/Ropho/Pirater/internal/model/film"
	user "github.com/Ropho/Pirater/internal/model/user"
)

type UserRepository interface {
	Create(*user.User) error
	FindByEmail(string) (*user.User, error)
	FindById(id int) (*user.User, error)
}

type FilmRepository interface {
	Create(films []film.Film) error
	FindByHash(hash uint32) (*film.Film, error)
	CountAllRows() (int, error)
	GetCarouselFilmsInfo(num int) ([]film.Film, error)
	GetNewFilmsInfo(num int) ([]film.Film, error)
}

type CookieRepository interface {
	Get(req *http.Request, sessionName string) (*sessions.Session, error)
	Delete(rsp http.ResponseWriter, req *http.Request, sessionName string) error
}
