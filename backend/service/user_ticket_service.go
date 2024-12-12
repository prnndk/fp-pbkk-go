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
		GetUserTicketById(ctx context.Context, user_ticket_id string) (dto.UserTicketResponse, error)
		DeleteUserTicket(ctx context.Context, user_ticket_id string) error
	}

	userTicketService struct {
		userTicketRepo repository.EventTicketRepository
		eventRepo      repository.EventRepository
		paymentRepo    repository.PembayaranRepository
	}
)

func NewUserTicketService(userTicketRepo repository.EventTicketRepository, eventRepo repository.EventRepository, paymentRepo repository.PembayaranRepository) UserTicketService {
	return &userTicketService{
		userTicketRepo: userTicketRepo,
		eventRepo:      eventRepo,
		paymentRepo:    paymentRepo,
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
		_, check, _ := s.paymentRepo.CheckPembayaranByTicketId(ctx, nil, userTicket.ID.String())
		dataUser := dto.UserResponse{
			ID:          userTicket.User.ID.String(),
			Name:        userTicket.User.Name,
			PhoneNumber: userTicket.User.PhoneNumber,
			Email:       userTicket.User.Email,
			Role:        userTicket.User.Role,
		}
		dataEvent := dto.EventResponseWithoutType{
			ID:       userTicket.Event.ID.String(),
			Name:     userTicket.Event.Name,
			Date:     userTicket.Event.Date,
			Pricing:  userTicket.Event.Pricing,
			IsActive: userTicket.Event.IsActive,
			Quota:    userTicket.Event.Quota,
		}
		userTicketResponse = append(userTicketResponse, dto.UserTicketResponse{
			ID:         userTicket.ID.String(),
			UserID:     userTicket.UserId,
			User:       dataUser,
			EventID:    userTicket.EventId,
			Event:      dataEvent,
			Quantity:   userTicket.Quantity,
			TotalPrice: userTicket.TotalPrice,
			IsPaid:     check,
		})
	}

	return userTicketResponse, nil
}

func (s *userTicketService) GetUserTicketById(ctx context.Context, user_ticket_id string) (dto.UserTicketResponse, error) {
	userTicket, err := s.userTicketRepo.GetTicketById(ctx, nil, user_ticket_id)
	if err != nil {
		return dto.UserTicketResponse{}, dto.ErrIdTicketNotFound
	}

	_, check, _ := s.paymentRepo.CheckPembayaranByTicketId(ctx, nil, userTicket.ID.String())

	dataUser := dto.UserResponse{
		ID:          userTicket.User.ID.String(),
		Name:        userTicket.User.Name,
		PhoneNumber: userTicket.User.PhoneNumber,
		Email:       userTicket.User.Email,
		Role:        userTicket.User.Role,
	}
	dataEvent := dto.EventResponseWithoutType{
		ID:       userTicket.Event.ID.String(),
		Name:     userTicket.Event.Name,
		Date:     userTicket.Event.Date,
		Pricing:  userTicket.Event.Pricing,
		IsActive: userTicket.Event.IsActive,
		Quota:    userTicket.Event.Quota,
	}
	return dto.UserTicketResponse{
		ID:         userTicket.ID.String(),
		UserID:     userTicket.UserId,
		User:       dataUser,
		EventID:    userTicket.EventId,
		Event:      dataEvent,
		Quantity:   userTicket.Quantity,
		TotalPrice: userTicket.TotalPrice,
		IsPaid:     check,
	}, nil
}

func (s *userTicketService) DeleteUserTicket(ctx context.Context, user_ticket_id string) error {
	tx := s.eventRepo.BeginTransaction()
	defer tx.Rollback()

	userTicket, err := s.userTicketRepo.GetTicketById(ctx, tx, user_ticket_id)
	if err != nil {
		return dto.ErrIdTicketNotFound
	}

	err = s.userTicketRepo.DeleteUserTicket(ctx, tx, user_ticket_id)
	if err != nil {
		return dto.ErrDeleteUserTicket
	}

	event, _, _ := s.eventRepo.CheckIfEventExist(ctx, tx, userTicket.EventId)
	event.Quota += userTicket.Quantity

	_, err = s.eventRepo.UpdateQuota(ctx, tx, event)
	if err != nil {
		return dto.ErrUpdateQuota
	}

	if err = tx.Commit().Error; err != nil {
		return dto.ErrDbTransactionInTicket
	}

	return nil
}
