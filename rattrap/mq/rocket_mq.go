package mq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"log"
	"mouse/rattrap/config"
)

//
// 存放MQ的发送和接受的包
//

var ProducerMQ rocketmq.Producer

func init() {
	producerInit()
}

func producerInit() {
	var err error
	ProducerMQ, err = rocketmq.NewProducer(
		producer.WithNameServer(config.Conf.RocketMQ.NameServer),
		producer.WithRetry(2),
		producer.WithGroupName(config.Conf.RocketMQ.GroupName),
	)
	if err != nil {
		log.Printf("Start RocketMQ Producer error!")
		log.Panic(err)
	}
	err = ProducerMQ.Start()
	if err != nil {
		log.Printf("Start RocketMQ Producer error!")
		log.Panic(err)
	}
	log.Printf("Start RocketMQ Producer!")
}
