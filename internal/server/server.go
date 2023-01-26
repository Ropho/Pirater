package server

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Ropho/Cinema1337/internal/config"
	"github.com/sirupsen/logrus"
)

func Empty(w http.ResponseWriter, r *http.Request) {
	logrus.Info("NEW REQUEST")
}

type Server struct {
	IP_Port string
	Router  *mux.Router
}

func NewServer(conf *config.Config) *Server {
	return &Server{
		IP_Port: conf.ServAddr + ":" + strconv.Itoa(conf.Port),
		Router:  mux.NewRouter(),
	}
}

func (serv *Server) Start() error {

	serv.Router.HandleFunc("/", HandleBase)

	return nil
}
