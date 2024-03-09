package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"projectIntern/internal/usecase"
	"projectIntern/pkg/response"
	"strconv"
)

type PlaceHandler struct {
	useCase usecase.PlaceUCItf
}

func NewPlaceHandler(useCase usecase.PlaceUCItf) *PlaceHandler {
	return &PlaceHandler{useCase: useCase}
}

func (pc PlaceHandler) GetByClass(ctx *gin.Context) {
	idQuery := ctx.Query("id")

	id, err := uuid.Parse(idQuery)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind request", err)
	}

	place, err := pc.useCase.GetClasses(id)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get beauty clinic by city", err)
	}

	response.Success(ctx, http.StatusOK, "successfully get beauty clinics by city", place)
}

func (pc PlaceHandler) GetByTreatment(ctx *gin.Context) {
	idQuery := ctx.Query("id")

	id, err := uuid.Parse(idQuery)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind request", err)
	}

	place, err := pc.useCase.GetTreatment(id)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get beauty clinic by city", err)
	}

	response.Success(ctx, http.StatusOK, "successfully get beauty clinics by city", place)
}

func (pc PlaceHandler) GetAll(ctx *gin.Context) {
	pageQuery := ctx.Query("page")

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind request", err)
	}

	places, err := pc.useCase.GetAll(page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get beauty clinic by city", err)
	}

	if len(places) == 0 {
		response.Error(ctx, http.StatusNotFound, "record not found", errors.New(""))
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get beauty clinics by city", places)
}

func (pc PlaceHandler) GetByCity(ctx *gin.Context) {
	pageQuery := ctx.Query("page")
	city := ctx.Query("city")

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind request", err)
	}

	places, err := pc.useCase.GetByCity(city, page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get beauty clinic by city", err)
	}

	if len(places) == 0 {
		response.Error(ctx, http.StatusNotFound, "record not found", errors.New(""))
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get beauty clinics by city", places)
}
