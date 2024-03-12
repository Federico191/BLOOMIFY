package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/internal/usecase"
	"projectIntern/model"
	"projectIntern/pkg/response"
	"strconv"
)

type PlaceHandler struct {
	useCase usecase.PlaceUCItf
}

func NewPlaceHandler(useCase usecase.PlaceUCItf) *PlaceHandler {
	return &PlaceHandler{useCase: useCase}
}

func (p PlaceHandler) GetAllBeautyClinic(ctx *gin.Context) {
	var filter model.FilterParam

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind query parameters", err)
		return
	}

	if err := ctx.ShouldBind(&filter); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind form data", err)
		return
	}

	pageQuery := ctx.Query("page")
	page, _ := strconv.Atoi(pageQuery)

	places, err := p.useCase.GetAllBeautyClinic(filter, page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get place by city", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get beauty clinic by city", places)
}

func (p PlaceHandler) GetAllSpaMassage(ctx *gin.Context) {
	var filter model.FilterParam

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind query parameters", err)
		return
	}

	if err := ctx.ShouldBind(&filter); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind form data", err)
		return
	}

	pageQuery := ctx.Query("page")
	page, _ := strconv.Atoi(pageQuery)

	places, err := p.useCase.GetAllSpaMassage(filter, page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get spa massage by city", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get spa massage by city", places)
}

func (p PlaceHandler) GetAllSalon(ctx *gin.Context) {
	var filter model.FilterParam

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind query parameters", err)
		return
	}

	if err := ctx.ShouldBind(&filter); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind form data", err)
		return
	}

	pageQuery := ctx.Query("page")
	page, _ := strconv.Atoi(pageQuery)

	places, err := p.useCase.GetAllSalon(filter, page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get salon by city", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get salon by city", places)
}

func (p PlaceHandler) GetAllFitnessCenter(ctx *gin.Context) {
	var filter model.FilterParam

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind query parameters", err)
		return
	}

	if err := ctx.ShouldBind(&filter); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind form data", err)
		return
	}

	pageQuery := ctx.Query("page")
	page, _ := strconv.Atoi(pageQuery)

	places, err := p.useCase.GetAllFitnessCenter(filter, page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get fitness center by city", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get fitness center by city", places)
}
