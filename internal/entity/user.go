package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID         uuid.UUID `gorm:"primaryKey;varchar(36)"`
	Email      string    `gorm:"size:50;not null;unique"`
	FullName   string    `gorm:"size:255; not null"`
	Avatar     string    `gorm:"size:255"`
	Password   string    `gorm:"size:100"`
	IsVerified bool      `gorm:"default:false"`
	CreatedAt  time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt  time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
}
