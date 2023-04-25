package conf

import (
	"go_package_scaffold/util"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {

	// 从本地读取环境变量
	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env.example file")
	}

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))
}
