package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/Ropho/Pirater/config"
	"github.com/Ropho/Pirater/internal/store/mainstore"
)

func NewServer(conf *config.Config, defaultlogger *log.Logger) (*Server, error) {

	logger, err := NewLogger(&conf.Log)
	if err != nil {
		return nil, fmt.Errorf("init logger from config error, using default")
	}

	store, err := mainstore.NewStore(conf, logger)
	if err != nil {
		return nil, fmt.Errorf("unbale init store: [%w]", err)
	}

	serv := &Server{
		IP_Port: fmt.Sprint(conf.Server.Addr, ":", strconv.Itoa(conf.Server.Port)),
		Router:  mux.NewRouter(),
		Store:   store,
		// Store:   sqlstore.NewStore(db),
		// SessionStore: newCookieStore([]byte(conf.Env.CookieKey)),
		Config: conf,
		Logger: logger,
	}

	return serv, nil
}

func (serv *Server) Start() error {

	serv.initHandlers()

	serv.Logger.Info("server Starting: ", serv.IP_Port)

	err := http.ListenAndServe(serv.IP_Port, serv.Router)
	if err != nil {
		return fmt.Errorf("server serve error: [%w]", err)
	}

	return nil
}

func (serv *Server) initHandlers() {

	serv.Router.Use(serv.setRequestId)
	serv.Router.Use(serv.logRequest)

	swaggerUrl := "doc.json"
	serv.Router.PathPrefix("/swagger").HandlerFunc(httpSwagger.Handler(
		httpSwagger.URL(swaggerUrl), //The url pointing to API definition
	))

	api := serv.Router.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/", serv.handleBase).Methods("GET")
	api.HandleFunc("/carousel", serv.handleGetCarousel()).Methods("GET")
	api.HandleFunc("/newFilms", serv.handleGetNewFilms()).Methods("GET")
	api.HandleFunc("/film/{hash}", serv.HandleGetCurrentFilm()).Methods("GET")

	api.HandleFunc("/users", serv.handleUsersCreate()).Methods("POST")
	api.HandleFunc("/sessions", serv.handleSessionsCreate()).Methods("POST")

	private := api.PathPrefix("/private").Subrouter()
	private.Use(serv.authenticateUser)
	private.HandleFunc("/whoami", serv.handleWhoami()).Methods("GET")

	admin := private.PathPrefix("/admin").Subrouter()
	admin.Use(serv.authorizeAdmin)
	admin.HandleFunc("/add/films", serv.handleAddFilms()).Methods("POST")
}
