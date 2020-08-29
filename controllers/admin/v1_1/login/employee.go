package ctrAdminV1_1Login

import (
	"go-easy-frame/controllers"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ctr.BaseController
}

func (this *Employee) Verify(ctx *gin.Context) {
	// 获取并验证参数

	// 调用逻辑处理 ...

	this.Success(ctx, map[string]interface{}{
		"verify": "123456",
	})
}

func (this *Employee) SendSms(ctx *gin.Context) {
	// 获取并验证参数

	// 调用逻辑处理 ...

	this.Success(ctx, map[string]interface{}{})
}

func (this *Employee) In(ctx *gin.Context) {
	// 获取并验证参数
	req := new(ReqEmployeeIn).GetReq(ctx)

	// 调用逻辑处理 ...

	// 返回
	this.Success(ctx, map[string]interface{}{
		"token": "***",
		"info":  map[string]interface{}{},
		"req":   req,
	})
}

func (this *Employee) Out(ctx *gin.Context) {
	// 获取并验证参数

	// 调用逻辑处理 ...

	this.Success(ctx, map[string]interface{}{})
}
