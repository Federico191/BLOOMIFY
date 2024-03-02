package usecase

import (
	"projectIntern/internal/repository"
	"projectIntern/pkg/jwt"
)

type UseCase struct {
	Auth AuthUseCaseItf
}

func Init(repo *repository.Repository, tokenMaker jwt.JWTMakerItf) *UseCase {
	return &UseCase{
		Auth: NewAuthUseCase(repo.User, tokenMaker),
	}
}
