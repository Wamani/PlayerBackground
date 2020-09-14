package data

type UserInfo struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Avatar string `json:"avatar"`
	Description string `json:"description"`
	UUID string `json:"uuid"`
}
type UserInfoRes struct {
	Data UserInfo `json:"data"`
	ErrorCode int `json:"error_code"`
	ErrorMsg string `json:"error_msg"`
}
type UserRegisterResData struct {
	ID string `json:"id"`
}
type UserRegisterRes struct {
	Data UserRegisterResData `json:"data"`
	ErrorCode int `json:"error_code"`
	ErrorMsg string `json:"error_msg"`
}
type UserLoginResData struct {
	User string `json:"user"`
	Session string `json:"session"`
}
type UserLoginRes struct {
	Data UserLoginResData `json:"data"`
}