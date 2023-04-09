package sqlstore

import (
	"fmt"

	user "github.com/Ropho/Pirater/internal/model/user"
)

type SqlUserRepository struct {
	store *SqlStore
}

func (r *SqlUserRepository) Create(u *user.User) error {

	tx, err := r.store.Db.Begin()
	if err != nil {
		return fmt.Errorf("begin Tx error: [%w]", err)
	}
	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO users(email, encr_pass) VALUES (?, ?)", u.Email, u.EncryptedPass)
	if err != nil {
		return fmt.Errorf("insert user into \"users\" db error: [%w]", err)
	}

	user_id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("unable to get last id: [%w]", err)
	}
	_, err = tx.Exec("INSERT INTO user_rights (user_id, right_id) VALUES (?, ?)", user_id, user.UserRights[u.Right])
	if err != nil {
		return fmt.Errorf("insert user into \"user_rights\" db error: [%w]", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit transaction error: [%w]", err)
	}

	return nil
}

func (r *SqlUserRepository) FindByEmail(email string) (*user.User, error) {

	u := &user.User{
		Email: email,
	}

	var right_id int
	err := r.store.Db.QueryRow("SELECT users.id, users.encr_pass, user_rights.right_id FROM users "+
		"JOIN user_rights ON users.id = user_rights.user_id "+
		"WHERE email = ?", email).Scan(
		&u.Id, &u.EncryptedPass, &right_id)
	if err != nil {
		return nil, fmt.Errorf("FIND USER BY EMAIL ERROR: [%w]", err)
	}

	err = r.store.Db.QueryRow("SELECT `right` FROM rights WHERE id = ?", right_id).
		Scan(&u.Right)
	if err != nil {
		return nil, fmt.Errorf("find user right error: [%w]", err)
	}

	return u, nil
}

func (r *SqlUserRepository) FindById(id int) (*user.User, error) {

	u := &user.User{
		Id: id,
	}
	var right_id int
	err := r.store.Db.QueryRow("SELECT users.email, users.encr_pass, user_rights.right_id FROM users "+
		"JOIN user_rights ON users.id = user_rights.user_id "+
		"WHERE id = ?", id).
		Scan(&u.Email, &u.EncryptedPass, &right_id)
	if err != nil {
		return nil, fmt.Errorf("find user by id error: [%w]", err)
	}

	err = r.store.Db.QueryRow("SELECT `right` FROM rights WHERE id = ?", right_id).
		Scan(&u.Right)
	if err != nil {
		return nil, fmt.Errorf("find user right error: [%w]", err)
	}

	return u, nil
}
