package response

import "time"

type PhoneDetectionResponse struct {
	Result Result `json:"result"`
	Uuid   string `json:"uuid"`
}

type Result struct {
	PhoneStatus int       `json:"phoneStatus"`
	Carrier     int       `json:"carrier"`
	Phone       int64     `json:"phone"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	CreateTime  time.Time `json:"createTime"`
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
