package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name      string  `gorm:"size:255;not null"`
	ProblemId uint    `gorm:"foreignKey:ID"`
	Problem   Problem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PhotoLink string  `gorm:"size:255"`
}
