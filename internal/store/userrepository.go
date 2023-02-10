package store

import (
	user "github.com/Ropho/Cinema/internal/model"
	"github.com/sirupsen/logrus"
)

type UserRepo struct {
	store *Store
}

func (r *UserRepo) Create(u *user.User) (*user.User, error) {

	stmt, err := r.store.Db.Prepare("INSERT INTO users(email, pass) VALUES (?, ?)")
	if err != nil {
		logrus.Fatal("PREPARE INSERT ERROR: ", err)
		return nil, err
	}

	_, err = stmt.Exec(u.Email, u.EncryptedPass)
	if err != nil {
		logrus.Error("EXEC INSERT ERROR: ", err)
		return nil, err
	}

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
	// logrus.Info(u)

	return u, nil
}

func (r *UserRepo) FindById(id int) (*user.User, error) {

	u := &user.User{
		Id: id,
	}

	err := r.store.Db.QueryRow("SELECT id, email, pass FROM users WHERE id = ?", id).Scan(&u.Id, &u.Email, &u.EncryptedPass)
	if err != nil {
		logrus.Error("FIND USER BY ID ERROR: ", err)
		return nil, err
	}

	return u, nil
}
