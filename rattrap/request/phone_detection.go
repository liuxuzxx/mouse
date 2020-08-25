package request

import (
	"github.com/google/uuid"
	"strconv"
)

type PhoneDetectionRequest struct {
	Number int64 `json:"number"`
}

type PhoneDetectionRemoteRequest struct {
	Uuid   string  `json:"uuid"`
	Phones []int64 `json:"phones"`
}

func (p *PhoneDetectionRequest) Convert() PhoneDetectionRemoteRequest {
	return PhoneDetectionRemoteRequest{
		Uuid:   strconv.Itoa(uuid.ClockSequence()),
		Phones: []int64{p.Number},
	}
}
