package usecase

import (
	"projectIntern/internal/repository"
	"projectIntern/internal/util/token"
)

type UseCase struct {
	Auth AuthUseCaseItf
}

func Init(repo *repository.Repository, tokenMaker token.Maker) *UseCase {
	return &UseCase{
		Auth: NewAuthUseCase(repo.User, tokenMaker),
	}
}
