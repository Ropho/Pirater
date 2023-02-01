package server

import (
	"encoding/json"
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
		s.error(w, r, http.StatusBadRequest, "BAD REQUEST")
		logrus.Fatal("DECODE BODY ERROR: ", err)
	}
	// fmt.Println(req)

	u := &model.User{
		Email: req.Email,
		Pass:  req.Pass,
	}

	u, err = s.Store.User().Create(u)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, "ALREADY CREATED")
		return
	}

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
		s.error(w, r, http.StatusBadRequest, "BAD REQUEST")
		logrus.Fatal("DECODE BODY ERROR: ", err)
	}

	u := &model.User{
		Email: req.Email,
		Pass:  req.Pass,
	}

	ans, err := s.Store.User().FindByEmail(u.Email)
	if err != nil {
		s.error(w, r, http.StatusUnauthorized, "INCORRECT PASS / EMAIL")
		return
	}

	if u.Pass == ans.Pass {
		s.respond(w, r, http.StatusAccepted, "LOGGED IN")
	} else {
		s.error(w, r, http.StatusUnauthorized, "INCORRECT PASS / EMAIL")
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, data string) {

	s.respond(w, r, code, responseErr(data))
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data string) {

	w.WriteHeader(code)

	w.Write([]byte(responseInfo(data)))
}

func responseErr(s string) string {
	return "\033[31m" + s + "\n\033[0m"
}

func responseInfo(s string) string {
	return "\033[34m" + s + "\n\033[0m"
}
