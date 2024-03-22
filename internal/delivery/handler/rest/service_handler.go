package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/internal/usecase"
	"projectIntern/model"
	"projectIntern/pkg/response"
	"strconv"
)

type ServiceHandler struct {
	useCase usecase.ServiceUCItf
}

func NewServiceHandler(useCase usecase.ServiceUCItf) *ServiceHandler {
	return &ServiceHandler{useCase: useCase}
}

func (s ServiceHandler) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind request", err)
		return
	}

	service, err := s.useCase.GetById(uint(id))
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get service", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get service", service)
}

func (s ServiceHandler) GetAllBeautyClinic(ctx *gin.Context) {
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

	places, err := s.useCase.GetAllBeautyClinic(filter, page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get beauty clinics", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get beauty clinics", places)
}

func (s ServiceHandler) GetAllSpaMassage(ctx *gin.Context) {
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

	places, err := s.useCase.GetAllSpaMassage(filter, page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get spa massages", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get spa beauties", places)
}

func (s ServiceHandler) GetAllSalon(ctx *gin.Context) {
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

	places, err := s.useCase.GetAllSalon(filter, page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get salons", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get salons", places)
}

func (s ServiceHandler) GetAllFitnessCenter(ctx *gin.Context) {
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

	places, err := s.useCase.GetAllFitnessCenter(filter, page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get fitness centers", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get fitness centers", places)
}
