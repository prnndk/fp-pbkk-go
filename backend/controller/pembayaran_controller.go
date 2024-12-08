package controller

import "github.com/prnndk/final-project-golang-pbkk/service"

type (
	PembayaranController interface {
	}

	pembayaranController struct {
		pembayaranService service.PembayaranService
	}
)

func NewPembayaranController(ps service.PembayaranService) PembayaranController {
	return &pembayaranController{
		pembayaranService: ps,
	}
}
