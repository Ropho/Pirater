package server

import (
	"encoding/json"
	"io"
	"net/http"

	_ "github.com/Ropho/Cinema/docs"
	"github.com/Ropho/Cinema/internal/model"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const sessionName = "KINOPOISK"

type Request struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

// HANDLE BASE godoc
// @Summary TEST
// @Description testing
// @Tags W/O AUTH
// @Router /api [get]
func (s *Server) handleBase(w http.ResponseWriter, r *http.Request) {

	_, err := io.WriteString(w, "BASE RESPONSE")
	if err != nil {
		logrus.Fatal("RESPONSE WRITE ERROR: ", err)
	}
}

// USER REGISTER godoc
// @Summary USER REGISTER
// @Tags W/O AUTH
// @Accept       json
// @Produce      json
// @Param input body Request true "User Registry"
// @Success      200  {string} string
// @Success      201  {string} string "Happily Registered"
// @Success      400  {string} string
// @Failure      422  {string}  string
// @Router /api/users [post]
func (s *Server) handleUsersCreate(w http.ResponseWriter, r *http.Request) {
	req := &Request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, "BAD REQUEST")
		logrus.Error("DECODE BODY ERROR: ", err)
		return
	}

	u := &model.User{
		Email: req.Email,
		Pass:  req.Pass,
	}

	if err := u.BeforeCreate(); err != nil {
		s.error(w, r, http.StatusBadRequest, "PASS / EMAIL VALIDATION FAIL")
		logrus.Error("BAD VALIDATION: ", err)
		return
	}

	_, err = s.Store.User().Create(u)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, "CREATE ERROR")
		logrus.Error("CREATE ERROR", err)
		return
	}

	s.respond(w, r, http.StatusCreated, "REGISTERED")
	logrus.Info("REGISTERED!!!")
}

// Session Create godoc
// @Summary USER LOG
// @Tags W/O AUTH
// @Accept       json
// @Produce      json
// @Param input body Request true "User log"
// @Success      200  {string} string
// @Success      201  {string} string "Happily Logged"
// @Success      400  {string} string
// @Failure      422  {string}  string
// @Router /api/sessions [post]
func (s *Server) handleSessionsCreate(w http.ResponseWriter, r *http.Request) {

	req := &Request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, "BAD REQUEST")
		logrus.Error("DECODE BODY ERROR: ", err)
		return
	}

	u := &model.User{
		Email: req.Email,
		Pass:  req.Pass,
	}

	err = model.Validate(u)
	if err != nil {
		logrus.Error("VALIDATION USER ERROR: ", err)
		return
	}

	// if err := u.CheckLogUserParam(); err != nil {
	// 	s.error(w, r, http.StatusBadRequest, "PASS / EMAIL VALIDATION FAIL")
	// 	logrus.Error("BAD VALIDATION: ", err)
	// 	return
	// }
	ans, err := s.Store.User().FindByEmail(u.Email)
	if err != nil {
		s.error(w, r, http.StatusUnauthorized, "INCORRECT PASS / EMAIL")
		logrus.Error("FIND BY EMAIL FAIl")
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(ans.EncryptedPass), []byte(u.Pass)); err != nil {

		logrus.Error("INCORRECT PASS")
		s.error(w, r, http.StatusUnauthorized, "INCORRECT PASS / EMAIL")

	} else {
		session, err := s.SessionStore.Get(r, sessionName)
		if err != nil {
			logrus.Error(err)
			s.error(w, r, http.StatusInternalServerError, "COOKIE FAIL")
			return
		}
		session.Values["user_id"] = ans.Id

		logrus.Info("STORED ID", ans.Id)

		if err := session.Save(r, w); err != nil {
			logrus.Error(err)
			s.error(w, r, http.StatusInternalServerError, "COOKIE FAIL")
			return
		}

		s.respond(w, r, http.StatusAccepted, "LOGGED IN")
	}
}

func (serv *Server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serv.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User).Email)
	}
}
