package main

import (
	"github.com/byyh/go-easy-frame/services/log"
	"github.com/byyh/go-easy-frame/services/queue"
	"runtime/debug"

	"fmt"
)

const MAX_FAILED_HANDLE_COUNT = 10 // 最多失败次数

// 队列消费逻辑基类
type QueueBase struct {
	IsHandleSuccess bool
	isSendSms       string
	FailedMsgIdMap  map[string]int // 失败的临时记录

	Body []byte

	queueName string
	mquri     string

	err error
}

func (this *QueueBase) LoadConfig() {
	this.FailedMsgIdMap = make(map[string]int)
}

// @extend interface ConsumeRecv，成功返回true失败返回false
func (this *QueueBase) Handle(body []byte, consume queue.ConsumeInterface) bool {
	log.New().Info("---接受到新消息-----")
	defer func() {
		if err := recover(); nil != err {
			log.New().Error("消费处理异常", err, fmt.Sprintf("%s", debug.Stack()))
		}
	}()

	this.IsHandleSuccess = true
	this.Body = body
	res := consume.HandleExec()
	log.New().Info("处理结果：", res, this.IsHandleSuccess)
	return this.IsHandleSuccess
}

// 异常处理，
// uniqueKey 标识消息的唯一标记，用于记录重复处理的次数。避免某个消息一直重复处理导致队列阻塞。
func (this *QueueBase) ExceptionHandle(err interface{}, uniqueKey string) {
	ob := log.New()
	ob.Error("消息处理失败,", err, string(this.Body))

	this.IsHandleSuccess = false

	if "0" == uniqueKey || "" == uniqueKey {
		ob.Error("error: 消息id == 0，消息解析错误，删除队列，请检查！！ 消息内容=", string(this.Body))
		this.IsHandleSuccess = true
	} else {
		if num, ok := this.FailedMsgIdMap[uniqueKey]; ok {
			if MAX_FAILED_HANDLE_COUNT < num {
				this.IsHandleSuccess = true
				ob.Error("error: 队列处理失败的次数大于 ", MAX_FAILED_HANDLE_COUNT,
					" 次，删除队列数据，请检查内容,uniqueKey =", uniqueKey,
					" ,消息内容=", string(this.Body))
				delete(this.FailedMsgIdMap, uniqueKey)
			} else {
				this.FailedMsgIdMap[uniqueKey] = num + 1
			}
		} else {
			this.FailedMsgIdMap[uniqueKey] = 1
		}
	}
}
