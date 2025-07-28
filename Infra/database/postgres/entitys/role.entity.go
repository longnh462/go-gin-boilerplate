package entitys

type RoleEntity struct {
	BaseEntity

	RoleId      string `json:"roleId" gorm:"uniqueIndex;primaryKey;not null;size:20"`
	Name        string `json:"name" gorm:"not null;size:50"`
	DisplayName string `json:"displayName" gorm:"size:100"`
	Description string `json:"description" gorm:"size:255"`
	IsActive    bool   `json:"isActive" gorm:"default:true"`

	// Relationships
	UserRoles []UserRole   `json:"userRoles,omitempty" gorm:"foreignKey:RoleID"`
	Users     []UserEntity `json:"users,omitempty" gorm:"many2many:user_roles;"`
}
