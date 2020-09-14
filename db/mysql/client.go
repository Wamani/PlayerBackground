package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Client struct {
	url 		string
	logger 	 	*logrus.Logger
	client      *gorm.DB
}

func (client *Client)Init (logger *logrus.Logger, url string) {
	client.logger = logger
	client.url = url
}

func (client *Client)GetClient()  *gorm.DB {
	if client.client == nil {
		db, err := gorm.Open("mysql", client.url)
		if err != nil{
			client.logger.Panicln("failed to connect to mysql" + err.Error())
		}
		if db == nil {
			client.logger.Panicln("failed to connect to mysql")
		}
		client.client = db
	}
	return client.client
}

