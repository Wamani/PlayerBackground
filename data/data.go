package data


// ErrorCode common error code
type ErrorCode struct {
	ECode    int    `json:"error_code"`
	Detail   string `json:"error_msg"`
	Internal string `json:"internal_detail,omitempty"` // internal detail info from other services
}
type Config struct {
	MusicPath string
	MysqlUrl string
}

type File struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}
type Response struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	FileInfos []File `json:"file_infos"`
}

