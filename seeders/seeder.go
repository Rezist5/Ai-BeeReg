package seeders

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/Rezist5/Ai-BeeReg/models"
)

func Run(db *gorm.DB) error {
	if err := seedRoles(db); err != nil {
		return err
	}
	return seedRootUser(db)
}

func seedRoles(db *gorm.DB) error {
	roles := []models.Role{
		{Name: "USER"},
		{Name: "ADMIN"},
		{Name: "ROOT"},
	}

	for _, role := range roles {
		if err := db.Where(models.Role{Name: role.Name}).FirstOrCreate(&role).Error; err != nil {
			return err
		}
		log.Printf("Role '%s' seeded\n", role.Name)
	}

	return nil
}

func seedRootUser(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		log.Println("Root user already exists, skipping seeding")
		return nil
	}

	rootUser := models.User{
		Email:    "root@example.com",
		Fullname: "Root User",
		Password: hashPassword("supersecretpassword"),
		RoleID:   3,
	}

	if err := db.Create(&rootUser).Error; err != nil {
		return err
	}
	log.Printf("User '%s' seeded\n", rootUser.Fullname)

	return nil
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	return string(hash)
}
