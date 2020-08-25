package dto

import (
	"time"
)

type PhoneNumberDetectionDto struct {
	Id          uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	BatchId     string    `gorm:"column:batch_id;type:varchar(100)"`
	PhoneNumber uint64    `gorm:"column:phone_number"`
	Status      uint      `gorm:"column:status"`
	CreateTime  time.Time `gorm:"column:crate_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
}
