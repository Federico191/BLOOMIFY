package entity

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Name string `gorm:"size:100; not null"`
}
