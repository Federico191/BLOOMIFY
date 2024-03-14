package entity

import (
	"time"
)

type Product struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"size:255;not null"`
	ProblemId uint      `gorm:"foreignKey:ID"`
	Problem   Problem   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PhotoLink string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
