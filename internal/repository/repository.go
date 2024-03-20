package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	User             UserRepoItf
	Service          ServiceRepoItf
	TreatmentReview  TreatmentReviewRepoItf
	DoctorReview     DoctorReviewRepoItf
	Category         CategoryRepoItf
	Place            PlaceRepoItf
	BookingTreatment BookingTreatmentRepoItf
	BookingDoctor    BookingDoctorRepoItf
	Doctor           DoctorRepoItf
}

func Init(db *gorm.DB) *Repository {
	return &Repository{
		User:             NewUserRepo(db),
		Service:          NewServiceRepo(db),
		TreatmentReview:  NewTreatmentReviewRepo(db),
		DoctorReview:     NewDoctorReviewRepo(db),
		Category:         NewCategoryRepo(db),
		Place:            NewPlaceRepo(db),
		BookingTreatment: NewBookingTreatmentRepo(db),
		BookingDoctor:    NewBookingDoctorRepo(db),
		Doctor:           NewDoctorRepo(db),
	}
}
