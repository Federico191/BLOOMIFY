package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/internal/usecase"
	"projectIntern/model"
	"projectIntern/pkg/customerrors"
	"projectIntern/pkg/response"
)

type DashboardHandler struct {
	product usecase.ProductUCItf
	service usecase.ServiceUCItf
}

func NewDashboardHandler(product usecase.ProductUCItf, service usecase.ServiceUCItf) *DashboardHandler {
	return &DashboardHandler{product: product, service: service}
}

func (d DashboardHandler) UserDashboard(ctx *gin.Context) {
	user := ctx.MustGet("user").(*model.UserResponse)
	var products []*model.ProductResponse
	var services []*model.ServiceResponseDashboard
	var err error

	if user.ProblemId == 0 {
		products, err = d.product.GetByTopRate()
		if err != nil {
			if errors.Is(err, customerrors.ErrRecordNotFound) {
				response.Error(ctx, http.StatusNotFound, "doctors not found", err)
				return
			}
			response.Error(ctx, http.StatusInternalServerError, "failed to get products", err)
			return
		}

		services, err = d.service.GetByTopRate()
		if err != nil {
			if errors.Is(err, customerrors.ErrRecordNotFound) {
				response.Error(ctx, http.StatusNotFound, "doctors not found", err)
				return
			}
			response.Error(ctx, http.StatusInternalServerError, "failed to get services", err)
			return
		}

	} else if user.ProblemId != 0 {
		products, err = d.product.GetByProblem(user.ProblemId)
		if err != nil {
			if errors.Is(err, customerrors.ErrRecordNotFound) {
				response.Error(ctx, http.StatusNotFound, "doctors not found", err)
				return
			}
			response.Error(ctx, http.StatusInternalServerError, "failed to get products", err)
			return
		}

		services, err = d.service.GetByProblem(user.ProblemId)
		if err != nil {
			if errors.Is(err, customerrors.ErrRecordNotFound) {
				response.Error(ctx, http.StatusNotFound, "doctors not found", err)
				return
			}
			response.Error(ctx, http.StatusInternalServerError, "failed to get services", err)
			return
		}
	}

	response.Success(ctx, http.StatusOK, "success get user dashboard", model.DashboardResponse{Products: products, Services: services})

}

func (d DashboardHandler) GuestDashboard(ctx *gin.Context) {
	products, err := d.product.GetByTopRate()
	if err != nil {
		if errors.Is(err, customerrors.ErrRecordNotFound) {
			response.Error(ctx, http.StatusNotFound, "doctors not found", err)
			return
		}
		response.Error(ctx, http.StatusInternalServerError, "failed to get products", err)
		return
	}

	services, err := d.service.GetByTopRate()
	if err != nil {
		if errors.Is(err, customerrors.ErrRecordNotFound) {
			response.Error(ctx, http.StatusNotFound, "doctors not found", err)
			return
		}
		response.Error(ctx, http.StatusInternalServerError, "failed to get services", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success get guest dashboard", model.DashboardResponse{Products: products, Services: services})
}
