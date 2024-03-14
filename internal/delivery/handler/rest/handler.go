package rest

import (
	"projectIntern/internal/usecase"
)

type Handler struct {
	User    *UserHandler
	Review  *ReviewHandler
	Service *ServiceHandler
}

func Init(useCase *usecase.UseCase) *Handler {
	return &Handler{
		User:    NewUserHandler(useCase.User),
		Review:  NewReviewHandler(useCase.Review),
		Service: NewServiceHandler(useCase.Service),
	}
}
