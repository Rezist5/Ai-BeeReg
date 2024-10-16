package models

import (
	"gorm.io/gorm"
)

type AdminCompany struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	CompanyID uint `json:"company_id"`
	AdminID   uint `json:"admin_id"`
}

func (ac *AdminCompany) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&AdminCompany{})
}
