package rest

import (
	"projectIntern/internal/usecase"
)

type Handler struct {
	Auth *AuthHandler
}

func Init(useCase *usecase.UseCase) *Handler {
	return &Handler{Auth: NewAuthHandler(useCase.Auth)}
}
