package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/internal/usecase"
	"projectIntern/model"
	"projectIntern/pkg/response"
)

type PersonalizationHandler struct {
	Personalization usecase.PersonalizationItf
	Product         usecase.ProductUCItf
}

func NewPersonalizationHandler(personalization usecase.PersonalizationItf) *PersonalizationHandler {
	return &PersonalizationHandler{Personalization: personalization}
}

func (p PersonalizationHandler) Analyze(ctx *gin.Context) {
	user := ctx.MustGet("user").(*model.UserResponse)

	var req model.PersonalizationReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind request", err)
		return
	}

	skinProblem, err := p.Personalization.Analyze(user.ID, req)
	if err != nil {
		if errors.Is(err, errors.New("skin problem not found")) {
			response.Error(ctx, http.StatusBadRequest, "failed to analyze", err)
			return
		}
		response.Error(ctx, http.StatusInternalServerError, "failed to analyze", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully analyze", skinProblem)
}
