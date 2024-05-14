package mocks

import (
	"database/sql"
	"xyz-books/constant"
	log "xyz-books/utils/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MockDBClient struct {
	db *sql.DB
}

func NewMockDBClient(d *sql.DB) *MockDBClient {
	return &MockDBClient{db: d}
}

func (i *MockDBClient) GetDBClient(c *gin.Context) *gorm.DB {

	gdb, err := gorm.Open(sqlite.Open("./sqlite.db"), &gorm.Config{})
	if err != nil {
		return nil
	}
	return gdb
}

func (i *MockDBClient) Close(c *gin.Context, result bool) {
	log.WithContext(c).Info(constant.LogStartMessage)
	defer log.WithContext(c).Info(constant.LogFinishMessage)

	val, exist := c.Keys["DB"]
	if exist {
		db := val.(*gorm.DB)
		if result {
			db.Commit()
		} else {
			db.Rollback()
		}
		delete(c.Keys, "DB")
	}
}
