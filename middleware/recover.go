package middleware

// 捕获全局异常

import (
	"go-easy-frame/services/atom"
	"net/http"
	"reflect"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if "*atom.MyError" == reflect.TypeOf(err).String() {
				// 代码自定义的异常
				ret := err.(*atom.MyError)

				httpStatus := ret.HttpStatus()
				if 0 == ret.HttpStatus() {
					httpStatus = http.StatusOK
				}

				c.JSON(httpStatus, gin.H{
					"code": ret.Code(),
					"msg":  ret.Error(),
					"data": nil,
				})
			} else {
				// 系统级别的异常，打印错误栈
				debug.PrintStack()

				c.JSON(http.StatusOK, gin.H{
					"code": "-10",
					"msg":  errorToString(err),
					"data": nil,
				})
			}

			//终止后续接口调用
			c.Abort()
		}
	}()

	// 后继
	c.Next()
}

// recover错误信息转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
