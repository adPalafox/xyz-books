package register

import (
	"xyz-books/interface/controller"
	"xyz-books/interface/presenter"
	"xyz-books/usecase"
)

func BooksInit() *controller.BooksController {
	booksPresenter := presenter.NewBooksPresenter()
	booksUseCase := usecase.NewBookUseCase(booksPresenter)
	booksController := controller.NewBookController(booksUseCase)
	return &booksController
}
