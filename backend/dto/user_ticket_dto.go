package dto

import (
	"errors"

	"github.com/prnndk/final-project-golang-pbkk/entity"
)

const ()

var (
	ErrUserTicketAlreadyExists = errors.New("user already have ticket for this event")
	ErrEventCannotBeFound      = errors.New("event cannot be found")
	ErrBuyTicket               = errors.New("failed to buy ticket")
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

	// UserLoginRequest struct {
	// 	Email    string `json:"email" form:"email" binding:"required"`
	// 	Password string `json:"password" form:"password" binding:"required"`
	// }

	// UserLoginResponse struct {
	// 	Token string `json:"token"`
	// 	Role  string `json:"role"`
	// }

	// UpdateStatusIsVerifiedRequest struct {
	// 	UserId     string `json:"user_id" form:"user_id" binding:"required"`
	// 	IsVerified bool   `json:"is_verified" form:"is_verified"`
	// }
)