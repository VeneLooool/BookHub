package usecase

import (
	"context"
	"errors"
	"github.com/VeneLooool/BookHub/internal/entity"
	"github.com/VeneLooool/BookHub/internal/service/usecase/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	errBookNotFound  = errors.New("book not found in database")
	errEmptyBookFile = errors.New("book file is empty")
	errEmptyFilePath = errors.New("book path is empty")
)

func TestBookService_CreateBook(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookStorage := mock.NewMockBookStorage(ctrl)
	mockFileManager := mock.NewMockFileManager(ctrl)
	bookService := NewBookService(mockBookStorage, mockFileManager)

	tests := []test{
		{
			name: "normal test",
			mock: func() {
				var (
					file entity.File
					book entity.Book
				)
				file.Name = book.Title + "_" + book.Author
				book.File = file
				mockFileManager.EXPECT().CreateFile(context.Background(), file).Return("", nil)
				mockBookStorage.EXPECT().CreateBook(context.Background(), int64(0), book).Return(int64(10), nil)
			},
			res: int64(10),
			err: nil,
		},
		{
			name: "empty book",
			mock: func() {
				var (
					file entity.File
					book entity.Book
				)
				file.Name = book.Title + "_" + book.Author
				mockFileManager.EXPECT().CreateFile(context.Background(), file).Return("", errEmptyBookFile)
			},
			res: int64(0),
			err: errEmptyBookFile,
		},
		{
			name: "storage error",
			mock: func() {
				var (
					file entity.File
					book entity.Book
				)
				file.Name = book.Title + "_" + book.Author
				book.File = file
				mockFileManager.EXPECT().CreateFile(context.Background(), file).Return("", nil)
				mockBookStorage.EXPECT().CreateBook(context.Background(), int64(0), book).Return(int64(0), errServStorageErr)
			},
			res: int64(0),
			err: errServStorageErr,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()

			res, err := bookService.CreateBook(context.Background(), int64(0), entity.Book{})
			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, errors.Unwrap(err), tc.err)
		})
	}
}

func TestBookService_GetBook(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookStorage := mock.NewMockBookStorage(ctrl)
	mockFileManager := mock.NewMockFileManager(ctrl)
	bookService := NewBookService(mockBookStorage, mockFileManager)

	tests := []test{
		{
			name: "normal test",
			mock: func() {
				mockBookStorage.EXPECT().GetBook(context.Background(), int64(0)).Return(entity.Book{}, nil)
			},
			res: entity.Book{},
			err: nil,
		},
		{
			name: "storage error",
			mock: func() {
				mockBookStorage.EXPECT().GetBook(context.Background(), int64(0)).Return(entity.Book{}, errServStorageErr)
			},
			res: entity.Book{},
			err: errServStorageErr,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()

			res, err := bookService.GetBook(context.Background(), int64(0))
			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, errors.Unwrap(err), tc.err)
		})
	}
}

func TestBookService_UpdateBook(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookStorage := mock.NewMockBookStorage(ctrl)
	mockFileManager := mock.NewMockFileManager(ctrl)
	bookService := NewBookService(mockBookStorage, mockFileManager)

	tests := []test{
		{
			name: "empty book",
			mock: func() {
				mockBookStorage.EXPECT().GetBook(context.Background(), int64(0)).Return(entity.Book{}, errBookNotFound)
			},
			res: entity.Book{},
			err: errBookNotFound,
		},
		{
			name: "storage error",
			mock: func() {
				mockBookStorage.EXPECT().GetBook(context.Background(), int64(0)).Return(entity.Book{}, nil)
				mockBookStorage.EXPECT().UpdateBook(context.Background(), entity.Book{}).Return(errServStorageErr)
			},
			res: entity.Book{},
			err: errServStorageErr,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()

			res, err := bookService.UpdateBook(context.Background(), entity.Book{})
			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, errors.Unwrap(err), tc.err)
		})
	}
}

func TestBookService_DeleteBook(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookStorage := mock.NewMockBookStorage(ctrl)
	mockFileManager := mock.NewMockFileManager(ctrl)
	bookService := NewBookService(mockBookStorage, mockFileManager)

	tests := []test{
		{
			name: "normal test",
			mock: func() {
				mockBookStorage.EXPECT().GetBook(context.Background(), int64(0)).Return(entity.Book{}, nil)
				mockFileManager.EXPECT().DeleteFile(context.Background(), "").Return(nil)
				mockBookStorage.EXPECT().DeleteBook(context.Background(), int64(0)).Return(nil)
			},
			err: nil,
		},
		{
			name: "empty book",
			mock: func() {
				mockBookStorage.EXPECT().GetBook(context.Background(), int64(0)).Return(entity.Book{}, nil)
				mockFileManager.EXPECT().DeleteFile(context.Background(), "").Return(errEmptyFilePath)
			},
			err: errEmptyFilePath,
		},
		{
			name: "storage error",
			mock: func() {
				mockBookStorage.EXPECT().GetBook(context.Background(), int64(0)).Return(entity.Book{}, nil)
				mockFileManager.EXPECT().DeleteFile(context.Background(), "").Return(nil)
				mockBookStorage.EXPECT().DeleteBook(context.Background(), int64(0)).Return(errServStorageErr)
			},
			err: errServStorageErr,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()

			err := bookService.DeleteBook(context.Background(), int64(0))
			require.ErrorIs(t, errors.Unwrap(err), tc.err)
		})
	}
}
