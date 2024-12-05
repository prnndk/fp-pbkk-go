package repository

import (
	"context"

	"github.com/prnndk/final-project-golang-pbkk/entity"
	"gorm.io/gorm"
)

type (
	EventRepository interface {
		FindEventById(ctx context.Context, tx *gorm.DB, eventId string) (entity.Event, error)
		CheckIfEventExist(ctx context.Context, tx *gorm.DB, eventId string) (entity.Event, bool, error)
	}
	eventRepository struct {
		db *gorm.DB
	}
)

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{
		db: db,
	}
}

func (r *eventRepository) FindEventById(ctx context.Context, tx *gorm.DB, eventId string) (entity.Event, error) {
	if tx == nil {
		tx = r.db
	}

	var event entity.Event
	if err := tx.WithContext(ctx).Where("id = ?", eventId).First(&event).Error; err != nil {
		return entity.Event{}, err
	}

	return event, nil
}

func (r *eventRepository) CheckIfEventExist(ctx context.Context, tx *gorm.DB, eventId string) (entity.Event, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var event entity.Event
	if err := tx.WithContext(ctx).Where("id = ?", eventId).First(&event).Error; err != nil {
		return entity.Event{}, false, err
	}

	return event, true, nil
}
