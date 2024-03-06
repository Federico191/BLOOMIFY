package entity

import (
	"github.com/google/uuid"
	"time"
)

type BeautyClinic struct {
	ID        uuid.UUID `gorm:"primaryKey;size:36"`
	Name      string    `gorm:"size:255;not null;unique"`
	Address   string    `gorm:"size:255;not null;unique"`
	City      string    `gorm:"size:255;not null"`
	Contact   string    `gorm:"size:100;not null"`
	Rating    float64   `gorm:"not null"`
	Hour      uint      `gorm:"not null"`
	PhotoLink string    `gorm:"size:200"`
	CreatedAt time.Time `gorm:"autoUpdateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}
