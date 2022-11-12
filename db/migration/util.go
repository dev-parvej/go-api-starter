package db

import (
	"fmt"
)

func MigrationTemplate(purpose string) string {
	return fmt.Sprintf(`package db

import (
	"gorm.io/gorm"
)

func (m Migrator) Up%s(db *gorm.DB) {
	
}

func (m Migrator) Down%s(db *gorm.DB) {
	
}`, purpose, purpose)
}

type Migration struct {
	Id        uint `json:"id" gorm:"column=id"`
	Migration string
	Batch     uint
}
