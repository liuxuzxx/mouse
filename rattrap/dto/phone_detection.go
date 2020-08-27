package dto

import (
	"time"
)

type PhoneNumberDetectionDto struct {
	Id          uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserId      int       `gorm:"column:user_id"`
	BatchId     string    `gorm:"column:batch_id;type:varchar(100)"`
	PhoneNumber int64     `gorm:"column:phone_number"`
	Status      int       `gorm:"column:status"`
	CreateTime  time.Time `gorm:"column:create_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
}
