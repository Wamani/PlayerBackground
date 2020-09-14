// yanxin.wang 2019.05.27
// 添加错误码

package data

// Success name success error code
type Success int

const (
	E_OK = 0 //OK
	E_INVALID_PARAMS = 1
	// user info
	E_USER_NOT_EXIST = 1001
	E_USER_EXIST = 1002
)
