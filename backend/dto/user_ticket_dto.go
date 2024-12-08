package dto

import (
	"errors"

	"github.com/prnndk/final-project-golang-pbkk/entity"
)

const (
	MESSAGE_FAILED_GET_USER_TICKET  = "failed get user ticket"
	MESSAGE_SUCCESS_GET_USER_TICKET = "success get user ticket"
	MESSAGE_FAILED_BUY_TICKET       = "failed to buy ticket"
	MESSAGE_SUCCESS_BUY_TICKET      = "success buying ticket"
)

var (
	ErrUserTicketAlreadyExists = errors.New("user already have ticket for this event")
	ErrEventCannotBeFound      = errors.New("event cannot be found")
	ErrIdTicketNotFound              = errors.New("ticket id not found")
	ErrBuyTicket               = errors.New("failed to buy ticket")
	ErrGetUserTicket           = errors.New("failed to get user ticket")
	ErrTotalPriceNotMatch      = errors.New("total price not match")
	ErrEventAlreadyClosed      = errors.New("event already closed")
	ErrQuotaNotEnough          = errors.New("quota not enough")
	ErrUpdateQuota             = errors.New("failed to update quota")
	ErrDbTransactionInTicket   = errors.New("error in transaction")
)

type (
	UserTicketCreateRequest struct {
		EventID    string `json:"event_id" form:"event_id" binding:"required"`
		Quantity   int    `json:"quantity" form:"quantity" binding:"required"`
		TotalPrice int    `json:"total_price" form:"total_price" binding:"required"`
	}

	UserTicketResponse struct {
		ID         string `json:"id"`
		EventID    string `json:"event_id"`
		UserID     string `json:"user_id"`
		Quantity   int    `json:"quantity"`
		TotalPrice int    `json:"total_price"`
	}

	UserTicketPaginationResponse struct {
		Data []UserResponse `json:"data"`
		PaginationResponse
	}

	GetAllUserTicketRepositoryResponse struct {
		Users []entity.User
		PaginationResponse
	}

	UserTicketUpdateRequest struct {
		Name        string `json:"name" form:"name"`
		PhoneNumber string `json:"phone_number" form:"phone_number"`
		Email       string `json:"email" form:"email"`
	}

	UserTicketUpdateResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
		Email       string `json:"email"`
	}
)
