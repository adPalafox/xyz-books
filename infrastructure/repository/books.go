package repository

import (
	"fmt"
	"net/http"
	"xyz-books/constant"
	"xyz-books/entity"
	"xyz-books/infrastructure/repository/dao"
	te "xyz-books/infrastructure/repository/to_entity"
	sq "xyz-books/infrastructure/sqlite"
	"xyz-books/usecase/dto"
	er "xyz-books/utils/error"
	lg "xyz-books/utils/logger"

	"github.com/gin-gonic/gin"
)

type BooksRepository struct {
	con sq.DBClientInterface
}

func NewBooksRepository(con sq.DBClientInterface) BooksRepository {
	return BooksRepository{con: con}
}

func (b BooksRepository) ListBooks(
	c *gin.Context, length int, page int, sort string, order string,
) (*[]entity.Book, int64, error) {
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	dbClient := b.con.GetDBClient(c)
	orderStr := fmt.Sprintf("%s %s", sort, order)

	var count int64
	err := dbClient.Model(&dao.Book{}).
		Count(&count).
		Error
	if err != nil {
		er.WithError(c,
			http.StatusInternalServerError,
			constant.ResponseInternalServerMessage)
		return nil, 0, err
	}

	var books []dao.Book
	err = dbClient.
		Preload("Publisher").
		Preload("Authors").
		Limit(length).
		Offset((page - 1) * length).
		Order(orderStr).
		Find(&books).
		Error
	if err != nil {
		er.WithError(c,
			http.StatusInternalServerError,
			constant.ResponseInternalServerMessage)
		return nil, 0, err
	}

	bookRecords := te.ToBooks(c, &books)

	return bookRecords, count, nil
}

func (b BooksRepository) GetBookByISBN(
	c *gin.Context, isbn13 string) (*entity.Book, error) {
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	dbClient := b.con.GetDBClient(c)

	var book dao.Book
	err := dbClient.
		Preload("Publisher").
		Preload("Authors").
		Where("isbn13 = ?", isbn13).
		First(&book).
		Error
	if err != nil {
		er.WithError(c,
			http.StatusInternalServerError,
			constant.ResponseInternalServerMessage)

		return nil, err
	}

	bookRecord := te.ToBook(c, &book)

	return bookRecord, nil
}

func (b BooksRepository) EditBook(
	c *gin.Context, in *dto.EditBookInput) error {
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	dbClient := b.con.GetDBClient(c)

	var publisher dao.Publisher
	dbf := dbClient.
		Where("id = ?", in.Book.PublisherID).
		First(&publisher)
	if dbf.Error != nil {
		lg.WithContext(c).Warn(dbf.Error)
		return er.WithContextError(
			c, http.StatusBadRequest, "publisher is invalid")
	}

	var existingBook dao.Book
	dbf = dbClient.
		Where("isbn13 = ?", in.Book.Isbn13).
		First(&existingBook)
	if dbf.Error != nil {
		lg.WithContext(c).Error(dbf.Error)
		return er.WithContextError(
			c, http.StatusInternalServerError,
			constant.ResponseInternalServerMessage)
	}

	existingBook.Title = in.Book.Title
	existingBook.Isbn10 = in.Book.Isbn10
	existingBook.Isbn13 = in.Book.Isbn13
	existingBook.ListPrice = in.Book.ListPrice
	existingBook.PublicationYear = in.Book.PublicationYear
	existingBook.PublisherID = publisher.ID
	existingBook.ImageUrl = in.Book.ImageUrl
	existingBook.Edition = in.Book.Edition

	dbu := dbClient.
		Where("isbn13 = ?", existingBook.Isbn13).
		Updates(&existingBook)
	if dbu.Error != nil {
		lg.WithContext(c).Error(dbu.Error)
		return er.WithContextError(
			c, http.StatusInternalServerError,
			constant.ResponseInternalServerMessage)
	}

	return nil
}
