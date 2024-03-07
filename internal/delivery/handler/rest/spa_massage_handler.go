package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/internal/model"
	"projectIntern/internal/usecase"
	"projectIntern/pkg/response"
	"strconv"
)

type SpaMassageHandler struct {
	useCase usecase.SpaMassageUCItf
}

func NewSpaMassageHandler(useCase usecase.SpaMassageUCItf) *SpaMassageHandler {
	return &SpaMassageHandler{useCase: useCase}
}

func (bc SpaMassageHandler) Create(ctx *gin.Context) {
	var req model.SpaMassageRequest

	if err := ctx.ShouldBindJSON(req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind request", err)
		return
	}

	beautyClinic, err := bc.useCase.Create(req)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to create to beauty clinic", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully create beauty clinic", beautyClinic)
}

func (bc SpaMassageHandler) GetByCity(ctx *gin.Context) {
	pageQuery := ctx.Query("page")
	city := ctx.Query("city")

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind request", err)
	}

	beautyClinics, err := bc.useCase.GetByCity(city, page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get beauty clinic by city", err)
	}

	if len(beautyClinics) == 0 {
		response.Error(ctx, http.StatusNotFound, "record not found", errors.New(""))
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get beauty clinics by city", beautyClinics)
}
