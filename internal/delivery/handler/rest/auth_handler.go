package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"projectIntern/internal/usecase"
	"projectIntern/model"
	"projectIntern/pkg/customerrors"
	"projectIntern/pkg/encode"
	"projectIntern/pkg/response"
)

type AuthHandler struct {
	userUC usecase.UserUCItf
}

func NewAuthHandler(userUC usecase.UserUCItf) *AuthHandler {
	return &AuthHandler{userUC: userUC}
}

func (a AuthHandler) Register(c *gin.Context) {
	var req model.UserRegister

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind request", err)
		return
	}

	user, err := a.userUC.Register(&req)
	if err != nil {
		if errors.Is(err, customerrors.ErrEmailAlreadyExists) || errors.Is(err, gorm.ErrDuplicatedKey) {
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

	token, err := a.userUC.Login(&req)
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

	user, err := a.userUC.GetByVerificationCode(code)
	if err != nil {
		return
	}

	if user.VerificationCode != code {
		response.Error(ctx, http.StatusBadRequest, "failed to verify", errors.New("verification code not same"))
		return
	}

	user.IsVerified = true

	err = a.userUC.VerifyEmail(user.ID)
	if err != nil {
		return
	}

	response.Success(ctx, http.StatusOK, "Successfully verification account", nil)
}

func (a AuthHandler) GetUser(ctx *gin.Context) {
	user, ok := ctx.Get("userId")
	if !ok {
		response.Error(ctx, http.StatusNotFound, "failed to get user", nil)
	}

	response.Success(ctx, http.StatusOK, "success get user", user)
}
