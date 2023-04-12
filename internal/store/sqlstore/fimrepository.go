package sqlstore

import (
	"fmt"

	film "github.com/Ropho/Pirater/internal/model/film"
	"github.com/sirupsen/logrus"
)

type SqlFilmRepository struct {
	store *SqlStore
}

func (r *SqlFilmRepository) Create(films []film.Film) error {

	tx, err := r.store.Db.Begin()
	if err != nil {
		return fmt.Errorf("begin Tx error: [%w]", err)
	}
	defer tx.Rollback()

	for i := 0; i < len(films); i++ {

		result, err := tx.Exec("INSERT INTO films (name, description, afisha_url, header_url, video_url, hash) "+
			"VALUES (?, ?, ?, ?, ?, ?)",
			films[i].Name, films[i].Description, films[i].AfishaUrl, films[i].HeaderUrl, films[i].VideoUrl, films[i].Hash)
		if err != nil {
			return fmt.Errorf("insert film [%v] into \"films\" db error: [%w]", films[i], err)
		}

		film_id, err := result.LastInsertId()
		if err != nil {
			return fmt.Errorf("unable to get last id: [%w]", err)
		}

		///////////////////////BATCH INSERT CATEGORIES
		command := "INSERT INTO film_categories (film_id, category_id) VALUES "
		var params []interface{}
		var category_id int
		for j := 0; j < len(films[i].Categories); j++ {

			err := tx.QueryRow("SELECT id FROM categories WHERE category = ?", films[i].Categories[j]).
				Scan(&category_id)
			if err != nil {
				return fmt.Errorf("find category error: [%v]", films[i].Categories[j])
			}

			params = append(params, film_id, category_id)
			command += "(?, ?)"
			if j != len(films[i].Categories)-1 {
				command += ",\n"
			}
		}
		_, err = tx.Exec(command, params...)
		if err != nil {
			return fmt.Errorf("EXEC INSERT ERROR: [%w]", err)
		}

	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit transaction error: [%w]", err)
	}

	return nil
}

func (r *SqlFilmRepository) FindByHash(hash uint32) (*film.Film, error) {

	f := &film.Film{
		Hash: hash,
	}

	err := r.store.Db.QueryRow("SELECT id, name, description, "+
		"video_url, header_url, afisha_url FROM films WHERE hash = ?", f.Hash).Scan(
		&f.Id, &f.Name, &f.Description, &f.VideoUrl, &f.HeaderUrl, &f.AfishaUrl)
	if err != nil {
		return nil, fmt.Errorf("find film by hash error: [%w]", err)
	}

	rows, err := r.store.Db.Query("select category from categories "+
		"WHERE id IN (SELECT film_categories.category_id FROM "+
		"film_categories JOIN films ON films.id "+
		"= film_categories.film_id WHERE films.hash = ?)", f.Hash)
	if err != nil {
		return nil, fmt.Errorf("query for categories error: [%w]", err)
	}
	defer rows.Close()
	var category string
	for rows.Next() {
		if err = rows.Scan(&category); err != nil {
			return nil, fmt.Errorf("unable to retrieve category: [%w]", err)
		}
		f.Categories = append(f.Categories, category)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows end error: [%w]", err)
	}

	logrus.Info(f.Categories)

	return f, nil
}

func (r *SqlFilmRepository) CountAllRows() (int, error) {

	var cnt int
	err := r.store.Db.QueryRow("SELECT COUNT(*) FROM films").Scan(&cnt)
	if err != nil {
		return 0, fmt.Errorf("count films error: [%w]", err)
	}

	return cnt, nil
}

func (r *SqlFilmRepository) GetCarouselFilmsInfo(num int) ([]film.Film, error) {

	rows, err := r.store.Db.Query(
		"SELECT name, hash, header_url FROM films "+
			"ORDER BY RAND() "+
			"LIMIT ?", num)
	if err != nil {
		return nil, fmt.Errorf("select random [%d] rows error [%w]", num, err)
	}
	defer rows.Close()

	films := make([]film.Film, 0, num)

	for rows.Next() {
		var tmpFilm film.Film
		if err := rows.Scan(
			&tmpFilm.Name, &tmpFilm.Hash, &tmpFilm.HeaderUrl); err != nil {
			return nil, fmt.Errorf("scan rows error: [%w]", err)
		}

		films = append(films, tmpFilm)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get random films error: [%w]", err)
	}

	return films, err
}

func (r *SqlFilmRepository) GetNewFilmsInfo(num int) ([]film.Film, error) {

	films := make([]film.Film, 0, num)

	rows, err := r.store.Db.Query(
		"SELECT name, hash, afisha_url FROM films "+
			"ORDER BY id "+
			"LIMIT ?", num)
	if err != nil {
		return nil, fmt.Errorf("select new [%d] rows error: [%w]", num, err)
	}
	defer rows.Close()

	for rows.Next() {
		var tmpFilm film.Film
		if err := rows.Scan(&tmpFilm.Name, &tmpFilm.Hash, &tmpFilm.AfishaUrl); err != nil {
			return nil, fmt.Errorf("scan rows error: [%w]", err)
		}

		films = append(films, tmpFilm)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get new films error: [%w]", err)
	}

	return films, err
}
