package service

import (
	"context"

	"github.com/prnndk/final-project-golang-pbkk/dto"
	"github.com/prnndk/final-project-golang-pbkk/entity"
	"github.com/prnndk/final-project-golang-pbkk/repository"
)

type (
	PembayaranService interface {
		CreatePembayaran(ctx context.Context, pembayaran dto.PembayaranCreateRequest) (dto.PembayaranCreateResponse, error)
	}

	pembayaranService struct {
		eventTicketRepo repository.EventTicketRepository
		pembayaranRepo  repository.PembayaranRepository
	}
)

func NewPembayaranService(pembayaranRepo repository.PembayaranRepository, eventTicketRepo repository.EventTicketRepository) PembayaranService {
	return &pembayaranService{
		pembayaranRepo:  pembayaranRepo,
		eventTicketRepo: eventTicketRepo,
	}
}

func (ps *pembayaranService) CreatePembayaran(ctx context.Context, pembayaran dto.PembayaranCreateRequest) (dto.PembayaranCreateResponse, error) {

	_, flag, _ := ps.eventTicketRepo.CheckUserTicketById(ctx, nil, pembayaran.TicketID)
	if !flag {
		return dto.PembayaranCreateResponse{}, dto.ErrIdTicketNotFound
	}

	pembayaranEntity := entity.Pembayaran{
		TicketId:         pembayaran.TicketID,
		MetodePembayaran: pembayaran.MetodePembayaran,
		BuktiBayar:       pembayaran.BuktiBayar,
	}

	createPembayaran, err := ps.pembayaranRepo.CreatePembayaran(ctx, nil, pembayaranEntity)
	if err != nil {
		return dto.PembayaranCreateResponse{}, dto.ErrCreatePembayaran
	}

	return dto.PembayaranCreateResponse{
		ID:               createPembayaran.ID.String(),
		TicketID:         createPembayaran.TicketId,
		MetodePembayaran: createPembayaran.MetodePembayaran,
		BuktiBayar:       createPembayaran.BuktiBayar,
	}, nil
}
