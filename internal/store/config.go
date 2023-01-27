package store

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	DbName string `yaml:"dbname"`
}

func NewConfig() Config {

	data, err := os.ReadFile("config/database.yaml")
	if err != nil {
		logrus.Fatal("READIN DB CONFIG ERROR", err)
	}

	conf := Config{}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		logrus.Fatal("UNMARSHAL DB CONFIG FILE ERROR", err)
	}

	return conf
}
