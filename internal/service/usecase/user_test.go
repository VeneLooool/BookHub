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
	errUserNotFound    = errors.New("user not found in database")
	errInternalServErr = errors.New("internal server error")
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func TestUserService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	oldUser := entity.User{ID: 1}
	newUser := entity.User{ID: 1, UserName: "123", Password: "1234", Desc: "12345"}

	mockUserStorage := mock.NewMockUserStorage(ctrl)
	gomock.InOrder(mockUserStorage.EXPECT().UpdateUser(context.Background(), newUser).Return(nil))
	gomock.InOrder(mockUserStorage.EXPECT().GetUser(context.Background(), oldUser.ID).Return(oldUser, nil))
	userService := NewUserService(mockUserStorage)

	user, err := userService.UpdateUser(context.Background(), newUser)
	if err != nil {
		t.Errorf("user.UpdateUser err: %v", err)
	}
	require.Equal(t, user, newUser)
}

func TestUserService_UpdateUser2(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStorage := mock.NewMockUserStorage(ctrl)
	userService := NewUserService(mockUserStorage)

	tests := []test{
		{
			name: "empty user",
			mock: func() {
				mockUserStorage.EXPECT().GetUser(context.Background(), int64(0)).Return(entity.User{}, errUserNotFound)
			},
			res: entity.User{},
			err: errUserNotFound,
		},
		{
			name: "storage error",
			mock: func() {
				mockUserStorage.EXPECT().GetUser(context.Background(), int64(0)).Return(entity.User{}, nil)
				mockUserStorage.EXPECT().UpdateUser(context.Background(), entity.User{}).Return(errInternalServErr)
			},
			res: entity.User{},
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()

			res, err := userService.UpdateUser(context.Background(), entity.User{})
			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
