package v1

import (
	"my_gin_blog/model"
	"my_gin_blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加文章
func AddArticle(c *gin.Context){
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code:=model.CreateArticle(&data)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

// 查询分类下的所有文章
func GetCateArt(c *gin.Context){
	pageSize,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum,_:=strconv.Atoi(c.Query("pagenum"))
	cid,_:=strconv.Atoi(c.Param("cid"))
	data,code,total:=model.GetCateArt(cid,pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
	

}


// 查询某用户下的所有文章
func GetUserArt(c *gin.Context){
	pageSize,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum,_:=strconv.Atoi(c.Query("pagenum"))
	uid,_:=strconv.Atoi(c.Param("uid"))
	data,code,total:=model.GetUserArt(uid,pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
	

}



// 查询单个文章信息
func GetArtInfo(c *gin.Context){
	id,_:=strconv.Atoi(c.Param("id"))
	data,code:=model.GetArtInfo(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}


// 查询文章列表
func GetArticles(c *gin.Context){
	pageSize,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum,_:=strconv.Atoi(c.Query("pagenum"))
	data,code,total:=model.GetArticles(pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),

	})
}

// 编辑文章
func EditArticle(c *gin.Context){
	var data model.Article
	c.ShouldBindJSON(&data)
	id,_:=strconv.Atoi(c.Param("id"))

	code:=model.EditArticle(id,&data)


	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})

}

// 删除分类
func DeleteArticle(c *gin.Context){
	id,_:=strconv.Atoi(c.Param("id"))
	code:=model.DeleteArticle(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}





// 查询推荐文章
func GetRecomArt(c *gin.Context){

	data,code,total:=model.GetRecomArt()
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),

	})
}







