package main

import (
	"github.com/Ropho/Cinema/config"
	"github.com/Ropho/Cinema/internal/server"

	log "github.com/sirupsen/logrus"
)

// @title KINOPOISK API
// @version 1.0
// @description U can access functions from here
// @schemes http https

// @host localhost:8080
// @BasePath /

var DefaultLogger *log.Logger

func init() {
	DefaultLogger = log.New()
}

func main() {

	conf, err := config.NewConfig(DefaultLogger)
	if err != nil {
		DefaultLogger.Panic("unable to init config: ", err)
	}

	serv, err := server.NewServer(conf, DefaultLogger)
	if err != nil {
		DefaultLogger.Panic("unable to init server: ", err)
	}

	if err := serv.Start(); err != nil {
		serv.Logger.Panic("unable to start server: ", err)
	}
}
