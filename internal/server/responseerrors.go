package server

import (
	"errors"
)

var (
	errValidation    = errors.New("email | pass validation fail")
	errIncorrectData = errors.New("email | pass incorrect")
)
