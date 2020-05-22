package mongodb

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

var globalS *mgo.Session
var log *logrus.Logger
var mgoConf MongoDBConf

// config structure
type MongoDBConf struct {
	Url            string
	Name           string
	TimeOut        int
	PoolLimit      int
	SessionTimeout int
}
