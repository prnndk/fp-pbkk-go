package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prnndk/final-project-golang-pbkk/dto"
	"github.com/prnndk/final-project-golang-pbkk/service"
	"github.com/prnndk/final-project-golang-pbkk/utils"
)

type (
	StorageController interface {
		UploadFile(ctx *gin.Context)
	}

	storageController struct {
		storageService service.StorageService
	}
)

func NewStorageController(ss service.StorageService) StorageController {
	return &storageController{
		storageService: ss,
	}
}

func (sc *storageController) UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	storage := dto.UploadFileRequest{
		File: file,
	}

	result, err := sc.storageService.HandleImageUpload(ctx.Request.Context(), storage)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_PEMBAYARAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_PEMBAYARAN, result)
	ctx.JSON(http.StatusOK, res)
}
