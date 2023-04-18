package teststore

// import (
// 	"database/sql"

// 	"github.com/Ropho/Pirater/internal/store"
// )

// type Store struct {
// 	Db       *sql.DB
// 	UserRepo *UserRepo
// }

// func NewStore(db *sql.DB) *Store {

// 	return &Store{
// 		Db: db,
// 	}
// }

// func (st *Store) User() store.UserRepository {
// 	if st.UserRepo != nil {
// 		return st.UserRepo
// 	}

// 	st.UserRepo = &UserRepo{
// 		store: st,
// 	}

// 	return st.UserRepo
// }
