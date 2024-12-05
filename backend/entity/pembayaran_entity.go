package entity

import "github.com/google/uuid"

type Pembayaran struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	TicketId         string    `json:"ticket_id" db:"ticket_id"`
	MetodePembayaran string    `json:"metode_pembayaran" db:"metode_pembayaran"`
	BuktiBayar       string    `json:"bukti_bayar" db:"bukti_bayar"`
	IsVerified       bool      `json:"is_verified" db:"is_verified"`

	Timestamp
}
