package rest

import (
	"projectIntern/internal/usecase"
)

type Handler struct {
	Auth   *AuthHandler
	Place  *PlaceHandler
	Review *ReviewHandler
}

func Init(useCase *usecase.UseCase) *Handler {
	return &Handler{
		Auth:   NewAuthHandler(useCase.User),
		Place:  NewPlaceHandler(useCase.Place),
		Review: NewReviewHandler(useCase.Review),
	}
}
