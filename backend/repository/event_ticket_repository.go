package repository

import (
	"context"

	"github.com/prnndk/final-project-golang-pbkk/entity"
	"gorm.io/gorm"
)

type (
	EventTicketRepository interface {
		UserBuyEvent(ctx context.Context, tx *gorm.DB, userTicket entity.UserTicket) (entity.UserTicket, error)
		GetUserTicketByUserId(ctx context.Context, tx *gorm.DB, userId string) ([]entity.UserTicket, error)
		CheckUserTicket(ctx context.Context, tx *gorm.DB, userId string, eventId string) (entity.UserTicket, bool, error)
		CheckUserTicketById(ctx context.Context, tx *gorm.DB, userTicketId string) (entity.UserTicket, bool, error)
		GetTicketById(ctx context.Context, tx *gorm.DB, userTicketId string) (entity.UserTicket, error)
	}
	eventTicketRepository struct {
		db *gorm.DB
	}
)

func NewEventTicketRepository(db *gorm.DB) EventTicketRepository {
	return &eventTicketRepository{
		db: db,
	}
}

func (r *eventTicketRepository) UserBuyEvent(ctx context.Context, tx *gorm.DB, userTicket entity.UserTicket) (entity.UserTicket, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&userTicket).Error; err != nil {
		return entity.UserTicket{}, err
	}

	return userTicket, nil
}

func (r *eventTicketRepository) GetUserTicketByUserId(ctx context.Context, tx *gorm.DB, userId string) ([]entity.UserTicket, error) {
	if tx == nil {
		tx = r.db
	}

	var userTickets []entity.UserTicket
	if err := tx.WithContext(ctx).Where("user_id = ?", userId).Preload("Event").Preload("User").Find(&userTickets).Error; err != nil {
		return []entity.UserTicket{}, err
	}

	return userTickets, nil
}

func (r *eventTicketRepository) CheckUserTicket(ctx context.Context, tx *gorm.DB, userId string, eventId string) (entity.UserTicket, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var userTicket entity.UserTicket
	if err := tx.WithContext(ctx).Where("user_id = ? AND event_id = ?", userId, eventId).First(&userTicket).Error; err != nil {
		return entity.UserTicket{}, false, err
	}

	return userTicket, true, nil
}

func (r *eventTicketRepository) CheckUserTicketById(ctx context.Context, tx *gorm.DB, userTicketId string) (entity.UserTicket, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var userTicket entity.UserTicket
	if err := tx.WithContext(ctx).Where("id = ?", userTicketId).First(&userTicket).Error; err != nil {
		return entity.UserTicket{}, false, err
	}

	return userTicket, true, nil
}

func (r *eventTicketRepository) GetTicketById(ctx context.Context, tx *gorm.DB, userTicketId string) (entity.UserTicket, error) {
	if tx == nil {
		tx = r.db
	}

	var userTicket entity.UserTicket
	if err := tx.WithContext(ctx).Where("id = ?", userTicketId).Preload("Event").Preload("User").First(&userTicket).Error; err != nil {
		return entity.UserTicket{}, err
	}

	return userTicket, nil
}
