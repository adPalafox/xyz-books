package repository

import (
	"net/http"
	"xyz-books/constant"
	"xyz-books/infrastructure/sqlite"
	er "xyz-books/utils/error"
	log "xyz-books/utils/logger"

	"github.com/gin-gonic/gin"
)

type PostProcessRepository struct {
}

func NewPostProcessRepository() PostProcessRepository {
	return PostProcessRepository{}
}

func (i PostProcessRepository) Finish(c *gin.Context, err *error) {
	defer log.WithContext(c).Info(constant.LogStartMessage)
	log.WithContext(c).Info(constant.LogStartMessage)

	if perr := recover(); perr != nil {
		log.WithContext(c).Error(perr)
		*err = er.WithContextError(
			c, http.StatusInternalServerError, constant.ResponseInternalServerMessage)
	}
	sqlite.Close(c, *err == nil)
}
