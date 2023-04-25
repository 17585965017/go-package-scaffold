package router

import (
	"go_package_scaffold/api/v1"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 路由配置
	v1 := r.Group("/api/v1")
	{
		v1.GET("/home", api.GetHome)
	}

	return r
}
