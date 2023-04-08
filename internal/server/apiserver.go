package server

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"net/http"
	"strconv"

	"github.com/Ropho/Cinema/config"

	"github.com/Ropho/Cinema/internal/store/sqlstore"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewServer(conf *config.Config, defaultlogger *log.Logger) (*Server, error) {

	logger, err := NewLogger(&conf.Log)
	if err != nil {
		logger = defaultlogger
		logger.Error("init logger from config error, using default")
	}

	db, err := newDb(&conf.DBase, logger)
	if err != nil {
		return nil, fmt.Errorf("init db error: [%w]", err)
	}

	serv := &Server{
		IP_Port:      fmt.Sprint(conf.Server.Addr, ":", strconv.Itoa(conf.Server.Port)),
		Router:       mux.NewRouter(),
		Store:        sqlstore.NewStore(db),
		SessionStore: newCookieStore([]byte(conf.Env.CookieKey)),
		Config:       conf,
		Logger:       logger,
		// SwaggerUrl:   conf.Api.SwaggerUrl,
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

	// serv.Router.Use(handlers.CORS(handlers.ExposedHeaders([]string{"SET-COOKIE"})))
	serv.Router.Use(serv.setRequestId)
	serv.Router.Use(serv.logRequest)
	serv.Router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	url := "doc.json"
	serv.Router.PathPrefix("/swagger").HandlerFunc(httpSwagger.Handler(
		httpSwagger.URL(url), //The url pointing to API definition
	))
	// serv.SwaggerUrl

	// serv.Router.Handle(videoDir, serv.handleBase(http.FileServer(http.Dir(videoDir))))
	// serv.Router.PathPrefix("/static/").Handler(serv.handleStatic(http.StripPrefix("/static/", http.FileServer(http.Dir(videoDir)))))

	api := serv.Router.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/", serv.handleBase).Methods("GET")
	// api.PathPrefix("/video/").Handler(http.StripPrefix("/video/", http.FileServer(http.Dir(videoDir))))
	// api.Handle("/", serv.handleBase(http.FileServer(http.Dir("./video"))))
	api.HandleFunc("/carousel", serv.handleGetCarousel()).Methods("GET")
	api.HandleFunc("/newFilms", serv.handleGetNewFilms()).Methods("GET")
	api.HandleFunc("/film/{hash}", serv.HandleGetCurrentFilm()).Methods("GET")
	// api.PathPrefix("/film").Handler(serv.HandleGetCurrentFilm()).Methods("GET")

	api.HandleFunc("/users", serv.handleUsersCreate()).Methods("POST")
	api.HandleFunc("/sessions", serv.handleSessionsCreate()).Methods("POST")

	api.HandleFunc("/add/films", serv.handleAddFilms()).Methods("POST")

	private := api.PathPrefix("/private").Subrouter()
	private.Use(serv.authenticateUser)
	private.HandleFunc("/whoami", serv.handleWhoami()).Methods("GET")

}
