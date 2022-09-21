package usecase

import (
	"bookhub/internal/entity"
	"bookhub/internal/service/usecase/mock"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var errBookNotFound = errors.New("repo not found in database")

func TestBookService_UpdateBook(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookStorage := mock.NewMockBookStorage(ctrl)
	bookService := NewBookService(mockBookStorage)

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
			require.ErrorIs(t, err, tc.err)
		})
	}
}
