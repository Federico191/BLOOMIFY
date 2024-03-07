package usecase

import (
	"projectIntern/internal/repository"
	"projectIntern/pkg/jwt"
)

type UseCase struct {
	Auth          AuthUseCaseItf
	User          UserUCItf
	BeautyClinic  BeautyClinicUCItf
	Salon         SalonUCItf
	SpaMassage    SpaMassageUCItf
	FitnessCenter FitnessCenterUCItf
}

func Init(repo *repository.Repository, tokenMaker jwt.JWTMakerItf) *UseCase {
	return &UseCase{
		Auth:          NewAuthUseCase(repo.User, tokenMaker),
		BeautyClinic:  NewBeautyClinicUC(repo.BeautyClinic),
		Salon:         NewSalonUC(repo.SalonRepo),
		SpaMassage:    NewSpaMassageUC(repo.SpaMassage),
		User:          NewUseUC(repo.User),
		FitnessCenter: NewFitnessCenterUc(repo.FitnessCenter),
	}
}
