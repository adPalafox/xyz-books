package sqlite

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DBClientInterface interface {
	GetDBClient(*gin.Context) *gorm.DB
}
