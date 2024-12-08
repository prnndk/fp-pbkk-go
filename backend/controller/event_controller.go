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
