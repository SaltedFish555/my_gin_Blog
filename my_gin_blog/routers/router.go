package routers

import (
	v1 "my_gin_blog/api/v1"
	"my_gin_blog/middleware"
	"my_gin_blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(){
	gin.SetMode(utils.AppMode)
	r:=gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	router_auth:=r.Group("api/v1/") // 需要鉴权的接口
	router_auth.Use(middleware.JwtToken())
	{
		// User模块的路由接口
		router_auth.PUT("user/:id",v1.EditUser)
		router_auth.DELETE("user/:id",v1.DeleteUser)
		
		// 分类模块的路由接口
		router_auth.POST("category/add",v1.AddCategory)
		router_auth.PUT("category/:id",v1.EditCategory)
		router_auth.DELETE("category/:id",v1.DeleteCategory)

		// 文章模块的路由接口
		router_auth.POST("article/add",v1.AddArticle)
		router_auth.PUT("article/:id",v1.EditArticle)
		router_auth.DELETE("article/:id",v1.DeleteArticle)
		

		//上传文件
		router_auth.POST("upload",v1.UpLoad)

	}

	router_public:=r.Group("api/v1/") // 公共接口
	{
		router_public.POST("user/add",v1.AddUser)
		router_public.GET("users",v1.GetUsers)
		router_public.GET("user/id/:username",v1.GetUID)

		router_public.GET("categories",v1.GetCategories)
		
		router_public.GET("articles",v1.GetArticles)
		router_public.GET("article/clist/:cid",v1.GetCateArt)
		router_public.GET("article/ulist/:uid",v1.GetUserArt)
		router_public.GET("article/recom",v1.GetRecomArt)


		router_public.GET("article/info/:id",v1.GetArtInfo)
		router_public.POST("login",v1.Login)

		router_public.GET("file/:filename",v1.Download)

	}


		
		r.Run(utils.HttpPort)
}