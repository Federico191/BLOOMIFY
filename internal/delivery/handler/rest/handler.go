package rest

import (
	"projectIntern/internal/usecase"
)

type Handler struct {
	Auth  *AuthHandler
	Place *PlaceHandler
}

func Init(useCase *usecase.UseCase) *Handler {
	return &Handler{
		Auth:  NewAuthHandler(useCase.Auth),
		Place: NewPlaceHandler(useCase.Place),
	}
}
