package v1

import (
	"go_package_scaffold/api/v1"

	"github.com/gin-gonic/gin"
)

func homeGroup(e *gin.RouterGroup) {
	userALLGroup := e.Group("/home")
	{
		userALLGroup.GET("/", api.GetHome)
	}
}
