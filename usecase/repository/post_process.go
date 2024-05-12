package repository

import (
	"github.com/gin-gonic/gin"
)

type PostProcessRepositoryInterface interface {
	Finish(*gin.Context, *error)
}
