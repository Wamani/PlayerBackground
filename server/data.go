package server

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var config Config

// ErrorCode common error code
type ErrorCode struct {
	ECode    int    `json:"error_code"`
	Detail   string `json:"error_msg"`
	Internal string `json:"internal_detail,omitempty"` // internal detail info from other services
}
type Config struct {
	MusicPath string
}

type File struct {
	Name   string `json:"name"`
	Author string `json:"name"`
}
type Response struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	FileInfos []File `json:"file_infos"`
}
