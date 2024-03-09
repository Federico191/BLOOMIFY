package entity

import (
	"github.com/google/uuid"
	"time"
)

type Place struct {
	ID         uuid.UUID `gorm:"primaryKey;varchar(36)"`
	Name       string    `gorm:"size:255;not null"`
	Address    string    `gorm:"size:255;not null;unique"`
	City       string    `gorm:"size:255;not null"`
	Contact    string    `gorm:"size:100;not null"`
	Hour       string    `gorm:"not null"`
	PhotoLink  string    `gorm:"size:200"`
	CategoryId uint      `gorm:"not null;"`
	Category   Category  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time `gorm:"autoUpdateTime:milli"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime:milli"`
	Service    []Service
	Review     []Review
}
