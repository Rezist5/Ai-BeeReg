package services

import (
	"errors"
	"os"
	"time"

	"github.com/Rezist5/Ai-BeeReg/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repo      *repository.AuthRepository
	blacklist map[string]struct{}
}

func NewAuthService(db *gorm.DB) *AuthService {
	repo := repository.NewAuthRepository(db)
	return &AuthService{
		repo:      repo,
		blacklist: make(map[string]struct{}),
	}
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("пользователь не найден")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("неверный пароль")
	}

	token, err := s.GenerateToken(user.ID, user.Email, user.RoleID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) GenerateToken(userID uint, email string, roleID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub":     userID,
		"email":   email,
		"role_id": roleID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (s *AuthService) Logout(token string) {
	s.blacklist[token] = struct{}{}
}

func (s *AuthService) IsTokenBlacklisted(token string) bool {
	_, exists := s.blacklist[token]
	return exists
}
