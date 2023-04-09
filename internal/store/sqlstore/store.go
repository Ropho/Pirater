package sqlstore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Ropho/Pirater/internal/store"
)

type SqlStore struct {
	Db       *sql.DB
	UserRepo *SqlUserRepository
	FilmRepo *SqlFilmRepository
}

func NewStore(db *sql.DB) *SqlStore {
	return &SqlStore{
		Db: db,
	}
}

func (st *SqlStore) User() store.UserRepository {
	if st.UserRepo != nil {
		return st.UserRepo
	}

	st.UserRepo = &SqlUserRepository{
		store: st,
	}
	return st.UserRepo
}

func (st *SqlStore) Film() store.FilmRepository {
	if st.FilmRepo != nil {
		return st.FilmRepo
	}

	st.FilmRepo = &SqlFilmRepository{
		store: st,
	}
	return st.FilmRepo
}
