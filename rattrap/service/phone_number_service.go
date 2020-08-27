package service

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"io/ioutil"
	"log"
	"mouse/rattrap/config"
	"mouse/rattrap/db"
	"mouse/rattrap/mq"
	"mouse/rattrap/redis"
	"mouse/rattrap/request"
	"mouse/rattrap/response"
	"net/http"
	"strconv"
	"strings"
)

const (
	contentType = "application/json"
	url         = "http://172.16.0.100:7890/api/v1.0/phone-detection/phoneStatus/detection"
	tableName   = "phone_detection"
	topic       = "phone"
)

type PhoneDetection interface {
	Detection(request request.PhoneDetectionRequest) response.PhoneDetectionResponse
}

type PhoneDetectionService struct {
	tag string
}

func (p *PhoneDetectionService) Detection(localRequest request.PhoneDetectionRequest) response.PhoneDetectionResponse {
	result := p.remoteDetection(localRequest.Convert())
	detectionResponse := result.Convert()
	p.persistenceDB(detectionResponse)
	p.messageQueue(detectionResponse)
	return detectionResponse
}

func (p *PhoneDetectionService) persistenceDB(detectionResponse response.PhoneDetectionResponse) {
	dto := detectionResponse.ConvertDto()
	db.Db.Table(tableName).Create(&dto)
}

func (p *PhoneDetectionService) messageQueue(detectionResponse response.PhoneDetectionResponse) {
	message := detectionResponse.ConvertMessage()
	jsonBytes, _ := json.Marshal(message)
	sendMessage := &primitive.Message{
		Topic: topic,
		Body:  jsonBytes,
	}
	sendMessage = sendMessage.WithTag(p.tag)
	sendResult, _ := mq.ProducerMQ.SendSync(context.Background(), sendMessage)
	log.Printf("查看发送消息的结果:%v\n", sendResult)
}

func (p *PhoneDetectionService) remoteDetection(request request.PhoneDetectionRemoteRequest) response.PhoneDetectionRemoteResponse {
	jsonStr, _ := json.Marshal(request)
	resp, err := http.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("访问出现错误了...")
		panic(err)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	remoteResponse := response.PhoneDetectionRemoteResponse{}
	_ = json.Unmarshal(result, &remoteResponse)
	return remoteResponse
}

//
// 号码检测的消费者
//

type PhoneDetectionConsumer struct {
	tag          string
	pushConsumer rocketmq.PushConsumer
}

func (p *PhoneDetectionConsumer) consumeMessage(messages []*primitive.MessageExt) {
	var message mq.PhoneDetectionMessage
	for _, value := range messages {
		_ = json.Unmarshal(value.Body, &message)
		key := strings.Join([]string{p.tag, strconv.Itoa(message.UserId), strconv.Itoa(message.Status)}, "-")
		redis.Incrby(key, 1)
		log.Printf("查看消费的消息对象:%v\n", message)
	}
}

func (p *PhoneDetectionConsumer) initPushConsumer() {
	pushConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer(config.Conf.RocketMQ.NameServer),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName(config.Conf.RocketMQ.ConsumerGroupName),
	)
	if err != nil {
		panic(err)
	}
	_ = pushConsumer.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		p.consumeMessage(ext)
		return consumer.ConsumeSuccess, nil
	})
	err = pushConsumer.Start()
	if err != nil {
		panic(err)
	}
	p.pushConsumer = pushConsumer
}

var PhoneDetectionServiceImpl PhoneDetectionService

func init() {
	start()
}

func start() {
	var tag = "phone-detection-api"
	PhoneDetectionServiceImpl = PhoneDetectionService{
		tag: tag,
	}

	phoneDetectionConsumer := PhoneDetectionConsumer{
		tag: tag,
	}
	phoneDetectionConsumer.initPushConsumer()
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 100
}
