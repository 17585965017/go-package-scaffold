package router

import (
	"go_package_scaffold/api/v1"
	"go_package_scaffold/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Session("setOnProduction"))
	r.Use(middleware.Cors())

	// 路由配置
	v1 := r.Group("/api/v1")
	{
		v1.GET("/home", api.GetHome)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)
	}

	return r
}
