package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/gin-toy-layout/config"
	"github.com/mittacy/gin-toy-layout/utils/bizUtil"
	"github.com/mittacy/gin-toy/core/singleton"
)

// InitDependents 初始化依赖资源
// @param path 配置文件路径
// @param port 端口
// @param env 运行环境
func InitDependents(path string, env string) {
	// conf
	ParseConfig("env", path)

	// gin run env
	gin.SetMode(bizUtil.AppEnvToGinEnv(env))

	// log
	InitLog()

	// database
	config.InitMysql()
	config.InitRedis()
	config.InitMongo()

	// 初始化单例
	singleton.Build()
}
