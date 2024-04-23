package v1

import (
	"my_gin_blog/model"
	"my_gin_blog/utils/errmsg"
	"my_gin_blog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询用户是否存在
func UserExist(c *gin.Context) {

}

var code int

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var msg string
	msg, code = validator.Validate(&data)
	// 数据验证失败，可能是账号或密码的长度不符合要求，给前端返回错误信息
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	// 数据验证成功，查询用户名是否已被占用，如果没有就创建新用户
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	// 给前端返回创建结果信息
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	data, total := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户信息
func EditUser(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	} else if code == errmsg.ERROR_USERNAME_USED {
		c.Abort() // 不再执行当前请求的后续处理程序
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 通过username获取uid
func GetUID(c *gin.Context) {
	username:= c.Param("username")
	uid, code := model.GetUID(username)
	// 给前端返回创建结果信息
	c.JSON(http.StatusOK, gin.H{
		"uid":     uid,
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
