package v1

import (
	"my_gin_blog/middleware"
	"my_gin_blog/model"
	"my_gin_blog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	var data model.User
	c.ShouldBindJSON(&data)
	var token string
	code:=model.CheckLogin(data.Username,data.Password)
	if code==errmsg.SUCCESS{
		token,_=middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"token":token,
	})
}

