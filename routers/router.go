package routers

import (
	"go-easy-frame/config"
	"go-easy-frame/middleware"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.Recover) // 自定义全局的异常检测

	cfg := config.GetEnv()
	if "debug" == cfg.GinMode || "test" == cfg.GinMode {
		pprof.Register(router) // 性能监控，用于开发测试环境
	}
}

func Get() *gin.Engine {
	// 登陆路由
	loginRouter()

	// 普通路由
	generalRouter()

	// 其他路由 ...这里引入调用函数

	return router
}
