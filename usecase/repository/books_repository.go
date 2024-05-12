package repository

import (
	"xyz-books/entity"
	"xyz-books/usecase/dto"

	"github.com/gin-gonic/gin"
)

type BooksRepositoryInterface interface {
	ListBooks(*gin.Context, int, int, string, string) (*[]entity.Book, error)
	GetBookByISBN(*gin.Context, string) (*entity.Book, error)
	EditBook(*gin.Context, *dto.EditBookInput) error
}
