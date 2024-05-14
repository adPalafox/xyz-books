package controller

import (
	"net/http"
	"xyz-books/constant"
	"xyz-books/usecase/dto"
	"xyz-books/usecase/port"
	er "xyz-books/utils/error"
	lg "xyz-books/utils/logger"
	"xyz-books/utils/util"

	"github.com/gin-gonic/gin"
)

type BooksController struct {
	booksInputPort port.BooksInputPort
}

func NewBookController(in port.BooksInputPort) BooksController {
	return BooksController{booksInputPort: in}
}

func (b BooksController) ListBooks(c *gin.Context) {
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	length, page, sort, order, err := util.ValidateListBook(c)
	if err != nil {
		return
	}

	_ = b.booksInputPort.ListBooks(c, &dto.ListBookInput{
		Length: length,
		Page:   page,
		Sort:   sort,
		Order:  order,
	})
}

func (b BooksController) GetBook(c *gin.Context) {
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	isbn13 := c.Param("id")
	_ = b.booksInputPort.GetBook(c, &dto.GetBookInput{
		Isbn13: isbn13})
}

func (b BooksController) EditBook(c *gin.Context) {
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	var bookRequirement dto.EditBookInput
	if err := c.ShouldBindJSON(&bookRequirement.Book); err != nil {
		lg.WithContext(c).Warn(err.Error())
		er.WithContextError(
			c, http.StatusBadRequest,
			constant.ResponseInvalidArgumentMessage,
		)
		return
	}

	book := bookRequirement.Book
	if err := util.ValidateBook(c, book); err != nil {
		return
	}

	isbn13 := c.Param("id")
	book.Isbn13 = &isbn13

	_ = b.booksInputPort.EditBook(c, &dto.EditBookInput{
		Book: book,
	})
}
