package usecase

import (
	"projectIntern/internal/repository"
	"projectIntern/pkg"
	"projectIntern/pkg/email"
	"projectIntern/pkg/jwt"
)

type UseCase struct {
	User    UserUCItf
	Place   PlaceUCItf
	Service ServiceItf
	Review  ReviewUCItf
}

func Init(repo *repository.Repository, tokenMaker jwt.JWTMakerItf, email email.EmailItf, itf pkg.SupabaseStorageItf) *UseCase {
	return &UseCase{
		User:    NewUseUC(repo.User, tokenMaker, email, itf),
		Place:   NewPlaceUC(repo.Place),
		Service: NewService(repo.Service),
		Review:  NewReviewUC(repo.Review),
	}
}
