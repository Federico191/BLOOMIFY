package usecase

import (
	"errors"
	"github.com/google/uuid"
	"log"
	"os"
	"projectIntern/internal/entity"
	"projectIntern/internal/repository"
	"projectIntern/model"
	"projectIntern/pkg/helper"
	"projectIntern/pkg/midtrans"
)

type BookingTreatmentUCItf interface {
	Create(id uuid.UUID, req model.BookingTreatmentRequest) (*model.BookingTreatmentResponse, error)
	Update(bookingId string) error
	GetByStatus(transactionId string) (*model.BookingTreatmentResponse, error)
	GetById(id string) (*entity.BookingTreatment, error)
}

type BookingTreatmentUC struct {
	repo    repository.BookingTreatmentRepoItf
	user    repository.UserRepoItf
	service repository.ServiceRepoItf
	mdtrans midtrans.MidTransItf
}

func NewBookingTreatmentUC(repo repository.BookingTreatmentRepoItf,
	user repository.UserRepoItf,
	service repository.ServiceRepoItf,
	mdtrans midtrans.MidTransItf) BookingTreatmentUCItf {
	return &BookingTreatmentUC{
		repo:    repo,
		user:    user,
		service: service,
		mdtrans: mdtrans,
	}
}

func (b BookingTreatmentUC) Create(id uuid.UUID, req model.BookingTreatmentRequest) (*model.BookingTreatmentResponse, error) {
	user, err := b.user.GetById(id)
	if err != nil {
		return nil, err
	}

	service, err := b.service.GetById(req.ServiceId)
	if err != nil {
		return nil, err
	}

	bookingId := uuid.NewString()

	transaction, err := b.mdtrans.CreateTransactionService(*service, req.PaymentMethod, bookingId, int64(service.Price), *user)
	if err != nil {
		return nil, err
	}
	log.Println(transaction.TransactionID)

	timeParse, err := helper.ParseDateTime(req.Day, req.Time)
	if err != nil {
		return nil, err
	}
	booking := &entity.BookingTreatment{
		ID:            bookingId,
		UserId:        user.ID,
		User:          *user,
		ServiceId:     service.ID,
		Service:       *service,
		PaymentMethod: req.PaymentMethod,
		TransactionId: transaction.TransactionID,
		GrossAmount:   int64(service.Price),
		Status:        "Pending",
		PaymentCode:   transaction.VaNumbers[0].VANumber,
		BookAt:        timeParse,
	}
	data, err := b.repo.GetByTimeId(service.ID, timeParse)
	if data != nil {
		return nil, errors.New("time already booked")
	}
	err = b.repo.Create(booking)
	if err != nil {
		return nil, err
	}

	response := &model.BookingTreatmentResponse{
		BookingId:     booking.ID,
		TransactionId: booking.TransactionId,
		UserFullName:  user.FullName,
		UserEmail:     user.Email,
		PlaceName:     service.Place.Name,
		ServiceName:   service.Name,
		PlaceAddress:  service.Place.Address,
		PaymentMethod: booking.PaymentMethod,
		PaymentCode:   booking.PaymentCode,
		GrossAmount:   booking.GrossAmount,
		Status:        booking.Status,
		BookAt:        booking.BookAt,
	}

	return response, nil
}

func (b BookingTreatmentUC) Update(bookingId string) error {
	checkResponse, err := b.mdtrans.Notification(bookingId)
	if err != nil {
		return err
	}

	booking, err := b.repo.GetById(bookingId)
	if err != nil {
		return err
	}

	mySignature := helper.Hash512(checkResponse.OrderID + checkResponse.StatusCode + checkResponse.GrossAmount + os.Getenv("SERVER_KEY"))

	if checkResponse.SignatureKey != mySignature {
		booking.Status = "Failure"
		return errors.New("signature not valid")
	}

	if checkResponse != nil {
		if checkResponse.TransactionStatus == "settlement" {
			booking.Status = "Success"
		} else if checkResponse.TransactionStatus == "cancel" || checkResponse.TransactionStatus == "expire" {
			booking.Status = "Failure"
		} else if checkResponse.TransactionStatus == "pending" {
			booking.Status = "Pending"
		}
	}

	if booking.Status == "Failure" {
		err = b.repo.Delete(booking.ID)
		if err != nil {
			return err
		}
	}

	err = b.repo.Update(booking)
	if err != nil {
		return err
	}

	return nil
}

func (b BookingTreatmentUC) GetByStatus(transactionId string) (*model.BookingTreatmentResponse, error) {
	booking, err := b.repo.GetByStatus(transactionId)
	if err != nil {
		return nil, err
	}

	response := &model.BookingTreatmentResponse{
		BookingId:     booking.ID,
		TransactionId: booking.TransactionId,
		UserFullName:  booking.User.FullName,
		UserEmail:     booking.User.Email,
		PlaceName:     booking.Service.Place.Name,
		ServiceName:   booking.Service.Name,
		PlaceAddress:  booking.Service.Place.Address,
		PaymentMethod: booking.PaymentMethod,
		PaymentCode:   booking.PaymentCode,
		GrossAmount:   booking.GrossAmount,
		Status:        booking.Status,
		BookAt:        booking.BookAt,
	}

	return response, nil
}

func (b BookingTreatmentUC) GetById(id string) (*entity.BookingTreatment, error) {
	booking, err := b.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return booking, nil
}
