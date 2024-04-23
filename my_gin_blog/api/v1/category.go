package v1

import (
	"my_gin_blog/model"
	"my_gin_blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



// 查询分类是否存在
func CategoryExist(c *gin.Context){

}

// 添加分类
func AddCategory(c *gin.Context){
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code:=model.CheckCategory(data.Name)
	if code ==errmsg.SUCCESS{
		model.CreateCategory(&data)
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

// 查询分类下的所有文章


// 查询分类列表
func GetCategories(c *gin.Context){
	pageSize,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum,_:=strconv.Atoi(c.Query("pagenum"))
	data,total:=model.GetCategories(pageSize,pageNum)
	code:=errmsg.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}

// 编辑分类名
func EditCategory(c *gin.Context){
	var data model.Category
	c.ShouldBindJSON(&data)
	id,_:=strconv.Atoi(c.Param("id"))
	code:=model.CheckCategory(data.Name)
	if code==errmsg.SUCCESS{
		model.EditCategory(id,&data)
	}else if code==errmsg.ERROR_CATENAME_USED{
		c.Abort() // 不再执行当前请求的后续处理程序
	}


	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})

}

// 删除分类
func DeleteCategory(c *gin.Context){
	id,_:=strconv.Atoi(c.Param("id"))
	code:=model.DeleteCategory(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}











