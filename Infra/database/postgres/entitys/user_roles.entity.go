package entitys

type UserRole struct {
    BaseEntity            
    
    UserID   uint `json:"userId" gorm:"not null;index"`
    RoleID   uint `json:"roleId" gorm:"not null;index"`
    IsActive bool `json:"isActive" gorm:"default:true"`
    
    // Relationships
    User     UserEntity `json:"user,omitempty" gorm:"foreignKey:UserID"`
    Role     RoleEntity `json:"role,omitempty" gorm:"foreignKey:RoleID"`
}