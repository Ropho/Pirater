package cookiestore

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/boj/redistore"

	"github.com/Ropho/Pirater/config"
	"github.com/Ropho/Pirater/internal/store"
)

type CookieStore struct {
	RediStore  *redistore.RediStore
	CookieRepo *CookieRepository
}

func NewCookieStore(conf *config.Config, logger *log.Logger) (*CookieStore, error) {

	redistore, err := newRediStore(conf, logger)
	if err != nil {
		return nil, fmt.Errorf("unable init cookie store: [%w]", err)
	}

	return &CookieStore{
		RediStore: redistore,
	}, nil
}

func newRediStore(conf *config.Config, logger *log.Logger) (*redistore.RediStore, error) {
	// Fetch new store.
	store, err := redistore.NewRediStore(conf.CookieDBase.MaxConn, "tcp", fmt.Sprintf("%s:%d", conf.CookieDBase.Addr, conf.CookieDBase.Port), conf.CookieDBase.Pass, []byte(conf.Env.CookieKey))
	if err != nil {
		return nil, err
	}
	// defer store.Close()

	// Change session storage configuration for MaxAge = 10 days.
	store.SetMaxAge(180)

	return store, nil
}

func (cs *CookieStore) Cookie() store.CookieRepository {
	if cs.CookieRepo != nil {
		return cs.CookieRepo
	}

	cs.CookieRepo = &CookieRepository{
		store: cs,
	}

	return cs.CookieRepo
}
