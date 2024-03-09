package mysql

import (
	"gorm.io/gorm"
	"log"
	"projectIntern/internal/entity"
)

func Migration(db *gorm.DB) {
	err := db.Migrator().DropTable(
		entity.User{},
	)
	if err != nil {
		log.Fatalf("failed drop table: %v", err)
	}

	err = db.AutoMigrate(entity.User{})
	if err != nil {
		log.Fatalf("failed to migrate : %v", err)
	}
}
