package main

import (
	"go_package_scaffold/conf"
	"go_package_scaffold/router"
	"os"
)

func main() {

	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := router.NewRouter()
	r.Run(os.Getenv("RUN_PORT"))
}
