package server

import (
	"net/http"

	_ "github.com/Ropho/Cinema/docs"

	"github.com/sirupsen/logrus"
)

// TEST FUNC godoc
// @Summary TESTING
// @Tags W/O AUTH
// @Router / [get]
func (s *Server) handleBase(w http.ResponseWriter, r *http.Request) {
	s.respond(w, r, http.StatusOK, "HELLO WORLD!")
	logrus.Info("HELLO WORLD!!!")
}

// func (s *Server) handleStatic(h http.Handler) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		logrus.Info("static file process")

// 		h.ServeHTTP(w, r)
// 	}
// }
