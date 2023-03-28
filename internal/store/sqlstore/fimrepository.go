package sqlstore

import (
	film "github.com/Ropho/Cinema/internal/model/film"
	"github.com/sirupsen/logrus"
)

type SqlFilmRepository struct {
	store *SqlStore
}

func (r *SqlFilmRepository) Create(films []film.Film) error {

	numOfFields := 7
	params := make([]interface{}, 0, len(films)*numOfFields)

	command := "INSERT INTO films (name, pic_url, desc_path, film_path, trailer_path, hash, rating) VALUES "

	for i := 0; i < len(films); i++ {

		params = append(params, films[i].Name, films[i].PicUrl, films[i].DescPath,
			films[i].FilmPath, films[i].TrailerPath, films[i].Hash, films[i].Rating)
		command += "(?, ?, ?, ?, ?, ?, ?)"
		if i != len(films)-1 {
			command += ",\n"
		}
	}

	logrus.Info(command)
	logrus.Info(params...)
	// stmt, err := r.store.Db.Prepare(command)
	// if err != nil {
	// 	logrus.Fatal("PREPARE INSERT ERROR: ", err)
	// 	return err
	// }

	_, err := r.store.Db.Exec(command, params...)
	if err != nil {
		logrus.Error("EXEC INSERT ERROR: ", err)
		return err
	}

	return nil
}

// func (r *SqlFilmRepository) FindById(id int) (*film.Film, error) {

// 	f := &film.Film{
// 		Id: id,
// 	}

// 	err := r.store.Db.QueryRow(
// 		"SELECT name, location, description, url FROM films WHERE id = ?", id).Scan(
// 		&f.Name, &f.FilmPath, &f.DescPath, &f.PicUrl)
// 	if err != nil {
// 		logrus.Error("find film by id error: ", err)
// 		return nil, err
// 	}

// 	return f, nil
// }

func (r *SqlFilmRepository) FindByName(name string) (*film.Film, error) {

	f := &film.Film{
		Name: name,
	}

	err := r.store.Db.QueryRow("SELECT id, hash, trailer_path, film_path, desc_path, pic_url, rating FROM films WHERE name = ?", f.Name).Scan(
		&f.Id, &f.Hash, &f.TrailerPath, &f.FilmPath, &f.DescPath, &f.PicUrl, &f.Rating)
	if err != nil {
		logrus.Error("find film by name error: ", err)
		return nil, err
	}

	return f, nil
}

func (r *SqlFilmRepository) CountAllRows() (int, error) {

	var cnt int

	err := r.store.Db.QueryRow("SELECT COUNT(*) FROM films").Scan(&cnt)
	if err != nil {
		logrus.Error("count films error: ", err)
		return 0, err
	}

	return cnt, nil
}

func (r *SqlFilmRepository) GetRandomFilms(num int) ([]film.Film, error) {

	rows, err := r.store.Db.Query(
		"SELECT name, hash, trailer_path, film_path, desc_path, pic_url, rating FROM films "+
			"ORDER BY RAND() "+
			"LIMIT ?", num)
	if err != nil {
		logrus.Error("select random ", num, " rows error: ", err)
		return nil, err
	}
	defer rows.Close()

	films := make([]film.Film, 0, num)

	for rows.Next() {
		var tmpFilm film.Film
		if err := rows.Scan(
			&tmpFilm.Name, &tmpFilm.Hash, &tmpFilm.TrailerPath, &tmpFilm.FilmPath,
			&tmpFilm.DescPath, &tmpFilm.PicUrl, &tmpFilm.Rating); err != nil {
			logrus.Error("scan rows error: ", err)
			return nil, err
		}

		films = append(films, tmpFilm)
	}
	if err := rows.Err(); err != nil {
		logrus.Error("get random films error: ", err)
		return nil, err
	}

	return films, err
}

func (r *SqlFilmRepository) GetNewFilms(num int) ([]film.Film, error) {

	films := make([]film.Film, 0, num)

	rows, err := r.store.Db.Query(
		"SELECT added, name, hash, trailer_path, film_path, desc_path, pic_url, rating FROM films "+
			"ORDER BY STR_TO_DATE(`added`,'%m/%d/%Y %h:%i:%s %p')"+
			"LIMIT ?", num)
	if err != nil {
		logrus.Error("select random ", num, " rows error: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tmpFilm film.Film
		if err := rows.Scan(&tmpFilm.Timestamp, &tmpFilm.Name, &tmpFilm.Hash, &tmpFilm.TrailerPath, &tmpFilm.FilmPath,
			&tmpFilm.DescPath, &tmpFilm.PicUrl, &tmpFilm.Rating); err != nil {
			logrus.Error("scan rows error: ", err)
			return nil, err
		}

		films = append(films, tmpFilm)
	}
	if err := rows.Err(); err != nil {
		logrus.Error("get new films error: ", err)
		return nil, err
	}

	return films, err
}
