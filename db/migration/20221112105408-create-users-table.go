package db

import (
	"github.com/dev-parvej/go-api-starter-sql/models"
	"gorm.io/gorm"
)

func (m Migrator) UpCreateUsersTable(db *gorm.DB) {
	db.Migrator().CreateTable(&models.User{})
}

func (m Migrator) DownCreateUsersTable(db *gorm.DB) {
	db.Migrator().DropTable(&models.User{})
}
