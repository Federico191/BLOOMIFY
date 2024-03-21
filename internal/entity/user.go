package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID                uuid.UUID          `gorm:"type:varchar(36);primaryKey"`
	Email             string             `gorm:"size:50;not null;unique"`
	FullName          string             `gorm:"size:255;not null"`
	PhotoLink         string             `gorm:"size:255"`
	Password          string             `gorm:"size:100;not null"`
	VerificationCode  string             `gorm:"size:30"`
	IsVerified        bool               `gorm:"default:false"`
	ProblemID         uint               `gorm:"foreignKey:ID"`
	Problem           Problem            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt         time.Time          `gorm:"autoCreateTime:milli"`
	UpdatedAt         time.Time          `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	TreatmentReviews  []TreatmentReview  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DoctorReviews     []DoctorReview     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BookingTreatments []BookingTreatment `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
