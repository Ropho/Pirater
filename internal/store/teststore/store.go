package teststore

import (
	"database/sql"

	"github.com/Ropho/Cinema/internal/store"
)

type Store struct {
	Db *sql.DB
	// Conf     Config
	UserRepo *UserRepo
}

func NewStore(db *sql.DB) *Store {

	return &Store{
		Db: db,
	}

	// var err error

	// st.Conf = NewConfig()

	// //"root:2280@/test"
	// // st.Db, err = sql.Open("mysql", st.Conf.User+":"+st.Conf.Pass+"@/"+st.Conf.DbName)
	// st.Db, err = sql.Open("mysql", "root:2280@tcp(127.0.0.1:3307)/test")

	// if err != nil {
	// 	logrus.Fatal("SQL OPEN ERROR: ", err)
	// }

	// err = st.Db.Ping()
	// if err != nil {
	// 	logrus.Fatal("CANT CONNECT TO DB: ", err)
	// }

	// st.Db.SetConnMaxLifetime(time.Minute * 3)
	// st.Db.SetMaxOpenConns(10)
	// st.Db.SetMaxIdleConns(10)

	// return &st
}

// Store.User().Create()
func (st *Store) User() store.UserRepository {
	if st.UserRepo != nil {
		return st.UserRepo
	}

	st.UserRepo = &UserRepo{
		store: st,
	}

	return st.UserRepo
}
