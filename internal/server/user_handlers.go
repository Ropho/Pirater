package server

import (
	"encoding/json"
	"net/http"

	user "github.com/Ropho/Cinema/internal/model/user"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserClientInfo struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

// USER REGISTER godoc
// @Summary USER REGISTER
// @Tags W/O AUTH
// @Accept       json
// @Produce      json
// @Param input body UserClientInfo true "User Registry"
// @Success      200  {string} string
// @Success      201  {string} string "Happily Registered"
// @Success      400  {string} string
// @Failure      422  {string}  string
// @Router /api/users [post]
func (s *Server) handleUsersCreate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &UserClientInfo{}

		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, "")
			logrus.Error("decode body error: ", err)
			return
		}

		u := &user.User{
			Email: req.Email,
			Pass:  req.Pass,
		}

		if err := u.BeforeCreate(); err != nil {
			s.error(w, r, http.StatusBadRequest, "PASS / EMAIL VALIDATION FAIL")
			logrus.Error("BAD VALIDATION: ", err)
			return
		}

		err = s.Store.User().Create(u)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, "CREATE ERROR")
			logrus.Error("CREATE ERROR", err)
			return
		}

		s.respond(w, r, http.StatusCreated, "REGISTERED")
		logrus.Info("REGISTERED!!!")
	}
}

// Session Create godoc
// @Summary USER LOG
// @Tags W/O AUTH
// @Accept       json
// @Produce      json
// @Param input body UserClientInfo true "User log"
// @Header 202 {string} Token "Set-Cookie"
// @Success      202  {string} string "Happily Logged"
// @Failure      400  {string} string
// @Failure      422  {string}  string
// @Router /api/sessions [post]
func (serv *Server) handleSessionsCreate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &UserClientInfo{}

		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			serv.error(w, r, http.StatusBadRequest, "")
			logrus.Error("DECODE BODY ERROR: ", err)
			return
		}

		u := &user.User{
			Email: req.Email,
			Pass:  req.Pass,
		}

		err = user.Validate(u)
		if err != nil {
			serv.error(w, r, http.StatusBadRequest, "VALIDATION ERROR")
			logrus.Error("VALIDATION USER ERROR: ", err)
			return
		}

		ans, err := serv.Store.User().FindByEmail(u.Email)
		if err != nil {
			serv.error(w, r, http.StatusUnauthorized, "INCORRECT PASS / EMAIL")
			logrus.Error("FIND BY EMAIL FAIl")
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(ans.EncryptedPass), []byte(u.Pass)); err != nil {

			logrus.Error("INCORRECT PASS")
			serv.error(w, r, http.StatusUnauthorized, "INCORRECT PASS / EMAIL")

		} else {
			session, err := serv.SessionStore.Get(r, serv.Config.Env.SessionName)
			if err != nil {
				logrus.Error(err)
				serv.error(w, r, http.StatusInternalServerError, "session get error")
				return
			}
			session.Values["user_id"] = ans.Id

			logrus.Info("STORED ID", ans.Id)

			if err := session.Save(r, w); err != nil {
				logrus.Error(err)
				serv.error(w, r, http.StatusInternalServerError, "session save error")
				return
			}

			serv.respond(w, r, http.StatusAccepted, "LOGGED IN")
		}
	}
}

func (serv *Server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serv.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*user.User).Email)
	}
}
