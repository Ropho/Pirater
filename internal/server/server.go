package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Ropho/Cinema/internal/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

type ctxKey int8

const ctxKeyUser ctxKey = iota

type Server struct {
	IP_Port      string
	Router       *mux.Router
	Store        *store.Store
	SessionStore sessions.Store
	SwaggerUrl   string
}

func NewServer() *Server {

	conf := NewConfig()

	sStore := sessions.NewCookieStore([]byte(conf.CookieKey))
	sStore.MaxAge(1000)

	serv := &Server{
		IP_Port:      conf.ServAddr + ":" + strconv.Itoa(conf.Port),
		Router:       mux.NewRouter(),
		Store:        store.NewStore(),
		SessionStore: sStore,
	}
	return serv
}

func (serv *Server) Start() error {

	serv.Router.PathPrefix("/swagger").HandlerFunc(httpSwagger.Handler(
		httpSwagger.URL(serv.SwaggerUrl), //The url pointing to API definition
	)).Methods("GET")

	serv.Router.HandleFunc("/", serv.handleBase)
	serv.Router.HandleFunc("/users", serv.handleUsersCreate).Methods("POST")
	serv.Router.HandleFunc("/sessions", serv.handleSessionsCreate).Methods("POST")

	private := serv.Router.PathPrefix("/private").Subrouter()
	private.Use(serv.authenticateUser)
	private.HandleFunc("/whoami", serv.handleWhoami()).Methods("GET")

	logrus.Info("SERVER STARTING\n")
	err := http.ListenAndServe(serv.IP_Port, serv.Router)
	if err != nil {
		logrus.Fatal("SERVER PROCESS ERROR: ", err)
	}
	logrus.Info("SERVER CLOSED UNECTEDLY\n")

	return err
}

func (serv *Server) Close() {
	serv.Store.Db.Close()
	logrus.Info("SERVER CLOSE...")
}

func (serv *Server) authenticateUser(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := serv.SessionStore.Get(r, sessionName)
		if err != nil {
			serv.error(w, r, http.StatusInternalServerError, "SESSION ERROR")
			logrus.Error("SESSION GET ERROR: ", err)
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

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, data string) {

	s.respond(w, r, code, responseErr(data))
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data string) {

	w.WriteHeader(code)

	w.Write([]byte(responseInfo(data)))
}

func responseErr(s string) string {
	return "\033[31m" + s + "\n\033[0m"
}

func responseInfo(s string) string {
	return "\033[34m" + s + "\n\033[0m"
}
