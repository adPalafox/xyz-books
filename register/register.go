package register

import (
	"xyz-books/infrastructure/repository"
	"xyz-books/infrastructure/sqlite"
	"xyz-books/interface/controller"
	"xyz-books/interface/presenter"
	"xyz-books/usecase"
)

func BooksInit() *controller.BooksController {
	dbClient := sqlite.NewDBClient()
	postProcess := repository.NewPostProcessRepository()
	booksRepository := repository.NewBooksRepository(dbClient)
	booksPresenter := presenter.NewBooksPresenter()
	booksUseCase := usecase.NewBookUseCase(booksRepository, postProcess, booksPresenter)
	booksController := controller.NewBookController(booksUseCase)
	return &booksController
}
