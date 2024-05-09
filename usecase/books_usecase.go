package usecase

import (
	"xyz-books/constant"
	"xyz-books/entity"
	"xyz-books/usecase/dto"
	"xyz-books/usecase/port"
	log "xyz-books/utils/logger"

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

func (i BooksUseCase) ListBooks(c *gin.Context) {
	// close db postprocess
	log.WithContext(c).Info(constant.LogStartMessage)
	defer log.WithContext(c).Info(constant.LogFinishMessage)

	books := []entity.Book{}

	i.booksPort.ListBooks(c, &dto.ListBooksOuput{
		Books: &books})
}

func (i BooksUseCase) GetBook(c *gin.Context, in *dto.GetBookInput) {
	// close db postprocess
	log.WithContext(c).Info(constant.LogStartMessage)
	defer log.WithContext(c).Info(constant.LogFinishMessage)

	book := entity.Book{}

	i.booksPort.GetBook(c, &dto.GetBookOutput{
		Book: &book})
}

func (i BooksUseCase) EditBook(c *gin.Context, in *dto.EditBookInput) {
	// close db postprocess
	log.WithContext(c).Info(constant.LogStartMessage)
	defer log.WithContext(c).Info(constant.LogFinishMessage)

	i.booksPort.EditBook(c)
}
