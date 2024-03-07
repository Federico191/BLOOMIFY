package entity

import (
	"github.com/google/uuid"
	"time"
)

type Salon struct {
	ID        uuid.UUID `gorm:"primaryKey;size:36"`
	Name      string    `gorm:"size:255;not null"`
	Address   string    `gorm:"size:255;not null;unique"`
	City      string    `gorm:"size:255;not null"`
	Contact   string    `gorm:"size:100;not null"`
	Hour      string    `gorm:"size:20;not null"`
	PhotoLink string    `gorm:"size:200"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}

type SalonTreatment struct {
	SalonId   uuid.UUID `gorm:"primaryKey;varchar(36)"`
	Salon     Salon     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name      string    `gorm:"size:255;not null"`
	Problem   string    `gorm:"size:255;not null"`
	Price     uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoUpdateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}
