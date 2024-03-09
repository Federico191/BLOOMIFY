package entity

import (
	"github.com/google/uuid"
	"time"
)

type Service struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	PlaceId   uuid.UUID `gorm:"primaryKey;varchar(36)"`
	Place     *Place    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name      string    `gorm:"size:255;not null"`
	Problem   string    `gorm:"size:255;not null"`
	Price     uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoUpdateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}
