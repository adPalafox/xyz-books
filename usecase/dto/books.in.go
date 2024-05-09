package dto

import "xyz-books/entity"

type GetBookInput struct {
	Id string
}

type EditBookInput struct {
	Book *entity.Book
}
