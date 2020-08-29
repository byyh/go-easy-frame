package ctr

import (
	"github.com/byyh/go-easy-frame/services/atom"

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
func (this *BaseController) Failed(ctx *gin.Context, code int, msg string) {
	ctx.JSON(200, gin.H{
		"code":    code,
		"message": msg,
		"data":    nil,
	})
}

// 返回非200
func (this *BaseController) OutHttpError(ctx *gin.Context, httpCode int, msg string) {
	ctx.JSON(httpCode, gin.H{
		"message": msg,
	})
}
