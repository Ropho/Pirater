package server

import (
	"net/http"
	"strconv"

	"github.com/Ropho/Cinema/config"

	"github.com/Ropho/Cinema/internal/store/sqlstore"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewServer(conf *config.Config) (*Server, error) {

	sStore := sessions.NewCookieStore([]byte(conf.Server.CookieKey))
	sStore.MaxAge(1000)

	db, err := newDb(&conf.DBase)
	if err != nil {
		logrus.Error("db init error: ", err)
		return nil, err
	}
	serv := &Server{
		IP_Port:      conf.Server.Addr + ":" + strconv.Itoa(conf.Server.Port),
		Router:       mux.NewRouter(),
		Store:        sqlstore.NewStore(db),
		SessionStore: sStore,
		Config:       conf,
	}
	return serv, nil
}

func (serv *Server) Start() error {

	serv.Router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	serv.Router.Use(serv.setRequestId)

	serv.Router.HandleFunc("/", serv.handleBase).Methods("GET")

	serv.Router.PathPrefix("/swagger").HandlerFunc(httpSwagger.Handler(
		httpSwagger.URL(serv.SwaggerUrl), //The url pointing to API definition
	)).Methods("GET")

	// serv.Router.Handle(videoDir, serv.handleBase(http.FileServer(http.Dir(videoDir))))
	serv.Router.PathPrefix("/static/").Handler(serv.handleStatic(http.StripPrefix("/static/", http.FileServer(http.Dir(videoDir)))))

	api := serv.Router.PathPrefix("/api/").Subrouter()
	// api.PathPrefix("/video/").Handler(http.StripPrefix("/video/", http.FileServer(http.Dir(videoDir))))
	// api.Handle("/", serv.handleBase(http.FileServer(http.Dir("./video"))))
	api.HandleFunc("/carousel", serv.handleGetCarousel()).Methods("GET")
	api.HandleFunc("/newFilms", serv.handleGetNewFilms()).Methods("GET")
	api.HandleFunc("/film", serv.HandleGetCurrentFilm()).Methods("GET")

	api.HandleFunc("/users", serv.handleUsersCreate).Methods("POST")
	api.HandleFunc("/sessions", serv.handleSessionsCreate).Methods("POST")

	private := api.PathPrefix("/private").Subrouter()
	private.Use(serv.authenticateUser)
	private.HandleFunc("/whoami", serv.handleWhoami()).Methods("GET")

	logrus.Info("server starting\n")
	err := http.ListenAndServe(serv.IP_Port, serv.Router)
	if err != nil {
		logrus.Error("server serve error: ", err)
		return err
	}

	logrus.Info("server closed unexpectedly\n")
	return err

}
