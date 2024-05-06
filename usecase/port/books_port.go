package port

import "github.com/gin-gonic/gin"

type BooksInputPort interface {
	ListBooks(*gin.Context) error
	GetBook(*gin.Context) error
}

type BooksOutputPort interface {
	ListBooks(*gin.Context)
	GetBook(*gin.Context)
}
