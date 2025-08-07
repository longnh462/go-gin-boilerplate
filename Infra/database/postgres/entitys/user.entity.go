package entitys

import "github.com/google/uuid"

type UserEntity struct {
	BaseEntity
	UserId uuid.UUID `json:"user_id" gorm:"primaryKey;uniqueIndex;not null; size: 36"`
	RoleId uuid.UUID `json:"role_id" gorm:"not null;size:36;index:idx_role_id"`
	Email      string `gorm:"uniqueIndex;not null;size:100;index:idx_email"`
	Username   string `json:"user_nm" gorm:"uniqueIndex;not null;size:50"`
	Password   string `json:"usr_pw" gorm:"not null;size:100"`
	FirstName  string `json:"first_nm" gorm:"not null;size:25"`
	LastName   string `json:"last_nm" gorm:"not null;size:50"`
	Phone      string `json:"usr_phone" gorm:"not null;size:11"`
	IsActive   bool   `json:"is_active" gorm:"default:false"`
	IsVerified bool   `json:"is_verified" gorm:"default:false"`

	// Relationships
	Role RoleEntity `json:"role,omitempty" gorm:"foreignKey:RoleId;references:RoleId;"`
	Sessions []SessionEntity `json:"sessions,omitempty" gorm:"foreignKey:UserId;references:UserId"`
}
