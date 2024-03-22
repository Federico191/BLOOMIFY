package usecase

import (
	"errors"
	"github.com/google/uuid"
	"os"
	"projectIntern/internal/entity"
	"projectIntern/internal/repository"
	"projectIntern/model"
	"projectIntern/pkg/helper"
	"projectIntern/pkg/midtrans"
)

type BookingDoctorUCItf interface {
	Create(id uuid.UUID, req model.BookingDoctorRequest) (*model.BookingDoctorResponse, error)
	Update(bookingId string) error
	GetByStatus(transactionId string) (*model.BookingDoctorResponse, error)
	GetById(id string) (*entity.BookingDoctor, error)
}

type BookingDoctorUC struct {
	repo    repository.BookingDoctorRepoItf
	user    repository.UserRepoItf
	doctor  repository.DoctorRepoItf
	mdtrans midtrans.MidTransItf
}

func NewBookingDoctorUC(repo repository.BookingDoctorRepoItf, user repository.UserRepoItf, doctor repository.DoctorRepoItf, mdtrans midtrans.MidTransItf) BookingDoctorUCItf {
	return &BookingDoctorUC{
		repo:    repo,
		user:    user,
		doctor:  doctor,
		mdtrans: mdtrans,
	}
}

func (b BookingDoctorUC) Create(id uuid.UUID, req model.BookingDoctorRequest) (*model.BookingDoctorResponse, error) {
	user, err := b.user.GetById(id)
	if err != nil {
		return nil, err
	}

	doctor, err := b.doctor.GetById(req.DoctorId)
	if err != nil {
		return nil, err
	}

	bookingId := "doc-" + uuid.NewString()

	transaction, err := b.mdtrans.CreateTransactionDoctor(*doctor, req.PaymentMethod, bookingId, int64(doctor.Price), *user)
	if err != nil {
		return nil, err
	}

	timeParse, err := helper.ParseDateTime(req.Day, req.Time)
	if err != nil {
		return nil, err
	}

	booking := &entity.BookingDoctor{
		ID:            bookingId,
		UserId:        user.ID,
		User:          *user,
		DoctorId:      doctor.ID,
		Doctor:        *doctor,
		PaymentMethod: req.PaymentMethod,
		TransactionId: transaction.TransactionID,
		GrossAmount:   int64(doctor.Price),
		PaymentCode:   transaction.VaNumbers[0].VANumber,
		Status:        "Pending",
		BookAt:        timeParse,
	}

	data, err := b.repo.GetByTimeId(doctor.ID, timeParse)
	if data != nil {
		return nil, errors.New("time already booked")
	}

	err = b.repo.Create(booking)
	if err != nil {
		return nil, err
	}

	response := &model.BookingDoctorResponse{
		BookingId:     bookingId,
		TransactionId: transaction.TransactionID,
		UserFullName:  user.FullName,
		UserEmail:     user.Email,
		DoctorName:    doctor.Name,
		Profession:    doctor.Profession.Name,
		PaymentMethod: booking.PaymentMethod,
		PaymentCode:   booking.PaymentCode,
		GrossAmount:   booking.GrossAmount,
		Status:        "Pending",
		BookAt:        booking.BookAt,
	}

	return response, nil

}

func (b BookingDoctorUC) Update(bookingId string) error {
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

func (b BookingDoctorUC) GetByStatus(transactionId string) (*model.BookingDoctorResponse, error) {
	booking, err := b.repo.GetByStatus(transactionId)
	if err != nil {
		return nil, err
	}

	response := &model.BookingDoctorResponse{
		BookingId:     booking.ID,
		TransactionId: booking.TransactionId,
		UserFullName:  booking.User.FullName,
		UserEmail:     booking.User.Email,
		DoctorName:    booking.Doctor.Name,
		Profession:    booking.Doctor.Profession.Name,
		PaymentMethod: booking.PaymentMethod,
		PaymentCode:   booking.PaymentCode,
		GrossAmount:   booking.GrossAmount,
		Status:        booking.Status,
		BookAt:        booking.BookAt,
	}

	return response, nil
}

func (b BookingDoctorUC) GetById(id string) (*entity.BookingDoctor, error) {
	booking, err := b.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return booking, nil
}
