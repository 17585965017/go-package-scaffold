package api

import (
	v1 "go_package_scaffold/router/api/v1"

	"github.com/gin-gonic/gin"
)

func LoaderRouterAPI(e *gin.Engine) {
	routerAPIGroup := e.Group("/api")
	{
		v1.LoadV1(routerAPIGroup)
	}
}
