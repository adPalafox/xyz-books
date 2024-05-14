CWDIR := $(dir $(lastword $(MAKEFILE_LIST)))

include $(CWDIR)/common.conf
-include $(CWDIR)/.env

COMMON_ENV := \
	PORT=${PORT} \

run-local:
	@env $(COMMON_ENV) go run cmd/server/main.go

run-migrate:
	@env $(COMMON_ENV) go run cmd/db/main.go

prepare-mock:
	mockgen -destination=mocks/port/mock_user_port.go -package=mocks -source=usecase/port/books_port.go BooksInputPort BooksOutputPort
	mockgen -destination=mocks/repository/mock_books_repository.go -package=mocks -source=usecase/repository/books_repository.go BooksRepositoryInterface
	mockgen -destination=mocks/repository/mock_post_process.go -package=mocks -source=usecase/repository/post_process.go PostProcessRepositoryInterface

unit-test:
	go test ./usecase ./infrastructure/repository ./interface/controller -coverprofile=coverage.out -v
	go tool cover -func coverage.out