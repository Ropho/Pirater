package sqlstore

import (
	"fmt"
	"time"

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

	result, err := tx.Exec("INSERT INTO users(email, encr_pass) VALUES (?, ?)",
		u.Email, u.EncryptedPass)
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
	var registered string
	var modified string
	err := r.store.Db.QueryRow("SELECT users.id, users.encr_pass, user_rights.right_id, "+
		"users.registered, users.modified "+
		"FROM users "+
		"JOIN user_rights ON users.id = user_rights.user_id "+
		"WHERE email = ?", email).Scan(
		&u.Id, &u.EncryptedPass, &right_id, &registered, &modified)
	if err != nil {
		return nil, fmt.Errorf("FIND USER BY EMAIL ERROR: [%w]", err)
	}
	u.Registered, err = ParseTime(registered)
	if err != nil {
		return nil, fmt.Errorf("unable to parse registered time: [%w]", err)
	}
	u.Modified, err = ParseTime(modified)
	if err != nil {
		return nil, fmt.Errorf("unable to parse modified time: [%w]", err)
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
	var registered string
	var modified string
	err := r.store.Db.QueryRow("SELECT users.email, users.encr_pass, user_rights.right_id, "+
		"users.registered, users.modified "+
		"FROM users "+
		"JOIN user_rights ON users.id = user_rights.user_id "+
		"WHERE id = ?", id).
		Scan(&u.Email, &u.EncryptedPass, &right_id, &registered, &modified)
	if err != nil {
		return nil, fmt.Errorf("find user by id error: [%w]", err)
	}
	u.Registered, err = ParseTime(registered)
	if err != nil {
		return nil, fmt.Errorf("unable to parse registered time: [%w]", err)
	}
	u.Modified, err = ParseTime(modified)
	if err != nil {
		return nil, fmt.Errorf("unable to parse modified time: [%w]", err)
	}

	err = r.store.Db.QueryRow("SELECT `right` FROM rights WHERE id = ?", right_id).
		Scan(&u.Right)
	if err != nil {
		return nil, fmt.Errorf("find user right error: [%w]", err)
	}

	return u, nil
}

func ParseTime(timeArg string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", timeArg)
}
