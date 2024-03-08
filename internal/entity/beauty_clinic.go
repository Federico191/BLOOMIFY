package entity

import (
	"github.com/google/uuid"
	"time"
)

type BeautyClinic struct {
	ID         uuid.UUID `gorm:"primaryKey;varchar(36)"`
	Name       string    `gorm:"size:255;not null"`
	Address    string    `gorm:"size:255;not null;unique"`
	City       string    `gorm:"size:255;not null"`
	Contact    string    `gorm:"size:100;not null"`
	Hour       string    `gorm:"not null"`
	PhotoLink  string    `gorm:"size:200"`
	CreatedAt  time.Time `gorm:"autoUpdateTime:milli"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime:milli"`
	Treatments []BeautyClinicTreatment
}

type BeautyClinicTreatment struct {
	BeautyClinicId uuid.UUID    `gorm:"primaryKey;varchar(36)"`
	BeautyClinic   BeautyClinic `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name           string       `gorm:"size:255;not null"`
	Problem        string       `gorm:"size:255;not null"`
	Price          uint         `gorm:"not null"`
	CreatedAt      time.Time    `gorm:"autoUpdateTime:milli"`
	UpdatedAt      time.Time    `gorm:"autoUpdateTime:milli"`
}
