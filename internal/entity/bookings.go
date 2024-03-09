package entity

import (
	"github.com/google/uuid"
	"time"
)

type Booking struct {
	UserId  uuid.UUID `gorm:"primaryKey"`
	PlaceId uuid.UUID `gorm:"primaryKey"`
	Day     time.Time `gorm:"type:Date;not null"`
	Time    time.Time `gorm:"type:Time;not null"`
	BookAt  time.Time `gorm:"not null"`
}
