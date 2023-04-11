package sqlstore

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"

	"github.com/Ropho/Pirater/config"
	"github.com/Ropho/Pirater/internal/store"
)

type SqlStore struct {
	Db       *sql.DB
	UserRepo *SqlUserRepository
	FilmRepo *SqlFilmRepository
}

func NewStore(conf *config.SqlDBConfig, logger *log.Logger) (*SqlStore, error) {

	sqlDb, err := newSqlDb(conf, logger)
	if err != nil {
		return nil, fmt.Errorf("unable init sql Db: [%w]", err)
	}

	return &SqlStore{
		Db: sqlDb,
	}, nil
}

func newSqlDb(conf *config.SqlDBConfig, logger *log.Logger) (*sql.DB, error) {

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Pass, conf.Addr, conf.Port, conf.Name)
	// db, err := sql.Open("mysql", "root:2280@tcp(127.0.0.1:3307)/test")
	db, err := sql.Open("mysql", url)
	if err != nil {
		logger.Errorf("sql db open error: [%w]", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Errorf("db connect error: [%w]", err)
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
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
