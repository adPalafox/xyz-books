package usecase

import (
	"xyz-books/constant"
	"xyz-books/usecase/dto"
	"xyz-books/usecase/port"
	"xyz-books/usecase/repository"
	lg "xyz-books/utils/logger"

	"github.com/gin-gonic/gin"
)

type BooksUseCase struct {
	booksRepository       repository.BooksRepositoryInterface
	postProcessRepository repository.PostProcessRepositoryInterface
	booksPort             port.BooksOutputPort
}

func NewBookUseCase(
	br repository.BooksRepositoryInterface,
	pp repository.PostProcessRepositoryInterface,
	bp port.BooksOutputPort,
) BooksUseCase {
	return BooksUseCase{
		booksRepository:       br,
		postProcessRepository: pp,
		booksPort:             bp,
	}
}

func (b BooksUseCase) ListBooks(c *gin.Context, in *dto.ListBookInput,
) (err error) {
	defer b.postProcessRepository.Finish(c, &err)
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	books, err := b.booksRepository.ListBooks(
		c, in.Length, in.Page, in.Sort, in.Order)
	if err != nil {
		return
	}

	b.booksPort.ListBooks(c, &dto.ListBooksOuput{
		Books: books})
	return
}

func (b BooksUseCase) GetBook(c *gin.Context, in *dto.GetBookInput,
) (err error) {
	defer b.postProcessRepository.Finish(c, &err)
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	book, err := b.booksRepository.GetBookByISBN(c, in.Isbn13)
	if err != nil {
		return
	}

	b.booksPort.GetBook(c, &dto.GetBookOutput{
		Book: book})
	return
}

func (b BooksUseCase) EditBook(c *gin.Context, in *dto.EditBookInput,
) (err error) {
	defer b.postProcessRepository.Finish(c, &err)
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	err = b.booksRepository.EditBook(c, in)
	if err != nil {
		return
	}

	b.booksPort.EditBook(c)
	return
}
