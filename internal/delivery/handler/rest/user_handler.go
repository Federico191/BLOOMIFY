package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"projectIntern/internal/usecase"
	"projectIntern/model"
	"projectIntern/pkg/customerrors"
	"projectIntern/pkg/encode"
	"projectIntern/pkg/jwt"
	"projectIntern/pkg/response"
)

type UserHandler struct {
	userUC usecase.UserUCItf
	jwt    jwt.JWTMakerItf
}

func NewUserHandler(userUC usecase.UserUCItf) *UserHandler {
	return &UserHandler{userUC: userUC}
}

func (a UserHandler) Register(c *gin.Context) {
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

func (a UserHandler) Login(c *gin.Context) {
	var req model.UserLogin

	log.Println("sebelum bind")
	if err := c.ShouldBind(&req); err != nil {
		log.Println("sesudah", err)
		response.Error(c, http.StatusBadRequest, "Failed to bind request", err)
		return
	}

	log.Println("sebelum token")
	token, err := a.userUC.Login(&req)
	if err != nil {
		log.Println("sesudah", err)
		if errors.Is(err, customerrors.ErrEmailInvalid) || errors.Is(err, customerrors.ErrPasswordInvalid) {
			log.Println("sesudah", err)
			response.Error(c, http.StatusUnauthorized, "failed to log in", err)
			return
		} else if errors.Is(err, customerrors.ErrNotVerified) {
			log.Println("sesudah", err)
			response.Error(c, http.StatusForbidden, "failed to log in", err)
			return
		}
		log.Println("sesudah", err)
		response.Error(c, http.StatusInternalServerError, "Failed to log in", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": token, "Message": "Success Login"})
}

func (a UserHandler) VerifyEmail(ctx *gin.Context) {
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

func (a UserHandler) UpdatePhoto(ctx *gin.Context) {
	photo, err := ctx.FormFile("photo")
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
	}

	err = a.userUC.UpdatePhoto(ctx, model.UserUploadPhoto{Photo: photo})
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to upload photo", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "success upload photo", photo)
}

func (a UserHandler) GetUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		response.Error(ctx, http.StatusNotFound, "failed to get user", nil)
	}

	response.Success(ctx, http.StatusOK, "success get user", user)
}
