package entitys

import "time"

type SessionEntity struct {
	BaseEntity
	SessionId  string    `json:"sessionId" gorm:"primaryKey;uniqueIndex;not null;size:20"`
	UserId     string    `json:"userId" gorm:"primaryKey;not null;index;size:20"`
	Token      string    `json:"token" gorm:"not null;size:255"`
	ExpirestAt time.Time `json:"expiresAt" gorm:"not null"`
	LastActive time.Time `json:"lastActive" gorm:"not null"`
	IPAddress  string    `json:"ipAddress" gorm:"not null;size:45"` // IPv4 or IPv6
	UserAgent  string    `json:"userAgent" gorm:"not null;size:255"`
	IsActive   bool      `json:"isActive" gorm:"default:true"`

	User UserEntity `json:"user,omitempty" gorm:"foreignKey:UserId;references:UserId"`
}
