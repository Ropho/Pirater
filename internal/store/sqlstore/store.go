package sqlstore

import (
	"database/sql"

	"github.com/Ropho/Cinema/internal/store"
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

// func CloseStore(db *sql.DB) error {
// 	if err := db.Close(); err != nil {
// 		logrus.Error("unable to close sql db")
// 		return err
// 	}
// 	return nil
// }
