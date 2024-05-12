package repository

import (
	"xyz-books/constant"
	"xyz-books/entity"
	"xyz-books/infrastructure/repository/dao"
	lg "xyz-books/utils/logger"

	"github.com/gin-gonic/gin"
)

func ToBook(c *gin.Context, in *dao.Book) *entity.Book {
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	authors := ToAuthors(c, &in.Authors)

	return &entity.Book{
		ID:              in.ID,
		Title:           in.Title,
		Isbn10:          in.Isbn10,
		Isbn13:          in.Isbn13,
		ListPrice:       in.ListPrice,
		PublicationYear: in.PublicationYear,
		Publisher:       in.Publisher.Name,
		ImageUrl:        in.ImageUrl,
		Edition:         in.Edition,
		Authors:         authors,
	}
}

func ToBooks(c *gin.Context, in *[]dao.Book) *[]entity.Book {
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	var bookRecords []entity.Book
	for _, b := range *in {

		authors := ToAuthors(c, &b.Authors)

		bookRecord := &entity.Book{
			ID:              b.ID,
			Title:           b.Title,
			Isbn10:          b.Isbn10,
			Isbn13:          b.Isbn13,
			ListPrice:       b.ListPrice,
			PublicationYear: b.PublicationYear,
			Publisher:       b.Publisher.Name,
			ImageUrl:        b.ImageUrl,
			Edition:         b.Edition,
			Authors:         authors,
		}

		bookRecords = append(bookRecords, *bookRecord)
	}

	return &bookRecords
}

func ToAuthors(c *gin.Context, in *[]dao.Author) *[]entity.Author {
	lg.WithContext(c).Info(constant.LogStartMessage)
	defer lg.WithContext(c).Info(constant.LogFinishMessage)

	var authors []entity.Author
	for _, a := range *in {
		authors = append(authors,
			entity.Author{
				ID:         a.ID,
				FirstName:  a.FirstName,
				LastName:   a.LastName,
				MiddleName: a.MiddleName,
			})
	}

	return &authors
}
