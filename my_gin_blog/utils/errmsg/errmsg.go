package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
)

// code = 1000... 用户模块的错误
const (
	ERROR_USERNAME_USED    = 1001 + iota // 用户名已被占用
	ERROR_PASSWORD_WRONG   = 1001 + iota
	ERROR_USER_NOT_EXIST   = 1001 + iota
	ERROR_TOKEN_NOT_EXIST  = 1001 + iota
	ERROR_TOKEN_RUNTIME    = 1001 + iota
	ERROR_TOKEN_WRONG      = 1001 + iota // token错误
	ERROR_TOKEN_TYPE_WRONG = 1001 + iota
	ERROR_USER_NO_RIGHT    = 1001 + iota
)

// code = 2000... 文章模块的错误
const (
	ERROR_ART_NOT_EXIST = 2001 + iota
)

// code = 3000... 分类模块的错误
const (
	ERROR_CATENAME_USED  = 3001 + iota
	ERROR_CATE_NOT_EXIST = 3001 + iota
)

// code = 4000... 文件上传和下载模块的错误
const (
	ERROR_FILE_NOT_EXIST = 4001 + iota
	ERROR_FILE           = 4001 + iota
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已被占用",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "TOEKN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_CATENAME_USED:    "该分类名已被占用",
	ERROR_ART_NOT_EXIST:    "文章不存在",
	ERROR_CATE_NOT_EXIST:   "该分类不存在",
	ERROR_USER_NO_RIGHT:    "该用户无管理权限",
	ERROR_FILE_NOT_EXIST:   "该文件不存在",
	ERROR_FILE:             "下载文件错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]

}
