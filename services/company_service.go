// services/company_service.go
package services

import (
	"errors"

	"github.com/Rezist5/Ai-BeeReg/models"
	"github.com/Rezist5/Ai-BeeReg/repository"
	"gorm.io/gorm"
)

type CompanyService struct {
	repo *repository.CompanyRepository
}

func NewCompanyService(db *gorm.DB) *CompanyService {
	return &CompanyService{repo: repository.NewCompanyRepository(db)}
}

func (s *CompanyService) CreateCompany(name, description string, images []string) (*models.Company, error) {
	company := models.Company{Name: name, Description: description, Images: images}
	if err := s.repo.Create(&company); err != nil {
		return nil, err
	}
	return &company, nil
}

func (s *CompanyService) GetAllCompanies() ([]models.Company, error) {
	return s.repo.FindAll()
}

func (s *CompanyService) GetCompanyByID(id uint) (*models.Company, error) {
	company, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("company not found")
	}
	return company, nil
}

func (s *CompanyService) UpdateCompany(company *models.Company, name, description string, images []string) (*models.Company, error) {
	if name != "" {
		company.Name = name
	}
	if description != "" {
		company.Description = description
	}
	if len(images) > 0 {
		company.Images = images
	}
	if err := s.repo.Update(company); err != nil {
		return nil, err
	}
	return company, nil
}

func (s *CompanyService) DeleteCompany(company *models.Company) error {
	return s.repo.Delete(company)
}

func (s *CompanyService) AssignAdmin(company *models.Company, admin *models.User) error {
	return s.repo.AssignAdmin(company, admin)
}

func (s *CompanyService) RemoveAdmin(company *models.Company, adminID uint) error {
	return s.repo.RemoveAdmin(company.ID, adminID)
}
