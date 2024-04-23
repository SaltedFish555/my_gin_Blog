package model

import (
	"fmt"
	"my_gin_blog/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb(){
	dsn:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db,err=gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err!=nil{
		fmt.Println("连接数据库失败，请检查参数",err)
	}

	// 迁移表，无表时候会创建表，顺序不能改，否则会因为外键约束无法创建
	db.AutoMigrate(&User{},&Category{},&Article{})

	sqlDB,err:=db.DB()
	if err!=nil{
		fmt.Println("获取底层数据库连接失败",err)
	}
	
	sqlDB.SetMaxIdleConns(10) //最大闲置连接数
	sqlDB.SetMaxOpenConns(100) //最大连接数
	sqlDB.SetConnMaxLifetime(10*time.Second) //连接的最大可复用时间
	
	// sqlDB.Close()


}