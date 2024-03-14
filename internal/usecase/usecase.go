package usecase

import (
	"projectIntern/internal/repository"
	"projectIntern/pkg/email"
	"projectIntern/pkg/jwt"
	"projectIntern/pkg/supabase"
)

type UseCase struct {
	User    UserUCItf
	Service ServiceItf
	Review  ReviewUCItf
}

func Init(repo *repository.Repository, tokenMaker jwt.JWTMakerItf, email email.EmailItf, itf supabase.SupabaseStorageItf) *UseCase {
	return &UseCase{
		User:    NewUseUC(repo.User, tokenMaker, email, itf),
		Service: NewService(repo.Service, repo.Category),
		Review:  NewReviewUC(repo.Review),
	}
}
