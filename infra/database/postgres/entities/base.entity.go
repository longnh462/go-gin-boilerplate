package entities

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	CreatedAt time.Time      `json:"create_at" gorm:"column:create_at;autoCreateTime"`
	UpdatedAt time.Time      `json:"update_at" gorm:"column:update_at;autoCreateTime"`
	CreatedBy string         `json:"create_by" gorm:"column:created_by;size:36;not null;type:uuid"`
	UpdatedBy string         `json:"update_usr" gorm:"column:updated_by;size:36;not null;type:uuid"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
