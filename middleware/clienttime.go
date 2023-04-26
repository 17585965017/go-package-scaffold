package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ClienttimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestHeaderClienttime := c.GetHeader("Clienttime")
		requestParamsClienttime := c.Query("clienttime")

		if len(requestHeaderClienttime) == 0 || len(requestParamsClienttime) == 0 {
			c.JSON(http.StatusGatewayTimeout, gin.H{
				"code": http.StatusGatewayTimeout,
				"msg":  "clienttime 缺失",
			})
			c.Abort()
			return
		}

		if requestHeaderClienttime != requestParamsClienttime {
			c.JSON(http.StatusGatewayTimeout, gin.H{
				"code": http.StatusGatewayTimeout,
				"msg":  "时间不一致",
			})
			c.Abort()
			return
		}

		if clientRequestTime, err := strconv.ParseInt(requestHeaderClienttime, 10, 64); err != nil {
			c.JSON(http.StatusGatewayTimeout, gin.H{
				"code": http.StatusGatewayTimeout,
				"msg":  "时间格式错误",
			})
			c.Abort()
			return
		} else {
			// 当前时间戳-客户端时间戳大于 10s
			if time.Now().Unix()-clientRequestTime > 10 {
				c.JSON(http.StatusGatewayTimeout, gin.H{
					"code": http.StatusGatewayTimeout,
					"msg":  "请求超时",
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
