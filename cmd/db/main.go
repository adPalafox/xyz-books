package main

import (
	"log"
	"xyz-books/constant"
	"xyz-books/infrastructure/repository/dao"
	"xyz-books/utils/util"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	log.Println(constant.LogStartMessage)
	defer log.Println(constant.LogFinishMessage)

	db, err := connectToDB()
	if err != nil {
		panic(constant.LogDatabaseFailureToConnect)
	}

	err = migrateDB(db)
	if err != nil {
		log.Printf("%v: %v", constant.LogDatabaseFailureToMigrate, err)
		return
	}

	log.Println(constant.LogDatabaseSuccessMigration)
}

func migrateDB(db *gorm.DB) error {
	log.Println(constant.LogDatabaseRunningMigration)

	seedBooks, seedPublishers, err := util.GetSeedData()
	if err != nil {
		return err
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202405111600",
			Migrate: func(d *gorm.DB) error {
				d.AutoMigrate(
					&dao.Publisher{},
					&dao.Book{},
					&dao.Author{},
				)
				d.Create(&seedBooks)
				d.Create(&seedPublishers)
				return nil
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable(
					&dao.Book{},
					&dao.Author{},
					&dao.Publisher{},
				)
			},
		},
	})
	if err := m.Migrate(); err != nil {
		return err
	}

	return nil
}

func connectToDB() (*gorm.DB, error) {
	dns := constant.ProgramDatabaseMigrationDirectory

	db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
