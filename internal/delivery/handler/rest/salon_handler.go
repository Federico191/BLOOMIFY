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

type SalonHandler struct {
	useCase usecase.SalonUCItf
}

func NewSalonHandler(useCase usecase.SalonUCItf) *SalonHandler {
	return &SalonHandler{useCase: useCase}
}

func (bc SalonHandler) Create(ctx *gin.Context) {
	var req model.SalonRequest

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

func (bc SalonHandler) GetByCity(ctx *gin.Context) {
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
