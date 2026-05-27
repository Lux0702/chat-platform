package service

import (
	"chat-platform-api/internal/auth/entity"
	"chat-platform-api/internal/auth/repository"
	"chat-platform-api/pkg/jwt"
	"chat-platform-api/pkg/utils"
	"errors"
	"time"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func (s *AuthService) Register(
	name string,
	email string,
	password string,
) error {

	existingUser, err := s.userRepo.FindByEmail(email)

	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	user := entity.User{
		Name:      name,
		Email:     email,
		Password:  hashed,
		CreatedAt: time.Now(),
	}
	return s.userRepo.Create(&user)
}

func (s *AuthService) Login(
	email string,
	password string,
) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}
	ok := utils.ComparePassword(
		password,
		user.Password,
	)
	if !ok {
		return "", errors.New("invalid password")
	}

	token, err := jwt.GenerateToken(
		user.ID.Hex(),
	)
	if err != nil {
		return "", err
	}
	return token, nil
}
func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repository.NewUserRepository(),
	}
}
