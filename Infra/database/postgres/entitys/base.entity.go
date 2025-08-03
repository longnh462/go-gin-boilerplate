package entitys

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	CreateAt    time.Time `json:"create_At" gorm:"autoCreateTime"`
	UpdateAt    time.Time `json:"update_At" gorm:"autoCreateTime"`
	CreateUsr   string    `json:"create_usr" gorm:"size:25;not null;default:'system'"`
	UpdateUsr   string    `json:"update_usr" gorm:"size:25;not null;default:'system'"`
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
}
