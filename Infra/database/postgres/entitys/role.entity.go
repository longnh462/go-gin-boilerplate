package entitys

import "github.com/google/uuid"

type RoleEntity struct {
	BaseEntity

	RoleId      uuid.UUID `json:"role_id" gorm:"uniqueIndex;primaryKey;not null;size:36"`
	Name        string `json:"role_nm" gorm:"not null;size:50"`
	DisplayName string `json:"role_display_nm" gorm:"size:100"`
	Description string `json:"role_dsc" gorm:"size:255"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`

	// Relationships
	Users     []UserEntity `json:"users,omitempty" gorm:"many2many:user_roles;foreignKey:RoleId;joinForeignKey:RoleId;References:UserId;joinReferences:UserId;"`
}
