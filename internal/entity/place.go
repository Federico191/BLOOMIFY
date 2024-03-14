package entity

import "time"

type Place struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	Name       string    `gorm:"size:255;not null"`
	Address    string    `gorm:"size:255;not null;unique"`
	City       string    `gorm:"size:255;not null"`
	Contact    string    `gorm:"size:100;not null"`
	Hour       string    `gorm:"size:20;not null"`
	CategoryId int       `gorm:"foreignKey:CategoryID"`
	Category   Category  `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Services   []Service `gorm:"foreignKey:PlaceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
