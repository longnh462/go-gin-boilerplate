package entitys

import (
	"time"

	"github.com/google/uuid"
)

type SessionEntity struct {
	BaseEntity
	SessionId  uuid.UUID    `json:"session_id" gorm:"primaryKey;uniqueIndex;not null;size:36"`
	UserId     uuid.UUID    `json:"user_id" gorm:"primaryKey;not null;index;size:36"`
	Token      string    `json:"token" gorm:"not null;size:255"`
	ExpirestAt time.Time `json:"expires_at" gorm:"not null"`
	LastActive time.Time `json:"last_active" gorm:"not null"`
	IPAddress  string    `json:"ip_address" gorm:"not null;size:45"` // IPv4 or IPv6
	UserAgent  string    `json:"user_agent" gorm:"not null;size:255"`
	IsActive   bool      `json:"is_active" gorm:"default:true"`

}
