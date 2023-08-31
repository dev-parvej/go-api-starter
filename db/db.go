package db

import (
	"fmt"

	"github.com/dev-parvej/go-api-starter-sql/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connection() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_HOST"),
		config.Get("DB_PORT"),
		config.Get("DB_NAME"),
	)

	dbInstance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return dbInstance
}

func Model(model interface{}) *gorm.DB {
	dbInstance := connection()

	return dbInstance.Model(model)
}

func Table(table string) *gorm.DB {
	dbInstance := connection()

	return dbInstance.Table(table)
}

func Query() *gorm.DB {
	return connection()
}
