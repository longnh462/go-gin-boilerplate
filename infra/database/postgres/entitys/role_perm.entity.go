package entitys

import "github.com/google/uuid"

type RolePermissionEntity struct {
	BaseEntity

	RolePermId uuid.UUID `json:"role_perm_id" gorm:"primaryKey; not null; size: 36"`
	RoleId     uuid.UUID `json:"role_id" gorm:"not null;size:36"`
	PermId     uuid.UUID `json:"user_id" gorm:"not null;size:36"`

	Role       RoleEntity       `json:"role,omitempty" gorm:"foreignKey: RoleId; references: RoleId; constraint: OnDelete: CASCADE, OnUpdate: CASCADE"`
	Permission PermissionEntity `json:"permission,omitempty" gorm:"foreignKey: PermId, references:PermId; constrain: OnDelete:CASCADE, OnUpdate:CASCADE"`
}
