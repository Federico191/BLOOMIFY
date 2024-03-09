package entity

import (
	"github.com/google/uuid"
	"time"
)

type Review struct {
	ID        uint      `gorm:"primaryKey"`
	PlaceId   uuid.UUID `gorm:"primaryKey"`
	Place     Place     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rating    float64   `gorm:"not null"`
	Review    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}
