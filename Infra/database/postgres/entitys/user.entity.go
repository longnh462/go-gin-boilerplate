package entitys

type UserEntity struct {
	BaseEntity

	UserId     string `json:"userId" gorm:"uniqueIndex;not null; size: 20"`
	Email      string `json:"email" gorm:"uniqueIndex;not null;size:50"`
	Username   string `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Password   string `json:"-" gorm:"not null;size:50"`
	FirstName  string `json:"firstName" gorm:"not null;size:25"`
	LastName   string `json:"lastName" gorm:"not null;size:50"`
	Phone      string `json:"phone" gorm:"not null;size:11"`
	IsActive   bool   `json:"isActive" gorm:"default:false"`
	IsVerified bool   `json:"isVerified" gorm:"default:false"`
}
