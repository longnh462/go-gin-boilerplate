package entities

import "github.com/google/uuid"

type UserRoleEntity struct {
	BaseEntity

	UserRoleId uuid.UUID `json:"user_role_id" gorm:"primaryKey;not null;size:36;type:uuid;default:uuid_generate_v4()"`
	UserId     uuid.UUID `json:"user_id" gorm:"not null;size:36"`
	RoleId     uuid.UUID `json:"role_id" gorm:"not null;size:36"`

	User UserEntity `json:"user,omitempty" gorm:"foreignKey:UserId;references:UserId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Role RoleEntity `json:"role,omitempty" gorm:"foreignKey:RoleId;references:RoleId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

func (UserRoleEntity) TableName() string { return "user_roles" }