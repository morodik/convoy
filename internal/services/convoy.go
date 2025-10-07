package services

import (
	"context"
	"errors"
	"time"

	"github.com/morodik/convoy/internal/models"
	"github.com/morodik/convoy/internal/repository"
)

type ConvoyService struct {
	repo *repository.ConvoyRepository
}

func NewConvoyService(repo *repository.ConvoyRepository) *ConvoyService {
	return &ConvoyService{repo: repo}
}

func (s *ConvoyService) CreateConvoy(ctx context.Context, userID uint, title string, startTime time.Time, endTime *time.Time, isPrivate bool) (*models.Convoy, error) {
	if title == "" {
		return nil, errors.New("название конвоя не может быть пустым")
	}
	if startTime.Before(time.Now()) {
		return nil, errors.New("время начала не может быть в прошлом")
	}
	if endTime != nil && endTime.Before(startTime) {
		return nil, errors.New("время окончания не может быть раньше времени начала")
	}

	convoy := &models.Convoy{
		Title:     title,
		StartTime: startTime,
		EndTime:   endTime,
		IsPrivate: isPrivate,
		CreatedBy: userID,
	}

	return s.repo.CreateConvoy(ctx, convoy)
}
