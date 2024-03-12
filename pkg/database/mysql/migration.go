package mysql

import (
	"gorm.io/gorm"
	"log"
	"projectIntern/internal/entity"
)

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(entity.User{},
		entity.Category{},
		entity.Payment{},
		entity.Booking{},
		entity.Problem{},
		entity.Place{},
		entity.Service{},
		entity.Review{},
		entity.ServiceDate{},
		entity.Product{},
	)
	if err != nil {
		log.Fatalf("failed to migrate : %v", err)
	}
}
