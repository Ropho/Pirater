package model

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type RightType string

const (
	Default   RightType = "DEFAULT"
	Premium   RightType = "PREMIUM"
	Moderator RightType = "MODERATOR"
	Admin     RightType = "ADMIN"
)

var UserRights = map[RightType]int{
	Default:   1,
	Premium:   2,
	Moderator: 3,
	Admin:     4,
}

type User struct {
	Id            int       `json:"id"`
	Email         string    `json:"email"`
	Pass          string    `json:"-"`
	EncryptedPass string    `json:"-"`
	Right         RightType `json:"right"`
	Registered    time.Time `json:"registered"`
	Modified      time.Time `json:"modified"`
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

	err := validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Pass, validation.By(requiredIf(u.EncryptedPass == "")),
			validation.Length(3, 20)))
	if err != nil {
		return fmt.Errorf("validation error: [%w]", err)
	}

	return nil
}

func (u *User) Sanitize() {
	u.Pass = ""
}

func EncryptPass(s string) (string, error) {

	data, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}
