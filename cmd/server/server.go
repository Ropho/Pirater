package main

import (
	"github.com/Ropho/Cinema/config"
	"github.com/Ropho/Cinema/internal/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

// @title KINOPOISK API
// @version 1.0
// @description U can access functions from here
// @schemes http https

// @host 192.168.31.100:8080
// @BasePath /

func init() {
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {

	conf, err := config.NewConfig()
	if err != nil {
		logrus.Panic("unable to init config: ", err)
	}

	/////////////////////////////////////////////////
	serv, err := server.NewServer(conf)
	if err != nil {
		logrus.Panic("unable to init server: ", err)
	}

	if err := serv.Start(); err != nil {
		logrus.Panic("unable to start server: ", err)
	}
}
