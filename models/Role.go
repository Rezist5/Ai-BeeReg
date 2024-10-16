package models

import (
	"gorm.io/gorm"
)

type Role struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"unique;not null" json:"name"`
	Users []User `gorm:"foreignKey:RoleID" json:"users"`
}

func (r *Role) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Role{})
}
