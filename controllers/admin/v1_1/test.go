package ctrAdminV1_1

import (
	"go-easy-frame/controllers"
	"go-easy-frame/services/atom"
	"log"

	"github.com/gin-gonic/gin"
)

type Test struct {
	ctr.BaseController
}

// 正常返回
func (this *Test) Test(ctx *gin.Context) {
	v, exts := ctx.Get("cur_user")
	if !exts {
		this.Failed(401, "用户不存在")
		return
	}

	var in ReqUserTest
	if err := ctx.ShouldBind(&in); err == nil {
		log.Println(err)
		this.Failed(401, "the body should be formA")
		return
	}

	this.Success(ctx, map[string]interface{}{
		"cur_user": v,
		"input":    in,
	})
}

// 测试系统级bug异常
func (this *Test) Test2(ctx *gin.Context) {
	var a = []int{1, 2, 3, 4, 5}
	a[6] = 5

	v, exts := ctx.Get("cur_user")
	if !exts {
		this.Failed(401, "用户不存在")
		return
	}

	var in ReqUserTest
	if err := ctx.ShouldBind(&in); err == nil {
		log.Println(err)
		this.Failed(401, "the body should be formA")
		return
	}

	this.Success(ctx, map[string]interface{}{
		"cur_user": v,
		"input":    in,
	})
}

// 测试返回抛出人工异常
func (this *Test) Test3(ctx *gin.Context) {
	panic(atom.NewMyError("测试返回抛出人工异常", 5000))

	v, exts := ctx.Get("cur_user")
	if !exts {
		this.Failed(4001, "用户不存在")
		return
	}

	var in ReqUserTest
	if err := ctx.ShouldBind(&in); err == nil {
		log.Println(err)
		this.Failed(401, "the body should be formA")
		return
	}

	this.Success(ctx, map[string]interface{}{
		"cur_user": v,
		"input":    in,
	})
}

// 测试返回非200http码
func (this *Test) Test4(ctx *gin.Context) {
	atom.OutHttpError(405, "测试返回非200http码")

}
