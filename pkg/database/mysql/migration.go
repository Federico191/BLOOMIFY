package mysql

import (
	"gorm.io/gorm"
	"log"
	"projectIntern/internal/entity"
)

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.Problem{},
		entity.User{},
		entity.Category{},
		entity.Profession{},
		entity.Doctor{},
		entity.BookingDoctor{},
		entity.BookingTreatment{},
		entity.Place{},
		entity.Service{},
		entity.TreatmentReview{},
		entity.DoctorReview{},
		entity.Product{},
	)
	if err != nil {
		log.Fatalf("failed to migrate : %v", err)
	}
}
