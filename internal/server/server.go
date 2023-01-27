package server

import (
	"net/http"
	"strconv"

	"github.com/Ropho/Cinema1337/internal/store"
	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

func Empty(w http.ResponseWriter, r *http.Request) {
	logrus.Info("NEW REQUEST")
}

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

	serv.Router.HandleFunc("/", HandleBase)

	return nil
}

func (serv *Server) Close() {
	serv.Store.Db.Close()
}
