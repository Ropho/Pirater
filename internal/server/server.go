package server

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/Ropho/Pirater/config"
	"github.com/Ropho/Pirater/internal/store"
)

type ctxKey int8

const (
	ctxKeyUser ctxKey = iota
	ctxKeyRequestId
)

type Server struct {
	IP_Port string
	Router  *mux.Router
	Store   store.Store
	Config  *config.Config
	// SessionStore sessions.Store
	Logger *log.Logger
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, data string) {
	s.respond(w, r, code, data)
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data string) {
	w.WriteHeader(code)
	w.Write([]byte(data))
}
