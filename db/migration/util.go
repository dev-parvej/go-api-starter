package db

import (
	"fmt"

	"github.com/dev-parvej/go-api-starter-sql/util"
)

func MigrationTemplate(purpose string) string {
	return fmt.Sprintf(`package db

import (
	"gorm.io/gorm"
)

func (m Migrator) Up%s(db *gorm.DB) {
	
}

func (m Migrator) Down%s(db *gorm.DB) {
	
}`, util.Title(purpose), util.Title(purpose))
}

type Migration struct {
	Id        uint `json:"id" gorm:"column=id"`
	Migration string
	Batch     uint
}
