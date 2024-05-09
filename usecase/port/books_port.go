package port

import (
	"xyz-books/usecase/dto"

	"github.com/gin-gonic/gin"
)

type BooksInputPort interface {
	ListBooks(*gin.Context)
	GetBook(*gin.Context, *dto.GetBookInput)
	EditBook(*gin.Context, *dto.EditBookInput)
}

type BooksOutputPort interface {
	ListBooks(*gin.Context, *dto.ListBooksOuput)
	GetBook(*gin.Context, *dto.GetBookOutput)
	EditBook(*gin.Context)
}
