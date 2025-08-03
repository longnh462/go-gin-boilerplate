package entitys

import "github.com/google/uuid"

type UserRole struct {
    BaseEntity

    UserId     uuid.UUID `json:"user_id" gorm:"not null;primaryKey;index;size:36"`
    RoleId     uuid.UUID `json:"role_id" gorm:"not null;primaryKey;index;size:36"`
    IsActive   bool   `json:"is_active" gorm:"default:true"`

    // Relationships
    User UserEntity `json:"users,omitempty" gorm:"foreignKey:UserId;references:UserId"`
    Role RoleEntity `json:"roles,omitempty" gorm:"foreignKey:RoleId;references:RoleId"`
}