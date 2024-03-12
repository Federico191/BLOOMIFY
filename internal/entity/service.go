package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	gorm.Model
	PlaceID     uuid.UUID `gorm:"size:36;foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name        string    `gorm:"size:255;not null"`
	ProblemId   uint      `gorm:"foreignKey:ID"`
	Problem     Problem   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price       uint      `gorm:"not null"`
	PhotoLink   string    `gorm:"size:255"`
	Description string    `gorm:"type:text"`
	ServiceDate []ServiceDate
}

type ServiceDate struct {
	ServiceId uint      `gorm:"primaryKey;foreignKey:ID"`
	Service   Service   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Date      time.Time `gorm:"type:DATE"`
	Time      time.Time `gorm:"type:TIME"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
}
