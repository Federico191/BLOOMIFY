package entity

import (
	"github.com/google/uuid"
	"time"
)

type Service struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	PlaceID     uuid.UUID `gorm:"size:36;foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name        string    `gorm:"size:255;not null"`
	Problem     string    `gorm:"size:255;not null"`
	Price       uint      `gorm:"not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoUpdateTime:milli"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:milli"`
}
