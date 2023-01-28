package main

import (
	"net/http"

	user "github.com/Ropho/Cinema1337/internal/model"

	"github.com/Ropho/Cinema1337/internal/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.InfoLevel)

	serv := server.NewServer()
	err := serv.Start()
	if err != nil {
		logrus.Fatal("SERVER INIT ERROR:", err)
	}

	u := &user.User{
		Id:    0,
		Email: "ded32@mail.ru",
		Pass:  "1488",
	}

	serv.Store.User().Create(u)
	// fmt.Println(u)

	http.ListenAndServe(serv.IP_Port, serv.Router)

	defer serv.Close()
}
