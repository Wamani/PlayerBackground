package server

import (
	"github.com/sirupsen/logrus"
)

func Init(logger *logrus.Logger, conf Config) error {
	log = logger
	config = conf
	return nil
}
