package util

import (
	"net/http"
	"strconv"
	"xyz-books/constant"
	"xyz-books/entity"
	er "xyz-books/utils/error"
	lg "xyz-books/utils/logger"

	"github.com/gin-gonic/gin"
)

func ValidateListBook(c *gin.Context) (int, int, string, string, error) {
	keys := []string{"length", "page", "sort"}
	for _, key := range keys {
		if _, ok := c.Request.URL.Query()[key]; !ok {
			lg.WithContext(c).Warn("Not passed key: ", key)
			return 0, 0, "", "", er.WithContextError(
				c, http.StatusBadRequest, constant.ResponseInvalidParameterMessage)
		}
	}

	intLength, err := strconv.Atoi(c.Request.URL.Query()["length"][0])
	if err != nil {
		lg.WithContext(c).Warn(err.Error())
		return 0, 0, "", "", er.WithContextError(
			c, http.StatusBadRequest, constant.ResponseInvalidParameterMessage)
	}
	intPage, err := strconv.Atoi(c.Request.URL.Query()["page"][0])
	if err != nil {
		lg.WithContext(c).Warn(err.Error())
		return 0, 0, "", "", er.WithContextError(
			c, http.StatusBadRequest, constant.ResponseInvalidParameterMessage)
	}
	sort := c.Request.URL.Query()["sort"][0]
	sortMap := map[string]bool{
		constant.SortByTitle:           true,
		constant.SortByListPrice:       true,
		constant.SortByPublicationYear: true,
	}

	_, exist := sortMap[sort]
	if !exist {
		lg.WithContext(c).Warn("Invalid sort key: ", sort)
		return 0, 0, "", "", er.WithContextError(
			c, http.StatusBadRequest, constant.ResponseInvalidParameterMessage)
	}

	order := c.Request.URL.Query()["order"][0]
	orderMap := map[string]bool{
		constant.OrderByAsc: true,
		constant.OrderByDsc: true,
	}

	_, exist = orderMap[order]
	if !exist {
		lg.WithContext(c).Warn("Invalid order key: ", sort)
		return 0, 0, "", "", er.WithContextError(
			c, http.StatusBadRequest, constant.ResponseInvalidParameterMessage)
	}

	return intLength, intPage, sort, order, nil
}

func ValidateBook(c *gin.Context, book *entity.Book) error {
	if book.Title == "" ||
		book.PublisherID == 0 ||
		book.ListPrice == 0 ||
		book.PublicationYear == 0 ||
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
		if a.ID == 0 {
			lg.WithContext(c).Warn(constant.ResponseInvalidArgumentMessage)
			return er.WithContextError(
				c,
				http.StatusBadRequest,
				constant.ResponseInvalidArgumentMessage)
		}
	}
	return nil
}
