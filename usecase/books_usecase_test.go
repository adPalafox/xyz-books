package usecase

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"xyz-books/entity"
	p_mock "xyz-books/mocks/port"
	r_mock "xyz-books/mocks/repository"
	"xyz-books/usecase/dto"
	"xyz-books/utils/context"
	util "xyz-books/utils/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestBooksUsecase_ListBooks(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBooksRepository := r_mock.NewMockBooksRepositoryInterface(mockCtrl)
	mockPostProcessRepository := r_mock.NewMockPostProcessRepositoryInterface(mockCtrl)
	mockBooksOutputPort := p_mock.NewMockBooksOutputPort(mockCtrl)

	type args struct {
		c  *gin.Context
		in *dto.ListBookInput
	}
	tests := []struct {
		name    string
		before  func(*testing.T, *gin.Context)
		args    args
		wantErr bool
	}{
		{
			name: "Normal system success listing books",
			before: func(t *testing.T, ctx *gin.Context) {
				mockBooksRepository.EXPECT().ListBooks(ctx, 5, 10, "title", "asc").
					Return(
						&[]entity.Book{{ID: 1}, {ID: 2}},
						int64(2), nil)
				mockPostProcessRepository.EXPECT().Finish(ctx, gomock.Any())
				mockBooksOutputPort.EXPECT().ListBooks(
					ctx, &dto.ListBooksOuput{
						Books: &[]entity.Book{
							{ID: 1},
							{ID: 2},
						}, TotalCount: 2,
					})
			},
			args: args{
				c: context.GetTestGinContext(httptest.NewRecorder()),
				in: &dto.ListBookInput{
					Length: 5,
					Page:   10,
					Sort:   "title",
					Order:  "asc",
				},
			},
			wantErr: false,
		},
		{
			name: "Failure system error returned at list book",
			before: func(t *testing.T, ctx *gin.Context) {
				mockBooksRepository.EXPECT().ListBooks(ctx, 5, 10, "title", "asc").
					Return(nil, int64(0), fmt.Errorf("Some error"))
				mockPostProcessRepository.EXPECT().Finish(ctx, gomock.Any())
			},
			args: args{
				c: context.GetTestGinContext(httptest.NewRecorder()),
				in: &dto.ListBookInput{
					Length: 5,
					Page:   10,
					Sort:   "title",
					Order:  "asc",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before(t, tt.args.c)
			}
			i := BooksUseCase{
				booksRepository:       mockBooksRepository,
				postProcessRepository: mockPostProcessRepository,
				booksPort:             mockBooksOutputPort,
			}
			if err := i.ListBooks(tt.args.c, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("BooksUseCase.ListBooks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBooksUsecase_GetBook(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBooksRepository := r_mock.NewMockBooksRepositoryInterface(mockCtrl)
	mockPostProcessRepository := r_mock.NewMockPostProcessRepositoryInterface(mockCtrl)
	mockBooksOutputPort := p_mock.NewMockBooksOutputPort(mockCtrl)

	type args struct {
		c  *gin.Context
		in *dto.GetBookInput
	}
	tests := []struct {
		name    string
		before  func(*testing.T, *gin.Context)
		args    args
		wantErr bool
	}{
		{
			name: "Normal system success getting book",
			before: func(t *testing.T, ctx *gin.Context) {
				mockBooksRepository.EXPECT().GetBookByISBN(ctx, "Test Isbn13").
					Return(&entity.Book{ID: 1}, nil)
				mockPostProcessRepository.EXPECT().Finish(ctx, gomock.Any())
				mockBooksOutputPort.EXPECT().GetBook(
					ctx, &dto.GetBookOutput{
						Book: &entity.Book{ID: 1}})
			},
			args: args{
				c: context.GetTestGinContext(httptest.NewRecorder()),
				in: &dto.GetBookInput{
					Isbn13: "Test Isbn13",
				},
			},
			wantErr: false,
		},
		{
			name: "Failure system error getting book",
			before: func(t *testing.T, ctx *gin.Context) {
				mockBooksRepository.EXPECT().GetBookByISBN(ctx, "Test Isbn13").
					Return(nil, fmt.Errorf("Some error"))
				mockPostProcessRepository.EXPECT().Finish(ctx, gomock.Any())
			},
			args: args{
				c: context.GetTestGinContext(httptest.NewRecorder()),
				in: &dto.GetBookInput{
					Isbn13: "Test Isbn13",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before(t, tt.args.c)
			}
			i := BooksUseCase{
				booksRepository:       mockBooksRepository,
				postProcessRepository: mockPostProcessRepository,
				booksPort:             mockBooksOutputPort,
			}
			if err := i.GetBook(tt.args.c, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("BooksUseCase.GetBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBooksUsecase_EditBook(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBooksRepository := r_mock.NewMockBooksRepositoryInterface(mockCtrl)
	mockPostProcessRepository := r_mock.NewMockPostProcessRepositoryInterface(mockCtrl)
	mockBooksOutputPort := p_mock.NewMockBooksOutputPort(mockCtrl)

	type args struct {
		c  *gin.Context
		in *dto.EditBookInput
	}

	testInput := &entity.Book{
		Title:  "Test Title",
		Isbn13: util.StringToPointer("Test Isbn13"),
	}

	tests := []struct {
		name    string
		before  func(*testing.T, *gin.Context)
		args    args
		wantErr bool
	}{
		{
			name: "Normal system success editing book",
			before: func(t *testing.T, ctx *gin.Context) {
				mockBooksRepository.EXPECT().EditBook(ctx, &dto.EditBookInput{Book: testInput}).
					Return(nil)
				mockPostProcessRepository.EXPECT().Finish(ctx, gomock.Any())
				mockBooksOutputPort.EXPECT().EditBook(ctx)
			},
			args: args{
				c: context.GetTestGinContext(httptest.NewRecorder()),
				in: &dto.EditBookInput{
					Book: &entity.Book{
						Title:  "Test Title",
						Isbn13: util.StringToPointer("Test Isbn13"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Failure system error editing book",
			before: func(t *testing.T, ctx *gin.Context) {
				mockBooksRepository.EXPECT().EditBook(ctx, &dto.EditBookInput{Book: testInput}).
					Return(fmt.Errorf("Some error"))
				mockPostProcessRepository.EXPECT().Finish(ctx, gomock.Any())
			},
			args: args{
				c: context.GetTestGinContext(httptest.NewRecorder()),
				in: &dto.EditBookInput{
					Book: &entity.Book{
						Title:  "Test Title",
						Isbn13: util.StringToPointer("Test Isbn13"),
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before(t, tt.args.c)
			}
			i := BooksUseCase{
				booksRepository:       mockBooksRepository,
				postProcessRepository: mockPostProcessRepository,
				booksPort:             mockBooksOutputPort,
			}
			if err := i.EditBook(tt.args.c, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("BooksUseCase.EditBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
