package entity

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"size:10; unique"`
	Email     string    `gorm:"size:50, not null"`
	FullName  string    `gorm:"size:255; not null"`
	Avatar    string    `gorm:"size:255"`
	Password  string    `gorm:"size:100"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
}
