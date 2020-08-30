package error_msg

const (
	SUCCESS = 200
	ERROR = 500
	// user error
	ERROR_USER_NOT_EXIST = 1001		//用户不存在
	ERROR_USER_IS_EXIST = 1002		//用户已存在
	ERROR_PASSWORD_WRONG = 1003	//密码错误
	ERROR_REQ_PARAM_ERROR = 1004	//请求参数错误
	ERROR_UPLOAD_AVATAR_ERROR = 1005 //头像上传失败
	ERROR_USERNAME_IS_EXIST = 1006	//用户名已存在
	ERROR_TOKEN_NOT_EXIST = 2000	//未携带token
	ERROR_TOKEN_EXPIRED = 2001	//token过期
	ERROR_TOKEN_NOT_AUTH = 2002 //token未授权
	ERROR_TOKEN_FORMAT_ERROR = 2003	//token格式错误
	ERROR_TOKEN_INVAILD = 2004 //token无效
)

var errorMsg  = map[int]string{
	SUCCESS: "OK",
	ERROR: "FAIL",
	ERROR_USER_NOT_EXIST:"用户不存在",
	ERROR_USER_IS_EXIST: "用户已存在",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_UPLOAD_AVATAR_ERROR:"头像上传失败",
	ERROR_USERNAME_IS_EXIST: "用户名已存在",
	ERROR_TOKEN_NOT_EXIST: "请求未携带token",
	ERROR_TOKEN_EXPIRED: "token已过期",
	ERROR_TOKEN_NOT_AUTH: "token未激活",
	ERROR_TOKEN_FORMAT_ERROR: "token格式错误",
	ERROR_TOKEN_INVAILD: "token无效",
	ERROR_REQ_PARAM_ERROR: "请求参数错误",
}

func GetErrorMsg(code int) string {
	return errorMsg[code]
}





