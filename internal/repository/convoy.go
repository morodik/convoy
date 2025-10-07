package repository

import (
	"context"

	"github.com/morodik/convoy/internal/models"
	"gorm.io/gorm"
)

type ConvoyRepository struct {
	db *gorm.DB
}

func NewConvoyRepository(db *gorm.DB) *ConvoyRepository {
	return &ConvoyRepository{db: db}
}

func (r *ConvoyRepository) CreateConvoy(ctx context.Context, convoy *models.Convoy) (*models.Convoy, error) {
	if err := r.db.WithContext(ctx).Create(convoy).Error; err != nil {
		return nil, err
	}
	return convoy, nil
}
