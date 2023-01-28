package store

import (
	"log"

	user "github.com/Ropho/Cinema1337/internal/model"
	"github.com/sirupsen/logrus"
)

type UserRepo struct {
	store *Store
}

func (r *UserRepo) Create(u *user.User) (*user.User, error) {

	stmt, err := r.store.Db.Prepare("INSERT INTO users(email, pass) VALUES (?, ?)")
	if err != nil {
		logrus.Fatal("PREPARE INSERT ERROR: ", err)
	}

	res, err := stmt.Exec(u.Email, u.Pass)
	if err != nil {
		logrus.Fatal("EXEC INSERT ERROR: ", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	return u, nil
}

func (r *UserRepo) FindByEmail(email string) (*user.User, error) {

	u := &user.User{
		Email: email,
	}

	err := r.store.Db.QueryRow("SELECT id, pass FROM users WHERE email = ?", email).Scan(&u.Id, &u.EncryptedPass)
	if err != nil {
		logrus.Error("FIND USER BY EMAIL ERROR: ", err)
		return nil, err
	}

	return u, nil
}
