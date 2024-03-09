package entity

import (
	"github.com/google/uuid"
	"time"
)

type Class struct {
	ID        uint      `gorm:"primaryKey"`
	PlaceId   uuid.UUID `gorm:"primaryKey"`
	Place     Place     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name      string    `gorm:"size:255;not null"`
	Day       time.Time `gorm:"type:Date;not null"`
	Time      time.Time `gorm:"type:Time;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}
