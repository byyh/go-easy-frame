package atom

import (
	"log"
	"net/http"
)

type MyError struct {
	errMsg     string
	errCode    int // 小于-10不打日志的错误，大于-10的需要打印日志
	httpStatus int // http的状态码，为0或200表示输出均为200，其他值表示需要返回该http状态码
}

// 只通过初始化状态吗返回通用错误信息
func NewMyErrorByCode(code int) (e *MyError) {
	e = new(MyError)
	e.errMsg = GetRetMsgByCode(code)
	e.errCode = code
	e.httpStatus = http.StatusOK
	return
}

// 初始化返回自定义code和msg的错误
func NewMyError(msg string, code int) (e *MyError) {
	e = new(MyError)
	e.errMsg = msg
	e.errCode = code
	e.httpStatus = http.StatusOK
	return
}

// 初始化返回非200的http状态码的错误
func NewMyErrorByHttpStatus(msg string, httpStatus int) (e *MyError) {
	e = new(MyError)
	e.errMsg = msg
	e.errCode = -100
	e.httpStatus = httpStatus
	return
}

func (e *MyError) Error() string {
	return e.errMsg
}

func (e *MyError) Code() int {
	return e.errCode
}

func (e *MyError) HttpStatus() int {
	return e.httpStatus
}

func CheckErr(err error, param ...interface{}) {
	if err != nil {
		log.Println("error: ", err, param) //deal error here
		panic(err)
	}
}

func CheckErrFailedPanic(err error, code int, msg string, param ...interface{}) {
	if err != nil {
		log.Println("error: ", err, msg, param) //deal error here
		panic(NewMyError(GetRetMsgByCode(code), code))
	}
}

// 全局返回失败，用于任意过程中的流程终止且返回失败
func Failed(code int, msg string) {
	panic(NewMyError(msg, code))
}

// 返回非200HTTP状态码，用于任意过程中的流程终止且返回失败
func OutHttpError(httpCode int, msg string) {
	panic(NewMyErrorByHttpStatus(msg, httpCode))
}
