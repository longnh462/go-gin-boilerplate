package migrations

import (
    "github.com/longnh462/go-gin-boilerplate/Infra/database/postgres/entitys"
    "gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &entitys.UserEntity{},
        &entitys.RoleEntity{},
        &entitys.UserRole{},
    )
}