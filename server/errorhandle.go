package server

// GetErrorCode error code get
func GetErrorCode(code int, detail string) ErrorCode {
	return ErrorCode{
		ECode:  code,
		Detail: detail,
	}
}
