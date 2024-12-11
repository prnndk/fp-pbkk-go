package entity

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Pricing  int       `json:"pricing"`
	IsActive bool      `json:"is_active" db:"is_active"`
	Quota    int       `json:"quota"`
	TypeID   string    `json:"type_id" db:"type_id"`
	Type     Type      `json:"type" gorm:"references:ID"`

	Timestamp
}

type EventWithTypeName struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Pricing  int       `json:"pricing"`
	IsActive bool      `json:"is_active" db:"is_active"`
	Quota    int       `json:"quota"`
	TypeID   string    `json:"type_id" db:"type_id"`
	Type     Type      `json:"type" gorm:"references:ID"`

	Timestamp
}

type QuotaEvent struct {
	Quota int `json:"quota"`
}
