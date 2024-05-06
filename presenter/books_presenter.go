package presenter

import (
	"net/http"
	"xyz-books/constant"

	"github.com/gin-gonic/gin"
)

type BooksPresenter struct {
}

func NewBooksPresenter() *BooksPresenter {
	return &BooksPresenter{}
}

func (i BooksPresenter) ListBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": constant.ResponseOKMessage,
	})
}

func (i BooksPresenter) GetBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": constant.ResponseOKMessage,
	})
}
