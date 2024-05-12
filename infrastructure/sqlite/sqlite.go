package sqlite

import (
	"time"
	"xyz-books/constant"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Client struct {
	GormClient *gorm.DB
}

func NewSQLClient() (*Client, error) {
	dns := constant.ProgramDatabaseAccessDirectory

	conn, err := gorm.Open(sqlite.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := conn.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	return &Client{GormClient: conn}, nil
}
