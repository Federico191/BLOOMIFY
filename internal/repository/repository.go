package repository

import "gorm.io/gorm"

type Repository struct {
	User          UserRepoItf
	BeautyClinic  BeautyClinicRepoItf
	SalonRepo     SalonRepoItf
	SpaMassage    SpaMassageRepoItf
	FitnessCenter FitnessCenterRepoItf
}

func Init(db *gorm.DB) *Repository {
	return &Repository{
		User:          NewUserRepo(db),
		BeautyClinic:  NewBeautyClinicRepo(db),
		SalonRepo:     NewSalonRepo(db),
		SpaMassage:    NewSpaMassageRepo(db),
		FitnessCenter: NewFitnessCenterRepo(db),
	}
}
