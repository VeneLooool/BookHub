package usecase

import (
	"bookhub/internal/entity"
	"context"
	"fmt"
)

type UserStorage interface {
	CreateUser(context.Context, entity.User) (int64, error)
	GetUser(context.Context, int64) (entity.User, error)
	GetReposForUser(ctx context.Context, userID int64) ([]entity.Repo, error)
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

func (us *UserService) GetReposForUser(ctx context.Context, userID int64) (repos []entity.Repo, err error) {
	repos, err = us.storage.GetReposForUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("GetReposForUser: %w", err)
	}
	return repos, nil
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
