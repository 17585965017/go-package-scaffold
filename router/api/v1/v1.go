package v1

import (
	"github.com/gin-gonic/gin"
)

func LoadV1(e *gin.RouterGroup) {
	routerV1Group := e.Group("/v1")
	{
		userGroup(routerV1Group)
		homeGroup(routerV1Group)
	}
}
