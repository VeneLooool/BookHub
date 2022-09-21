package usecase

import (
	"bookhub/internal/entity"
	"context"
)

type UserStorage interface {
	CreateUser(context.Context, entity.User) (int64, error)
	GetUser(context.Context, int64) (entity.User, error)
	UpdateUser(context.Context, entity.User) error
	DeleteUser(context.Context, int64) error
}

type UserService struct {
	storage UserStorage
}

func NewUserService(userStorage UserStorage) *UserService {
	return &UserService{
		storage: userStorage,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user entity.User) (int64, error) {
	return us.storage.CreateUser(ctx, user)
}
func (us *UserService) GetUser(ctx context.Context, userID int64) (entity.User, error) {
	return us.storage.GetUser(ctx, userID)
}

func (us *UserService) updateUser(oldUser, newUser entity.User) entity.User {
	if newUser.UserName != "" {
		oldUser.UserName = newUser.UserName
	}
	if newUser.Password != "" {
		oldUser.Password = newUser.Password
	}
	if newUser.Desc != "" {
		oldUser.Desc = newUser.Desc
	}
	return oldUser
}
func (us *UserService) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	oldUser, err := us.storage.GetUser(ctx, user.ID)
	if err != nil {
		return entity.User{}, err
	}
	
	user = us.updateUser(oldUser, user)
	err = us.storage.UpdateUser(ctx, user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (us *UserService) DeleteUser(ctx context.Context, userID int64) error {
	return us.storage.DeleteUser(ctx, userID)
}
