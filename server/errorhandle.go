package server

import "../data"

// GetErrorCode error code get
func GetErrorCode(code int, detail string) data.ErrorCode {
	return data.ErrorCode{
		ECode:  code,
		Detail: detail,
	}
}
