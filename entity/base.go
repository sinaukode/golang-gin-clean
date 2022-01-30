package entity

import (
	"time"
)

type BaseModel struct {
	Id        uint       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"createdAt" `
	UpdatedAt time.Time  `json:"updatedAt" `
	DeletedAt *time.Time `json:"deletedAt"`
}
