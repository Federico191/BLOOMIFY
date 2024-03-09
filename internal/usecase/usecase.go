package usecase

import (
	"projectIntern/internal/repository"
	"projectIntern/pkg/email"
	"projectIntern/pkg/jwt"
)

type UseCase struct {
	Auth    AuthUseCaseItf
	User    UserUCItf
	Place   PlaceUCItf
	Service ServiceItf
	Class   ClassUCItf
	Review  ReviewUCItf
}

func Init(repo *repository.Repository, tokenMaker jwt.JWTMakerItf, email email.EmailItf) *UseCase {
	return &UseCase{
		Auth:    NewAuthUseCase(repo.User, tokenMaker, email),
		User:    NewUseUC(repo.User),
		Place:   NewPlaceUC(repo.Place),
		Service: NewService(repo.Service),
		Class:   NewClassUc(repo.Class),
		Review:  NewReviewUC(repo.Review),
	}
}
