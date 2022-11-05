package usecase

import (
	"bookhub/internal/entity"
	"bookhub/internal/storage"
	"context"
	"fmt"
)

type UserService struct {
	storage storage.UserStorage
}

func NewUserService(userStorage storage.UserStorage) *UserService {
	return &UserService{
		storage: userStorage,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user entity.User) (ID int64, err error) {
	ID, err = us.storage.CreateUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %w", err)
	}
	return ID, nil
}
func (us *UserService) GetUser(ctx context.Context, userID int64) (user entity.User, err error) {
	user, err = us.storage.GetUser(ctx, userID)
	if err != nil {
		return entity.User{}, fmt.Errorf("GetUser: %w", err)
	}
	return user, nil
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
func (us *UserService) UpdateUser(ctx context.Context, user entity.User) (updatedUser entity.User, err error) {
	oldUser, err := us.storage.GetUser(ctx, user.ID)
	if err != nil {
		return entity.User{}, fmt.Errorf("GetUser: %w", err)
	}

	user = us.updateUser(oldUser, user)
	err = us.storage.UpdateUser(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("UpdateUser: %w", err)
	}
	return user, nil
}

func (us *UserService) DeleteUser(ctx context.Context, userID int64) (err error) {
	if err = us.storage.DeleteUser(ctx, userID); err != nil {
		return fmt.Errorf("DeleteUser: %w", err)
	}
	return nil
}
