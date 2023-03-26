package sqlstore

import (
	model "github.com/Ropho/Cinema/internal/model/user"
	"github.com/sirupsen/logrus"
)

type SqlUserRepository struct {
	store *SqlStore
}

func (r *SqlUserRepository) Create(u *model.User) error {

	stmt, err := r.store.Db.Prepare("INSERT INTO users(email, encr_pass) VALUES (?, ?)")
	if err != nil {
		logrus.Panic("PREPARE INSERT ERROR: ", err)
		return err
	}

	_, err = stmt.Exec(u.Email, u.EncryptedPass)
	if err != nil {
		logrus.Error("EXEC INSERT ERROR: ", err)
		return err
	}

	return nil
}

func (r *SqlUserRepository) FindByEmail(email string) (*model.User, error) {

	u := &model.User{
		Email: email,
	}

	err := r.store.Db.QueryRow("SELECT id, encr_pass FROM users WHERE email = ?", email).Scan(&u.Id, &u.EncryptedPass)
	if err != nil {
		logrus.Error("FIND USER BY EMAIL ERROR: ", err)
		return nil, err
	}
	// logrus.Info(u)

	return u, nil
}

func (r *SqlUserRepository) FindById(id int) (*model.User, error) {

	u := &model.User{
		Id: id,
	}

	err := r.store.Db.QueryRow("SELECT id, email, encr_pass FROM users WHERE id = ?", id).Scan(&u.Id, &u.Email, &u.EncryptedPass)
	if err != nil {
		logrus.Error("FIND USER BY ID ERROR: ", err)
		return nil, err
	}

	return u, nil
}
