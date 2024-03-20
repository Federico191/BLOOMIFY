package entity

import (
	"github.com/google/uuid"
	"time"
)

type Doctor struct {
	ID           uuid.UUID      `gorm:"primaryKey;size:36"`
	Name         string         `gorm:"size:100;not null"`
	Price        int            `gorm:"not null"`
	ProfessionId uint           `gorm:"foreignKey:ID"`
	Profession   Profession     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	STRNumber    string         `gorm:"size:100;not null;unique"`
	BirthDate    time.Time      `gorm:"type:date;not null"`
	Alumnus      string         `gorm:"size:100;not null"`
	PracticeSite string         `gorm:"size:100;not null"`
	City         string         `gorm:"size:100;not null"`
	PhotoLink    string         `gorm:"size:255"`
	Rating       float64        `gorm:"-"`
	Reviews      []DoctorReview `gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time      `gorm:"autoCreateTime:milli"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime:milli"`
}
