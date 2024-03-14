package entity

import (
	"github.com/google/uuid"
	"time"
)

type Booking struct {
	ID          string    `gorm:"primaryKey;size:10"`
	UserId      uuid.UUID `gorm:"primaryKey"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ServiceId   uint      `gorm:"primaryKey"`
	Service     Service   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PaymentId   uint      `gorm:"foreignKey:PaymentID"`
	Payment     Payment   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GrossAmount int       `gorm:"not null"`
	Status      string    `gorm:"size:255"`
	BookAt      time.Time `gorm:"autoCreateTime"`
}
