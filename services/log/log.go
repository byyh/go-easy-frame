/*
 * 日志采集对接
 * 接入单个或多个数据库均需要在这里配置
 *
 * 注意：目前数据库初始化包括 gorm 库的初始化，以及 beego 的orm 初始化
 * 在调用的时候可以根据需要调用相关的库
 * 调用gorm的时候请采用 GormDb() 获取db，db会自动
 */

package log

import (
	"fmt"
	"log"

	"go-easy-frame/config"

	"github.com/byyh/go/com"
	"github.com/fluent/fluent-logger-golang/fluent"
)

type FluentLog struct {
	Logger *fluent.Fluent
}

var (
	ftLog       *FluentLog
	isFluentLog bool
)

func New() *FluentLog {
	return ftLog
}

func InitLog() {
	cfg := config.GetEnv()
	var err error
	isFluentLog = cfg.Isfluentlog

	if !isFluentLog {
		return
	}

	log.Println("begin init log")
	ftLog = new(FluentLog)

	port := cfg.Fluent.Port

	ftLog.Logger, err = fluent.New(fluent.Config{
		FluentPort: int(port),
		FluentHost: cfg.Fluent.Host,
		MaxRetry:   100,
		Async:      false,
	})
	if nil != err {
		log.Println("init fluent log failed", err)
	}
}

func (this *FluentLog) Error(args ...interface{}) {
	this.Write("error", args)
}

func (this *FluentLog) Warn(args ...interface{}) {
	this.Write("warn", args)
}

func (this *FluentLog) Info(args ...interface{}) {
	this.Write("info", args)
}

func (this *FluentLog) Debug(args ...interface{}) {
	this.Write("debug", args)
}

func (this *FluentLog) Write(level string, args ...interface{}) {
	if 0 >= len(args) {
		log.Println("没有日志需要输出", level)
		return
	}

	if !isFluentLog {
		log.Println(level, args)

		return
	}

	//log.Println("log-p=", level, args)
	if nil == ftLog || nil == ftLog.Logger {
		InitLog()
	}
	tag := config.GetEnv().Fluent.Tag + "-" + level

	mp := make(map[string]interface{})
	mp["tag"] = tag
	mp["level"] = level
	mp["log_time"] = new(com.Time).Now()

	var newArgs []string
	num := len(args)
	for i := 0; i < num; i++ {
		tmp := args[i].([]interface{})
		for _, arg := range tmp {
			var str string
			if "string" != com.Typeof(arg) {
				str = fmt.Sprintf("%s", arg)
			} else {
				str = arg.(string)
			}

			newArgs = append(newArgs, str)
		}
	}

	mp["data"] = newArgs

	go this.Post(tag, mp)
}

func (this *FluentLog) Post(tag string, mp map[string]interface{}) {
	defer func() {
		if err := recover(); nil != err {
			log.Println("fluent lib occur exception: ", err)
		}
	}()
	//log.Println("post data===", tag, mp)
	err := this.Logger.Post(tag, mp)
	if nil != err {
		log.Println("post fluent log failed", err)
		log.Println("err-data===", mp["data"])
	}
	//log.Println("发送完毕", mp["data"])
}
