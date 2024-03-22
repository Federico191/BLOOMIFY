package entity

import (
	"github.com/google/uuid"
	"time"
)

type BookingTreatment struct {
	ID            string    `gorm:"primaryKey;size:100"`
	UserId        uuid.UUID `gorm:"primaryKey"`
	User          User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ServiceId     uint      `gorm:"primaryKey"`
	Service       Service   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PaymentMethod string    `gorm:"size:10;not null"`
	TransactionId string    `gorm:"size:100;not null"`
	GrossAmount   int64     `gorm:"not null"`
	PaymentCode   string    `gorm:"size:20;not null"`
	Status        string    `gorm:"not null"`
	BookAt        time.Time `gorm:"type:TIMESTAMP"`
}

type BookingDoctor struct {
	ID            string    `gorm:"primaryKey;size:100"`
	UserId        uuid.UUID `gorm:"primaryKey"`
	User          User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DoctorId      uuid.UUID `gorm:"primaryKey"`
	Doctor        Doctor    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PaymentMethod string    `gorm:"size:10;not null"`
	TransactionId string    `gorm:"size:100;not null"`
	GrossAmount   int64     `gorm:"not null"`
	PaymentCode   string    `gorm:"size:20;not null"`
	Status        string    `gorm:"not null"`
	BookAt        time.Time `gorm:"type:TIMESTAMP"`
}
