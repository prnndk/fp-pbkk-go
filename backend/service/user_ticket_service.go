package service

import (
	"context"

	"github.com/prnndk/final-project-golang-pbkk/dto"
	"github.com/prnndk/final-project-golang-pbkk/entity"
	"github.com/prnndk/final-project-golang-pbkk/repository"
)

type (
	UserTicketService interface {
		UserBuyTicket(ctx context.Context, req dto.UserTicketCreateRequest, userId string) (dto.UserTicketResponse, error)
	}

	userTicketService struct {
		userTicketRepo repository.EventTicketRepository
		eventRepo      repository.EventRepository
	}
)

func NewUserTicketService(userTicketRepo repository.EventTicketRepository) UserTicketService {
	return &userTicketService{
		userTicketRepo: userTicketRepo,
	}
}

func (s *userTicketService) UserBuyTicket(ctx context.Context, req dto.UserTicketCreateRequest, userId string) (dto.UserTicketResponse, error) {
	_, flag, _ := s.userTicketRepo.CheckUserTicket(ctx, nil, userId, req.EventID)
	if flag {
		return dto.UserTicketResponse{}, dto.ErrUserTicketAlreadyExists
	}

	_, flag, _ = s.eventRepo.CheckIfEventExist(ctx, nil, req.EventID)
	if !flag {
		return dto.UserTicketResponse{}, dto.ErrEventCannotBeFound
	}

	userTicket := entity.UserTicket{
		UserId:     userId,
		EventId:    req.EventID,
		Quantity:   req.Quantity,
		TotalPrice: req.TotalPrice,
	}

	buyTicket, err := s.userTicketRepo.UserBuyEvent(ctx, nil, userTicket)
	if err != nil {
		return dto.UserTicketResponse{}, dto.ErrBuyTicket
	}

	return dto.UserTicketResponse{
		ID:         buyTicket.ID.String(),
		UserID:     buyTicket.UserId,
		EventID:    buyTicket.EventId,
		Quantity:   buyTicket.Quantity,
		TotalPrice: buyTicket.TotalPrice,
	}, nil
}
