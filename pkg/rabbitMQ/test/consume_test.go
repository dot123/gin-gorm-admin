package test

import (
	"fmt"
	"github.com/dot123/gin-gorm-admin/pkg/rabbitMQ"
	"sync"
	"testing"
)

func TestConsume(t *testing.T) {
	initConsumerabbitmq()
	Consume()
}

func Consume() {
	nomrl := &rabbitMQ.ConsumeReceive{
		ExchangeName: "testChange31", //队列名称
		ExchangeType: rabbitMQ.EXCHANGE_TYPE_DIRECT,
		Route:        "",
		QueueName:    "testQueue31",
		IsTry:        true,  //是否重试
		IsAutoAck:    false, //自动消息确认
		MaxReTry:     5,     //最大重试次数
		EventFail: func(code int, e error, data []byte) {
			fmt.Printf("error:%s", e)
		},
		EventSuccess: func(data []byte, header map[string]interface{}, retryClient rabbitMQ.RetryClientInterface) bool { //如果返回true 则无需重试
			_ = retryClient.Ack()
			fmt.Printf("data:%s\n", string(data))
			return true
		},
	}
	instanceConsumePool.RegisterConsumeReceive(nomrl)
	err := instanceConsumePool.RunConsume()
	if err != nil {
		fmt.Println(err)
	}
}

var onceConsumePool sync.Once
var instanceConsumePool *rabbitMQ.RabbitMQPool

func initConsumerabbitmq() *rabbitMQ.RabbitMQPool {
	onceConsumePool.Do(func() {
		instanceConsumePool = rabbitMQ.NewConsumePool()
		//instanceConsumePool.SetMaxConsumeChannel(100)
		//err := instanceConsumePool.Connect("127.0.0.1", 5672, "admin", "admin")
		err := instanceConsumePool.ConnectVirtualHost("127.0.0.1", 5672, "temptest", "test123456", "/temptest1")
		if err != nil {
			fmt.Println(err)
		}
	})
	return instanceConsumePool
}
