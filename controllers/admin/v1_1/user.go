package ctrAdminV1_1

import (
	"go-easy-frame/controllers"

	"github.com/gin-gonic/gin"
)

type User struct {
	ctr.BaseController
}

func (this *User) Info(ctx *gin.Context) {
	v, exts := ctx.Get("cur_user")
	if !exts {
		this.Failed(400, "用户不存在")
	}

	this.Success(ctx, v)
}
