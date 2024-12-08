package repository

import "gorm.io/gorm"

type (
	PembayaranRepository interface {
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
