package entity

import (
	"time"
)

type Service struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	PlaceID     uint      `gorm:"size:36"`
	Place       Place     `gorm:"foreignKey:PlaceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name        string    `gorm:"size:255;not null"`
	ProblemId   uint      `gorm:"foreignKey:ID"`
	Problem     Problem   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price       int       `gorm:"not null"`
	PhotoLink   string    `gorm:"size:255"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	ServiceDate []ServiceDate
	Reviews     []Review `gorm:"foreignKey:ServiceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AvgRating   float64
	Category    string
}

type ServiceDate struct {
	ServiceId uint      `gorm:"primaryKey;foreignKey:ID"`
	Service   Service   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Date      time.Time `gorm:"type:DATE"`
	Time      time.Time `gorm:"type:TIME"`
	IsBook    bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
}
