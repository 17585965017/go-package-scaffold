package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHome 首页的测试接口
func GetHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "这是主页",
	})
}
