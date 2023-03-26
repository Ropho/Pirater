package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/Ropho/Cinema/config"

	"github.com/Ropho/Cinema/internal/store"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
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
}

func newDb(conf *config.DBaseConfig) (*sql.DB, error) {

	//"root:2280@/test"
	port := fmt.Sprintf("%d", conf.DbPort)
	url := conf.DbUser + ":" + conf.DbPass + "@tcp(" + conf.DbAddr + ":" + port + ")/" + conf.DbName
	// db, err := sql.Open("mysql", "root:2280@tcp(127.0.0.1:3307)/test")
	db, err := sql.Open("mysql", url)
	if err != nil {
		logrus.Error("sql db open error: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logrus.Error("db connect error: ", err)
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func (serv *Server) authenticateUser(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := serv.SessionStore.Get(r, sessionName)
		if err != nil {
			serv.error(w, r, http.StatusInternalServerError, "SESSION ERROR")
			logrus.Error("session get error: ", err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			serv.error(w, r, http.StatusNetworkAuthenticationRequired, "UNABLE TO AUTH")
			logrus.Error("GET USER_ID ERROR: ", err)
			return
		}

		// logrus.Info(id)

		u, err := serv.Store.User().FindById(id.(int))
		if err != nil {
			serv.error(w, r, http.StatusNetworkAuthenticationRequired, "UNABLE TO AUTH")
			logrus.Error("FIND USER BY ID ERROR: ", err)
			return
		}

		logrus.Info("AUTHENTICATE USER GOOD")

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *Server) setRequestId(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("Request ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestId, id)))
	})
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, data string) {
	s.respond(w, r, code, data)
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data string) {
	w.WriteHeader(code)
	w.Write([]byte(data))
}
