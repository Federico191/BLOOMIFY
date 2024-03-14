package model

import (
	"github.com/google/uuid"
	"time"
)

type BookingRequest struct {
	ID          string    `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	ServiceId   uint      `json:"service_id"`
	Day         time.Time `json:"day"`
	Time        time.Time `json:"time"`
	PaymentId   uint      `json:"payment_id"`
	GrossAmount int       `json:"gross_amount"`
}
