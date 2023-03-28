package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/Ropho/Cinema/config"

	"github.com/Ropho/Cinema/internal/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
)

type ctxKey int8

const (
	videoDir = "./video"

	ctxKeyUser ctxKey = iota
	ctxKeyRequestId
)

type Server struct {
	IP_Port      string
	Router       *mux.Router
	Store        store.Store
	Config       *config.Config
	SessionStore sessions.Store
	SwaggerUrl   string
	Logger       *log.Logger
}

func newDb(conf *config.DBaseConfig, logger *log.Logger) (*sql.DB, error) {

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.DbUser, conf.DbPass, conf.DbAddr, conf.DbPort, conf.DbName)
	// db, err := sql.Open("mysql", "root:2280@tcp(127.0.0.1:3307)/test")
	db, err := sql.Open("mysql", url)
	if err != nil {
		logger.Error("sql db open error: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Error("db connect error: ", err)
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, data string) {
	s.respond(w, r, code, data)
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data string) {
	w.WriteHeader(code)
	w.Write([]byte(data))
}
