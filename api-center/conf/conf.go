package conf

import (
	"api/cache"
	"api/model"
	"api/util"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	MYSQL_DSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
		HOST, DBNAME)

	model.Database(MYSQL_DSN)
	cache.Redis()
}
