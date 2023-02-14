package main

import (
	"github.com/Ropho/Cinema/internal/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

// @title KINOPOISK API
// @version 1.0
// @description U can access functions from here
// @schemes http https

// @host localhost:8080
// @BasePath /

func main() {

	logrus.Info("SUCCESFULLY IN HERE")

	logrus.SetLevel(logrus.InfoLevel)

	/////////////////////////////////////////////////
	serv := server.NewServer()
	if serv.Start() != nil {
		logrus.Exit(1)
	}

	defer serv.Close()
}
