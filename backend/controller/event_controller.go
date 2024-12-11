package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prnndk/final-project-golang-pbkk/dto"
	"github.com/prnndk/final-project-golang-pbkk/service"
	"github.com/prnndk/final-project-golang-pbkk/utils"
)

type (
	EventController interface {
		GetAllEvent(ctx *gin.Context)
		GetSingleEvent(ctx *gin.Context)
		UpdateQuotaEvent(ctx *gin.Context)
	}

	eventController struct {
		eventService service.EventService
	}
)

func NewEventController(eventService service.EventService) EventController {
	return &eventController{
		eventService: eventService,
	}
}

func (ec *eventController) GetAllEvent(ctx *gin.Context) {
	result, err := ec.eventService.GetAllEvent(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ALL_EVENT, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	response := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_ALL_EVENT,
		Data:    result,
	}

	ctx.JSON(http.StatusOK, response)
}

func (ec *eventController) GetSingleEvent(ctx *gin.Context) {
	eventId := ctx.Param("id")
	result, err := ec.eventService.GetSingleEvent(ctx.Request.Context(), eventId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_SINGLE_EVENT, err.Error(), nil)
		if err == dto.ErrEventCannotBeFound {
			ctx.JSON(http.StatusNotFound, res)
		} else {
			ctx.JSON(http.StatusInternalServerError, res)
		}
		return
	}

	response := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_SINGLE_EVENT,
		Data:    result,
	}

	ctx.JSON(http.StatusOK, response)
}

func (ec *eventController) UpdateQuotaEvent(ctx *gin.Context) {
	eventId := ctx.Param("id")
	var quotaRequest dto.QuotaResponse
	if err := ctx.ShouldBindJSON(&quotaRequest); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_QUOTA, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := ec.eventService.UpdateQuoteEvent(ctx.Request.Context(), eventId, quotaRequest.Quota)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_QUOTA, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	response := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_UPDATE_QUOTA,
		Data:    result,
	}

	ctx.JSON(http.StatusOK, response)
}
