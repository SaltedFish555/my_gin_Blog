package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowOrigins:  []string{"*"}, //允许所有域名访问
			AllowMethods:  []string{"*"}, //允许所有方法
			AllowHeaders:  []string{"Origin","Authorization"},
			ExposeHeaders: []string{"Content-length"},
			MaxAge:        12 * time.Hour,
		},
	)
}


