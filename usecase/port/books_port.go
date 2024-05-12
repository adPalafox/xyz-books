package port

import (
	"xyz-books/usecase/dto"

	"github.com/gin-gonic/gin"
)

type BooksInputPort interface {
	ListBooks(*gin.Context, *dto.ListBookInput) error
	GetBook(*gin.Context, *dto.GetBookInput) error
	EditBook(*gin.Context, *dto.EditBookInput) error
}

type BooksOutputPort interface {
	ListBooks(*gin.Context, *dto.ListBooksOuput)
	GetBook(*gin.Context, *dto.GetBookOutput)
	EditBook(*gin.Context)
}
