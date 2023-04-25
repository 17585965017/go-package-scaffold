package v1

import (
	"go_package_scaffold/api/v1"

	"github.com/gin-gonic/gin"
)

func userGroup(e *gin.RouterGroup) {
	userALLGroup := e.Group("/user")
	{
		userALLGroup.POST("/register", api.UserRegister)
		userALLGroup.POST("/login", api.UserLogin)
	}
}
