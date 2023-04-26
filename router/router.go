package router

import (
	"go_package_scaffold/middleware"
	routerAPI "go_package_scaffold/router/api"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.ClienttimeMiddleware())
	r.Use(middleware.TimeoutMiddleware())
	r.Use(middleware.Session("setOnProduction"))
	r.Use(middleware.Cors())

	routerAPI.LoaderRouterAPI(r) // 加载 api 接口路由

	return r
}
