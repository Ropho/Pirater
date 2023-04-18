package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/xfrr/goffmpeg/transcoder"

	film "github.com/Ropho/Pirater/internal/model/film"
	"github.com/Ropho/Pirater/internal/utils"
)

// Session Create godoc
// @Summary GET CAROUSEL
// @Tags W/O AUTH
// @Produce json
// @Param  count query int true "number of films required"
// @Success      200  {array} server.handleGetCarousel.CarouselFilmInfo
// @Failure      500  {string}  string
// @Router /carousel [get]
func (serv *Server) handleGetCarousel() http.HandlerFunc {

	type CarouselFilmInfo struct {
		Hash      uint32 `json:"hash"`
		Name      string `json:"name"`
		HeaderUrl string `json:"header_url"`
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
// @Param  count query int true "number of films required"
// @Success      200  {array} server.handleGetNewFilms.NewFilmInfo
// @Failure      500  {string}  string
// @Router /newFilms [get]
func (serv *Server) handleGetNewFilms() http.HandlerFunc {

	type NewFilmInfo struct {
		Hash      uint32 `json:"hash"`
		Name      string `json:"name"`
		AfishaUrl string `json:"afisha_url"`
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

// Upload film godoc
// @Summary FILM UPLOAD
// @Tags ADMIN
// @Accept       mpfd
// @Param films formData server.handleFilmUpload.AddFilmInfo true "film info"
// @Param video formData file true "video"
// @Param header formData file true "header"
// @Param afisha formData file true "afisha"
// @Success      200  {string} string "Films added"
// @Failure      405  {string}  string
// @Failure      422  {string}  string
// @Router /private/admin/film/upload [post]
func (serv *Server) handleFilmUpload() http.HandlerFunc {

	//used for input params
	type AddFilmInfo struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Categories  []string `json:"categories"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		var err error

		err = r.ParseMultipartForm(20 * 1024 * 1024)
		if err != nil {
			serv.Logger.Errorf("max size overflow: [%w]", err)
			serv.error(w, r, http.StatusBadRequest, "max size overflow")
		}
		////////////////////////////////////////////////
		var filmInfo film.Film
		filmInfo.Name = r.FormValue("name")
		filmInfo.Description = r.FormValue("description")
		categories := r.FormValue("categories")
		filmInfo.Categories = strings.Split(categories, ",")

		filmInfo.Hash, err = utils.Hash([]byte(filmInfo.Name))
		if err != nil {
			serv.Logger.Errorf("hash algo error: [%w]", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		filmUrl := fmt.Sprintf("shared/%v", filmInfo.Hash)
		filmDir := fmt.Sprintf("./shared/%v", filmInfo.Hash)

		if err = os.Mkdir(filmDir, 0766); err != nil && !os.IsExist(err) {
			serv.Logger.Errorf("unable to create dir: [%w]", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		//////////////////////////////////////////////////////////////
		if err = serv.getFile(w, r, filmDir, "video"); err != nil {
			return
		}
		////////////////////////////////////////////
		// Create new instance of transcoder
		trans := new(transcoder.Transcoder)

		// Initialize transcoder passing the input file path and output file path
		err = trans.Initialize(filmDir+"/video", filmDir+"/video.m3u8")
		if err, ok := err.(*exec.ExitError); !ok && err != nil {
			serv.Logger.Errorf("unable to init trans: [%w]", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		// SET FRAME RATE TO MEDIAFILE
		// trans.MediaFile().SetFrameRate(70)
		// SET ULTRAFAST PRESET TO MEDIAFILE
		// trans.MediaFile().SetPreset("ultrafast")
		// Start transcoder process to check progress
		done := trans.Run(true)

		// This channel is used to wait for the transcoding process to end
		err = <-done
		if err, ok := err.(*exec.ExitError); !ok && err != nil {
			serv.Logger.Errorf("unable to convert video: [%w]", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		webServerInfo := fmt.Sprintf("%s:%d", serv.Config.Server.ProxyAddr, serv.Config.Server.ProxyPort)
		filmInfo.VideoUrl = fmt.Sprintf("http://%s/%s/%s", webServerInfo, filmUrl, "video.m3u8")
		/////////////////////////////////////////
		if err = serv.getFile(w, r, filmDir, "header"); err != nil {
			return
		}
		filmInfo.HeaderUrl = fmt.Sprintf("http://%s/%s/%s", webServerInfo, filmUrl, "header")

		if err = serv.getFile(w, r, filmDir, "afisha"); err != nil {
			return
		}
		filmInfo.AfishaUrl = fmt.Sprintf("http://%s/%s/%s", webServerInfo, filmUrl, "afisha")

		///////////////////////////////////////////
		films := []film.Film{filmInfo}
		err = serv.Store.Film().Create(films)
		if err != nil {
			serv.Logger.Errorf("create film error: [%w]", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		serv.respond(w, r, http.StatusAccepted, "film was successfully added")
	}
}

func (serv *Server) getFile(w http.ResponseWriter, r *http.Request, dirName string, fileName string) error {

	file, _, err := r.FormFile(fileName)
	if err != nil {
		serv.Logger.Errorf("form file error: [%w]", err)
		serv.error(w, r, http.StatusInternalServerError, "")
	}
	defer file.Close()

	resfile, err := os.Create(dirName + "/" + fileName)
	if err != nil {
		serv.Logger.Errorf("create file error: [%w]", err)
		serv.error(w, r, http.StatusInternalServerError, "")
		return err
	}
	defer resfile.Close()

	err = os.Chmod(dirName+"/"+fileName, 0666)
	if err != nil {
		serv.Logger.Errorf("chmod file error: [%w]", err)
		serv.error(w, r, http.StatusInternalServerError, "")
	}

	if _, err := io.Copy(resfile, file); err != nil {
		serv.Logger.Errorf("copy file error: [%w]", err)
		serv.error(w, r, http.StatusInternalServerError, "")
		return err
	}

	return nil
}

// Delete film godoc
// @Summary Delete Film
// @Tags ADMIN
// @Param  hash query uint32 true "number of films required"
// @Success      200  {string} string
// @Failure      500  {string}  string
// @Router /private/admin/film/delete [delete]
func (serv *Server) handleFilmDelete() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		hash, err := strconv.ParseUint(mux.Vars(r)["hash"], 10, 32)
		if err != nil {
			serv.Logger.Error("request for film wrong hash: [%w]", err)
			serv.error(w, r, http.StatusBadRequest, "bad hash")
			return
		}

		err = serv.Store.Film().DeleteByHash(uint32(hash))
		if err != nil {
			serv.Logger.Errorf("get random films error: [%w]", err)
			serv.error(w, r, http.StatusInternalServerError, "")
			return
		}

		serv.respond(w, r, http.StatusOK, "successfully deleted")
	}
}
