package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string `gorm:"unique;not null" json:"email"`
	Fullname string `gorm:"not null" json:"fullname"`
	Password string `json:"password"`
	RoleID   uint   `gorm:"not null" json:"role_id"`
	Role     Role   `gorm:"foreignKey:RoleID" json:"role"`
}

func (u *User) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
