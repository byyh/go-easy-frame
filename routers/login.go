package routers

import (
	"go-easy-frame/controllers/admin/v1_1/login"
)

// 登陆相关
func loginRouter() {

	login := router.Group("/admin/v1.1/login")
	{
		login.GET("/verify", (&ctrAdminV1_1Login.Employee{}).Verify)
		login.POST("/send/sms", (&ctrAdminV1_1Login.Employee{}).SendSms)
		login.POST("/in", (&ctrAdminV1_1Login.Employee{}).In)
		login.POST("/out", (&ctrAdminV1_1Login.Employee{}).Out)
	}

}
