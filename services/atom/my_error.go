package atom

import (
	"log"
)

type MyError struct {
	errMsg  string
	errCode int // 小于-10不打日志的错误，大于-10的需要打印日志
}

func NewMyErrorByCode(code int) (e *MyError) {
	e = new(MyError)
	e.errMsg = GetRetMsgByCode(code)
	e.errCode = code
	return
}

func NewMyError(msg string, code int) (e *MyError) {
	e = new(MyError)
	e.errMsg = msg
	e.errCode = code
	return
}

func (e *MyError) Error() string {
	return e.errMsg
}

func (e *MyError) Code() int {
	return e.errCode
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
