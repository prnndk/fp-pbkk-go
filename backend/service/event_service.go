package service

import (
	"context"

	"github.com/prnndk/final-project-golang-pbkk/dto"
	"github.com/prnndk/final-project-golang-pbkk/repository"
)

type (
	EventService interface {
		GetAllEvent(ctx context.Context) ([]dto.GetAllEventResponse, error)
		GetSingleEvent(ctx context.Context, eventId string) (dto.GetAllEventResponse, error)
	}

	eventService struct {
		eventRepo repository.EventRepository
	}
)

func NewEventService(eventRepo repository.EventRepository) EventService {
	return &eventService{
		eventRepo: eventRepo,
	}
}

func (es *eventService) GetAllEvent(ctx context.Context) ([]dto.GetAllEventResponse, error) {

	data, err := es.eventRepo.FindAllEvent(ctx, nil)
	if err != nil {
		return []dto.GetAllEventResponse{}, dto.ErrGettingAllEvent
	}

	var eventResponse []dto.GetAllEventResponse
	for _, event := range data {
		eventResponse = append(eventResponse, dto.GetAllEventResponse{
			ID:       event.ID.String(),
			Name:     event.Name,
			Date:     event.Date,
			Pricing:  event.Pricing,
			IsActive: event.IsActive,
			Quota:    event.Quota,
			Type: dto.TypeResponse{
				ID:   event.Type.ID.String(),
				Name: event.Type.Name,
			},
		})
	}

	return eventResponse, nil
}

func (es *eventService) GetSingleEvent(ctx context.Context, eventId string) (dto.GetAllEventResponse, error) {
	event, err := es.eventRepo.FindEventById(ctx, nil, eventId)
	if err != nil {
		return dto.GetAllEventResponse{}, dto.ErrEventCannotBeFound
	}

	return dto.GetAllEventResponse{
		ID:       event.ID.String(),
		Name:     event.Name,
		Date:     event.Date,
		Pricing:  event.Pricing,
		IsActive: event.IsActive,
		Quota:    event.Quota,
		Type: dto.TypeResponse{
			ID:   event.Type.ID.String(),
			Name: event.Type.Name,
		},
	}, nil
}
