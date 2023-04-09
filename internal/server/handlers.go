package server

import (
	"net/http"

	_ "github.com/Ropho/Pirater/docs"
)

// TEST FUNC godoc
// @Summary TESTING
// @Tags W/O AUTH
// @Router / [get]
func (serv *Server) handleBase(w http.ResponseWriter, r *http.Request) {
	serv.respond(w, r, http.StatusOK, "HELLO WORLD!")
	serv.Logger.Info("HELLO WORLD!!!")
}
