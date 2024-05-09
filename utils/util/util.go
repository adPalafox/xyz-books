package util

import (
	"net/http"
	"xyz-books/constant"
	"xyz-books/entity"
	er "xyz-books/utils/error"
	log "xyz-books/utils/logger"

	"github.com/gin-gonic/gin"
)

func ValidateBook(c *gin.Context, book *entity.Book) error {
	if book.Title == "" ||
		book.Publisher == "" ||
		book.ListPrice == 0 ||
		book.PublicationYear == 0 ||
		len(*book.Authors) == 0 ||
		(book.Isbn10 == nil && book.Isbn13 == nil) {
		return er.WithContextError(
			c,
			http.StatusBadRequest,
			constant.ResponseInvalidArgumentMessage)
	}
	return nil
}

func ValidateAuthors(c *gin.Context, authors *[]entity.Author) error {
	for _, a := range *authors {
		if a.FirstName == "" ||
			a.LastName == "" {
			log.WithContext(c).Warn(constant.ResponseInvalidArgumentMessage)
			return er.WithContextError(
				c,
				http.StatusBadRequest,
				constant.ResponseInvalidArgumentMessage)
		}
	}
	return nil
}
