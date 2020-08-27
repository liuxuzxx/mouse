package request

import (
	uuid "github.com/iris-contrib/go.uuid"
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
		Uuid:   uuid.Must(uuid.NewV4()).String(),
		Phones: []int64{p.Number},
	}
}
