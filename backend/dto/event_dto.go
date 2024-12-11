package dto

import (
	"errors"
	"time"
)

const (
	MESSAGE_FAILED_GET_ALL_EVENT     = "Failed to get all event"
	MESSAGE_SUCCESS_GET_ALL_EVENT    = "Success get all event"
	MESSAGE_FAILED_GET_SINGLE_EVENT  = "Failed to get single event"
	MESSAGE_SUCCESS_GET_SINGLE_EVENT = "Success get single event"
	MESSAGE_FAILED_UPDATE_QUOTA      = "Failed to update quota"
	MESSAGE_SUCCESS_UPDATE_QUOTA     = "Success update quota"
)

var (
	ErrGettingAllEvent    = errors.New("failed to get all event")
	ErrGettingSingleEvent = errors.New("failed to get single event")
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

	EventResponseWithoutType struct {
		ID       string    `json:"id"`
		Name     string    `json:"name"`
		Date     time.Time `json:"date"`
		Pricing  int       `json:"pricing"`
		IsActive bool      `json:"is_active"`
		Quota    int       `json:"quota"`
	}
	QuotaResponse struct {
		Quota int `json:"quota"`
	}
)
