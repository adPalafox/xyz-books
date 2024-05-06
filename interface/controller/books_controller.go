package controller

import (
	"xyz-books/usecase/port"

	"github.com/gin-gonic/gin"
)

type BooksController struct {
	booksInputPort port.BooksInputPort
}

func NewBookController(in port.BooksInputPort) BooksController {
	return BooksController{booksInputPort: in}
}

func (i BooksController) ListBooks(c *gin.Context) {
	_ = i.booksInputPort.ListBooks(c)
}

func (i BooksController) GetBook(c *gin.Context) {
	_ = i.booksInputPort.GetBook(c)
}
