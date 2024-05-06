package usecase

import (
	"xyz-books/usecase/port"

	"github.com/gin-gonic/gin"
)

type BooksUseCase struct {
	booksPort port.BooksOutputPort
}

func NewBookUseCase(
	bp port.BooksOutputPort,
) BooksUseCase {
	return BooksUseCase{
		booksPort: bp,
	}
}

func (i BooksUseCase) ListBooks(c *gin.Context) error {
	i.booksPort.ListBooks(c)
	return nil
}

func (i BooksUseCase) GetBook(c *gin.Context) error {
	i.booksPort.GetBook(c)
	return nil
}
