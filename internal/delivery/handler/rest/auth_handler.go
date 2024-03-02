package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/internal/model"
	"projectIntern/internal/usecase"
)

type AuthHandler struct {
	userUC usecase.AuthUseCaseItf
}

func NewAuthHandler(userUC usecase.AuthUseCaseItf) *AuthHandler {
	return &AuthHandler{userUC: userUC}
}

func (a AuthHandler) Register(c *gin.Context) {
	var req model.UserRegister

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
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
