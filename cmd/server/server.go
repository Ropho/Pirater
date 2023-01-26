package main

import (
	"net/http"

	"github.com/Ropho/Cinema1337/internal/config"
	"github.com/Ropho/Cinema1337/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.InfoLevel)

	conf := config.NewConfig()
	serv := server.NewServer(conf)

	// fmt.Println(serv)

	err := serv.Start()
	if err != nil {
		logrus.Fatal("SERVER INIT ERROR")
	}

	http.ListenAndServe(serv.IP_Port, serv.Router)

}
