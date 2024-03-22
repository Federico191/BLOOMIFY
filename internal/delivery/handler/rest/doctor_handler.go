package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"projectIntern/internal/usecase"
	"projectIntern/model"
	"projectIntern/pkg/customerrors"
	"projectIntern/pkg/response"
	"strconv"
)

type DoctorHandler struct {
	doctorUC usecase.DoctorUCItf
}

func NewDoctorHandler(doctorUC usecase.DoctorUCItf) *DoctorHandler {
	return &DoctorHandler{doctorUC: doctorUC}
}

func (d DoctorHandler) GetAll(ctx *gin.Context) {
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

	places, err := d.doctorUC.GetAll(filter, page)
	if err != nil {
		if errors.Is(err, customerrors.ErrRecordNotFound) {
			response.Error(ctx, http.StatusNotFound, "doctors not found", err)
			return
		}
		response.Error(ctx, http.StatusInternalServerError, "failed to get doctors", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get doctors", places)
}

func (d DoctorHandler) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind request", err)
		return
	}

	doctor, err := d.doctorUC.GetById(id)
	if err != nil {
		if errors.Is(err, customerrors.ErrRecordNotFound) {
			response.Error(ctx, http.StatusNotFound, "doctor not found", err)
			return
		}
		response.Error(ctx, http.StatusInternalServerError, "failed to get doctor", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get doctor", doctor)
}
