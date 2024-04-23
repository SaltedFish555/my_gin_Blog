package fileload

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"my_gin_blog/utils"
	"my_gin_blog/utils/errmsg"
	"net/http"
	"os"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var AccessKey=utils.AccessKey
var SecretKey=utils.SecretKey
var Bucket=utils.Bucket
var QiniuUrl=utils.QiniuServer

func UploadFile(file multipart.File,fileName string,fileSize int64)(downloadUrl string,code int){
	putPolicy:=storage.PutPolicy{
		Scope:Bucket,
	}
	mac:=qbox.NewMac(AccessKey,SecretKey)
	upToken:=putPolicy.UploadToken(mac)
	region,_:=storage.GetRegion(AccessKey,Bucket)
	cfg:=storage.Config{
		Region:region,
		UseCdnDomains:false,
		UseHTTPS:false,
	}
	putExtra:=storage.PutExtra{}
	formUploader:=storage.NewFormUploader(&cfg)
	key:=fileName //文件名
	ret:=storage.PutRet{}
	err:=formUploader.Put(context.Background(),&ret,upToken,key,file,fileSize,&putExtra)
	if err!=nil{
		fmt.Println("上传错误",err)
		return "",errmsg.ERROR
	}
	downloadUrl=QiniuUrl+"/"+ret.Key

	return downloadUrl,errmsg.SUCCESS
	



}


func DownloadFile(fileName string,dir string) (code int){
	downloadUrl:=QiniuUrl+"/"+fileName
	response,err:=http.Get(downloadUrl)
	if err!=nil{
		fmt.Println("获取下载链接响应错误",err)
		return errmsg.ERROR_FILE_NOT_EXIST
	}
	defer response.Body.Close()
	path:=dir+"/"+fileName
	file,err:=os.Create(path)
	if err!=nil{
		fmt.Println("创建目标文件错误",err)
		return errmsg.ERROR_FILE
	}
	defer file.Close()

	_,err=io.Copy(file,response.Body)
	if err!=nil{
		fmt.Println("将文件内容从响应体复制到本地错误",err)
		return errmsg.ERROR_FILE
	}
	return errmsg.SUCCESS
}



















