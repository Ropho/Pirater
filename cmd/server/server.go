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

// @host 192.168.31.100:80
// localhost:80
// 10.10.132.79:80

// @BasePath /api/

var DefaultLogger *log.Logger

func init() {
	DefaultLogger = log.New()
}

func main() {

	conf, err := config.NewConfig()
	if err != nil {
		DefaultLogger.Fatalf("unable to init config: [%w]", err)
	}

	serv, err := server.NewServer(conf, DefaultLogger)
	if err != nil {
		DefaultLogger.Fatalf("unable to init server: [%w]", err)
	}

	if err := serv.Start(); err != nil {
		serv.Logger.Panicf("unable to start server: [%w]", err)
	}
}
