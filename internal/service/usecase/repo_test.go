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

var (
	errRepoNotFound   = errors.New("repo not found in database")
	errServStorageErr = errors.New("internal server error")
)

func TestRepoService_UpdateRepo(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepoStorage := mock.NewMockRepoStorage(ctrl)
	repoService := NewRepoService(mockRepoStorage)

	tests := []test{
		{
			name: "empty repo",
			mock: func() {
				mockRepoStorage.EXPECT().GetRepo(context.Background(), int64(0)).Return(entity.Repo{}, errRepoNotFound)
			},
			res: entity.Repo{},
			err: errRepoNotFound,
		},
		{
			name: "storage error",
			mock: func() {
				mockRepoStorage.EXPECT().GetRepo(context.Background(), int64(0)).Return(entity.Repo{}, nil)
				mockRepoStorage.EXPECT().UpdateRepo(context.Background(), entity.Repo{}).Return(errServStorageErr)
			},
			res: entity.Repo{},
			err: errServStorageErr,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()

			res, err := repoService.UpdateRepo(context.Background(), entity.Repo{})
			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
