package entitys

type UserRole struct {
    BaseEntity

    UserRoleId string `json:"userRoleId" gorm:"primaryKey;not null;size:20"`
    UserID     string `json:"userId" gorm:"not null;index;size:20"`
    RoleID     string `json:"roleId" gorm:"not null;index;size:20"`
    IsActive   bool   `json:"isActive" gorm:"default:true"`

    // Relationships
    User UserEntity `json:"user,omitempty" gorm:"foreignKey:UserID;references:UserId"`
    Role RoleEntity `json:"role,omitempty" gorm:"foreignKey:RoleID;references:RoleId"`
}