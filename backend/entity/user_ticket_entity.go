package entity

import "github.com/google/uuid"

type UserTicket struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserId     string    `json:"type_id" db:"user_id"`
	EventId    string    `json:"event_id" db:"event_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice int       `json:"total_price" db:"total_price"`
}
