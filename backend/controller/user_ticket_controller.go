package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prnndk/final-project-golang-pbkk/dto"
	"github.com/prnndk/final-project-golang-pbkk/service"
	"github.com/prnndk/final-project-golang-pbkk/utils"
)

type (
	UserTicketController interface {
		BuyTicket(ctx *gin.Context)
	}

	userTicketController struct {
		userTicketService service.UserTicketService
	}
)

func NewUserTicketController(uts service.UserTicketService) UserTicketController {
	return &userTicketController{
		userTicketService: uts,
	}
}

func (c *userTicketController) BuyTicket(ctx *gin.Context) {
	var ticketRequest dto.UserTicketCreateRequest

	if err := ctx.ShouldBind(&ticketRequest); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
}

func (c *userTicketController) GetUserTicket(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// result, err := c.userTicketService.GetAllUserTicketWithPagination(ctx.Request.Context(), req)
	// if err != nil {
	// 	res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_TICKET, err.Error(), nil)
	// 	ctx.JSON(http.StatusBadRequest, res)
	// 	return
	// }

	// res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER_TICKET, result)
	// ctx.JSON(http.StatusOK, res)
}
