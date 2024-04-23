package main

import (
	"fmt"
	"my_gin_blog/model"
	"my_gin_blog/routers"
)

func main(){
	fmt.Scan()
	fmt.Println("博客后端开始运行")

	// 引用数据库
	model.InitDb()
	routers.InitRouter()
}

