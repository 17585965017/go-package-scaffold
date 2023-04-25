package conf

import (
	"go_package_scaffold/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Init 初始化配置项
func Init() {

	// 设置日志级别
	util.BuildLogger("debug")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
