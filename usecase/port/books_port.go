package port

import "github.com/gin-gonic/gin"

type BooksInputPort interface {
	GetBook(*gin.Context) error
}

type BooksOutputPort interface {
	GetBook(*gin.Context)
}
