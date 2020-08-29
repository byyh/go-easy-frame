package main

import (
	"fmt"
	"go-easy-frame/services/db"
	"go-easy-frame/services/log"
	"os"
	"reflect"
)

var (
	classMap map[string]interface{}
)

// Go 不支持直接从字符串创建结构体的反射
func InitClassMap() {
	classMap = make(map[string]interface{})

	classMap["MsgPush"] = &MsgPush{} // 消息推送
}

// 启动示例：  consumer MsgPush Run
//        消费者可执行文件名   任务名  启动函数
func main() {
	num := len(os.Args)
	if 2 > num {
		fmt.Println("param must more than 2")
		return
	}

	db.InitDb()
	log.InitLog()

	log.New().Info("begin:", os.Args[1], os.Args[2])

	InitClassMap()

	class, ok := classMap[os.Args[1]]
	if !ok {
		log.New().Error("class not exists, class name: ", os.Args[1])
		return
	}

	fmt.Println(ok, class)
	Call(class)
}

func Call(class interface{}, args ...interface{}) {
	defer func() {
		if err := recover(); nil != err {
			log.New().Error("error: handle recover", err)
			panic(err)
		}
	}()

	fmt.Println("begin <", os.Args[1], os.Args[2], args, "> run ...")
	var param []reflect.Value

	if 0 < len(args) {
		num := len(args)
		for i := 0; i < num; {
			param = append(param, reflect.ValueOf(args[i]))
			i++
		}
	} else {
		// 取命令行传入的参数
		num := len(os.Args)

		for i := 0; i < num-3; {
			param = append(param, reflect.ValueOf(os.Args[i+3]))
			i++
		}
	}

	cls := reflect.ValueOf(class).MethodByName(os.Args[2])
	if cls.IsNil() {
		log.New().Error("func name not exists, name: ", os.Args[2])
	}

	log.New().Debug("param==", param, len(param))
	cls.Call(param)

	log.New().Info("run over,exit: <", os.Args[1], os.Args[2], args, ">")
}
