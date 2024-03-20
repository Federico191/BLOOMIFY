package usecase

import (
	"projectIntern/internal/repository"
	"projectIntern/pkg/email"
	"projectIntern/pkg/jwt"
	"projectIntern/pkg/midtrans"
	"projectIntern/pkg/supabase"
)

type UseCase struct {
	User             UserUCItf
	Service          ServiceItf
	Review           TreatmentReviewUCItf
	BookingTreatment BookingTreatmentUCItf
	BookingDoctor    BookingDoctorUCItf
	Doctor           DoctorUCItf
}

func Init(repo *repository.Repository, tokenMaker jwt.JWTMakerItf, email email.EmailItf, itf supabase.SupabaseStorageItf, transItf midtrans.MidTransItf) *UseCase {
	return &UseCase{
		User:             NewUseUC(repo.User, tokenMaker, email, itf),
		Service:          NewService(repo.Service, repo.Category),
		Review:           NewReviewUC(repo.TreatmentReview),
		BookingTreatment: NewBookingTreatmentUC(repo.BookingTreatment, repo.User, repo.Service, transItf),
		BookingDoctor:    NewBookingDoctorUC(repo.BookingDoctor, repo.User, repo.Doctor, transItf),
		Doctor:           NewDoctorUC(repo.Doctor),
	}
}
