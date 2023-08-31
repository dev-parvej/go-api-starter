package db

import (
	"github.com/dev-parvej/go-api-starter-sql/models"
	"gorm.io/gorm"
)

func (m Migrator) UpCreateRefreshTokensTable(db *gorm.DB) {
	db.Migrator().CreateTable(&models.RefreshToken{})
}

func (m Migrator) DownCreateRefreshTokensTable(db *gorm.DB) {
	db.Migrator().DropTable(&models.RefreshToken{})
}
