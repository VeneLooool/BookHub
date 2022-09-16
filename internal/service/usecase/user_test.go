package usecase

import (
	"bookhub/internal/entity"
	"bookhub/internal/service/usecase/mock"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

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
