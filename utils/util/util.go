package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"xyz-books/constant"
	"xyz-books/entity"
	"xyz-books/infrastructure/repository/dao"
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

func GetSeedData() (
	seedBooks *[]dao.Book,
	seedPublishers *[]dao.Publisher,
	err error) {
	rawBooks, err := os.ReadFile(constant.ProgramSeedBookDirectory)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading seed data file: %w", err)
	}
	err = json.Unmarshal(rawBooks, &seedBooks)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing seed data: %w", err)
	}

	rawPublishers, err := os.ReadFile(constant.ProgramSeedPublisherDirectory)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading seed data file: %w", err)
	}
	err = json.Unmarshal(rawPublishers, &seedPublishers)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing seed data: %w", err)
	}

	return
}

func StringToPointer(s string) *string {
	return &s
}
