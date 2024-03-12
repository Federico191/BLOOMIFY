package entity

import (
	"gorm.io/gorm"
)

type Place struct {
	gorm.Model
	Name       string   `gorm:"size:255;not null"`
	Address    string   `gorm:"size:255;not null;unique"`
	City       string   `gorm:"size:255;not null"`
	Contact    string   `gorm:"size:100;not null"`
	Hour       string   `gorm:"size:20;not null"`
	PhotoLink  string   `gorm:"size:255"`
	CategoryId int      `gorm:"foreignKey:CategoryID"`
	Category   Category `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Service    []Service
	Review     []Review
}
