package mq

import (
	"time"
)

type PhoneDetectionMessage struct {
	UserId     int       `json:"userId"`
	Number     int64     `json:"number"`
	Status     int       `json:"status"`
	CreateDate time.Time `json:"createTime"`
}