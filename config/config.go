package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Addr      string `yaml:"addr"`
	Port      int    `yaml:"port"`
	CookieKey string `yaml:"cookie_key"`
}

type DBaseConfig struct {
	DbUser string `yaml:"db_user"`
	DbPass string `yaml:"db_pass"`
	DbName string `yaml:"db_name"`
	DbAddr string `yaml:"db_addr"`
	DbPort int    `yaml:"db_port"`
}

type ApiConfig struct {
	SwaggerUrl string `yaml:"swagger_url"`
}

type EnvVar struct {
	SessionName string
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	DBase  DBaseConfig  `yaml:"db"`
	Api    ApiConfig    `yaml:"api"`
	Env    EnvVar
}

func NewConfig(logger *log.Logger) (*Config, error) {

	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		logger.Error("read config error: ", err)
		return nil, err
	}

	conf := Config{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		logger.Error("unable to parse config: ", err)
		return nil, err
	}

	err = getEnvVar(&conf)
	if err != nil {
		logger.Error("unable to parse env var: ", err)
		return nil, err
	}

	return &conf, nil
}

func getEnvVar(conf *Config) error {

	err := godotenv.Load("./config/key.env")
	if err != nil {
		log.Error("error loading .env file:", err)
		return err
	}

	conf.Env.SessionName = os.Getenv("sessionName")

	return nil
}
