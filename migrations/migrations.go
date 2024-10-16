package migrations

import (
	"log"

	"github.com/Rezist5/Ai-BeeReg/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := (&models.Role{}).Migrate(db)
	if err != nil {
		log.Fatalf("Failed to migrate Role model: %v", err)
	}

	err = (&models.User{}).Migrate(db)
	if err != nil {
		log.Fatalf("Failed to migrate User model: %v", err)
	}

	err = (&models.Company{}).Migrate(db)
	if err != nil {
		log.Fatalf("Failed to migrate Company model: %v", err)
	}

	err = (&models.AdminCompany{}).Migrate(db)
	if err != nil {
		log.Fatalf("Failed to migrate AdminCompany model: %v", err)
	}

	log.Println("Database migration completed successfully")
}
