package ctr

import (
	"go-easy-frame/services/atom"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

// 返回成功
func (this *BaseController) Success(ctx *gin.Context, res interface{}) {
	ctx.JSON(200, gin.H{
		"code":    atom.Success,
		"message": atom.GetRetMsgByCode(atom.Success),
		"data":    res,
	})
}

// 返回失败
func (this *BaseController) Failed(code int, msg string) {
	panic(atom.NewMyError(msg, code))
}

// 返回非200HTTP状态码
func (this *BaseController) OutHttpError(httpCode int, msg string) {
	panic(atom.NewMyErrorByHttpStatus(msg, httpCode))
}
