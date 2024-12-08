package dto

import "errors"

const (
	MESSAGE_FAILED_CREATE_PEMBAYARAN  = "Failed to create pembayaran"
	MESSAGE_SUCCESS_CREATE_PEMBAYARAN = "Success create pembayaran"
)

var (
	ErrCreatePembayaran = errors.New("failed to create pembayaran")
)

type (
	PembayaranCreateRequest struct {
		TicketID         string `json:"ticket_id" form:"ticket_id" binding:"required"`
		MetodePembayaran string `json:"metode_pembayaran" form:"metode_pembayaran" binding:"required"`
		BuktiBayar       string `json:"bukti_bayar" form:"bukti_bayar" binding:"required"`
	}

	PembayaranCreateResponse struct {
		ID               string `json:"id"`
		TicketID         string `json:"ticket_id"`
		MetodePembayaran string `json:"metode_pembayaran"`
		BuktiBayar       string `json:"bukti_bayar"`
	}
)
