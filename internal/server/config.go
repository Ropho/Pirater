package server

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ServAddr  string `yaml:"serv_addr"`
	Port      int    `yaml:"port"`
	CookieKey string `yaml:"cookie_key"`
}

func NewConfig() *Config {

	data, err := os.ReadFile("./config/server.yaml")
	if err != nil {
		logrus.Fatal("READING CONFIG FILE ERROR")
	}

	conf := Config{}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		logrus.Fatal("PARSE CONFIG error: ", err)
	}

	return &conf
}
