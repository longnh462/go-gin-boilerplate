package entitys

import "github.com/google/uuid"

type PermissionEntity struct {
	BaseEntity
	PermId   uuid.UUID `json:"permission_id" gorm:"primaryKey;not null;size:36"`
	PermName string    `json:"permission_name" gorm:"not null;size:255"`
	PerDesc  string    `json:"description" gorm:"size:500"`
	IsActive bool      `json:"is_active" gorm:"default:true"`

	// Relationships
	Roles []RoleEntity `json:"roles,omitempty" gorm:"many2many:role_permissions;"`
}
