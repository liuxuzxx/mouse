package response

import (
	"math/rand"
	"mouse/rattrap/dto"
	"mouse/rattrap/mq"
	"time"
)

type PhoneDetectionResponse struct {
	Result Result `json:"result"`
	Uuid   string `json:"uuid"`
}

type Result struct {
	UserId      int       `json:"userId"`
	PhoneStatus int       `json:"phoneStatus"`
	Carrier     int       `json:"carrier"`
	Phone       int64     `json:"phone"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	CreateTime  time.Time `json:"createTime"`
}

func (p *PhoneDetectionResponse) ConvertMessage() mq.PhoneDetectionMessage {
	return mq.PhoneDetectionMessage{
		UserId:     p.Result.UserId,
		Number:     p.Result.Phone,
		Status:     p.Result.PhoneStatus,
		CreateDate: time.Time{},
	}
}

func (p *PhoneDetectionResponse) ConvertDto() dto.PhoneNumberDetectionDto {
	return dto.PhoneNumberDetectionDto{
		UserId:      p.Result.UserId,
		BatchId:     p.Uuid,
		PhoneNumber: p.Result.Phone,
		Status:      p.Result.PhoneStatus,
		CreateTime:  p.Result.CreateTime,
		UpdateTime:  p.Result.CreateTime,
	}
}

type PhoneDetectionRemoteResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

type Data struct {
	Uuid   string        `json:"uuid"`
	Phones []PhoneResult `json:"phones"`
}

type PhoneResult struct {
	PhoneStatus int    `json:"phoneStatus"`
	Carrier     int    `json:"carrier"`
	Phone       int64  `json:"phone"`
	Province    string `json:"province"`
	City        string `json:"city"`
}

func (r *PhoneResult) convert() Result {
	return Result{
		UserId:      rand.Intn(100) + 1,
		PhoneStatus: r.PhoneStatus,
		Carrier:     r.Carrier,
		Phone:       r.Phone,
		Province:    r.Province,
		City:        r.City,
		CreateTime:  time.Now(),
	}
}

func (p *PhoneDetectionRemoteResponse) Convert() PhoneDetectionResponse {
	result := make([]Result, 0)
	for _, value := range p.Data.Phones {
		result = append(result, value.convert())
	}
	return PhoneDetectionResponse{
		Result: result[0],
		Uuid:   p.Data.Uuid,
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
