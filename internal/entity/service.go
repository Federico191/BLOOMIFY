package entity

import (
	"time"
)

type Service struct {
	ID          uint              `gorm:"primaryKey;autoIncrement"`
	PlaceID     uint              `gorm:"size:36"`
	Place       Place             `gorm:"foreignKey:PlaceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name        string            `gorm:"size:255;not null"`
	ProblemId   uint              `gorm:"foreignKey:ID"`
	Problem     Problem           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price       int               `gorm:"not null"`
	PhotoLink   string            `gorm:"size:255"`
	Description string            `gorm:"type:text"`
	CreatedAt   time.Time         `gorm:"autoCreateTime"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime"`
	Reviews     []TreatmentReview `gorm:"foreignKey:ServiceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AvgRating   float64           `gorm:"-"`
}
