package dto

import "xyz-books/entity"

type ListBooksOuput struct {
	Books *[]entity.Book
}

type GetBookOutput struct {
	Book *entity.Book
}
