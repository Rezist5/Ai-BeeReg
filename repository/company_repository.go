// repository/company_repository.go
package repository

import (
	"github.com/Rezist5/Ai-BeeReg/models"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	DB *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{DB: db}
}

func (r *CompanyRepository) Create(company *models.Company) error {
	return r.DB.Create(company).Error
}

func (r *CompanyRepository) FindAll() ([]models.Company, error) {
	var companies []models.Company
	err := r.DB.Find(&companies).Error
	return companies, err
}

func (r *CompanyRepository) FindByID(id uint) (*models.Company, error) {
	var company models.Company
	err := r.DB.First(&company, id).Error
	return &company, err
}

func (r *CompanyRepository) Update(company *models.Company) error {
	return r.DB.Save(company).Error
}

func (r *CompanyRepository) Delete(company *models.Company) error {
	return r.DB.Delete(company).Error
}

func (r *CompanyRepository) AssignAdmin(company *models.Company, admin *models.User) error {
	return r.DB.Model(company).Association("Admins").Append(admin)
}
func (r *CompanyRepository) RemoveAdmin(companyID uint, adminID uint) error {
	return r.DB.Where("company_id = ? AND admin_id = ?", companyID, adminID).Delete(&models.AdminCompany{}).Error
}
