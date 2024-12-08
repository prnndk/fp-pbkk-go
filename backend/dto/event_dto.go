package dto

import (
	"errors"
	"time"
)

const (
	MESSAGE_FAILED_GET_ALL_EVENT  = "Failed to get all event"
	MESSAGE_SUCCESS_GET_ALL_EVENT = "Success get all event"
)

var (
	ErrGettingAllEvent = errors.New("failed to get all event")
)

type (
	GetAllEventResponse struct {
		ID       string       `json:"id"`
		Name     string       `json:"name"`
		Date     time.Time    `json:"date"`
		Pricing  int          `json:"pricing"`
		IsActive bool         `json:"is_active"`
		Quota    int          `json:"quota"`
		Type     TypeResponse `json:"type"`
	}
)
