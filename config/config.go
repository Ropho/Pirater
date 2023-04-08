package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
}

type DBaseConfig struct {
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
}

// type ApiConfig struct {
// SwaggerUrl string `yaml:"swagger_url"`
// }

type EnvVar struct {
	SessionName string `yaml:"session_name"`
	CookieKey   string `yaml:"cookie_key"`
}

type LogConfig struct {
	Level string `yaml:"level"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	DBase  DBaseConfig  `yaml:"db"`
	// Api    ApiConfig    `yaml:"api"`
	Log LogConfig `yaml:"log"`
	Env EnvVar
}

func NewConfig() (*Config, error) {

	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("read config error: [%w]", err)
	}

	conf := Config{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: [%w]", err)
	}

	err = getEnvVar(&conf)
	if err != nil {
		return nil, fmt.Errorf("unable to parse env var: [%w]", err)
	}

	return &conf, nil
}

func getEnvVar(conf *Config) error {

	err := godotenv.Load("./config/key.env")
	if err != nil {
		return fmt.Errorf("error loading .env file: [%w]", err)
	}

	conf.Env.SessionName = os.Getenv("sessionName")
	conf.Env.CookieKey = os.Getenv("cookieKey")

	return nil
}
