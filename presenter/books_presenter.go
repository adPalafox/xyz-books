package presenter

import (
	"net/http"
	"xyz-books/constant"
	"xyz-books/usecase/dto"

	"github.com/gin-gonic/gin"
)

type BooksPresenter struct {
}

func NewBooksPresenter() *BooksPresenter {
	return &BooksPresenter{}
}

func (i BooksPresenter) ListBooks(c *gin.Context, in *dto.ListBooksOuput) {
	c.JSON(http.StatusOK, gin.H{
		"data": in.Books,
	})
}

func (i BooksPresenter) GetBook(c *gin.Context, in *dto.GetBookOutput) {
	c.JSON(http.StatusOK, gin.H{
		"data": in.Book,
	})
}

func (i BooksPresenter) EditBook(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": constant.ResponseOKMessage,
	})
}
