package server

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func HandleBase(w http.ResponseWriter, r *http.Request) {

	_, err := io.WriteString(w, "BASE RESPONSE")
	if err != nil {
		logrus.Error("RESPONSE WRITE ERROR")
	}

}
