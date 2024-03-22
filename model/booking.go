package model

import (
	"github.com/google/uuid"
	"time"
)

type BookingDoctorRequest struct {
	DoctorId      uuid.UUID `json:"doctor_id" binding:"required"`
	Day           string    `json:"day" binding:"required"`
	Time          string    `json:"time" binding:"required"`
	PaymentMethod string    `json:"payment_method" binding:"required"`
}

type BookingTreatmentRequest struct {
	ServiceId     uint   `json:"service_id" binding:"required"`
	Day           string `json:"day" binding:"required"`
	Time          string `json:"time" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}

type BookingTreatmentResponse struct {
	BookingId     string    `json:"booking_id"`
	TransactionId string    `json:"transaction_id"`
	UserFullName  string    `json:"user_full_name"`
	UserEmail     string    `json:"user_email"`
	PlaceName     string    `json:"place_name"`
	ServiceName   string    `json:"service_name"`
	PlaceAddress  string    `json:"service_address"`
	PaymentMethod string    `json:"payment_method"`
	PaymentCode   string    `json:"payment_code"`
	GrossAmount   int64     `json:"gross_amount"`
	Status        string    `json:"status"`
	BookAt        time.Time `json:"book_at"`
}

type BookingDoctorResponse struct {
	BookingId     string    `json:"booking_id"`
	TransactionId string    `json:"transaction_id"`
	UserFullName  string    `json:"user_full_name"`
	UserEmail     string    `json:"user_email"`
	DoctorName    string    `json:"doctor_name"`
	Profession    string    `json:"profession"`
	PaymentMethod string    `json:"payment_method"`
	PaymentCode   string    `json:"payment_code"`
	GrossAmount   int64     `json:"gross_amount"`
	Status        string    `json:"status"`
	BookAt        time.Time `json:"book_at"`
}

type BookingStatus struct {
	TransactionId string `json:"transaction_id"`
}
