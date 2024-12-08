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
		GetUserTicket(ctx *gin.Context)
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
	userId := ctx.MustGet("user_id").(string)

	if err := ctx.ShouldBind(&ticketRequest); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userTicketService.UserBuyTicket(ctx, ticketRequest, userId)

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_BUY_TICKET, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_BUY_TICKET, result)
	ctx.JSON(http.StatusCreated, res)

}

func (c *userTicketController) GetUserTicket(ctx *gin.Context) {
	var req dto.PaginationRequest
	userId := ctx.MustGet("user_id").(string)
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userTicketService.GetUserTicket(ctx.Request.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_TICKET, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER_TICKET, result)
	ctx.JSON(http.StatusOK, res)
}
