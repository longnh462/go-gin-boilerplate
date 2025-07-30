package entitys

import "time"

type SessionEntity struct {
	BaseEntity
	SessionId  string
	UserId     string
	Token      string
	ExpirestAt time.Time
	LastActive time.Time
	IPAddress  string
	UserAgent  string
	IsActive   bool

	User UserEntity
}
