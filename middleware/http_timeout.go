package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func testResponse(c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, gin.H{
		"code": http.StatusGatewayTimeout,
		"msg":  "timeout",
	})
}

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(10*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}
