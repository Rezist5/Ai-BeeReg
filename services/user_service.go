package services

import (
	"errors"

	"github.com/Rezist5/Ai-BeeReg/models"
	"github.com/Rezist5/Ai-BeeReg/repository"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{repo: repository.NewUserRepository(db)}
}

func (s *UserService) CreateUser(email, fullname, password, roleName string) (*models.User, error) {
	var role models.Role
	if err := s.repo.DB.Where("name = ?", roleName).First(&role).Error; err != nil {
		return nil, errors.New("role not found")
	}

	user := &models.User{
		Email:    email,
		Fullname: fullname,
		Password: password,
		RoleID:   role.ID,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAllUsers(role string) ([]models.User, error) {
	if role == "ROOT" {
		return s.repo.FindAll()
	} else if role == "ADMIN" {
		return s.repo.FindByRoleID(1)
	} else {
		return nil, errors.New("access denied")
	}
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) UpdateUser(id uint, email, fullname, password string) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Email = email
	user.Fullname = fullname
	user.Password = password
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(id uint) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(user)
}
