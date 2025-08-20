package entities

import "github.com/google/uuid"

type UserEntity struct {
	BaseEntity
	UserId     uuid.UUID `json:"user_id" gorm:"primaryKey;uniqueIndex;not null;size:36;type:uuid;default:uuid_generate_v4()"`
	Email      string    `gorm:"uniqueIndex;not null;size:100"`
	Username   string    `json:"user_nm" gorm:"uniqueIndex;not null;size:50"`
	Password   string    `json:"usr_pw" gorm:"not null;size:100"`
	FirstName  string    `json:"first_nm" gorm:"not null;size:25"`
	LastName   string    `json:"last_nm" gorm:"not null;size:50"`
	Phone      string    `json:"usr_phone" gorm:"not null;size:11"`
	IsActive   bool      `json:"is_active" gorm:"default:false"`
	IsVerified bool      `json:"is_verified" gorm:"default:false"`

	// Relationships
	UserRole []UserRoleEntity `json:"user_role,omitempty" gorm:"foreignKey:UserId;references:UserId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Sessions []SessionEntity `json:"sessions,omitempty" gorm:"foreignKey:UserId;references:UserId"`
}

func (UserEntity) TableName() string { return "users" }
