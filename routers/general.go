package routers

import (
	"go-easy-frame/controllers/admin/v1_1"
	"go-easy-frame/middleware"
)

//
func generalRouter() {

	// 测试
	test := router.Group("/admin/")
	test.Use(middleware.LoginVerify)
	{
		test2 := test.Group("/v1.1/test")
		{
			test2.GET("/test", (&ctrAdminV1_1.Test{}).Test)
			test2.GET("/test2", (&ctrAdminV1_1.Test{}).Test2)
			test2.GET("/test3", (&ctrAdminV1_1.Test{}).Test3)
			test2.GET("/test4", (&ctrAdminV1_1.Test{}).Test4)
		}
		exam := test.Group("/v1.1/exam")
		{
			exam.POST("/db", (&ctrAdminV1_1.Exam{}).Db)
			exam.GET("/redis", (&ctrAdminV1_1.Exam{}).Redis)
		}
	}
	// 用户信息
	user := router.Group("/admin/v1.1/user").Use(middleware.LoginVerify)
	{
		user.GET("/info", (&ctrAdminV1_1.User{}).Info)
	}
}
