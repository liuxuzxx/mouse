package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mouse/rattrap/request"
	"mouse/rattrap/response"
	"net/http"
	"time"
)

const (
	contentType = "application/json"
	url         = "http://172.16.0.100:7890/api/v1.0/phone-detection/phoneStatus/detection"
)

type PhoneDetection interface {
	Detection(request request.PhoneDetectionRequest) response.PhoneDetectionResponse
}

type PhoneDetectionService struct {
}

func (p *PhoneDetectionService) Detection(localRequest request.PhoneDetectionRequest) response.PhoneDetectionResponse {
	result := p.remoteDetection(localRequest.Convert())
	return result.Convert()
}

func (p *PhoneDetectionService) remoteDetection(request request.PhoneDetectionRemoteRequest) response.PhoneDetectionRemoteResponse {
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(request)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
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

var PhoneDetectionServiceImpl PhoneDetectionService

func init() {
	PhoneDetectionServiceImpl = PhoneDetectionService{}
}
