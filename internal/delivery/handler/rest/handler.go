package rest

import (
	"projectIntern/internal/usecase"
)

type Handler struct {
	User   *UserHandler
	Place  *PlaceHandler
	Review *ReviewHandler
}

func Init(useCase *usecase.UseCase) *Handler {
	return &Handler{
		User:   NewUserHandler(useCase.User),
		Place:  NewPlaceHandler(useCase.Place),
		Review: NewReviewHandler(useCase.Review),
	}
}
