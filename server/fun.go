package server

import (
	"../data"
	"../db/mysql"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var config data.Config
var client mysql.Client

func Init(logger *logrus.Logger, conf data.Config) error {
	log = logger
	config = conf
	client.Init(log, conf.MysqlUrl)
	client.GetClient().SingularTable(true)
	return nil
}
