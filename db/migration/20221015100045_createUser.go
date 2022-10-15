package db

import (
	"github.com/dev-parvej/go-api-starter/models"
	"gorm.io/gorm"
)

func (m Migrator) UpCreateuser(db *gorm.DB) {
	db.Migrator().CreateTable(&models.User{})
}

func (m Migrator) DownCreateuser(db *gorm.DB) {
	db.Migrator().DropTable(&models.User{})
}
