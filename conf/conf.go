package conf

import (
	"go_package_scaffold/util"
	"os"
)

// Init 初始化配置项
func Init() {

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))
}
