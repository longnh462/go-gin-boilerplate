package entities

import "github.com/google/uuid"

type PermissionEntity struct {
	BaseEntity
	PermId   uuid.UUID `json:"permission_id" gorm:"primaryKey;not null;size:36;type:uuid;default:uuid_generate_v4()"`
	PermName string    `json:"permission_name" gorm:"not null;size:255;uniqueIndex"`
	PermDesc string    `json:"perm_desc" gorm:"size:500"`
	IsActive bool      `json:"is_active" gorm:"default:true"`

	// Relationships
	RolePermissions []RolePermissionEntity `json:"role_permissions,omitempty" gorm:"foreignKey:PermId;references:PermId"`
}

func (PermissionEntity) TableName() string { return "permissions" }
