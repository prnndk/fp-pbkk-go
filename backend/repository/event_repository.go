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
		FindEventByType(ctx context.Context, tx *gorm.DB, typeId string) ([]entity.Event, error)
		FindAllEvent(ctx context.Context, tx *gorm.DB) ([]entity.Event, error)
		UpdateQuota(ctx context.Context, tx *gorm.DB, event entity.Event) (entity.Event, error)
		BeginTransaction() *gorm.DB
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

func (r *eventRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
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

func (r *eventRepository) FindEventByType(ctx context.Context, tx *gorm.DB, typeId string) ([]entity.Event, error) {
	if tx == nil {
		tx = r.db
	}

	var event []entity.Event
	if err := tx.WithContext(ctx).Where("type_id = ?", typeId).Find(&event).Error; err != nil {
		return []entity.Event{}, err
	}

	return event, nil
}

func (r *eventRepository) FindAllEvent(ctx context.Context, tx *gorm.DB) ([]entity.Event, error) {
	if tx == nil {
		tx = r.db
	}

	var events []entity.Event
	if err := tx.WithContext(ctx).Preload("Type").Find(&events).Error; err != nil {
		return []entity.Event{}, err
	}

	return events, nil
}

func (r *eventRepository) UpdateQuota(ctx context.Context, tx *gorm.DB, event entity.Event) (entity.Event, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Model(&event).Update("quota", event.Quota).Error; err != nil {
		return entity.Event{}, err
	}

	return event, nil
}
