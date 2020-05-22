package server

import (
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetList(c *gin.Context) {
	var response Response
	path := config.MusicPath
	log.Debug("current request path: " + path)
	fileList, err := GetAllFiles(path)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, GetErrorCode(-1, err.Error()))
		return
	}
	var file File
	var files []File
	for _, ele := range fileList {
		file.Name = ele
		files = append(files, file)
	}
	response.ErrorCode = E_OK
	response.ErrorMsg = "success"
	response.FileInfos = files
	c.JSON(http.StatusOK, response)
}

func GetFile(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		log.Error("invalid path")
		c.JSON(http.StatusBadRequest, GetErrorCode(1, "invalid path"))
		return
	}
	stat, err := os.Stat(path)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, GetErrorCode(1, err.Error()))
		return
	}
	if stat.IsDir() {
		log.Warning("invalid file path: " + path)
		c.JSON(http.StatusBadRequest, GetErrorCode(1, "invalid file path: "+path))
		return
	}
	c.Header("Content-Type", "audio/mp3; charset=utf-8")
	c.File(path)

}

func Upload(c *gin.Context) {
	err := c.Request.ParseMultipartForm(200000)
	if err != nil {
		log.Fatal(err)
	}
	// 获取表单
	form := c.Request.MultipartForm
	// 获取参数upload后面的多个文件名，存放到数组files里面，
	files := form.File["files"]

	// 遍历数组，每取出一个file就拷贝一次
	for i, _ := range files {
		fileName := files[i].Filename
		log.Info("get file name " + fileName)

		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Error(err.Error())
			c.String(http.StatusCreated, " failed to get "+fileName+"success! \n")
			continue
		}
		fileName = strings.ReplaceAll(fileName, " ", "")
		lastindex := strings.LastIndex(fileName, "-")
		savePath := ""
		if lastindex == -1 {
			log.Warning("failed to split file name to find author")
			savePath = config.MusicPath + "/" + "others" + "/"

		} else {
			author := fileName[:lastindex]
			savePath = config.MusicPath + "/" + author + "/"
		}
		// create dir if dir not exist
		_, err = os.Stat(savePath)
		if err != nil {
			_ = os.MkdirAll(savePath, os.ModePerm)
		}
		log.Debug("saving " + fileName + "to path " + savePath)
		savePath += fileName
		out, err := os.Create(savePath)
		defer out.Close()
		if err != nil {
			log.Error(err.Error())
			c.String(http.StatusCreated, " failed to create "+savePath+"\n")
			continue
		}

		_, err = io.Copy(out, file)
		if err != nil {
			log.Error(err.Error())
			c.String(http.StatusCreated, " failed to copy "+savePath+"\n")
			continue
		}
		log.Debug("saved " + savePath)

		c.String(http.StatusCreated, "upload "+fileName+" success! \n")
	}

}
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}
