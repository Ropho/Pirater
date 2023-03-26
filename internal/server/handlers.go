package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/Ropho/Cinema/docs"
	user "github.com/Ropho/Cinema/internal/model/user"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const sessionName = "KINOPOISK"

type UserInfo struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

// TEST FUNC godoc
// @Summary TESTING
// @Tags W/O AUTH
// @Router / [get]
func (s *Server) handleBase(w http.ResponseWriter, r *http.Request) {
	s.respond(w, r, http.StatusOK, "HELLO WORLD!")
	logrus.Info("HELLO WORLD!!!")
}

func (s *Server) handleStatic(h http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("static file process")

		h.ServeHTTP(w, r)
	}
}

// USER REGISTER godoc
// @Summary USER REGISTER
// @Tags W/O AUTH
// @Accept       json
// @Produce      json
// @Param input body UserInfo true "User Registry"
// @Success      200  {string} string
// @Success      201  {string} string "Happily Registered"
// @Success      400  {string} string
// @Failure      422  {string}  string
// @Router /api/users [post]
func (s *Server) handleUsersCreate(w http.ResponseWriter, r *http.Request) {
	req := &UserInfo{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, "BAD REQUEST")
		logrus.Error("DECODE BODY ERROR: ", err)
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

// Session Create godoc
// @Summary USER LOG
// @Tags W/O AUTH
// @Accept       json
// @Produce      json
// @Param input body UserInfo true "User log"
// @Success      202  {string} string "Happily Logged"
// @Success      400  {string} string
// @Failure      422  {string}  string
// @Router /api/sessions [post]
func (s *Server) handleSessionsCreate(w http.ResponseWriter, r *http.Request) {

	req := &UserInfo{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, "BAD REQUEST")
		logrus.Error("DECODE BODY ERROR: ", err)
		return
	}

	u := &user.User{
		Email: req.Email,
		Pass:  req.Pass,
	}

	err = user.Validate(u)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, "VALIDATION ERROR")
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
		serv.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*user.User).Email)
	}
}

// Session Create godoc
// @Summary GET CAROUSEL
// @Tags W/O AUTH
// @Produce json
// @Param  count query string true "number of films required"
// @Success      200  {array} server.handleGetCarousel.CarouselFilmInfo
// @Failure      500  {string}  string
// @Router /api/carousel [get]
func (serv *Server) handleGetCarousel() http.HandlerFunc {

	type CarouselFilmInfo struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		PicUrl string `json:"url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		if !r.URL.Query().Has("count") {
			logrus.Error("count of film not given")
			serv.error(w, r, http.StatusBadRequest, "count of films not given")
			return
		}
		carouselNum, err := strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			logrus.Error("unable to get number from given number of films")
			serv.error(w, r, http.StatusBadRequest, "bad given number of films")
			return
		}
		if carouselNum < 0 {
			logrus.Error("carousel film count negative: ")
			serv.error(w, r, http.StatusBadRequest, "")
			return
		}

		var films []CarouselFilmInfo

		// ADD BATCH HERE / TRANSACTION
		////////////////////////////////////////
		dbFilms, err := serv.Store.Film().GetRandomFilms(carouselNum)
		if err != nil {
			logrus.Error("get random films error: ", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		for i := 0; i < len(dbFilms); i++ {
			carouselFilm := CarouselFilmInfo{
				Id:     i,
				Name:   dbFilms[i].Name,
				PicUrl: dbFilms[i].PicUrl,
			}
			films = append(films, carouselFilm)
		}

		ans, err := json.Marshal(films)
		if err != nil {
			logrus.Error("unable to marshal film: ", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}
		w.Write(ans)
	}
}

// Session Create godoc
// @Summary GET New Films
// @Tags W/O AUTH
// @Produce json
// @Param  count query string true "number of films required"
// @Success      200  {array} server.handleGetNewFilms.NewFilmInfo
// @Failure      500  {string}  string
// @Router /api/newFilms [get]
func (serv *Server) handleGetNewFilms() http.HandlerFunc {

	type NewFilmInfo struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		PicUrl      string `json:"url"`
		Description string `json:"description"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if !r.URL.Query().Has("count") {
			logrus.Error("count of film not given")
			serv.error(w, r, http.StatusBadRequest, "count of films not given")
			return
		}
		newFilmNum, err := strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			logrus.Error("unable to get number from given number of films")
			serv.error(w, r, http.StatusBadRequest, "bad given number of films")
			return
		}

		if newFilmNum < 0 {
			logrus.Error("number of new films negative")
			serv.error(w, r, http.StatusBadRequest, "bad given number of films")
		}

		var films []NewFilmInfo

		// ADD BATCH HERE / TRANSACTION
		////////////////////////////////////////
		dbFilms, err := serv.Store.Film().GetNewFilms(newFilmNum)
		if err != nil {
			logrus.Error("get new films error: ", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		for i := 0; i < newFilmNum; i++ {
			newFilm := NewFilmInfo{
				Id:          i,
				Name:        dbFilms[i].Name,
				PicUrl:      dbFilms[i].PicUrl,
				Description: dbFilms[i].Description,
			}
			films = append(films, newFilm)
		}

		ans, err := json.Marshal(films)
		if err != nil {
			logrus.Error("unable to marshal film: ", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}
		w.Write(ans)

	}
}

// Session Create godoc
// @Summary GET Current Film
// @Tags W/O AUTH
// @Produce json
// @Param name query string true "Film name"
// @Success      200  {object} model.Film
// @Failure      500  {string}  string
// @Router /api/film [get]
func (serv *Server) HandleGetCurrentFilm() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if !r.URL.Query().Has("name") {
			logrus.Error("request for film w/o name")
			serv.error(w, r, http.StatusBadRequest, "no film provided")
			return
		}

		name := r.URL.Query().Get("name")

		film, err := serv.Store.Film().FindByName(name)
		if err != nil {
			logrus.Error("unable to find film with its name: ", err)
			serv.error(w, r, http.StatusInternalServerError, "")
		}

		ans, err := json.Marshal(film)
		if err != nil {
			logrus.Error("unable to marshal film: ", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}
		w.Write(ans)
	}
}

// func (serv *Server)
