package sqlstore

import (
	film "github.com/Ropho/Cinema/internal/model/film"
	"github.com/sirupsen/logrus"
)

type SqlFilmRepository struct {
	store *SqlStore
}

// func (r *SqlFilmRepository) Create(u *film.Film) error {

// 	stmt, err := r.store.Db.Prepare("INSERT INTO films(email, pass) VALUES (?, ?)")
// 	if err != nil {
// 		logrus.Fatal("PREPARE INSERT ERROR: ", err)
// 		return err
// 	}

// 	_, err = stmt.Exec(u.Email, u.EncryptedPass)
// 	if err != nil {
// 		logrus.Error("EXEC INSERT ERROR: ", err)
// 		return err
// 	}

// 	return nil
// }

func (r *SqlFilmRepository) FindById(id int) (*film.Film, error) {

	f := &film.Film{
		Id: id,
	}

	err := r.store.Db.QueryRow(
		"SELECT name, location, description, url FROM films WHERE id = ?", id).Scan(
		&f.Name, &f.FilmPath, &f.Description, &f.PicUrl)
	if err != nil {
		logrus.Error("find film by id error: ", err)
		return nil, err
	}

	return f, nil
}

func (r *SqlFilmRepository) FindByName(name string) (*film.Film, error) {

	f := &film.Film{
		Name: name,
	}

	err := r.store.Db.QueryRow("SELECT id, location, description, url FROM films WHERE name = ?", f.Name).Scan(
		&f.Id, &f.FilmPath, &f.Description, &f.PicUrl)
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
		"SELECT * FROM films "+
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
		if err := rows.Scan(&tmpFilm.Id, &tmpFilm.Name, &tmpFilm.PicUrl,
			&tmpFilm.Category, &tmpFilm.Description, &tmpFilm.FilmPath, &tmpFilm.Rights); err != nil {
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
		"SELECT * FROM films "+
			"ORDER BY id "+
			"LIMIT ?", num)
	if err != nil {
		logrus.Error("select random ", num, " rows error: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tmpFilm film.Film
		if err := rows.Scan(&tmpFilm.Id, &tmpFilm.Name, &tmpFilm.PicUrl,
			&tmpFilm.Category, &tmpFilm.Description, &tmpFilm.FilmPath, &tmpFilm.Rights); err != nil {
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
