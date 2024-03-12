package entity

import (
	"github.com/google/uuid"
	"time"
)

type Booking struct {
	ID          string    `gorm:"primaryKey;size:10"`
	UserId      uuid.UUID `gorm:"primaryKey"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PlaceId     uuid.UUID `gorm:"primaryKey"`
	Place       Place     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Day         time.Time `gorm:"type:Date;not null"`
	Time        time.Time `gorm:"type:Time;not null"`
	PaymentId   uint      `gorm:"foreignKey:PaymentID"`
	Payment     Payment   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GrossAmount int       `gorm:"not null"`
	BookAt      time.Time `gorm:"not null"`
}
