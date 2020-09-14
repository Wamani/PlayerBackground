package mongodb

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func Init(mongoConf MongoDBConf, logger *logrus.Logger) error {
	log = logger
	log.Infoln("connecting to mongodb")
	mgoConf = mongoConf
	dialInfo, err := mgo.ParseURL(mongoConf.Url)
	if err != nil {
		return err
	}
	dialInfo.Source = mongoConf.Name
	dialInfo.PoolLimit = mongoConf.PoolLimit
	dialInfo.Timeout = time.Duration(mgoConf.TimeOut) * time.Second
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Errorf("create session: %s\n", err)
		return err
	}
	if s == nil {
		log.Errorln("failed to connect to mongodb server")
		return errors.New("failed to connect to mongodb server")
	}
	log.Infoln("mongodb config : ", mgoConf)
	s.SetSocketTimeout(time.Duration(mongoConf.SessionTimeout) * time.Second)
	globalS = s
	log.Infoln("connected to mongodb")
	return nil

}

func Connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	ms := globalS.Copy()
	c := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c
}

func FindOne(collection string, query, selector, result interface{}) error {
	ms, c := Connect(mgoConf.Name, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).One(result)
}

func Find(collection string, query, selector, result interface{}) error {
	ms, c := Connect(mgoConf.Name, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).All(result)
}
func FindIDs(collection string, ids []string, result interface{}) error {
	ms, c := Connect(mgoConf.Name, collection)
	defer ms.Close()

	return c.Find(bson.M{"id": bson.M{"$in": ids}}).All(result)
}
func IsEmpty(collection string) (bool, error) {
	ms, c := Connect(mgoConf.Name, collection)
	defer ms.Close()
	count, err := c.Count()

	return count == 0, err
}

func Update(collection string, selector, update interface{}) error {
	ms, c := Connect(mgoConf.Name, collection)
	defer ms.Close()

	return c.Update(selector, update)
}

func InsertElems(collection string, elems []interface{}) error {
	ms, c := Connect(mgoConf.Name, collection)
	defer ms.Close()

	return c.Insert(elems...)
}

func DeleteElems(collection string, selector []string) error {
	ms, c := Connect(mgoConf.Name, collection)
	defer ms.Close()

	_, err := c.RemoveAll(bson.M{"id": bson.M{"$in": selector}})
	return err
	//return c.Remove(bson.M{"dev_id":bson.M{"$in":devIds,
	//	"dev_sup_id":bson.M{"$in":devSupIds},"ch":bson.M{"$in":ch},
	//	"area_type":bson.M{"$in":areaType}}})
}
