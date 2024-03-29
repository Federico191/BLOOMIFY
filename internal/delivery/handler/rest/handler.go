package rest

import (
	"projectIntern/internal/usecase"
)

type Handler struct {
	User            *UserHandler
	Service         *ServiceHandler
	Booking         *BookingHandler
	Doctor          *DoctorHandler
	Personalization *PersonalizationHandler
	Product         *ProductHandler
	Dashboard       *DashboardHandler
}

func Init(useCase *usecase.UseCase) *Handler {
	return &Handler{
		User:            NewUserHandler(useCase.User),
		Service:         NewServiceHandler(useCase.Service),
		Booking:         NewBookingHandler(useCase.BookingTreatment, useCase.BookingDoctor),
		Doctor:          NewDoctorHandler(useCase.Doctor),
		Personalization: NewPersonalizationHandler(useCase.Personalization),
		Product:         NewProductHandler(useCase.Product),
		Dashboard:       NewDashboardHandler(useCase.Product, useCase.Service),
	}
}
