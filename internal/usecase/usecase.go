package usecase

import (
	"projectIntern/internal/repository"
	"projectIntern/pkg/email"
	"projectIntern/pkg/jwt"
)

type UseCase struct {
	User    UserUCItf
	Place   PlaceUCItf
	Service ServiceItf
	Class   ClassUCItf
	Review  ReviewUCItf
}

func Init(repo *repository.Repository, tokenMaker jwt.JWTMakerItf, email email.EmailItf) *UseCase {
	return &UseCase{
		User:    NewUseUC(repo.User, tokenMaker, email),
		Place:   NewPlaceUC(repo.Place),
		Service: NewService(repo.Service),
		Class:   NewClassUc(repo.Class),
		Review:  NewReviewUC(repo.Review),
	}
}
