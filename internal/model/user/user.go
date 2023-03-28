package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id            int       `json:"id"`
	Email         string    `json:"email"`
	Pass          string    `json:"pass,omitempty"`
	EncryptedPass string    `json:"-"`
	Registered    time.Time `json:"registered,omitempty"`
}

func (u *User) BeforeCreate() error {

	err := Validate(u)
	if err != nil {
		logrus.Error("VALIDATION USER ERROR: ", err)
		return err
	}

	enc, err := EncryptPass(u.Pass)
	if err != nil {
		logrus.Error("ENCRYPT PASS ERROR: ", err)
		return err
	}
	u.EncryptedPass = enc

	u.Sanitize()

	return nil
}

func Validate(u *User) error {

	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Pass, validation.By(requiredIf(u.EncryptedPass == "")),
			validation.Length(3, 20)))
}

func (u *User) Sanitize() {
	u.Pass = ""
}

func EncryptPass(s string) (string, error) {

	// r := rand.New(rand.NewSource(99))
	// passLen := r.Int() % 32

	data, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}
