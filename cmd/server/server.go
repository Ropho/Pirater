package main

import (
	"github.com/Ropho/Cinema1337/internal/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.InfoLevel)

	serv := server.NewServer()
	if serv.Start() != nil {
		logrus.Exit(1)
	}

	// u := &user.User{
	// 	Id:    0,
	// 	Email: "ded32@mail.ru",
	// 	Pass:  "1488",
	// }

	// serv.Store.User().Create(u)
	// fmt.Println(u)

	defer serv.Close()
}
