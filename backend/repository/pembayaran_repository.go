package repository

import (
	"context"

	"github.com/prnndk/final-project-golang-pbkk/entity"
	"gorm.io/gorm"
)

type (
	PembayaranRepository interface {
		CreatePembayaran(ctx context.Context, tx *gorm.DB, pembayaran entity.Pembayaran) (entity.Pembayaran, error)
	}

	pembayaranRepository struct {
		db *gorm.DB
	}
)

func NewPembayaranRepository(db *gorm.DB) PembayaranRepository {
	return &pembayaranRepository{
		db: db,
	}
}

func (pr *pembayaranRepository) CreatePembayaran(ctx context.Context, tx *gorm.DB, pembayaran entity.Pembayaran) (entity.Pembayaran, error) {
	if tx == nil {
		tx = pr.db
	}

	if err := tx.WithContext(ctx).Create(&pembayaran).Error; err != nil {
		return entity.Pembayaran{}, err
	}

	return pembayaran, nil
}
