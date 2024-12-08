package service

import (
	"context"
	"time"

	"github.com/prnndk/final-project-golang-pbkk/dto"
	"github.com/prnndk/final-project-golang-pbkk/entity"
	"github.com/prnndk/final-project-golang-pbkk/repository"
)

type (
	UserTicketService interface {
		UserBuyTicket(ctx context.Context, req dto.UserTicketCreateRequest, userId string) (dto.UserTicketResponse, error)
		GetUserTicket(ctx context.Context, user_id string) ([]dto.UserTicketResponse, error)
	}

	userTicketService struct {
		userTicketRepo repository.EventTicketRepository
		eventRepo      repository.EventRepository
	}
)

func NewUserTicketService(userTicketRepo repository.EventTicketRepository, eventRepo repository.EventRepository) UserTicketService {
	return &userTicketService{
		userTicketRepo: userTicketRepo,
		eventRepo:      eventRepo,
	}
}

func (s *userTicketService) UserBuyTicket(ctx context.Context, req dto.UserTicketCreateRequest, userId string) (dto.UserTicketResponse, error) {

	tx := s.eventRepo.BeginTransaction()

	defer tx.Rollback()

	event, flag_event, _ := s.eventRepo.CheckIfEventExist(ctx, tx, req.EventID)
	if !flag_event {
		return dto.UserTicketResponse{}, dto.ErrEventCannotBeFound
	}

	_, flag, _ := s.userTicketRepo.CheckUserTicket(ctx, tx, userId, req.EventID)
	if flag {
		return dto.UserTicketResponse{}, dto.ErrUserTicketAlreadyExists
	}

	if event.Pricing*req.Quantity != req.TotalPrice {
		return dto.UserTicketResponse{}, dto.ErrTotalPriceNotMatch

	}

	if event.Date.After(time.Now()) {
		return dto.UserTicketResponse{}, dto.ErrEventAlreadyClosed
	}

	if event.Quota < req.Quantity {
		return dto.UserTicketResponse{}, dto.ErrQuotaNotEnough
	}

	userTicket := entity.UserTicket{
		UserId:     userId,
		EventId:    req.EventID,
		Quantity:   req.Quantity,
		TotalPrice: req.TotalPrice,
	}

	buyTicket, err := s.userTicketRepo.UserBuyEvent(ctx, tx, userTicket)
	if err != nil {
		return dto.UserTicketResponse{}, dto.ErrBuyTicket
	}

	event.Quota = event.Quota - req.Quantity

	_, err = s.eventRepo.UpdateQuota(ctx, tx, event)
	if err != nil {
		return dto.UserTicketResponse{}, dto.ErrUpdateQuota
	}

	if err = tx.Commit().Error; err != nil {
		return dto.UserTicketResponse{}, dto.ErrDbTransactionInTicket
	}

	return dto.UserTicketResponse{
		ID:         buyTicket.ID.String(),
		UserID:     buyTicket.UserId,
		EventID:    buyTicket.EventId,
		Quantity:   buyTicket.Quantity,
		TotalPrice: buyTicket.TotalPrice,
	}, nil
}

func (s *userTicketService) GetUserTicket(ctx context.Context, user_id string) ([]dto.UserTicketResponse, error) {
	userTickets, err := s.userTicketRepo.GetUserTicketByUserId(ctx, nil, user_id)
	if err != nil {
		return []dto.UserTicketResponse{}, dto.ErrGetUserTicket
	}

	var userTicketResponse []dto.UserTicketResponse
	for _, userTicket := range userTickets {
		userTicketResponse = append(userTicketResponse, dto.UserTicketResponse{
			ID:         userTicket.ID.String(),
			UserID:     userTicket.UserId,
			EventID:    userTicket.EventId,
			Quantity:   userTicket.Quantity,
			TotalPrice: userTicket.TotalPrice,
		})
	}

	return userTicketResponse, nil
}
