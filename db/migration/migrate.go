package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/dev-parvej/go-api-starter-sql/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MIGRATE            = "db:migrate"
	CREATE_MIGRATION   = "db:create_migration"
	ROLLBACK_MIGRATION = "db:rollback"
)

func MigrateConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_MIGRATION_HOST"),
		config.Get("DB_PORT"),
		config.Get("DB_NAME"),
	)

	dbInstance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return dbInstance
}

func Migrate(action string) {
	path, _ := os.Getwd()
	folderName := path + "/db/migration/"

	if action == MIGRATE {
		files, _ := ioutil.ReadDir(folderName)
		migrateDatabase(MigrateConnection(), files)
	} else if action == CREATE_MIGRATION {
		purpose := os.Args[2]
		if purpose == "" {
			log.Fatal("Purpose can not be empty")
		}
		fileName := folderName + time.Now().Format("20060102150405") + "_" + purpose + ".go"
		createMigration(purpose, fileName)
	} else if action == ROLLBACK_MIGRATION {
		if !MigrateConnection().Migrator().HasTable("migrations") {
			log.Default().Println("No migration to rollback")
			return
		}

		migrations := []Migration{}
		MigrateConnection().Raw("SELECT * FROM migrations where `batch`= (select Max(`batch`) from migrations) Order By id desc").Find(&migrations)
		rollbackMigration(MigrateConnection(), migrations)
	}
}

func migrateDatabase(dbInstance *gorm.DB, files []os.FileInfo) {
	if !dbInstance.Migrator().HasTable("migrations") {
		dbInstance.Exec("CREATE TABLE migrations (id int NOT NULL AUTO_INCREMENT, migration VARCHAR(255), batch int, CONSTRAINT PK_ PRIMARY KEY (id));")
		log.Default().Println("Created migration table")
	}

	maxMigration := Migration{}

	dbInstance.Table("migrations").Order("id DESC").First(&maxMigration)
	batch := maxMigration.Batch

	for _, file := range files {
		if file.Name() == "migrate.go" || file.Name() == "util.go" || file.Name() == "migrator.go" {
			continue
		}
		migration := Migration{}
		dbInstance.Table("migrations").Where("migration = ?", file.Name()).Scan(&migration)

		if migration.Id != 0 {
			continue
		}

		splitted := strings.Split(file.Name(), "_")
		functionName := "Up" + cases.Title(language.English).String(strings.ReplaceAll(splitted[1], ".go", ""))

		migrator := Migrator{}

		log.Default().Println("Migrating", file.Name())
		err := reflect.ValueOf(migrator).MethodByName(functionName).Call([]reflect.Value{reflect.ValueOf(dbInstance)})
		if err != nil {
			log.Fatal(err)
		}
		migration.Batch = batch + 1
		migration.Migration = file.Name()

		dbInstance.Table("migrations").Create(&migration)

		log.Default().Println("Migrated", file.Name())
	}
}

func rollbackMigration(dbInstance *gorm.DB, files []Migration) {
	for _, file := range files {
		splitted := strings.Split(file.Migration, "_")
		functionName := "Down" + cases.Title(language.English).String(strings.ReplaceAll(splitted[1], ".go", ""))

		migrator := Migrator{}
		log.Default().Println("Roll backing migration", file.Migration)
		err := reflect.ValueOf(migrator).MethodByName(functionName).Call([]reflect.Value{reflect.ValueOf(dbInstance)})

		if err != nil {
			log.Fatal(err)
		}

		dbInstance.Table("migrations").Where("migration=?", file.Migration).Delete(&file)

		log.Default().Println("Roll backed successfully", file.Migration)
	}
}

func createMigration(purpose string, fileName string) {
	migrationTemplate := MigrationTemplate(purpose)

	err := ioutil.WriteFile(
		fileName,
		[]byte(migrationTemplate),
		os.ModePerm,
	)

	if err != nil {
		log.Fatal("Failed to create migration")
		os.Exit(1)
	}

	fmt.Printf("Migration %s roll backed successfully", fileName)
}
