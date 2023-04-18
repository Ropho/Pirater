package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

const (
	ConfigDir = "./config/"
	EnvFile   = "key.env"
)

type ServerConfig struct {
	Addr      string `yaml:"addr"`
	Port      int    `yaml:"port"`
	ProxyAddr string `yaml:"proxy_addr"`
	ProxyPort int    `yaml:"proxy_port"`
}

type SqlDBConfig struct {
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
}

type CookieDBConfig struct {
	Pass    string `yaml:"pass"`
	Addr    string `yaml:"addr"`
	Port    int    `yaml:"port"`
	MaxConn int    `yaml:"max_connections"`
}

type EnvVar struct {
	SessionName string `yaml:"session_name"`
	CookieKey   string `yaml:"cookie_key"`
}

type LogConfig struct {
	Level string `yaml:"level"`
}

type Config struct {
	Server      ServerConfig   `yaml:"server"`
	SqlDBase    SqlDBConfig    `yaml:"sql_db"`
	CookieDBase CookieDBConfig `yaml:"cookie_db"`
	Log         LogConfig      `yaml:"log"`
	Env         EnvVar
}

func NewConfig(configFile string) (*Config, error) {

	data, err := os.ReadFile(ConfigDir + configFile)
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

	err := godotenv.Load(ConfigDir + EnvFile)
	if err != nil {
		return fmt.Errorf("error loading .env file: [%w]", err)
	}

	conf.Env.SessionName = os.Getenv("sessionName")
	conf.Env.CookieKey = os.Getenv("cookieKey")

	return nil
}
