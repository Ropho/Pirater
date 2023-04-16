package server

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	user "github.com/Ropho/Pirater/internal/model/user"
)

type UserInputInfo struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type UserOutputInfo struct {
	Email      string `json:"email"`
	Right      string `json:"right"`
	Registered string `json:"registered,omitempty"`
	Modified   string `json:"modified,omitempty"`
}

func makeOutputInfo(u *user.User) *UserOutputInfo {

	outputInfo := &UserOutputInfo{
		Email: u.Email,
		Right: string(u.Right),
	}

	registered := u.Registered.Format("2006-01-02 15:04:05")
	if registered != "0001-01-01 00:00:00" {
		outputInfo.Registered = registered
	}

	modified := u.Modified.Format("2006-01-02 15:04:05")
	if modified != "0001-01-01 00:00:00" {
		outputInfo.Modified = modified
	}

	return outputInfo
}

// USER REGISTER godoc
// @Summary USER REGISTER
// @Tags W/O AUTH
// @Accept       json
// @Produce      json
// @Param user body UserInputInfo true "User Registry"
// @Success      200  {object} UserOutputInfo
// @Failure      400  {string} string
// @Failure 500 {string} string
// @Router /users [post]
func (serv *Server) handleUsersCreate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &UserInputInfo{}

		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			serv.Logger.Errorf("decode body error: [%w]", err)
			serv.error(w, r, http.StatusBadRequest, "")
			return
		}

		u := &user.User{
			Email: req.Email,
			Pass:  req.Pass,
			Right: user.Default,
		}

		if err := u.BeforeCreate(); err != nil {
			serv.error(w, r, http.StatusBadRequest, errValidation.Error())
			serv.Logger.Errorf("BAD VALIDATION: [%w], %v", err, u)
			return
		}

		err = serv.Store.User().Create(u)
		if err != nil {
			serv.Logger.Errorf("unable to create user: [%w]", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		ans, err := json.Marshal(makeOutputInfo(u))
		if err != nil {
			serv.Logger.Errorf("unable to marshal userInfo: [%w], %v", err, ans)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}
		w.Write(ans)

		serv.Logger.Info("REGISTERED!!!", u)
	}
}

// Session Create godoc
// @Summary USER LOG IN
// @Tags W/O AUTH
// @Accept       json
// @Produce      json
// @Param input body UserInputInfo true "User log in"
// @Success      202  {string} string "Happily Logged"
// @Failure      400  {string} string
// @Failure 500 {string} string
// @Router /sessions [post]
func (serv *Server) handleSessionsCreate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &UserInputInfo{}

		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			serv.Logger.Errorf("decode body error: [%w], %v", err, r.Body)
			serv.error(w, r, http.StatusBadRequest, "")
			return
		}

		u := &user.User{
			Email: req.Email,
			Pass:  req.Pass,
		}

		err = user.Validate(u)
		if err != nil {
			serv.Logger.Errorf("VALIDATION USER ERROR: [%w], %v", err, u)
			serv.error(w, r, http.StatusBadRequest, errValidation.Error())
			return
		}

		ans, err := serv.Store.User().FindByEmail(u.Email)
		if err != nil {
			serv.error(w, r, http.StatusUnauthorized, errIncorrectData.Error())
			serv.Logger.Errorf("FIND BY EMAIL FAIl: [%w], %v", err, u)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(ans.EncryptedPass), []byte(u.Pass)); err != nil {

			serv.Logger.Error("INCORRECT PASS")
			serv.error(w, r, http.StatusUnauthorized, errIncorrectData.Error())

		} else {
			session, err := serv.Store.Cookie().Get(r, serv.Config.Env.SessionName)
			if err != nil {
				serv.Logger.Errorf("unable to get session: [%w]", err)
				serv.error(w, r, http.StatusInternalServerError, "")
				return
			}
			session.Values["user_id"] = ans.Id

			if err := session.Save(r, w); err != nil {
				serv.Logger.Errorf("unable to save session: [%w]", err)
				serv.error(w, r, http.StatusInternalServerError, "")
				return
			}

			ansByte, err := json.Marshal(makeOutputInfo(ans))
			if err != nil {
				serv.Logger.Errorf("unable to marshal userInfo: [%w], %v", err, ans)
				serv.error(w, r, http.StatusInternalServerError, "")
				return
			}
			w.Write(ansByte)

			serv.Logger.Info("Logged in!!!", u)
		}
	}
}

// Session End godoc
// @Summary USER LOG OUT
// @Tags AUTH
// @Success      202  {string} string "Happily Logged Out"
// @Failure      400  {string} string
// @Failure 500 {string} string
// @Router /private/out [get]
func (serv *Server) handleSessionsEnd() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		err := serv.Store.Cookie().Delete(w, r, serv.Config.Env.SessionName)
		if err != nil {
			serv.Logger.Errorf("unable to delete session: [%w]", err)
			serv.error(w, r, http.StatusInternalServerError, "")
		}

		serv.respond(w, r, http.StatusOK, "LOGGED OUT")
	}
}

// WHOAMI godoc
// @Summary WHOAMI
// @Tags AUTH
// @Success      200  {object} UserOutputInfo
// @Failure 500 {string} string
// @Router /private/whoami [get]
func (serv *Server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, ok := r.Context().Value(ctxKeyUser).(*user.User)
		if !ok {
			serv.Logger.Error("unable to get user from context")
			serv.error(w, r, http.StatusInternalServerError, "")
		}

		ans, err := json.Marshal(makeOutputInfo(u))
		if err != nil {
			serv.Logger.Errorf("unable to marshal userInfo: [%w], %v", err, ans)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(ans)

		serv.Logger.Infof("handle whoami: [%v]", u)
	}
}
