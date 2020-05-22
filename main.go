/**
* @brief	 main code for broker
* @author    yanxin.wang
* @email     w_aman@qq.com
* @version   1.0
* @date      2019-05-27
 */
package main

import (
	"./conf"
	"./server"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
)

var log = logrus.New()

var maxRotateSize int32 //64M default

type logFileWriter struct {
	file *os.File
	//write count
	size int64
}

var logFileNumber int

func (p *logFileWriter) Write(data []byte) (n int, err error) {

	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.file == nil {
		return 0, errors.New("file not opened")
	}
	n, e := p.file.Write(data)
	p.size += int64(n)
	//文件最大 maxRotateSize M
	if p.size > 1024*1024*int64(maxRotateSize) {
		logFileNumber++
		logFileNumber = logFileNumber % 10
		if err := p.file.Close(); err != nil {
			fmt.Println("log file close error", err.Error())
		}
		fmt.Println("log file is full")
		err := os.Rename("./logs/broker.log", "./logs/"+strconv.Itoa(
			logFileNumber)+".log")
		if err != nil {
			fmt.Println(err.Error())
		}
		p.file, _ = os.OpenFile("./logs/broker"+".log",
			os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0755)
		p.size = 0
	}
	return n, e
}

func init() {
	logFileNumber = 0
	//load config file
	err := configor.New(&configor.Config{Debug: false}).Load(
		&conf.BrokerConfig, "./conf/config.json")
	if err != nil {
		log.Fatalln("load config file error", err.Error())
	}
	//config logger
	_, err = os.Stat("./logs")
	if err != nil {
		if ok := os.IsExist(err); !ok {
			err = os.Mkdir("./logs", os.ModePerm)
			if err != nil {
				log.Warnln(err.Error())
			}
		}
	}
	file, err := os.OpenFile("./logs/broker"+".log",
		os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0755)
	if err != nil {
		log.Errorln("log init failed ", err.Error())
	}
	if file == nil {
		log.Errorln("file is nil")
		return
	}
	info, err := file.Stat()
	if err != nil {
		log.Fatal(err.Error())
	}
	maxRotateSize = conf.BrokerConfig.Logger.MaxRotateSize
	fileWriter := logFileWriter{file, info.Size()}
	//set log level and style
	loglevel, err := logrus.ParseLevel(conf.BrokerConfig.Logger.LogLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			//repoPath := fmt.Sprintf("%s", os.Getenv("GOPATH"))
			fileName := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d",
				fileName, f.Line)
		},
	})
	log.SetLevel(loglevel)
	log.SetOutput(io.MultiWriter(&fileWriter, os.Stdout))
	// init http server
	err = server.Init(log, conf.BrokerConfig.Server)
	if err != nil {
		log.Fatalln("server", err.Error())
		return
	}
	log.Infoln("server init success")
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(ginlogrus.Logger(log), gin.Recovery())
	// 音乐相关
	router.GET("/", func(context *gin.Context) {
		context.Header("Content-Type", "text/html")
		context.File("./upload.html")
	})
	router.GET("install", func(context *gin.Context) {
		name := context.Query("name")
		context.Header("Content-Type", "application/vnd.android.package-archive")
		if name == "nana" {
			context.FileAttachment("./app-release.apk", "app-release.apk")
		}
		context.FileAttachment("./player.apk", "player.apk")
	})
	router.GET("list", server.GetList)
	router.GET("file", server.GetFile)
	router.POST("upload", server.Upload)
	address := "0.0.0.0" + ":" + conf.BrokerConfig.PORT
	log.Infoln("Listening and serving HTTP on", address)
	err := router.Run(address)
	if err != nil {
		log.Fatalln("error", err.Error())
	}
}
