package models

import (
	"gorm.io/gorm"
)

type Company struct {
	ID          uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string   `gorm:"not null" json:"name"`
	Description string   `json:"description"`
	Images      []string `gorm:"type:json" json:"images"`
	Admins      []User   `gorm:"many2many:admin_companies;foreignKey:ID;joinForeignKey:CompanyID;References:ID;joinReferences:AdminID" json:"admins"`
}

func (c *Company) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Company{})
}
