package migrations

import (
	"github.com/mmfalcao/go-basic-api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
