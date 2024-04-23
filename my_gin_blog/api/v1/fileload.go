package v1

import (
	"fmt"
	"my_gin_blog/fileload"
	"my_gin_blog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)
func UpLoad(c *gin.Context){
	file,fileHeader,_:=c.Request.FormFile("file")
	fileName:=fileHeader.Filename
	fileSize:=fileHeader.Size
	fmt.Println("fileName:",fileName)
	downloadUrl,code:=fileload.UploadFile(file,fileName,fileSize)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"download_url":downloadUrl,		
	})
}

func Download(c *gin.Context){
	fileName:=c.Param("filename")
	dir:="./"
	code:=fileload.DownloadFile(fileName,dir)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"filename":fileName,		
	})

}

