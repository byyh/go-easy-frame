package routers

import (
	"github.com/byyh/go-easy-frame/controllers/admin/v1_1/login"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()

}

func Get() *gin.Engine {
	// 登陆相关
	login := router.Group("/admin/v1.1/login")
	{
		login.GET("/verify", (&ctrAdminV1_1Login.Employee{}).Verify)
		login.POST("/send/sms", (&ctrAdminV1_1Login.Employee{}).SendSms)
		login.POST("/in", (&ctrAdminV1_1Login.Employee{}).In)
		login.POST("/out", (&ctrAdminV1_1Login.Employee{}).Out)
	}

	return router
}
