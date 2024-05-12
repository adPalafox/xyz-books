package dto

import "xyz-books/entity"

type ListBookInput struct {
	Page   int
	Length int
	Sort   string
	Order  string
}

type GetBookInput struct {
	Isbn13 string
}

type EditBookInput struct {
	Book *entity.Book
}
