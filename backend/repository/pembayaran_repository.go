package repository

import (
	"context"

	"github.com/prnndk/final-project-golang-pbkk/entity"
	"gorm.io/gorm"
)

type (
	PembayaranRepository interface {
		CreatePembayaran(ctx context.Context, tx *gorm.DB, pembayaran entity.Pembayaran) (entity.Pembayaran, error)
		CheckPembayaranByTicketId(ctx context.Context, tx *gorm.DB, ticketId string) (entity.Pembayaran, bool, error)
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

func (pr *pembayaranRepository) CheckPembayaranByTicketId(ctx context.Context, tx *gorm.DB, ticketId string) (entity.Pembayaran, bool, error) {
	if tx == nil {
		tx = pr.db
	}

	var pembayaran entity.Pembayaran
	if err := tx.WithContext(ctx).Where("ticket_id = ?", ticketId).First(&pembayaran).Error; err != nil {
		return entity.Pembayaran{}, false, err
	}

	return pembayaran, true, nil
}
