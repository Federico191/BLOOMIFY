package handler

import (
	"github.com/go-playground/validator/v10"
	"projectIntern/internal/usecase"
)

type Handler struct {
	Auth *AuthHandler
}

func Init(useCase *usecase.UseCase, validator *validator.Validate) *Handler {
	return &Handler{Auth: NewAuthHandler(useCase.Auth, validator)}
}
