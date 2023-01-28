package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id            int
	Email         string
	Pass          string
	EncryptedPass string
}

func (u *User) BeforeCreate() error {

	err := validate(u)
	if err != nil {
		logrus.Error("VALIDATION USER ERROR: ", err)
		return err
	}

	enc, err := encryptPass(u.Pass)
	if err != nil {
		logrus.Error("ENCRYPT PASS ERROR: ", err)
		return err
	}
	u.EncryptedPass = enc

	return nil
}

func validate(u *User) error {

	return validation.ValidateStruct(u, validation.Field(u.Email, validation.Required, is.Email),
		validation.Field(u.Pass, validation.By(requiredIf(u.EncryptedPass == "")), validation.Length(4, 20)))
}

func encryptPass(s string) (string, error) {

	data, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}
