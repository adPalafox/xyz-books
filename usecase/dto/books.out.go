package dto

import "xyz-books/entity"

type ListBooksOuput struct {
	Books      *[]entity.Book
	TotalCount int64
}

type GetBookOutput struct {
	Book *entity.Book
}
