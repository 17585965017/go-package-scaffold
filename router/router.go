package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 路由配置
	v1 := r.Group("/api/v1")
	{
		v1.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	return r
}
