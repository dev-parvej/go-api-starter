package db

import (
	"github.com/dev-parvej/go-api-starter/models"
	"gorm.io/gorm"
)

func (m Migrator) UpUserprofile(db *gorm.DB) {
	db.Migrator().CreateTable(&models.Profile{})
}

func (m Migrator) DownUserprofile(db *gorm.DB) {
	db.Migrator().DropTable(&models.Profile{})
}
