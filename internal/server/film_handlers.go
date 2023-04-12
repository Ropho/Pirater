package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	film "github.com/Ropho/Pirater/internal/model/film"
	"github.com/Ropho/Pirater/internal/utils"
)

// Session Create godoc
// @Summary GET CAROUSEL
// @Tags W/O AUTH
// @Produce json
// @Param  count query string true "number of films required"
// @Success      200  {array} server.handleGetCarousel.CarouselFilmInfo
// @Failure      500  {string}  string
// @Router /carousel [get]
func (serv *Server) handleGetCarousel() http.HandlerFunc {

	type CarouselFilmInfo struct {
		Hash      uint32 `json:"hash"`
		Name      string `json:"name"`
		HeaderUrl string `json:"pic_url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		if !r.URL.Query().Has("count") {
			serv.Logger.Error("count of film not given")
			serv.error(w, r, http.StatusBadRequest, "count of films not given")
			return
		}
		carouselNum, err := strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			serv.Logger.Error("unable to get number from given number of films")
			serv.error(w, r, http.StatusBadRequest, "bad given number of films")
			return
		}
		if carouselNum < 0 {
			serv.Logger.Error("carousel film count negative")
			serv.error(w, r, http.StatusBadRequest, "")
			return
		}

		var films []CarouselFilmInfo

		dbFilms, err := serv.Store.Film().GetCarouselFilmsInfo(carouselNum)
		if err != nil {
			serv.Logger.Errorf("get random films error: [%w]", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		for i := 0; i < len(dbFilms); i++ {
			carouselFilm := CarouselFilmInfo{
				Hash:      dbFilms[i].Hash,
				Name:      dbFilms[i].Name,
				HeaderUrl: dbFilms[i].HeaderUrl,
			}
			films = append(films, carouselFilm)
		}

		ans, err := json.Marshal(films)
		if err != nil {
			serv.Logger.Errorf("unable to marshal film: [%w]", err)
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
// @Router /newFilms [get]
func (serv *Server) handleGetNewFilms() http.HandlerFunc {

	type NewFilmInfo struct {
		Hash      uint32 `json:"hash"`
		Name      string `json:"name"`
		AfishaUrl string `json:"pic_url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if !r.URL.Query().Has("count") {
			serv.Logger.Error("count of film not given")
			serv.error(w, r, http.StatusBadRequest, "count of films not given")
			return
		}
		newFilmNum, err := strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			serv.Logger.Error("unable to get number from given number of films")
			serv.error(w, r, http.StatusBadRequest, "bad given number of films")
			return
		}

		if newFilmNum < 0 {
			serv.Logger.Error("number of new films negative")
			serv.error(w, r, http.StatusBadRequest, "bad given number of films")
		}

		var films []NewFilmInfo

		////////////////////////////////////////
		dbFilms, err := serv.Store.Film().GetNewFilmsInfo(newFilmNum)
		if err != nil {
			serv.Logger.Errorf("get new films error: [%w]", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		for i := 0; i < len(dbFilms); i++ {
			newFilm := NewFilmInfo{
				Name:      dbFilms[i].Name,
				AfishaUrl: dbFilms[i].AfishaUrl,
				Hash:      dbFilms[i].Hash,
			}
			films = append(films, newFilm)
		}

		ans, err := json.Marshal(films)
		if err != nil {
			serv.Logger.Errorf("unable to marshal film: [%w]", err)
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
// @Param hash path uint32 true "Film Hash"
// @Success      200  {object} server.handleGetCurrentFilm.GetFilmInfo
// @Failure      500  {string}  string
// @Router /film/{hash} [get]
func (serv *Server) handleGetCurrentFilm() http.HandlerFunc {

	type GetFilmInfo struct {
		Name        string   `json:"name"`
		Hash        uint32   `json:"hash"`
		Description string   `json:"description"`
		Categories  []string `json:"categories"`
		VideoUrl    string   `json:"video_url"`
		HeaderUrl   string   `json:"header_url"`
		AfishaUrl   string   `json:"afisha_url"`
		//CadreUrl []string
	}
	return func(w http.ResponseWriter, r *http.Request) {

		hash, err := strconv.ParseUint(mux.Vars(r)["hash"], 10, 32)
		if err != nil {
			serv.Logger.Error("request for film wrong hash: [%w]", err)
			serv.error(w, r, http.StatusBadRequest, "bad hash")
			return
		}

		film, err := serv.Store.Film().FindByHash(uint32(hash))
		if err != nil {
			serv.Logger.Errorf("unable to find film with its name: [%w]", err)
			serv.error(w, r, http.StatusInternalServerError, "")
		}

		filmInfo := &GetFilmInfo{
			Name:        film.Name,
			Hash:        film.Hash,
			Description: film.Description,
			Categories:  film.Categories,
			VideoUrl:    film.VideoUrl,
			HeaderUrl:   film.HeaderUrl,
			AfishaUrl:   film.AfishaUrl,
		}

		ans, err := json.Marshal(filmInfo)
		if err != nil {
			serv.Logger.Errorf("unable to marshal film: [%w]", err)
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
// @Failure      405  {string}  string
// @Failure      422  {string}  string
// @Router /private/admin/add/films [post]
func (serv *Server) handleAddFilms() http.HandlerFunc {

	type AddFilmInfo struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Categories  []string `json:"categories"`
		VideoUrl    string   `json:"video_url"`
		HeaderUrl   string   `json:"header_url"`
		AfishaUrl   string   `json:"afisha_url"`
		//CadreUrl []string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := []AddFilmInfo{}
		films := []film.Film{}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			serv.Logger.Errorf("DECODE BODY ERROR: [%w]", err)
			serv.error(w, r, http.StatusBadRequest, "BAD REQUEST")
			return
		}

		for i := 0; i < len(req); i++ {

			hash := utils.Hash([]byte(req[i].Name))
			if err != nil {
				serv.Logger.Errorf("hash algo error: [%w]", err)
				serv.error(w, r, http.StatusInternalServerError, "")
				return
			}

			films = append(films, film.Film{
				Name:        req[i].Name,
				Hash:        hash,
				Description: req[i].Description,
				Categories:  req[i].Categories,
				VideoUrl:    req[i].VideoUrl,
				HeaderUrl:   req[i].HeaderUrl,
				AfishaUrl:   req[i].AfishaUrl,
			})
		}

		err = serv.Store.Film().Create(films)
		if err != nil {
			serv.Logger.Errorf("create film error: [%w]", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		serv.respond(w, r, http.StatusAccepted, "films were successfully added")
	}

}
