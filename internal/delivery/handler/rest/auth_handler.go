package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/internal/model"
	"projectIntern/internal/usecase"
	"projectIntern/pkg/customerrors"
	"projectIntern/pkg/encode"
	"projectIntern/pkg/response"
)

type AuthHandler struct {
	authUC usecase.AuthUseCaseItf
}

func NewAuthHandler(authUC usecase.AuthUseCaseItf) *AuthHandler {
	return &AuthHandler{authUC: authUC}
}

func (a AuthHandler) Register(c *gin.Context) {
	var req model.UserRegister

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind request", err)
		return
	}

	user, err := a.authUC.Register(&req)
	if err != nil {
		if errors.Is(err, customerrors.ErrEmailAlreadyExists) {
			response.Error(c, http.StatusConflict, "failed to create user", err)
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	response.Success(c, http.StatusCreated, "Success Register user", user)
}

func (a AuthHandler) Login(c *gin.Context) {
	var req model.UserLogin

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind request", err)
		return
	}

	token, err := a.authUC.Login(&req)
	if err != nil {
		if errors.Is(err, customerrors.ErrEmailInvalid) || errors.Is(err, customerrors.ErrPasswordInvalid) {
			response.Error(c, http.StatusUnauthorized, "failed to log in", err)
			return
		} else if errors.Is(err, customerrors.ErrNotVerified) {
			response.Error(c, http.StatusForbidden, "failed to log in", err)
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to log in", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": token, "Message": "Success Login"})
}

func (a AuthHandler) VerifyEmail(ctx *gin.Context) {
	codeQuery := ctx.Param("code")

	code, err := encode.Decode(codeQuery)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind request", err)
		return
	}

	user, err := a.authUC.GetByVerificationCode(code)
	if err != nil {
		return
	}

	if user.VerificationCode != code {
		response.Error(ctx, http.StatusBadRequest, "failed to verify", errors.New("verification code not same"))
		return
	}

	user.IsVerified = true

	err = a.authUC.VerifyEmail(user.ID)
	if err != nil {
		return
	}

	response.Success(ctx, http.StatusOK, "Successfully verification account", nil)
}
