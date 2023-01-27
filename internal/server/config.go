package server

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ServAddr string `yaml:"serv_addr"`
	Port     int    `yaml:"port"`
}

func NewConfig() *Config {

	data, err := os.ReadFile("./config/server.yaml")
	if err != nil {
		logrus.Error("READING CONFIG FILE ERROR")
	}

	conf := Config{}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("PARSE CONFIG error: %v", err)
	}

	// fmt.Println(conf)

	return &conf
}
