package server

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Ropho/Pirater/config"
)

func NewLogger(conf *config.LogConfig) (*log.Logger, error) {

	logger := log.New()
	logger.SetFormatter(new(log.TextFormatter))
	level, err := log.ParseLevel(conf.Level)
	if err != nil {
		return nil, fmt.Errorf("unable to config logger : [%w]", err)
	}

	logger.SetLevel(level)

	return logger, nil
}
