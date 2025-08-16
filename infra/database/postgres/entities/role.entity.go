package entities

import "github.com/google/uuid"

type RoleEntity struct {
	BaseEntity
	RoleId   uuid.UUID `json:"role_id" gorm:"uniqueIndex;primaryKey;not null;size:36;type:uuid;default:uuid_generate_v4()"`
	RoleName string    `json:"role_nm" gorm:"not null;size:50;uniqueIndex"`
	RoleDesc string    `json:"role_dsc" gorm:"size:255"`
	IsActive bool      `json:"is_active" gorm:"default:true"`

	// Relationships
	Users           []UserEntity           `json:"users,omitempty" gorm:"foreignKey:RoleId"`
	RolePermissions []RolePermissionEntity `json:"role_permissions,omitempty" gorm:"foreignKey:RoleId;references:RoleId"`
}

func (RoleEntity) TableName() string { return "roles" }