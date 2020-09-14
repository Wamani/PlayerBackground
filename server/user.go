package server

import (
	"../data"
	"encoding/json"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func Register(c *gin.Context) {
	log.Info("register ")
	userInfo := data.UserInfo{}
	err := c.BindJSON(&userInfo)
	if err != nil {
		log.Error("get register info failed" + err.Error())
		c.JSON(http.StatusBadRequest, GetErrorCode(data.E_INVALID_PARAMS, "invalid params"))
		return
	}
	if userInfo.Email == "" || userInfo.Password == "" {
		log.Warnln("empty email or password")
		c.JSON(http.StatusBadRequest, GetErrorCode(data.E_INVALID_PARAMS, "invalid user info"))
		return
	}
	// convert param to string for log
	paramString, err := json.Marshal(userInfo)
	if err != nil {
		log.Warnln(err.Error())
	}
	log.Debugln("request info", string(paramString))
	var userInfos []data.UserInfo
	client.GetClient().Where("email = ?", userInfo.Email).Table("user_info").First(&userInfos)
	if len(userInfos) > 0 {
		c.JSON(http.StatusOK, GetErrorCode(data.E_USER_EXIST, "user already exist"))
		return
	}
	uuidV4, _ := uuid.NewV4()
	userInfo.UUID = uuidV4.String()
	client.GetClient().Create(&userInfo)

	userRegisterResData := data.UserRegisterResData{}
	userRegisterResData.ID = userInfo.UUID
	userRegisterRes := data.UserRegisterRes{}
	userRegisterRes.ErrorMsg = "success"
	userRegisterRes.ErrorCode = data.E_OK
	userRegisterRes.Data = userRegisterResData
	// log response info
	resString, err := json.Marshal(userRegisterRes)
	if err != nil {
		log.Warnln(err.Error())
	}

	log.Debugln("response", string(resString))
	c.JSON(http.StatusOK, userRegisterRes)
}

func Login(c *gin.Context) {
	userLoginResData := data.UserLoginResData{}
	userLoginResData.Session = "12345667"
	userLoginResData.User = "xiaoxin"
	userLoginRes := data.UserLoginRes{}
	userLoginRes.Data = userLoginResData
	c.JSON(http.StatusOK, userLoginRes)
}

func GetUserInfo(c *gin.Context) {
	userId, err := c.Params.Get("id")
	if !err {
		c.JSON(http.StatusBadRequest, GetErrorCode(-1, "MISSING USER ID"))
		return
	}
	log.Info(userId)
	var userInfos []data.UserInfo
	client.GetClient().Where("id = ?", userId).Table("user_info").First(&userInfos)
	userInfoRes := data.UserInfoRes{}
	if len(userInfos) == 0 {
		c.JSON(http.StatusOK, GetErrorCode(data.E_USER_NOT_EXIST, "user not exist"))
		return
	}
	userInfoRes.ErrorCode = data.E_OK
	userInfoRes.ErrorMsg = "success"
	userInfoRes.Data = userInfos[0]
	c.JSON(http.StatusOK, userInfoRes)
}
