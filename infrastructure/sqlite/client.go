package sqlite

import (
	"xyz-books/constant"
	log "xyz-books/utils/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DBClient struct {
	dbClient *Client
}

func NewDBClient() *DBClient {
	return &DBClient{}
}

func (i *DBClient) getDBclient() *Client {
	if i.dbClient != nil {
		return i.dbClient
	}

	m, err := NewSQLClient()
	if err != nil {
		panic(err)
	}
	i.dbClient = m
	return i.dbClient
}

// To set the clientdb withing the context using existing connection
func (i *DBClient) GetDBClient(c *gin.Context) *gorm.DB {
	log.WithContext(c).Info(constant.LogStartMessage)
	defer log.WithContext(c).Info(constant.LogFinishMessage)

	_, exist_new := c.Keys["NEW_DB"]
	if exist_new {
		newClient := i.getDBclient().GormClient.WithContext(c).Begin()
		c.Set("DB", newClient)
		return newClient
	}
	val, exist := c.Keys["DB"]
	if exist {
		return val.(*gorm.DB)
	}
	newClient := i.getDBclient().GormClient.WithContext(c).Begin()
	c.Set("DB", newClient)
	return newClient
}

func Close(c *gin.Context, result bool) {
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
