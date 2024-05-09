package controller

import (
	"net/http"
	"xyz-books/constant"
	"xyz-books/usecase/dto"
	"xyz-books/usecase/port"
	er "xyz-books/utils/error"
	log "xyz-books/utils/logger"
	"xyz-books/utils/util"

	"github.com/gin-gonic/gin"
)

type BooksController struct {
	booksInputPort port.BooksInputPort
}

func NewBookController(in port.BooksInputPort) BooksController {
	return BooksController{booksInputPort: in}
}

func (i BooksController) ListBooks(c *gin.Context) {
	log.WithContext(c).Info(constant.LogStartMessage)
	defer log.WithContext(c).Info(constant.LogFinishMessage)

	i.booksInputPort.ListBooks(c)
}

func (i BooksController) GetBook(c *gin.Context) {
	log.WithContext(c).Info(constant.LogStartMessage)
	defer log.WithContext(c).Info(constant.LogFinishMessage)

	id := c.Param("id")
	i.booksInputPort.GetBook(c, &dto.GetBookInput{
		Id: id})
}

func (i BooksController) EditBook(c *gin.Context) {
	log.WithContext(c).Info(constant.LogStartMessage)
	defer log.WithContext(c).Info(constant.LogFinishMessage)

	var bookRequirement dto.EditBookInput
	if err := c.ShouldBindJSON(&bookRequirement.Book); err != nil {
		log.WithContext(c).Warn(err.Error())
		er.WithContextError(
			c, http.StatusBadRequest, constant.ResponseInvalidArgumentMessage)
		return
	}

	book := bookRequirement.Book
	if err := util.ValidateBook(c, book); err != nil {
		return
	}

	if err := util.ValidateAuthors(c, book.Authors); err != nil {
		return
	}

	i.booksInputPort.EditBook(c, &dto.EditBookInput{
		Book: book})
}
