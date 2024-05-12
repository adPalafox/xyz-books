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
) (*[]entity.Book, error) {
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	dbClient := b.con.GetDBClient(c)
	orderStr := fmt.Sprintf("%s %s", sort, order)

	var books []dao.Book
	err := dbClient.
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

		return nil, err
	}

	bookRecords := te.ToBooks(c, &books)

	return bookRecords, nil
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

func (b BooksRepository) EditBook(c *gin.Context, in *dto.EditBookInput) error {
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	dbClient := b.con.GetDBClient(c)

	// var author dao.Author
	// dbf := db.Where("caller = ?", in).First(&caller)
	// if dbf.Error != nil {
	// 	log.WithContext(c).Warn(dbf.Error)
	// 	return er.WithContextError(
	// 		c, http.StatusBadRequest, "caller is invalid")
	// }

	newBookRecord := dao.Book{
		Title:           in.Book.Title,
		Isbn10:          in.Book.Isbn10,
		Isbn13:          in.Book.Isbn13,
		ListPrice:       in.Book.ListPrice,
		PublicationYear: in.Book.PublicationYear,
		// PublisherID: ,
		ImageUrl: in.Book.ImageUrl,
		Edition:  in.Book.Edition,
		// Authors: ,
	}

	var book dao.Book
	dbf := dbClient.Find(&book, in.Book.ID)
	if dbf.Error != nil {
		lg.WithContext(c).Error(dbf.Error)
		return er.WithContextError(
			c, http.StatusInternalServerError, constant.ResponseInternalServerMessage)
	}

	dbc := dbClient.Updates(&newBookRecord)
	if dbc.Error != nil {
		lg.WithContext(c).Error(dbc.Error)
		return er.WithContextError(
			c, http.StatusInternalServerError, constant.ResponseInternalServerMessage)
	}
	if dbc.RowsAffected == 0 {
		lg.WithContext(c).Warn(constant.ResponseNotfoundMessage)
		return er.WithContextError(
			c, http.StatusNotFound, constant.ResponseNotfoundMessage)
	}

	return nil
}
