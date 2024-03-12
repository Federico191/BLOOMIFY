package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID  uuid.UUID `gorm:"size:36;foreignKey:ID"`
	User    User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PlaceID uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rating  int       `gorm:"not null"`
	Review  string    `gorm:"not null"`
}
