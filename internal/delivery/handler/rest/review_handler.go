package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/internal/usecase"
	"projectIntern/model"
	"projectIntern/pkg/response"
)

type ReviewHandler struct {
	ReviewUC usecase.ReviewUCItf
}

func NewReviewHandler(ReviewUC usecase.ReviewUCItf) *ReviewHandler {
	return &ReviewHandler{ReviewUC: ReviewUC}
}

func (r ReviewHandler) Create(ctx *gin.Context) {
	var req model.ReviewRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind request", err)
		return
	}

	review, err := r.ReviewUC.Create(req)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get review", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "successfully create review", review)

}
