package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/internal/usecase"
	"projectIntern/model"
	"projectIntern/pkg/response"
)

type ProductHandler struct {
	product usecase.ProductUCItf
}

func NewProductHandler(product usecase.ProductUCItf) *ProductHandler {
	return &ProductHandler{product: product}
}

func (p ProductHandler) GetByProblem(ctx *gin.Context) {
	user := ctx.MustGet("user").(*model.UserResponse)

	products, err := p.product.GetByProblem(user.ProblemId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get products", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully get products", products)
}
