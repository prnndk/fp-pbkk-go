package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prnndk/final-project-golang-pbkk/dto"
	"github.com/prnndk/final-project-golang-pbkk/service"
	"github.com/prnndk/final-project-golang-pbkk/utils"
)

type (
	PembayaranController interface {
		CreatePembayaran(ctx *gin.Context)
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

func (pc *pembayaranController) CreatePembayaran(ctx *gin.Context) {
	var pembayaran dto.PembayaranCreateRequest
	if err := ctx.ShouldBind(&pembayaran); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := pc.pembayaranService.CreatePembayaran(ctx.Request.Context(), pembayaran)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_PEMBAYARAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_PEMBAYARAN, result)
	ctx.JSON(http.StatusOK, res)

}

