package rest

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/internal/usecase"
	"projectIntern/model"
	"projectIntern/pkg/customerrors"
	"projectIntern/pkg/response"
)

type BookingHandler struct {
	treatment usecase.BookingTreatmentUCItf
	doctor    usecase.BookingDoctorUCItf
}

func NewBookingHandler(useCase usecase.BookingTreatmentUCItf, doctor usecase.BookingDoctorUCItf) *BookingHandler {
	return &BookingHandler{treatment: useCase, doctor: doctor}
}

func (b BookingHandler) CreateBookingTreatment(ctx *gin.Context) {
	user := ctx.MustGet("user").(*model.UserResponse)

	var req model.BookingTreatmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind request", err)
		return
	}

	booking, err := b.treatment.Create(user.ID, req)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to create booking", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "successfully create booking", booking)
}

func (b BookingHandler) CreateBookingDoctor(ctx *gin.Context) {
	user := ctx.MustGet("user").(*model.UserResponse)

	var req model.BookingDoctorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind request", err)
		return
	}

	booking, err := b.doctor.Create(user.ID, req)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to create booking", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "successfully create booking", booking)
}

func (b BookingHandler) Update(ctx *gin.Context) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&notificationPayload)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to decode payload", err)
		return
	}

	orderId, exist := notificationPayload["order_id"].(string)
	if !exist {
		response.Error(ctx, http.StatusNotFound, "failed to found order_id", errors.New("order_id not found"))
		return
	}

	doctor, err := b.doctor.GetById(orderId)
	if err != nil && !errors.Is(err, customerrors.ErrRecordNotFound) {
		response.Error(ctx, http.StatusInternalServerError, "failed to get booking", err)
		return
	}

	treatment, err := b.treatment.GetById(orderId)
	if err != nil && !errors.Is(err, customerrors.ErrRecordNotFound) {
		response.Error(ctx, http.StatusInternalServerError, "failed to get booking", err)
		return
	}

	if treatment == nil && doctor == nil {
		response.Error(ctx, http.StatusNotFound, "failed to found order_id", errors.New("order_id not found"))
		return
	}

	if treatment != nil {
		err = b.treatment.Update(orderId)
		if err != nil {
			response.Error(ctx, http.StatusInternalServerError, "failed to update order", err)
			return
		}
	} else {
		err = b.doctor.Update(orderId)
		if err != nil {
			response.Error(ctx, http.StatusInternalServerError, "failed to update order", err)
			return
		}
	}

	response.Success(ctx, http.StatusOK, "success update order", nil)
}

func (b BookingHandler) GetByStatus(ctx *gin.Context) {
	transactionId := ctx.Param("id")

	booking, err := b.treatment.GetByStatus(transactionId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get booking", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get booking", booking)
}
