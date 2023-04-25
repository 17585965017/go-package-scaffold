package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping 状态检查页面
func GetHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "这是主页",
	})
}
