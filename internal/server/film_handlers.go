package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	film "github.com/Ropho/Cinema/internal/model/film"

	"github.com/Ropho/Cinema/internal/utils"
	"github.com/sirupsen/logrus"
)

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
				Description: dbFilms[i].DescPath,
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

// Session Create godoc
// @Summary ADD FILMS
// @Tags ADMIN
// @Accept       json
// @Produce      json
// @Param films body []server.handleAddFilms.AddFilmInfo true "films info"
// @Success      200  {string} string "Films added"
// @Failure      422  {string}  string
// @Router /api/add/films [post]
// /api/admin/add/films [post]
func (serv *Server) handleAddFilms() http.HandlerFunc {

	type AddFilmInfo struct {
		Name        string   `json:"name"`
		PicUrl      string   `json:"url"`
		DescPath    string   `json:"description,omitempty"`
		FilmPath    string   `json:"film,omitempty"`
		TrailerPath string   `json:"trailer,omitempty"`
		Categories  []string `json:"categories,omitempty"`
		Rights      []string `json:"rights,omitempty"`
		Rating      int      `json:"rating"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := []AddFilmInfo{}
		films := []film.Film{}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			serv.error(w, r, http.StatusBadRequest, "BAD REQUEST")
			logrus.Error("DECODE BODY ERROR: ", err)
			return
		}

		for i := 0; i < len(req); i++ {

			hash := utils.Hash([]byte(req[i].Name))
			if err != nil {
				logrus.Error("hash algo error", err)
				serv.error(w, r, http.StatusInternalServerError, "")
				return
			}

			films = append(films, film.Film{
				Name:        req[i].Name,
				PicUrl:      req[i].PicUrl,
				Hash:        hash,
				DescPath:    req[i].DescPath,
				FilmPath:    req[i].FilmPath,
				TrailerPath: req[i].TrailerPath,
				Categories:  req[i].Categories,
				Rights:      req[i].Rights,
				Rating:      req[i].Rating,
			})
		}
		err = serv.Store.Film().Create(films)
		if err != nil {
			logrus.Error("create error", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		serv.respond(w, r, http.StatusAccepted, "films were successfully added")
	}

}
