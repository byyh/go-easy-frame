package queue

import (
	"errors"
	"log"
	"time"

	"github.com/byyh/go/com"
	"github.com/streadway/amqp"
)

type RabbitMqClient struct {
	conn    *amqp.Connection
	Channel *amqp.Channel

	QueueName    string
	ExchangeName string
	ExchangeType string
	RoutingKey   string
	MqUrl        string

	mustExitEvent bool

	count int
	err   error
}

// 单个消费者最多处理多少个消息后退出
const MAX_HANDLE_MSG_NUM_BY_ONE_CONSUME = 10000000

// 消费处理失败暂停的毫秒数量
const MAX_HANDLE_CONSUME_MSG_FAILED_MILLSECOND = 200

var ()

func (this *RabbitMqClient) CheckErr(err error, msg string) {
	if nil == err {
		return
	}

	log.Println(msg, err)
	panic(err)
}

// 调用发送需要先初始化
func (this *RabbitMqClient) InitPublishMq() {
	if this.Channel == nil {
		this.MqConnect()
	}

	this.Bind()
}

func (this *RabbitMqClient) MqConnect() {
	this.CheckParameter()

	this.conn, this.err = amqp.Dial(this.MqUrl)
	this.CheckErr(this.err, "failed to connect tp rabbitmq")

	this.Channel, this.err = this.conn.Channel()
	this.CheckErr(this.err, "failed to open a channel")
}

func (this *RabbitMqClient) Close() {
	this.Channel.Close()
	this.conn.Close()

	this.Channel = nil
}

func (this *RabbitMqClient) ConsumeNoAck(receve ConsumeInterface) {
	if this.Channel == nil {
		this.MqConnect()
	}

	this.ConnectCheckTimer()

	msgs, err := this.Channel.Consume(this.QueueName, "", true /*true*/, false, false, false, nil)
	this.CheckErr(err, "")

	forwait := make(chan bool)

	go func() {
		for d := range msgs {
			if this.mustExitEvent {
				log.Println("recv mustExitEvent")
				forwait <- true
			}

			this.count++
			if MAX_HANDLE_MSG_NUM_BY_ONE_CONSUME < this.count {
				log.Println("handle msg count gt ", MAX_HANDLE_MSG_NUM_BY_ONE_CONSUME, ",ready quit")
				forwait <- true
			}

			log.Println("--------------------------")
			log.Println(new(com.Time).Now(), ",no: ", this.count, "接受到消息: ", d)

			bl := receve.Handle(d.Body, receve)
			if bl {
				log.Println("handle consume error", ", 消费内容=", d)
				continue
			}
		}
	}()

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")

	<-forwait
}

// 需要应答的消费
func (this *RabbitMqClient) ConsumeMustAck(receve ConsumeInterface) {
	if this.Channel == nil {
		this.MqConnect()
	}

	this.Bind()
	this.ConnectCheckTimer()

	err := this.Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	this.CheckErr(err, "Failed to set QoS")

	msgs, err := this.Channel.Consume(this.QueueName, "", false /*autoAck*/, false, false, false, nil)
	this.CheckErr(err, "")

	forwait := make(chan bool)

	go func() {
		for d := range msgs {
			if this.mustExitEvent {
				log.Println("recv mustExitEvent")
				forwait <- true
			}

			this.count++
			if MAX_HANDLE_MSG_NUM_BY_ONE_CONSUME < this.count {
				log.Println("handle msg count gt ", MAX_HANDLE_MSG_NUM_BY_ONE_CONSUME, ",ready quit")
				forwait <- true
			}

			log.Println("--------------------------")
			log.Println(new(com.Time).Now(), ",no: ", this.count, "接受到消息: ", d)

			bl := receve.Handle(d.Body, receve)
			if !bl {
				log.Println(new(com.Time).Now(), ",rabbitmq-client-consume handle error,", ", 消费内容=", d)
				this.Channel.Nack(d.DeliveryTag, false, true)
				time.Sleep(time.Duration(MAX_HANDLE_CONSUME_MSG_FAILED_MILLSECOND) * time.Millisecond)
				continue
			}

			log.Println(new(com.Time).Now(), ",wait 3 s,", d)
			//time.Sleep(time.Duration(500) * time.Millisecond)

			this.Channel.Ack(d.DeliveryTag, false)
			log.Println(new(com.Time).Now(), ",ack , wait 3 s, ", d.DeliveryTag)

			//time.Sleep(time.Duration(500) * time.Millisecond)

		}
	}()

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")

	<-forwait
}

func (this *RabbitMqClient) ConnectCheckTimer() {
	c := this.conn.NotifyClose(make(chan *amqp.Error))

	go func() {
		log.Println("wait..")
		select {
		case e := <-c:
			log.Println("connect is close by server", e.Reason)
			//this.Close()
			this.mustExitEvent = true
		}
	}()
}

func (this *RabbitMqClient) Publish(data string) {

	err := this.Channel.Publish(
		this.ExchangeName, // exchange
		this.RoutingKey,   // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(data),
			DeliveryMode: amqp.Transient, // 1=non-persistent/Transient, 2=Persistent
			Priority:     0})             // 0-9
	this.CheckErr(err, "Failed to publish a message")
}

func (this *RabbitMqClient) CheckParameter() {
	if "" == this.MqUrl {
		panic(errors.New("MqUrl is not allow empty"))
	}
}

func (this *RabbitMqClient) Bind() {
	err := this.Channel.ExchangeDeclare(
		this.ExchangeName, // name of the exchange
		this.ExchangeType, // type
		true,              // durable
		false,             // delete when complete
		false,             // internal
		false,             // noWait
		nil,               // arguments
	)
	this.CheckErr(err, "Exchange Declare failed:")

	queue, err := this.Channel.QueueDeclare(
		this.QueueName, // name of the queue
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // noWait
		nil,            // arguments
	)
	this.CheckErr(err, "Queue Declare failed:")

	err = this.Channel.QueueBind(
		queue.Name,        // namethis.QueueName of the queue
		this.RoutingKey,   // bindingKey
		this.ExchangeName, // sourceExchange
		false,             // noWait
		nil,               // arguments
	)
	this.CheckErr(err, "Queue Bind failed:")
}

func (this *RabbitMqClient) GetChannel() *amqp.Channel {
	return this.Channel
}
