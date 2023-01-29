package server

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Ropho/Cinema1337/internal/model"
	"github.com/sirupsen/logrus"
)

func (s *Server) handleBase(w http.ResponseWriter, r *http.Request) {

	_, err := io.WriteString(w, "BASE RESPONSE")
	if err != nil {
		logrus.Fatal("RESPONSE WRITE ERROR: ", err)
	}
}

func (s *Server) handleUsersCreate(w http.ResponseWriter, r *http.Request) {

	type Request struct {
		Email string `json:"email"`
		Pass  string `json:"pass"`
	}
	req := &Request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		logrus.Fatal("DECODE BODY ERROR: ", err)
	}
	// fmt.Println(req)

	u := &model.User{
		Email: req.Email,
		Pass:  req.Pass,
	}

	u, err = s.Store.User().Create(u)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}

	// u.Sanitize()

	s.respond(w, r, http.StatusCreated, "REGISTERED")

}

func (s *Server) handleSessionsCreate(w http.ResponseWriter, r *http.Request) {

	type Request struct {
		Email string `json:"email"`
		Pass  string `json:"pass"`
	}
	req := &Request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		logrus.Fatal("DECODE BODY ERROR: ", err)
	}

	u := &model.User{
		Email: req.Email,
		Pass:  req.Pass,
	}

	logrus.Info(u.Pass)
	ans, err := s.Store.User().FindByEmail(u.Email)
	if err != nil {
		logrus.Fatal("FIND BY EMAIL ERROR: ", err)
	}
	logrus.Info("ALL GOOD")

	logrus.Info(ans.Pass)

	if u.Pass == ans.Pass {
		logrus.Info("ALL GOOD")
		s.respond(w, r, http.StatusAccepted, "LOGGED IN")

	} else {
		s.error(w, r, http.StatusUnauthorized, errors.New("WRONG PASS OR EMAIL"))
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {

	s.respond(w, r, code, err.Error())
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data string) {

	w.WriteHeader(code)

	w.Write([]byte(data))
}
