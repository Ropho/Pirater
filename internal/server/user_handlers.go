package server

import (
	"encoding/json"
	"net/http"

	user "github.com/Ropho/Pirater/internal/model/user"
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
// @Param user body UserClientInfo true "User Registry"
// @Success      200  {string} string
// @Success      201  {string} string "Happily Registered"
// @Success      400  {string} string
// @Failure      422  {string}  string
// @Router /users [post]
func (serv *Server) handleUsersCreate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &UserClientInfo{}

		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			serv.error(w, r, http.StatusBadRequest, "")
			serv.Logger.Errorf("decode body error: [%w]", err)
			return
		}

		u := &user.User{
			Email: req.Email,
			Pass:  req.Pass,
			Right: user.Default,
		}

		if err := u.BeforeCreate(); err != nil {
			serv.error(w, r, http.StatusBadRequest, "PASS / EMAIL VALIDATION FAIL")
			serv.Logger.Errorf("BAD VALIDATION: [%w]", err)
			return
		}

		err = serv.Store.User().Create(u)
		if err != nil {
			serv.error(w, r, http.StatusUnprocessableEntity, "CREATE ERROR")
			serv.Logger.Errorf("CREATE ERROR: [%w]", err)
			return
		}

		serv.respond(w, r, http.StatusCreated, "REGISTERED")
		serv.Logger.Info("REGISTERED!!!")
	}
}

// Session Create godoc
// @Summary USER LOG
// @Tags W/O AUTH
// @Accept       json
// @Produce      json
// @Param input body UserClientInfo true "User log"
// @Success      202  {string} string "Happily Logged"
// @Failure      400  {string} string
// @Failure      422  {string}  string
// @Router /sessions [post]
func (serv *Server) handleSessionsCreate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &UserClientInfo{}

		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			serv.error(w, r, http.StatusBadRequest, "")
			serv.Logger.Errorf("DECODE BODY ERROR: [%w]", err)
			return
		}

		u := &user.User{
			Email: req.Email,
			Pass:  req.Pass,
		}

		err = user.Validate(u)
		if err != nil {
			serv.error(w, r, http.StatusBadRequest, "VALIDATION ERROR")
			serv.Logger.Errorf("VALIDATION USER ERROR: [%w]", err)
			return
		}

		ans, err := serv.Store.User().FindByEmail(u.Email)
		if err != nil {
			serv.error(w, r, http.StatusUnauthorized, "INCORRECT PASS / EMAIL")
			serv.Logger.Error("FIND BY EMAIL FAIl")
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(ans.EncryptedPass), []byte(u.Pass)); err != nil {

			serv.Logger.Error("INCORRECT PASS")
			serv.error(w, r, http.StatusUnauthorized, "INCORRECT PASS / EMAIL")

		} else {
			session, err := serv.SessionStore.Get(r, serv.Config.Env.SessionName)
			if err != nil {
				serv.Logger.Error(err)
				serv.error(w, r, http.StatusInternalServerError, "session get error")
				return
			}
			session.Values["user_id"] = ans.Id

			if err := session.Save(r, w); err != nil {
				serv.Logger.Error(err)
				serv.error(w, r, http.StatusInternalServerError, "session save error")
				return
			}

			serv.respond(w, r, http.StatusAccepted, "LOGGED IN")
		}
	}
}

// WHOAMI godoc
// @Summary WHOAMI
// @Tags AUTH
// @Success      200  {string} string
// @Router /private/whoami [get]
func (serv *Server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serv.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*user.User).Email)
	}
}
