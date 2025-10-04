package services

import (
	"context"
	"fmt"
	"regexp"

	"github.com/morodik/convoy/internal/models"
	"github.com/morodik/convoy/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserProfile(ctx context.Context, userID uint) (*models.User, error) {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

func (s *UserService) UpdateUsername(ctx context.Context, userID uint, username string) (*models.User, error) {
	// валидация
	if len(username) < 3 || len(username) > 50 {
		return nil, fmt.Errorf("username должен быть от 3 до 50 символов")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(username) {
		return nil, fmt.Errorf("username может содержать только буквы, цифры и подчёркивания")
	}

	user, err := s.repo.UpdateUsername(ctx, userID, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
