package main

import (
	"github.com/byyh/go-easy-frame/config"
	"github.com/byyh/go-easy-frame/services/queue"
)

// 消息推送队列消费
type MsgPush struct {
	QueueBase
}

// 消费启动入口
func (this *MsgPush) Run() {

	this.CreateRabbitMq().ConsumeMustAck(this)
}

func (this *MsgPush) Init() {

}

func (this *MsgPush) CreateRabbitMq() *queue.RabbitMqClient {
	cfg := config.GetEnv()

	mq := &queue.RabbitMqClient{
		MqUrl:        cfg.RabbitmqPush.AmqpUri,
		QueueName:    cfg.RabbitmqPush.Queuename,
		ExchangeName: cfg.RabbitmqPush.Exchange,
		ExchangeType: cfg.RabbitmqPush.ExchangeType,
		RoutingKey:   cfg.RabbitmqPush.RoutingKey,
	}

	return mq
}

func (this *MsgPush) HandleExec() bool {
	defer func() {
		if err := recover(); nil != err {
			// 捕获异常处理逻辑 ...
		}
	}()

	// 正常处理逻辑 ...

	return true
}
