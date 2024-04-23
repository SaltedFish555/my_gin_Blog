package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwyKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwyKey = file.Section("server").Key("JwyKey").MustString("89js82js72")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("blog_admin")
	DbPassword = file.Section("database").Key("DbPassword").MustString("178422")
	DbName = file.Section("database").Key("DbName").MustString("my_gin_blog")
}

func LoadQiniu(file *ini.File){
	AccessKey  =file.Section("qiniu").Key("AccessKey").String()
	SecretKey  =file.Section("qiniu").Key("SecretKey").String()
	Bucket     =file.Section("qiniu").Key("Bucket").String()
	QiniuServer=file.Section("qiniu").Key("QiniuServer").String()
}
