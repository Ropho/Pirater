package main

import (
	"flag"
	"time"
	_ "time/tzdata"

	"github.com/Ropho/Pirater/config"
	"github.com/Ropho/Pirater/internal/server"

	log "github.com/sirupsen/logrus"
)

// @title KINOPOISK API
// @version 1.0
// @description U can access functions from here
// @schemes http https

// // @host web_server:80
// @host 192.168.31.100:80
// //@host 10.5.0.10:8080
// // @host 172.19.0.5:8080

//// @host nef_proj-web_server-1:80
// // @host 172.22.0.5:80

// //@host 192.168.31.100:80
// //@host 10.55.132.27

// @BasePath /api/

var DefaultLogger *log.Logger

func init() {
	DefaultLogger = log.New()

	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		DefaultLogger.Fatalf("unable to set time: [%w]", err)
	}
	time.Local = loc
}

func main() {

	configFile := flag.String("config", "config.yaml", "input custom config")
	flag.Parse()

	conf, err := config.NewConfig(*configFile)
	if err != nil {
		DefaultLogger.Fatalf("unable to init config: [%w]", err)
	}

	serv, err := server.NewServer(conf, DefaultLogger)
	if err != nil {
		DefaultLogger.Fatalf("unable to init server: [%w]", err)
	}

	if err := serv.Start(); err != nil {
		serv.Logger.Fatalf("unable to start server: [%w]", err)
	}
}
