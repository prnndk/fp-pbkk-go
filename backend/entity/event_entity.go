package entity

import "github.com/google/uuid"

type Event struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name     string    `json:"name"`
	Date     string    `json:"date"`
	Pricing  string    `json:"pricing"`
	IsActive bool      `json:"is_active" db:"is_active"`
	Quota    int       `json:"quota"`
	TypeID   string `json:"type_id" db:"type_id"`

	Timestamp
}
