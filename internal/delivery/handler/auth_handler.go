package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"projectIntern/internal/model"
	"projectIntern/internal/usecase"
)

type AuthHandler struct {
	userUC    usecase.AuthUseCaseItf
	validator *validator.Validate
}

func NewAuthHandler(userUC usecase.AuthUseCaseItf, validator *validator.Validate) *AuthHandler {
	return &AuthHandler{userUC: userUC, validator: validator}
}

func (a AuthHandler) Register(c *gin.Context) {
	var req model.UserRegister

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err := a.validator.Struct(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := a.userUC.Register(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user, "message": "Success Register user"})
}

func (a AuthHandler) Login(c *gin.Context) {
	var req model.UserLogin

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err := a.validator.Struct(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	token, err := a.userUC.Login(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": token, "Message": "Success Login"})
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
