package mainstore

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Ropho/Pirater/config"
	"github.com/Ropho/Pirater/internal/store"
	cookie "github.com/Ropho/Pirater/internal/store/cookiestore"
	sql "github.com/Ropho/Pirater/internal/store/sqlstore"
)

type MainStore struct {
	SqlStore    *sql.SqlStore
	CookieStore *cookie.CookieStore
}

func NewStore(conf *config.Config, logger *log.Logger) (store.Store, error) {

	cookieStore, err := cookie.NewCookieStore(conf, logger)
	if err != nil {
		return nil, fmt.Errorf("unable init Store: [%w]", err)
	}

	sqlStore, err := sql.NewStore(&conf.SqlDBase, logger)
	if err != nil {
		return nil, fmt.Errorf("unable init Store: [%w]", err)
	}

	return &MainStore{
		SqlStore:    sqlStore,
		CookieStore: cookieStore,
	}, nil
}

func (ms *MainStore) User() store.UserRepository {
	return ms.SqlStore.User()
}
func (ms *MainStore) Film() store.FilmRepository {
	return ms.SqlStore.Film()
}
func (ms *MainStore) Cookie() store.CookieRepository {
	return ms.CookieStore.Cookie()
}
