package service

import "github.com/prnndk/final-project-golang-pbkk/repository"

type (
	PembayaranService interface {
	}

	pembayaranService struct {
		pembayaranRepo repository.PembayaranRepository
	}
)

func NewPembayaranService(pembayaranRepo repository.PembayaranRepository) PembayaranService {
	return &pembayaranService{
		pembayaranRepo: pembayaranRepo,
	}
}
