package model

import (
	"github.com/google/uuid"
	"time"
)

type BookingDoctorRequest struct {
	DoctorId      uuid.UUID `json:"doctor_id"`
	Day           string    `json:"day"`
	Time          string    `json:"time"`
	PaymentMethod string    `json:"payment_method"`
}

type BookingTreatmentRequest struct {
	ServiceId     uint   `json:"service_id"`
	Day           string `json:"day"`
	Time          string `json:"time"`
	PaymentMethod string `json:"payment_method"`
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
