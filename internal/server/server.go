package server

import (
	"net/http"
	"strconv"

	"github.com/Ropho/Cinema1337/internal/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	IP_Port string
	Router  *mux.Router
	Store   *store.Store
}

func NewServer() *Server {

	conf := NewConfig()

	return &Server{
		IP_Port: conf.ServAddr + ":" + strconv.Itoa(conf.Port),
		Router:  mux.NewRouter(),
		Store:   store.NewStore(),
	}
}

func (serv *Server) Start() error {

	serv.Router.HandleFunc("/", serv.handleBase)
	serv.Router.HandleFunc("/users", serv.handleUsersCreate).Methods("POST")
	serv.Router.HandleFunc("/sessions", serv.handleSessionsCreate).Methods("POST")

	err := http.ListenAndServe(serv.IP_Port, serv.Router)
	if err != nil {
		logrus.Fatal("SERVER PROCESS ERROR: ", err)
	}

	return err
}

func (serv *Server) Close() {
	serv.Store.Db.Close()
}
